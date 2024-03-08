// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0--rc3
// source: artist_service.proto

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
	ArtistService_CreatePlaylist_FullMethodName = "/mpdb.ArtistService/CreatePlaylist"
	ArtistService_DeletePlaylist_FullMethodName = "/mpdb.ArtistService/DeletePlaylist"
	ArtistService_AddSong_FullMethodName        = "/mpdb.ArtistService/AddSong"
	ArtistService_DeleteSong_FullMethodName     = "/mpdb.ArtistService/DeleteSong"
	ArtistService_CheckFollowers_FullMethodName = "/mpdb.ArtistService/CheckFollowers"
)

// ArtistServiceClient is the client API for ArtistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtistServiceClient interface {
	CreatePlaylist(ctx context.Context, in *CreatePlaylistRequest, opts ...grpc.CallOption) (*CreatePlaylistResponse, error)
	DeletePlaylist(ctx context.Context, in *DeletePlaylistRequest, opts ...grpc.CallOption) (*DeletePlaylistResponse, error)
	AddSong(ctx context.Context, in *AddSongRequest, opts ...grpc.CallOption) (*AddSongResponse, error)
	DeleteSong(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*DeleteSongResponse, error)
	CheckFollowers(ctx context.Context, in *CheckFollowersRequest, opts ...grpc.CallOption) (*CheckFollowersResponse, error)
}

type artistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArtistServiceClient(cc grpc.ClientConnInterface) ArtistServiceClient {
	return &artistServiceClient{cc}
}

func (c *artistServiceClient) CreatePlaylist(ctx context.Context, in *CreatePlaylistRequest, opts ...grpc.CallOption) (*CreatePlaylistResponse, error) {
	out := new(CreatePlaylistResponse)
	err := c.cc.Invoke(ctx, ArtistService_CreatePlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) DeletePlaylist(ctx context.Context, in *DeletePlaylistRequest, opts ...grpc.CallOption) (*DeletePlaylistResponse, error) {
	out := new(DeletePlaylistResponse)
	err := c.cc.Invoke(ctx, ArtistService_DeletePlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) AddSong(ctx context.Context, in *AddSongRequest, opts ...grpc.CallOption) (*AddSongResponse, error) {
	out := new(AddSongResponse)
	err := c.cc.Invoke(ctx, ArtistService_AddSong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) DeleteSong(ctx context.Context, in *DeleteSongRequest, opts ...grpc.CallOption) (*DeleteSongResponse, error) {
	out := new(DeleteSongResponse)
	err := c.cc.Invoke(ctx, ArtistService_DeleteSong_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) CheckFollowers(ctx context.Context, in *CheckFollowersRequest, opts ...grpc.CallOption) (*CheckFollowersResponse, error) {
	out := new(CheckFollowersResponse)
	err := c.cc.Invoke(ctx, ArtistService_CheckFollowers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtistServiceServer is the server API for ArtistService service.
// All implementations must embed UnimplementedArtistServiceServer
// for forward compatibility
type ArtistServiceServer interface {
	CreatePlaylist(context.Context, *CreatePlaylistRequest) (*CreatePlaylistResponse, error)
	DeletePlaylist(context.Context, *DeletePlaylistRequest) (*DeletePlaylistResponse, error)
	AddSong(context.Context, *AddSongRequest) (*AddSongResponse, error)
	DeleteSong(context.Context, *DeleteSongRequest) (*DeleteSongResponse, error)
	CheckFollowers(context.Context, *CheckFollowersRequest) (*CheckFollowersResponse, error)
	mustEmbedUnimplementedArtistServiceServer()
}

// UnimplementedArtistServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArtistServiceServer struct {
}

func (UnimplementedArtistServiceServer) CreatePlaylist(context.Context, *CreatePlaylistRequest) (*CreatePlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlaylist not implemented")
}
func (UnimplementedArtistServiceServer) DeletePlaylist(context.Context, *DeletePlaylistRequest) (*DeletePlaylistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlaylist not implemented")
}
func (UnimplementedArtistServiceServer) AddSong(context.Context, *AddSongRequest) (*AddSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSong not implemented")
}
func (UnimplementedArtistServiceServer) DeleteSong(context.Context, *DeleteSongRequest) (*DeleteSongResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSong not implemented")
}
func (UnimplementedArtistServiceServer) CheckFollowers(context.Context, *CheckFollowersRequest) (*CheckFollowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckFollowers not implemented")
}
func (UnimplementedArtistServiceServer) mustEmbedUnimplementedArtistServiceServer() {}

// UnsafeArtistServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtistServiceServer will
// result in compilation errors.
type UnsafeArtistServiceServer interface {
	mustEmbedUnimplementedArtistServiceServer()
}

func RegisterArtistServiceServer(s grpc.ServiceRegistrar, srv ArtistServiceServer) {
	s.RegisterService(&ArtistService_ServiceDesc, srv)
}

func _ArtistService_CreatePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).CreatePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_CreatePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).CreatePlaylist(ctx, req.(*CreatePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_DeletePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePlaylistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).DeletePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_DeletePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).DeletePlaylist(ctx, req.(*DeletePlaylistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_AddSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).AddSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_AddSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).AddSong(ctx, req.(*AddSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_DeleteSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).DeleteSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_DeleteSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).DeleteSong(ctx, req.(*DeleteSongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_CheckFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckFollowersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).CheckFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_CheckFollowers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).CheckFollowers(ctx, req.(*CheckFollowersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtistService_ServiceDesc is the grpc.ServiceDesc for ArtistService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtistService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mpdb.ArtistService",
	HandlerType: (*ArtistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlaylist",
			Handler:    _ArtistService_CreatePlaylist_Handler,
		},
		{
			MethodName: "DeletePlaylist",
			Handler:    _ArtistService_DeletePlaylist_Handler,
		},
		{
			MethodName: "AddSong",
			Handler:    _ArtistService_AddSong_Handler,
		},
		{
			MethodName: "DeleteSong",
			Handler:    _ArtistService_DeleteSong_Handler,
		},
		{
			MethodName: "CheckFollowers",
			Handler:    _ArtistService_CheckFollowers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "artist_service.proto",
}
