// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: currencies/v1/currencies.proto

package currenciesconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	v1 "currencies-exchange/gen/currencies/v1"
	errors "errors"
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
	// ExchangeServiceName is the fully-qualified name of the ExchangeService service.
	ExchangeServiceName = "currencies.v1.ExchangeService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ExchangeServiceExchangeProcedure is the fully-qualified name of the ExchangeService's Exchange
	// RPC.
	ExchangeServiceExchangeProcedure = "/currencies.v1.ExchangeService/Exchange"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	exchangeServiceServiceDescriptor        = v1.File_currencies_v1_currencies_proto.Services().ByName("ExchangeService")
	exchangeServiceExchangeMethodDescriptor = exchangeServiceServiceDescriptor.Methods().ByName("Exchange")
)

// ExchangeServiceClient is a client for the currencies.v1.ExchangeService service.
type ExchangeServiceClient interface {
	Exchange(context.Context, *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error)
}

// NewExchangeServiceClient constructs a client for the currencies.v1.ExchangeService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewExchangeServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ExchangeServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &exchangeServiceClient{
		exchange: connect.NewClient[v1.ExchangeRequest, v1.ExchangeResponse](
			httpClient,
			baseURL+ExchangeServiceExchangeProcedure,
			connect.WithSchema(exchangeServiceExchangeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// exchangeServiceClient implements ExchangeServiceClient.
type exchangeServiceClient struct {
	exchange *connect.Client[v1.ExchangeRequest, v1.ExchangeResponse]
}

// Exchange calls currencies.v1.ExchangeService.Exchange.
func (c *exchangeServiceClient) Exchange(ctx context.Context, req *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error) {
	return c.exchange.CallUnary(ctx, req)
}

// ExchangeServiceHandler is an implementation of the currencies.v1.ExchangeService service.
type ExchangeServiceHandler interface {
	Exchange(context.Context, *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error)
}

// NewExchangeServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewExchangeServiceHandler(svc ExchangeServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	exchangeServiceExchangeHandler := connect.NewUnaryHandler(
		ExchangeServiceExchangeProcedure,
		svc.Exchange,
		connect.WithSchema(exchangeServiceExchangeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/currencies.v1.ExchangeService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ExchangeServiceExchangeProcedure:
			exchangeServiceExchangeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedExchangeServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedExchangeServiceHandler struct{}

func (UnimplementedExchangeServiceHandler) Exchange(context.Context, *connect.Request[v1.ExchangeRequest]) (*connect.Response[v1.ExchangeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("currencies.v1.ExchangeService.Exchange is not implemented"))
}