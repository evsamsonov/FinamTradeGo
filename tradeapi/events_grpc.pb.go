// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: grpc/tradeapi/v1/events.proto

package tradeapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Events_GetEvents_FullMethodName = "/grpc.tradeapi.v1.Events/GetEvents"
)

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventsClient interface {
	// Event Service sends events after explicit subscription.
	// Сервис событий. Отправляет события после вызова соответствующих методов подписки.
	GetEvents(ctx context.Context, opts ...grpc.CallOption) (Events_GetEventsClient, error)
}

type eventsClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsClient(cc grpc.ClientConnInterface) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) GetEvents(ctx context.Context, opts ...grpc.CallOption) (Events_GetEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Events_ServiceDesc.Streams[0], Events_GetEvents_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsGetEventsClient{stream}
	return x, nil
}

type Events_GetEventsClient interface {
	Send(*SubscriptionRequest) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventsGetEventsClient struct {
	grpc.ClientStream
}

func (x *eventsGetEventsClient) Send(m *SubscriptionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventsGetEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventsServer is the server API for Events service.
// All implementations must embed UnimplementedEventsServer
// for forward compatibility
type EventsServer interface {
	// Event Service sends events after explicit subscription.
	// Сервис событий. Отправляет события после вызова соответствующих методов подписки.
	GetEvents(Events_GetEventsServer) error
	mustEmbedUnimplementedEventsServer()
}

// UnimplementedEventsServer must be embedded to have forward compatible implementations.
type UnimplementedEventsServer struct {
}

func (UnimplementedEventsServer) GetEvents(Events_GetEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedEventsServer) mustEmbedUnimplementedEventsServer() {}

// UnsafeEventsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventsServer will
// result in compilation errors.
type UnsafeEventsServer interface {
	mustEmbedUnimplementedEventsServer()
}

func RegisterEventsServer(s grpc.ServiceRegistrar, srv EventsServer) {
	s.RegisterService(&Events_ServiceDesc, srv)
}

func _Events_GetEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventsServer).GetEvents(&eventsGetEventsServer{stream})
}

type Events_GetEventsServer interface {
	Send(*Event) error
	Recv() (*SubscriptionRequest, error)
	grpc.ServerStream
}

type eventsGetEventsServer struct {
	grpc.ServerStream
}

func (x *eventsGetEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventsGetEventsServer) Recv() (*SubscriptionRequest, error) {
	m := new(SubscriptionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Events_ServiceDesc is the grpc.ServiceDesc for Events service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Events_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.tradeapi.v1.Events",
	HandlerType: (*EventsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvents",
			Handler:       _Events_GetEvents_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/tradeapi/v1/events.proto",
}
