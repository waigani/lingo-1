// Code generated by protoc-gen-go.
// source: query.proto
// DO NOT EDIT!

/*
Package query is a generated protocol buffer package.

It is generated from these files:
	query.proto

It has these top-level messages:
	QueryRequest
	QueryReply
*/
package query

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The request message containing the CLQL source code.
type QueryRequest struct {
	Clql string `protobuf:"bytes,1,opt,name=clql" json:"clql,omitempty"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}

// The query response.
type QueryReply struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *QueryReply) Reset()         { *m = QueryReply{} }
func (m *QueryReply) String() string { return proto.CompactTextString(m) }
func (*QueryReply) ProtoMessage()    {}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "query.QueryRequest")
	proto.RegisterType((*QueryReply)(nil), "query.QueryReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Query service

type QueryClient interface {
	// Sends a greeting
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error)
}

type queryClient struct {
	cc *grpc.ClientConn
}

func NewQueryClient(cc *grpc.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryReply, error) {
	out := new(QueryReply)
	err := grpc.Invoke(ctx, "/query.Query/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Query service

type QueryServer interface {
	// Sends a greeting
	Query(context.Context, *QueryRequest) (*QueryReply, error)
}

func RegisterQueryServer(s *grpc.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(QueryServer).Query(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "query.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Query_Query_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
