// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: coffee_shop.proto

package coffeeshop_protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CoffeeShop_GetMenu_FullMethodName        = "/coffeeshop.CoffeeShop/GetMenu"
	CoffeeShop_PlaceOrder_FullMethodName     = "/coffeeshop.CoffeeShop/PlaceOrder"
	CoffeeShop_GetOrderStatus_FullMethodName = "/coffeeshop.CoffeeShop/GetOrderStatus"
)

// CoffeeShopClient is the client API for CoffeeShop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoffeeShopClient interface {
	GetMenu(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Menu], error)
	PlaceOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Receipt, error)
	GetOrderStatus(ctx context.Context, in *Receipt, opts ...grpc.CallOption) (*OrderStatus, error)
}

type coffeeShopClient struct {
	cc grpc.ClientConnInterface
}

func NewCoffeeShopClient(cc grpc.ClientConnInterface) CoffeeShopClient {
	return &coffeeShopClient{cc}
}

func (c *coffeeShopClient) GetMenu(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Menu], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CoffeeShop_ServiceDesc.Streams[0], CoffeeShop_GetMenu_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[MenuRequest, Menu]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CoffeeShop_GetMenuClient = grpc.ServerStreamingClient[Menu]

func (c *coffeeShopClient) PlaceOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Receipt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Receipt)
	err := c.cc.Invoke(ctx, CoffeeShop_PlaceOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coffeeShopClient) GetOrderStatus(ctx context.Context, in *Receipt, opts ...grpc.CallOption) (*OrderStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderStatus)
	err := c.cc.Invoke(ctx, CoffeeShop_GetOrderStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoffeeShopServer is the server API for CoffeeShop service.
// All implementations must embed UnimplementedCoffeeShopServer
// for forward compatibility.
type CoffeeShopServer interface {
	GetMenu(*MenuRequest, grpc.ServerStreamingServer[Menu]) error
	PlaceOrder(context.Context, *Order) (*Receipt, error)
	GetOrderStatus(context.Context, *Receipt) (*OrderStatus, error)
	mustEmbedUnimplementedCoffeeShopServer()
}

// UnimplementedCoffeeShopServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCoffeeShopServer struct{}

func (UnimplementedCoffeeShopServer) GetMenu(*MenuRequest, grpc.ServerStreamingServer[Menu]) error {
	return status.Errorf(codes.Unimplemented, "method GetMenu not implemented")
}
func (UnimplementedCoffeeShopServer) PlaceOrder(context.Context, *Order) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedCoffeeShopServer) GetOrderStatus(context.Context, *Receipt) (*OrderStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderStatus not implemented")
}
func (UnimplementedCoffeeShopServer) mustEmbedUnimplementedCoffeeShopServer() {}
func (UnimplementedCoffeeShopServer) testEmbeddedByValue()                    {}

// UnsafeCoffeeShopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoffeeShopServer will
// result in compilation errors.
type UnsafeCoffeeShopServer interface {
	mustEmbedUnimplementedCoffeeShopServer()
}

func RegisterCoffeeShopServer(s grpc.ServiceRegistrar, srv CoffeeShopServer) {
	// If the following call pancis, it indicates UnimplementedCoffeeShopServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CoffeeShop_ServiceDesc, srv)
}

func _CoffeeShop_GetMenu_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MenuRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoffeeShopServer).GetMenu(m, &grpc.GenericServerStream[MenuRequest, Menu]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CoffeeShop_GetMenuServer = grpc.ServerStreamingServer[Menu]

func _CoffeeShop_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoffeeShopServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoffeeShop_PlaceOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoffeeShopServer).PlaceOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoffeeShop_GetOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Receipt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoffeeShopServer).GetOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoffeeShop_GetOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoffeeShopServer).GetOrderStatus(ctx, req.(*Receipt))
	}
	return interceptor(ctx, in, info, handler)
}

// CoffeeShop_ServiceDesc is the grpc.ServiceDesc for CoffeeShop service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoffeeShop_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coffeeshop.CoffeeShop",
	HandlerType: (*CoffeeShopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceOrder",
			Handler:    _CoffeeShop_PlaceOrder_Handler,
		},
		{
			MethodName: "GetOrderStatus",
			Handler:    _CoffeeShop_GetOrderStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMenu",
			Handler:       _CoffeeShop_GetMenu_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "coffee_shop.proto",
}
