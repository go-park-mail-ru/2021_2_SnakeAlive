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
	DeleteTrip(ctx context.Context, in *TripRequest, opts ...grpc.CallOption) (*Users, error)
	GetTripsByUser(ctx context.Context, in *ByUserRequest, opts ...grpc.CallOption) (*Trips, error)
	GetAlbum(ctx context.Context, in *AlbumRequest, opts ...grpc.CallOption) (*Album, error)
	AddAlbum(ctx context.Context, in *ModifyAlbumRequest, opts ...grpc.CallOption) (*Album, error)
	UpdateAlbum(ctx context.Context, in *ModifyAlbumRequest, opts ...grpc.CallOption) (*Album, error)
	DeleteAlbum(ctx context.Context, in *AlbumRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	SightsByTrip(ctx context.Context, in *SightsRequest, opts ...grpc.CallOption) (*Sights, error)
	GetAlbumsByUser(ctx context.Context, in *ByUserRequest, opts ...grpc.CallOption) (*Albums, error)
	AddTripUser(ctx context.Context, in *AddTripUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ShareLink(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*Link, error)
	AddUserByLink(ctx context.Context, in *AddByShareRequest, opts ...grpc.CallOption) (*Link, error)
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

func (c *tripServiceClient) DeleteTrip(ctx context.Context, in *TripRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/DeleteTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) GetTripsByUser(ctx context.Context, in *ByUserRequest, opts ...grpc.CallOption) (*Trips, error) {
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

func (c *tripServiceClient) GetAlbumsByUser(ctx context.Context, in *ByUserRequest, opts ...grpc.CallOption) (*Albums, error) {
	out := new(Albums)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/GetAlbumsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) AddTripUser(ctx context.Context, in *AddTripUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/AddTripUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) ShareLink(ctx context.Context, in *ShareRequest, opts ...grpc.CallOption) (*Link, error) {
	out := new(Link)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/ShareLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) AddUserByLink(ctx context.Context, in *AddByShareRequest, opts ...grpc.CallOption) (*Link, error) {
	out := new(Link)
	err := c.cc.Invoke(ctx, "/services.trip_service.TripService/AddUserByLink", in, out, opts...)
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
	DeleteTrip(context.Context, *TripRequest) (*Users, error)
	GetTripsByUser(context.Context, *ByUserRequest) (*Trips, error)
	GetAlbum(context.Context, *AlbumRequest) (*Album, error)
	AddAlbum(context.Context, *ModifyAlbumRequest) (*Album, error)
	UpdateAlbum(context.Context, *ModifyAlbumRequest) (*Album, error)
	DeleteAlbum(context.Context, *AlbumRequest) (*empty.Empty, error)
	SightsByTrip(context.Context, *SightsRequest) (*Sights, error)
	GetAlbumsByUser(context.Context, *ByUserRequest) (*Albums, error)
	AddTripUser(context.Context, *AddTripUserRequest) (*empty.Empty, error)
	ShareLink(context.Context, *ShareRequest) (*Link, error)
	AddUserByLink(context.Context, *AddByShareRequest) (*Link, error)
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
func (UnimplementedTripServiceServer) DeleteTrip(context.Context, *TripRequest) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrip not implemented")
}
func (UnimplementedTripServiceServer) GetTripsByUser(context.Context, *ByUserRequest) (*Trips, error) {
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
func (UnimplementedTripServiceServer) GetAlbumsByUser(context.Context, *ByUserRequest) (*Albums, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlbumsByUser not implemented")
}
func (UnimplementedTripServiceServer) AddTripUser(context.Context, *AddTripUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTripUser not implemented")
}
func (UnimplementedTripServiceServer) ShareLink(context.Context, *ShareRequest) (*Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareLink not implemented")
}
func (UnimplementedTripServiceServer) AddUserByLink(context.Context, *AddByShareRequest) (*Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserByLink not implemented")
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
	in := new(ByUserRequest)
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
		return srv.(TripServiceServer).GetTripsByUser(ctx, req.(*ByUserRequest))
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

func _TripService_GetAlbumsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).GetAlbumsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/GetAlbumsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).GetAlbumsByUser(ctx, req.(*ByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_AddTripUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTripUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).AddTripUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/AddTripUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).AddTripUser(ctx, req.(*AddTripUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_ShareLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).ShareLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/ShareLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).ShareLink(ctx, req.(*ShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_AddUserByLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddByShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).AddUserByLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.trip_service.TripService/AddUserByLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).AddUserByLink(ctx, req.(*AddByShareRequest))
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
		{
			MethodName: "GetAlbumsByUser",
			Handler:    _TripService_GetAlbumsByUser_Handler,
		},
		{
			MethodName: "AddTripUser",
			Handler:    _TripService_AddTripUser_Handler,
		},
		{
			MethodName: "ShareLink",
			Handler:    _TripService_ShareLink_Handler,
		},
		{
			MethodName: "AddUserByLink",
			Handler:    _TripService_AddUserByLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services/trip/api.proto",
}
