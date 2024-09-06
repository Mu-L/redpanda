// Copyright 2024 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0
#pragma once

#include "iceberg/partition.h"
#include "json/_include_first.h"
#include "json/document.h"
#include "json/stringbuffer.h"
#include "json/writer.h"

namespace iceberg {

partition_field parse_partition_field(const json::Value&);
partition_spec parse_partition_spec(const json::Value&);

} // namespace iceberg

namespace json {

void rjson_serialize(
  json::Writer<json::StringBuffer>& w, const iceberg::partition_field& m);
void rjson_serialize(
  json::Writer<json::StringBuffer>& w, const iceberg::partition_spec& m);

} // namespace json