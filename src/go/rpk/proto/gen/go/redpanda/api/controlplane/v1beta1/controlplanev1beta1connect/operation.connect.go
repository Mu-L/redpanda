// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: redpanda/api/controlplane/v1beta1/operation.proto

package controlplanev1beta1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1beta1 "github.com/redpanda-data/redpanda/src/go/rpk/proto/gen/go/redpanda/api/controlplane/v1beta1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// OperationServiceName is the fully-qualified name of the OperationService service.
	OperationServiceName = "redpanda.api.controlplane.v1beta1.OperationService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OperationServiceGetOperationProcedure is the fully-qualified name of the OperationService's
	// GetOperation RPC.
	OperationServiceGetOperationProcedure = "/redpanda.api.controlplane.v1beta1.OperationService/GetOperation"
	// OperationServiceListOperationsProcedure is the fully-qualified name of the OperationService's
	// ListOperations RPC.
	OperationServiceListOperationsProcedure = "/redpanda.api.controlplane.v1beta1.OperationService/ListOperations"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	operationServiceServiceDescriptor              = v1beta1.File_redpanda_api_controlplane_v1beta1_operation_proto.Services().ByName("OperationService")
	operationServiceGetOperationMethodDescriptor   = operationServiceServiceDescriptor.Methods().ByName("GetOperation")
	operationServiceListOperationsMethodDescriptor = operationServiceServiceDescriptor.Methods().ByName("ListOperations")
)

// OperationServiceClient is a client for the redpanda.api.controlplane.v1beta1.OperationService
// service.
type OperationServiceClient interface {
	GetOperation(context.Context, *connect.Request[v1beta1.GetOperationRequest]) (*connect.Response[v1beta1.Operation], error)
	ListOperations(context.Context, *connect.Request[v1beta1.ListOperationsRequest]) (*connect.Response[v1beta1.ListOperationsResponse], error)
}

// NewOperationServiceClient constructs a client for the
// redpanda.api.controlplane.v1beta1.OperationService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOperationServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OperationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &operationServiceClient{
		getOperation: connect.NewClient[v1beta1.GetOperationRequest, v1beta1.Operation](
			httpClient,
			baseURL+OperationServiceGetOperationProcedure,
			connect.WithSchema(operationServiceGetOperationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listOperations: connect.NewClient[v1beta1.ListOperationsRequest, v1beta1.ListOperationsResponse](
			httpClient,
			baseURL+OperationServiceListOperationsProcedure,
			connect.WithSchema(operationServiceListOperationsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// operationServiceClient implements OperationServiceClient.
type operationServiceClient struct {
	getOperation   *connect.Client[v1beta1.GetOperationRequest, v1beta1.Operation]
	listOperations *connect.Client[v1beta1.ListOperationsRequest, v1beta1.ListOperationsResponse]
}

// GetOperation calls redpanda.api.controlplane.v1beta1.OperationService.GetOperation.
func (c *operationServiceClient) GetOperation(ctx context.Context, req *connect.Request[v1beta1.GetOperationRequest]) (*connect.Response[v1beta1.Operation], error) {
	return c.getOperation.CallUnary(ctx, req)
}

// ListOperations calls redpanda.api.controlplane.v1beta1.OperationService.ListOperations.
func (c *operationServiceClient) ListOperations(ctx context.Context, req *connect.Request[v1beta1.ListOperationsRequest]) (*connect.Response[v1beta1.ListOperationsResponse], error) {
	return c.listOperations.CallUnary(ctx, req)
}

// OperationServiceHandler is an implementation of the
// redpanda.api.controlplane.v1beta1.OperationService service.
type OperationServiceHandler interface {
	GetOperation(context.Context, *connect.Request[v1beta1.GetOperationRequest]) (*connect.Response[v1beta1.Operation], error)
	ListOperations(context.Context, *connect.Request[v1beta1.ListOperationsRequest]) (*connect.Response[v1beta1.ListOperationsResponse], error)
}

// NewOperationServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOperationServiceHandler(svc OperationServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	operationServiceGetOperationHandler := connect.NewUnaryHandler(
		OperationServiceGetOperationProcedure,
		svc.GetOperation,
		connect.WithSchema(operationServiceGetOperationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	operationServiceListOperationsHandler := connect.NewUnaryHandler(
		OperationServiceListOperationsProcedure,
		svc.ListOperations,
		connect.WithSchema(operationServiceListOperationsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/redpanda.api.controlplane.v1beta1.OperationService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OperationServiceGetOperationProcedure:
			operationServiceGetOperationHandler.ServeHTTP(w, r)
		case OperationServiceListOperationsProcedure:
			operationServiceListOperationsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOperationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOperationServiceHandler struct{}

func (UnimplementedOperationServiceHandler) GetOperation(context.Context, *connect.Request[v1beta1.GetOperationRequest]) (*connect.Response[v1beta1.Operation], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.controlplane.v1beta1.OperationService.GetOperation is not implemented"))
}

func (UnimplementedOperationServiceHandler) ListOperations(context.Context, *connect.Request[v1beta1.ListOperationsRequest]) (*connect.Response[v1beta1.ListOperationsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.controlplane.v1beta1.OperationService.ListOperations is not implemented"))
}