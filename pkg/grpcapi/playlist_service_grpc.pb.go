// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0--rc3
// source: playlist_service.proto

package grpcapi

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
	PlaylistService_ViewSongsByPlaylist_FullMethodName     = "/mpdb.PlaylistService/ViewSongsByPlaylist"
	PlaylistService_GetPlaylistByTitle_FullMethodName      = "/mpdb.PlaylistService/GetPlaylistByTitle"
	PlaylistService_GetPlaylistById_FullMethodName         = "/mpdb.PlaylistService/GetPlaylistById"
	PlaylistService_GetPlaylistCreationTime_FullMethodName = "/mpdb.PlaylistService/GetPlaylistCreationTime"
)

// PlaylistServiceClient is the client API for PlaylistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaylistServiceClient interface {
	ViewSongsByPlaylist(ctx context.Context, in *ViewSongsByPlaylistRequest, opts ...grpc.CallOption) (*ViewSongsByPlaylistResponse, error)
	GetPlaylistByTitle(ctx context.Context, in *GetPlaylistByTitleRequest, opts ...grpc.CallOption) (*GetPlaylistByTitleResponse, error)
	GetPlaylistById(ctx context.Context, in *GetPlaylistByIdRequest, opts ...grpc.CallOption) (*GetPlaylistByIdResponse, error)
	GetPlaylistCreationTime(ctx context.Context, in *GetPlaylistCreationTimeRequest, opts ...grpc.CallOption) (*GetPlaylistCreationTimeResponse, error)
}

type playlistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaylistServiceClient(cc grpc.ClientConnInterface) PlaylistServiceClient {
	return &playlistServiceClient{cc}
}

func (c *playlistServiceClient) ViewSongsByPlaylist(ctx context.Context, in *ViewSongsByPlaylistRequest, opts ...grpc.CallOption) (*ViewSongsByPlaylistResponse, error) {
	out := new(ViewSongsByPlaylistResponse)
	err := c.cc.Invoke(ctx, PlaylistService_ViewSongsByPlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) GetPlaylistByTitle(ctx context.Context, in *GetPlaylistByTitleRequest, opts ...grpc.CallOption) (*GetPlaylistByTitleResponse, error) {
	out := new(GetPlaylistByTitleResponse)
	err := c.cc.Invoke(ctx, PlaylistService_GetPlaylistByTitle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) GetPlaylistById(ctx context.Context, in *GetPlaylistByIdRequest, opts ...grpc.CallOption) (*GetPlaylistByIdResponse, error) {
	out := new(GetPlaylistByIdResponse)
	err := c.cc.Invoke(ctx, PlaylistService_GetPlaylistById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistServiceClient) GetPlaylistCreationTime(ctx context.Context, in *GetPlaylistCreationTimeRequest, opts ...grpc.CallOption) (*GetPlaylistCreationTimeResponse, error) {
	out := new(GetPlaylistCreationTimeResponse)
	err := c.cc.Invoke(ctx, PlaylistService_GetPlaylistCreationTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaylistServiceServer is the server API for PlaylistService service.
// All implementations must embed UnimplementedPlaylistServiceServer
// for forward compatibility
type PlaylistServiceServer interface {
	ViewSongsByPlaylist(context.Context, *ViewSongsByPlaylistRequest) (*ViewSongsByPlaylistResponse, error)
	GetPlaylistByTitle(context.Context, *GetPlaylistByTitleRequest) (*GetPlaylistByTitleResponse, error)
	GetPlaylistById(context.Context, *GetPlaylistByIdRequest) (*GetPlaylistByIdResponse, error)
	GetPlaylistCreationTime(context.Context, *GetPlaylistCreationTimeRequest) (*GetPlaylistCreationTimeResponse, error)
	mustEmbedUnimplementedPlaylistServiceServer()
}

// UnimplementedPlaylistServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlaylistServiceServer struct {
}

func (UnimplementedPlaylistServiceServer) ViewSongsByPlaylist(context.Context, *ViewSongsByPlaylistRequest) (*ViewSongsByPlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewSongsByPlaylist not implemented")
}
func (UnimplementedPlaylistServiceServer) GetPlaylistByTitle(context.Context, *GetPlaylistByTitleRequest) (*GetPlaylistByTitleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylistByTitle not implemented")
}
func (UnimplementedPlaylistServiceServer) GetPlaylistById(context.Context, *GetPlaylistByIdRequest) (*GetPlaylistByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylistById not implemented")
}
func (UnimplementedPlaylistServiceServer) GetPlaylistCreationTime(context.Context, *GetPlaylistCreationTimeRequest) (*GetPlaylistCreationTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylistCreationTime not implemented")
}
func (UnimplementedPlaylistServiceServer) mustEmbedUnimplementedPlaylistServiceServer() {}

// UnsafePlaylistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaylistServiceServer will
// result in compilation errors.
type UnsafePlaylistServiceServer interface {
	mustEmbedUnimplementedPlaylistServiceServer()
}

func RegisterPlaylistServiceServer(s grpc.ServiceRegistrar, srv PlaylistServiceServer) {
	s.RegisterService(&PlaylistService_ServiceDesc, srv)
}

func _PlaylistService_ViewSongsByPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewSongsByPlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).ViewSongsByPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlaylistService_ViewSongsByPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).ViewSongsByPlaylist(ctx, req.(*ViewSongsByPlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_GetPlaylistByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaylistByTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).GetPlaylistByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlaylistService_GetPlaylistByTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).GetPlaylistByTitle(ctx, req.(*GetPlaylistByTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_GetPlaylistById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaylistByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).GetPlaylistById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlaylistService_GetPlaylistById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).GetPlaylistById(ctx, req.(*GetPlaylistByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlaylistService_GetPlaylistCreationTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaylistCreationTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServiceServer).GetPlaylistCreationTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlaylistService_GetPlaylistCreationTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServiceServer).GetPlaylistCreationTime(ctx, req.(*GetPlaylistCreationTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlaylistService_ServiceDesc is the grpc.ServiceDesc for PlaylistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlaylistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mpdb.PlaylistService",
	HandlerType: (*PlaylistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ViewSongsByPlaylist",
			Handler:    _PlaylistService_ViewSongsByPlaylist_Handler,
		},
		{
			MethodName: "GetPlaylistByTitle",
			Handler:    _PlaylistService_GetPlaylistByTitle_Handler,
		},
		{
			MethodName: "GetPlaylistById",
			Handler:    _PlaylistService_GetPlaylistById_Handler,
		},
		{
			MethodName: "GetPlaylistCreationTime",
			Handler:    _PlaylistService_GetPlaylistCreationTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playlist_service.proto",
}
