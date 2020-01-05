#pragma once

#include "model/fundamental.h"
#include "model/record.h"
#include "model/record_batch_reader.h"
#include "model/timestamp.h"
#include "storage2/common.h"
#include "storage2/detail/index.h"
#include "storage2/indices.h"

namespace storage {

class segment;
class log_manager;

/**
 * Represents a single topic-log in the storage subsystem.
 *
 * A log is a logical unit of storage where writes and reads
 * are ordered. Data is stored in a series of segment files that contain
 * contigious list of record_batch-es.
 */
class log {
public:
    /**
     * This enum describes the direction of trimming in a log.
     * It is used as a parameter to log::trim.
     */
    enum class trim_direction {
        /**
         * Will remove all record batches from the first available record
         * until the given key.
         */
        from_beginning,

        /**
         * Will remove all record batches starting at the given offset
         * until the end of the log.
         */
        until_end
    };

public: // ro getters
    /**
     * Returns the namespace-topic-partition identifier of this log.
     */
    const model::ntp& ntp() const;

public: // ro access
    template<
      typename Key,
      typename = detail::is_indexable<Key, segment_indices>>
    seastar::input_stream<char> read(Key key) const;

public: // rw access
    /**
     * Removes record batches from one of the ends of the log.
     *
     * Depending on the trim direction, this function will either
     * have the semantics of filesystem truncate() or can be used
     * for purging old records from the log.
     *
     * The range may or may not span more than one segment.
     */
    template<
      typename Key,
      typename = detail::is_indexable<Key, segment_indices>>
    future<trim_result> trim(
      Key at,
      model::term_id term,
      trim_direction direction = trim_direction::until_end);

    /**
     * Writes a record batch to the log.
     *
     * This method will append the record batch to the current active segment
     * and possibly roll to a new segment if the max length is reached.
     *
     * This method also assigns offsets to the record batch, and returns the
     * assigned range in the returned promise.
     */
    future<std::vector<append_result>> append(model::record_batch_reader&&);

    /**
     * Flushes in-memory record batches to disk.
     * After this method completes (assuming no new writes happen in the
     * meantime), the committed offset should be equal to dirty offset.
     */
    seastar::future<flush_result> flush();

    /**
     * Implicit flush + release all held resources.
     *
     * This will close all file handles that are open, namely
     * the active segment file, active index, etc.
     *
     * This will also emplicitly seal the current active segment.
     */
    seastar::future<flush_result> close();

private: // async construction
    /**
     * Opens a log and prepares it for IO operations.
     *
     * Opening a log involves:
     *   - enumerating available segments
     *   - finding and storing in memory the map of key-ranges to segments (for
     *     index cache)
     *   - loading segment caches into memory
     */
    static future<log> open(const log_manager&, model::ntp, io_priority_class);

private:
    class impl;
    shared_ptr<impl> _impl;

private:
    log(shared_ptr<impl>);
    friend class log_manager;
};

} // namespace storage
