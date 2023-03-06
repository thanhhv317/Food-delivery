// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: restaurant_like_service.proto

package proto

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

// RestaurantLikeServiceClient is the client API for RestaurantLikeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestaurantLikeServiceClient interface {
	GetRestaurantLikeStat(ctx context.Context, in *RestaurantLikeStatRequest, opts ...grpc.CallOption) (*RestaurantLikeStatResponse, error)
}

type restaurantLikeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestaurantLikeServiceClient(cc grpc.ClientConnInterface) RestaurantLikeServiceClient {
	return &restaurantLikeServiceClient{cc}
}

func (c *restaurantLikeServiceClient) GetRestaurantLikeStat(ctx context.Context, in *RestaurantLikeStatRequest, opts ...grpc.CallOption) (*RestaurantLikeStatResponse, error) {
	out := new(RestaurantLikeStatResponse)
	err := c.cc.Invoke(ctx, "/demo.RestaurantLikeService/GetRestaurantLikeStat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestaurantLikeServiceServer is the server API for RestaurantLikeService service.
// All implementations must embed UnimplementedRestaurantLikeServiceServer
// for forward compatibility
type RestaurantLikeServiceServer interface {
	GetRestaurantLikeStat(context.Context, *RestaurantLikeStatRequest) (*RestaurantLikeStatResponse, error)
	mustEmbedUnimplementedRestaurantLikeServiceServer()
}

// UnimplementedRestaurantLikeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRestaurantLikeServiceServer struct {
}

func (UnimplementedRestaurantLikeServiceServer) GetRestaurantLikeStat(context.Context, *RestaurantLikeStatRequest) (*RestaurantLikeStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurantLikeStat not implemented")
}
func (UnimplementedRestaurantLikeServiceServer) mustEmbedUnimplementedRestaurantLikeServiceServer() {}

// UnsafeRestaurantLikeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestaurantLikeServiceServer will
// result in compilation errors.
type UnsafeRestaurantLikeServiceServer interface {
	mustEmbedUnimplementedRestaurantLikeServiceServer()
}

func RegisterRestaurantLikeServiceServer(s grpc.ServiceRegistrar, srv RestaurantLikeServiceServer) {
	s.RegisterService(&RestaurantLikeService_ServiceDesc, srv)
}

func _RestaurantLikeService_GetRestaurantLikeStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantLikeStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantLikeServiceServer).GetRestaurantLikeStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.RestaurantLikeService/GetRestaurantLikeStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantLikeServiceServer).GetRestaurantLikeStat(ctx, req.(*RestaurantLikeStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RestaurantLikeService_ServiceDesc is the grpc.ServiceDesc for RestaurantLikeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RestaurantLikeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.RestaurantLikeService",
	HandlerType: (*RestaurantLikeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRestaurantLikeStat",
			Handler:    _RestaurantLikeService_GetRestaurantLikeStat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "restaurant_like_service.proto",
}
