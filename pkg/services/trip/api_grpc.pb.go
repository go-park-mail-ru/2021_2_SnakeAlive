// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package trip_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TripServiceClient is the client API for TripService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TripServiceClient interface {
	GetTrip(ctx context.Context, in *TripRequest, opts ...grpc.CallOption) (*Trip, error)
	AddTrip(ctx context.Context, in *ModifyTripRequest, opts ...grpc.CallOption) (*Trip, error)
	UpdateTrip(ctx context.Context, in *ModifyTripRequest, opts ...grpc.CallOption) (*Trip, error)
	DeleteTrip(ctx context.Context, in *TripRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetTripsByUser(ctx context.Context, in *TripByUserRequest, opts ...grpc.CallOption) (*Trips, error)
	GetAlbum(ctx context.Context, in *AlbumRequest, opts ...grpc.CallOption) (*Album, error)
	AddAlbum(ctx context.Context, in *ModifyAlbumRequest, opts ...grpc.CallOption) (*Album, error)
	UpdateAlbum(ctx context.Context, in *ModifyAlbumRequest, opts ...grpc.CallOption) (*Album, error)
	DeleteAlbum(ctx context.Context, in *AlbumRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SightsByTrip(ctx context.Context, in *SightsRequest, opts ...grpc.CallOption) (*Sights, error)
}

type tripServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTripServiceClient(cc grpc.ClientConnInterface) TripServiceClient {
	return &tripServiceClient{cc}
}

func (c *tripServiceClient) GetTrip(ctx context.Context, in *TripRequest, opts ...grpc.CallOption) (*Trip, error) {
	out := new(Trip)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/GetTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) AddTrip(ctx context.Context, in *ModifyTripRequest, opts ...grpc.CallOption) (*Trip, error) {
	out := new(Trip)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/AddTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) UpdateTrip(ctx context.Context, in *ModifyTripRequest, opts ...grpc.CallOption) (*Trip, error) {
	out := new(Trip)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/UpdateTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) DeleteTrip(ctx context.Context, in *TripRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/DeleteTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) GetTripsByUser(ctx context.Context, in *TripByUserRequest, opts ...grpc.CallOption) (*Trips, error) {
	out := new(Trips)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/GetTripsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) GetAlbum(ctx context.Context, in *AlbumRequest, opts ...grpc.CallOption) (*Album, error) {
	out := new(Album)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/GetAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) AddAlbum(ctx context.Context, in *ModifyAlbumRequest, opts ...grpc.CallOption) (*Album, error) {
	out := new(Album)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/AddAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) UpdateAlbum(ctx context.Context, in *ModifyAlbumRequest, opts ...grpc.CallOption) (*Album, error) {
	out := new(Album)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/UpdateAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) DeleteAlbum(ctx context.Context, in *AlbumRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/DeleteAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) SightsByTrip(ctx context.Context, in *SightsRequest, opts ...grpc.CallOption) (*Sights, error) {
	out := new(Sights)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/SightsByTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TripServiceServer is the server API for TripService service.
// All implementations must embed UnimplementedTripServiceServer
// for forward compatibility
type TripServiceServer interface {
	GetTrip(context.Context, *TripRequest) (*Trip, error)
	AddTrip(context.Context, *ModifyTripRequest) (*Trip, error)
	UpdateTrip(context.Context, *ModifyTripRequest) (*Trip, error)
	DeleteTrip(context.Context, *TripRequest) (*empty.Empty, error)
	GetTripsByUser(context.Context, *TripByUserRequest) (*Trips, error)
	GetAlbum(context.Context, *AlbumRequest) (*Album, error)
	AddAlbum(context.Context, *ModifyAlbumRequest) (*Album, error)
	UpdateAlbum(context.Context, *ModifyAlbumRequest) (*Album, error)
	DeleteAlbum(context.Context, *AlbumRequest) (*empty.Empty, error)
	SightsByTrip(context.Context, *SightsRequest) (*Sights, error)
	mustEmbedUnimplementedTripServiceServer()
}

// UnimplementedTripServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTripServiceServer struct {
}

func (UnimplementedTripServiceServer) GetTrip(context.Context, *TripRequest) (*Trip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrip not implemented")
}
func (UnimplementedTripServiceServer) AddTrip(context.Context, *ModifyTripRequest) (*Trip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTrip not implemented")
}
func (UnimplementedTripServiceServer) UpdateTrip(context.Context, *ModifyTripRequest) (*Trip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrip not implemented")
}
func (UnimplementedTripServiceServer) DeleteTrip(context.Context, *TripRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrip not implemented")
}
func (UnimplementedTripServiceServer) GetTripsByUser(context.Context, *TripByUserRequest) (*Trips, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTripsByUser not implemented")
}
func (UnimplementedTripServiceServer) GetAlbum(context.Context, *AlbumRequest) (*Album, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlbum not implemented")
}
func (UnimplementedTripServiceServer) AddAlbum(context.Context, *ModifyAlbumRequest) (*Album, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAlbum not implemented")
}
func (UnimplementedTripServiceServer) UpdateAlbum(context.Context, *ModifyAlbumRequest) (*Album, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAlbum not implemented")
}
func (UnimplementedTripServiceServer) DeleteAlbum(context.Context, *AlbumRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlbum not implemented")
}
func (UnimplementedTripServiceServer) SightsByTrip(context.Context, *SightsRequest) (*Sights, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SightsByTrip not implemented")
}
func (UnimplementedTripServiceServer) mustEmbedUnimplementedTripServiceServer() {}

// UnsafeTripServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TripServiceServer will
// result in compilation errors.
type UnsafeTripServiceServer interface {
	mustEmbedUnimplementedTripServiceServer()
}

func RegisterTripServiceServer(s grpc.ServiceRegistrar, srv TripServiceServer) {
	s.RegisterService(&TripService_ServiceDesc, srv)
}

func _TripService_GetTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).GetTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/GetTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).GetTrip(ctx, req.(*TripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_AddTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).AddTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/AddTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).AddTrip(ctx, req.(*ModifyTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_UpdateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).UpdateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/UpdateTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).UpdateTrip(ctx, req.(*ModifyTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_DeleteTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).DeleteTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/DeleteTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).DeleteTrip(ctx, req.(*TripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_GetTripsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TripByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).GetTripsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/GetTripsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).GetTripsByUser(ctx, req.(*TripByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_GetAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).GetAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/GetAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).GetAlbum(ctx, req.(*AlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_AddAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).AddAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/AddAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).AddAlbum(ctx, req.(*ModifyAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_UpdateAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).UpdateAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/UpdateAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).UpdateAlbum(ctx, req.(*ModifyAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_DeleteAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).DeleteAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/DeleteAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).DeleteAlbum(ctx, req.(*AlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_SightsByTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).SightsByTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/SightsByTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).SightsByTrip(ctx, req.(*SightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TripService_ServiceDesc is the grpc.ServiceDesc for TripService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TripService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.trip_service.TripService",
	HandlerType: (*TripServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTrip",
			Handler:    _TripService_GetTrip_Handler,
		},
		{
			MethodName: "AddTrip",
			Handler:    _TripService_AddTrip_Handler,
		},
		{
			MethodName: "UpdateTrip",
			Handler:    _TripService_UpdateTrip_Handler,
		},
		{
			MethodName: "DeleteTrip",
			Handler:    _TripService_DeleteTrip_Handler,
		},
		{
			MethodName: "GetTripsByUser",
			Handler:    _TripService_GetTripsByUser_Handler,
		},
		{
			MethodName: "GetAlbum",
			Handler:    _TripService_GetAlbum_Handler,
		},
		{
			MethodName: "AddAlbum",
			Handler:    _TripService_AddAlbum_Handler,
		},
		{
			MethodName: "UpdateAlbum",
			Handler:    _TripService_UpdateAlbum_Handler,
		},
		{
			MethodName: "DeleteAlbum",
			Handler:    _TripService_DeleteAlbum_Handler,
		},
		{
			MethodName: "SightsByTrip",
			Handler:    _TripService_SightsByTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services/trip/api.proto",
}
