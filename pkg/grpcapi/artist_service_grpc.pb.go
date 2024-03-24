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
	ArtistService_CreateArtist_FullMethodName                = "/mpdb.ArtistService/CreateArtist"
	ArtistService_ReadArtistById_FullMethodName              = "/mpdb.ArtistService/ReadArtistById"
	ArtistService_ReadArtistByName_FullMethodName            = "/mpdb.ArtistService/ReadArtistByName"
	ArtistService_ReadArtistBioById_FullMethodName           = "/mpdb.ArtistService/ReadArtistBioById"
	ArtistService_ReadArtistFollowerCountById_FullMethodName = "/mpdb.ArtistService/ReadArtistFollowerCountById"
	ArtistService_ReadArtistLikesCountById_FullMethodName    = "/mpdb.ArtistService/ReadArtistLikesCountById"
	ArtistService_UpdateArtistName_FullMethodName            = "/mpdb.ArtistService/UpdateArtistName"
	ArtistService_UpdateArtistBio_FullMethodName             = "/mpdb.ArtistService/UpdateArtistBio"
	ArtistService_UpdateArtistFollowerCount_FullMethodName   = "/mpdb.ArtistService/UpdateArtistFollowerCount"
	ArtistService_UpdateArtistLikesCount_FullMethodName      = "/mpdb.ArtistService/UpdateArtistLikesCount"
	ArtistService_DeleteArtistById_FullMethodName            = "/mpdb.ArtistService/DeleteArtistById"
)

// ArtistServiceClient is the client API for ArtistService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtistServiceClient interface {
	// CREATE REQUEST
	CreateArtist(ctx context.Context, in *CreateArtistRequest, opts ...grpc.CallOption) (*CreateArtistResponse, error)
	// READ REQUEST
	ReadArtistById(ctx context.Context, in *ReadArtistByIdRequest, opts ...grpc.CallOption) (*ReadArtistByIdResponse, error)
	ReadArtistByName(ctx context.Context, in *ReadArtistByNameRequest, opts ...grpc.CallOption) (*ReadArtistByNameResponse, error)
	ReadArtistBioById(ctx context.Context, in *ReadArtistBioRequest, opts ...grpc.CallOption) (*ReadArtistBioResponse, error)
	ReadArtistFollowerCountById(ctx context.Context, in *ReadArtistFollowerCountRequest, opts ...grpc.CallOption) (*ReadArtistFollowerCountResponse, error)
	ReadArtistLikesCountById(ctx context.Context, in *ReadArtistLikesCountRequest, opts ...grpc.CallOption) (*ReadArtistLikesCountResponse, error)
	// UPDATE REQUEST
	UpdateArtistName(ctx context.Context, in *UpdateArtistNameRequest, opts ...grpc.CallOption) (*UpdateArtistNameResponse, error)
	UpdateArtistBio(ctx context.Context, in *UpdateArtistBioRequest, opts ...grpc.CallOption) (*UpdateArtistBioResponse, error)
	UpdateArtistFollowerCount(ctx context.Context, in *UpdateArtistFollowerCountRequest, opts ...grpc.CallOption) (*UpdateArtistFollowerCountResponse, error)
	UpdateArtistLikesCount(ctx context.Context, in *UpdateArtistLikesCountRequest, opts ...grpc.CallOption) (*UpdateArtistLikesCountResponse, error)
	// DELETE REQUEST
	DeleteArtistById(ctx context.Context, in *DeleteArtistByIdRequest, opts ...grpc.CallOption) (*DeleteArtistByIdResponse, error)
}

type artistServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArtistServiceClient(cc grpc.ClientConnInterface) ArtistServiceClient {
	return &artistServiceClient{cc}
}

func (c *artistServiceClient) CreateArtist(ctx context.Context, in *CreateArtistRequest, opts ...grpc.CallOption) (*CreateArtistResponse, error) {
	out := new(CreateArtistResponse)
	err := c.cc.Invoke(ctx, ArtistService_CreateArtist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) ReadArtistById(ctx context.Context, in *ReadArtistByIdRequest, opts ...grpc.CallOption) (*ReadArtistByIdResponse, error) {
	out := new(ReadArtistByIdResponse)
	err := c.cc.Invoke(ctx, ArtistService_ReadArtistById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) ReadArtistByName(ctx context.Context, in *ReadArtistByNameRequest, opts ...grpc.CallOption) (*ReadArtistByNameResponse, error) {
	out := new(ReadArtistByNameResponse)
	err := c.cc.Invoke(ctx, ArtistService_ReadArtistByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) ReadArtistBioById(ctx context.Context, in *ReadArtistBioRequest, opts ...grpc.CallOption) (*ReadArtistBioResponse, error) {
	out := new(ReadArtistBioResponse)
	err := c.cc.Invoke(ctx, ArtistService_ReadArtistBioById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) ReadArtistFollowerCountById(ctx context.Context, in *ReadArtistFollowerCountRequest, opts ...grpc.CallOption) (*ReadArtistFollowerCountResponse, error) {
	out := new(ReadArtistFollowerCountResponse)
	err := c.cc.Invoke(ctx, ArtistService_ReadArtistFollowerCountById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) ReadArtistLikesCountById(ctx context.Context, in *ReadArtistLikesCountRequest, opts ...grpc.CallOption) (*ReadArtistLikesCountResponse, error) {
	out := new(ReadArtistLikesCountResponse)
	err := c.cc.Invoke(ctx, ArtistService_ReadArtistLikesCountById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) UpdateArtistName(ctx context.Context, in *UpdateArtistNameRequest, opts ...grpc.CallOption) (*UpdateArtistNameResponse, error) {
	out := new(UpdateArtistNameResponse)
	err := c.cc.Invoke(ctx, ArtistService_UpdateArtistName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) UpdateArtistBio(ctx context.Context, in *UpdateArtistBioRequest, opts ...grpc.CallOption) (*UpdateArtistBioResponse, error) {
	out := new(UpdateArtistBioResponse)
	err := c.cc.Invoke(ctx, ArtistService_UpdateArtistBio_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) UpdateArtistFollowerCount(ctx context.Context, in *UpdateArtistFollowerCountRequest, opts ...grpc.CallOption) (*UpdateArtistFollowerCountResponse, error) {
	out := new(UpdateArtistFollowerCountResponse)
	err := c.cc.Invoke(ctx, ArtistService_UpdateArtistFollowerCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) UpdateArtistLikesCount(ctx context.Context, in *UpdateArtistLikesCountRequest, opts ...grpc.CallOption) (*UpdateArtistLikesCountResponse, error) {
	out := new(UpdateArtistLikesCountResponse)
	err := c.cc.Invoke(ctx, ArtistService_UpdateArtistLikesCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) DeleteArtistById(ctx context.Context, in *DeleteArtistByIdRequest, opts ...grpc.CallOption) (*DeleteArtistByIdResponse, error) {
	out := new(DeleteArtistByIdResponse)
	err := c.cc.Invoke(ctx, ArtistService_DeleteArtistById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtistServiceServer is the server API for ArtistService service.
// All implementations must embed UnimplementedArtistServiceServer
// for forward compatibility
type ArtistServiceServer interface {
	// CREATE REQUEST
	CreateArtist(context.Context, *CreateArtistRequest) (*CreateArtistResponse, error)
	// READ REQUEST
	ReadArtistById(context.Context, *ReadArtistByIdRequest) (*ReadArtistByIdResponse, error)
	ReadArtistByName(context.Context, *ReadArtistByNameRequest) (*ReadArtistByNameResponse, error)
	ReadArtistBioById(context.Context, *ReadArtistBioRequest) (*ReadArtistBioResponse, error)
	ReadArtistFollowerCountById(context.Context, *ReadArtistFollowerCountRequest) (*ReadArtistFollowerCountResponse, error)
	ReadArtistLikesCountById(context.Context, *ReadArtistLikesCountRequest) (*ReadArtistLikesCountResponse, error)
	// UPDATE REQUEST
	UpdateArtistName(context.Context, *UpdateArtistNameRequest) (*UpdateArtistNameResponse, error)
	UpdateArtistBio(context.Context, *UpdateArtistBioRequest) (*UpdateArtistBioResponse, error)
	UpdateArtistFollowerCount(context.Context, *UpdateArtistFollowerCountRequest) (*UpdateArtistFollowerCountResponse, error)
	UpdateArtistLikesCount(context.Context, *UpdateArtistLikesCountRequest) (*UpdateArtistLikesCountResponse, error)
	// DELETE REQUEST
	DeleteArtistById(context.Context, *DeleteArtistByIdRequest) (*DeleteArtistByIdResponse, error)
	mustEmbedUnimplementedArtistServiceServer()
}

// UnimplementedArtistServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArtistServiceServer struct {
}

func (UnimplementedArtistServiceServer) CreateArtist(context.Context, *CreateArtistRequest) (*CreateArtistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArtist not implemented")
}
func (UnimplementedArtistServiceServer) ReadArtistById(context.Context, *ReadArtistByIdRequest) (*ReadArtistByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadArtistById not implemented")
}
func (UnimplementedArtistServiceServer) ReadArtistByName(context.Context, *ReadArtistByNameRequest) (*ReadArtistByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadArtistByName not implemented")
}
func (UnimplementedArtistServiceServer) ReadArtistBioById(context.Context, *ReadArtistBioRequest) (*ReadArtistBioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadArtistBioById not implemented")
}
func (UnimplementedArtistServiceServer) ReadArtistFollowerCountById(context.Context, *ReadArtistFollowerCountRequest) (*ReadArtistFollowerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadArtistFollowerCountById not implemented")
}
func (UnimplementedArtistServiceServer) ReadArtistLikesCountById(context.Context, *ReadArtistLikesCountRequest) (*ReadArtistLikesCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadArtistLikesCountById not implemented")
}
func (UnimplementedArtistServiceServer) UpdateArtistName(context.Context, *UpdateArtistNameRequest) (*UpdateArtistNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArtistName not implemented")
}
func (UnimplementedArtistServiceServer) UpdateArtistBio(context.Context, *UpdateArtistBioRequest) (*UpdateArtistBioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArtistBio not implemented")
}
func (UnimplementedArtistServiceServer) UpdateArtistFollowerCount(context.Context, *UpdateArtistFollowerCountRequest) (*UpdateArtistFollowerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArtistFollowerCount not implemented")
}
func (UnimplementedArtistServiceServer) UpdateArtistLikesCount(context.Context, *UpdateArtistLikesCountRequest) (*UpdateArtistLikesCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArtistLikesCount not implemented")
}
func (UnimplementedArtistServiceServer) DeleteArtistById(context.Context, *DeleteArtistByIdRequest) (*DeleteArtistByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArtistById not implemented")
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

func _ArtistService_CreateArtist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArtistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).CreateArtist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_CreateArtist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).CreateArtist(ctx, req.(*CreateArtistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_ReadArtistById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadArtistByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).ReadArtistById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_ReadArtistById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).ReadArtistById(ctx, req.(*ReadArtistByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_ReadArtistByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadArtistByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).ReadArtistByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_ReadArtistByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).ReadArtistByName(ctx, req.(*ReadArtistByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_ReadArtistBioById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadArtistBioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).ReadArtistBioById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_ReadArtistBioById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).ReadArtistBioById(ctx, req.(*ReadArtistBioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_ReadArtistFollowerCountById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadArtistFollowerCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).ReadArtistFollowerCountById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_ReadArtistFollowerCountById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).ReadArtistFollowerCountById(ctx, req.(*ReadArtistFollowerCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_ReadArtistLikesCountById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadArtistLikesCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).ReadArtistLikesCountById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_ReadArtistLikesCountById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).ReadArtistLikesCountById(ctx, req.(*ReadArtistLikesCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_UpdateArtistName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArtistNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).UpdateArtistName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_UpdateArtistName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).UpdateArtistName(ctx, req.(*UpdateArtistNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_UpdateArtistBio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArtistBioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).UpdateArtistBio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_UpdateArtistBio_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).UpdateArtistBio(ctx, req.(*UpdateArtistBioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_UpdateArtistFollowerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArtistFollowerCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).UpdateArtistFollowerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_UpdateArtistFollowerCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).UpdateArtistFollowerCount(ctx, req.(*UpdateArtistFollowerCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_UpdateArtistLikesCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArtistLikesCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).UpdateArtistLikesCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_UpdateArtistLikesCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).UpdateArtistLikesCount(ctx, req.(*UpdateArtistLikesCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_DeleteArtistById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtistByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).DeleteArtistById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtistService_DeleteArtistById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).DeleteArtistById(ctx, req.(*DeleteArtistByIdRequest))
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
			MethodName: "CreateArtist",
			Handler:    _ArtistService_CreateArtist_Handler,
		},
		{
			MethodName: "ReadArtistById",
			Handler:    _ArtistService_ReadArtistById_Handler,
		},
		{
			MethodName: "ReadArtistByName",
			Handler:    _ArtistService_ReadArtistByName_Handler,
		},
		{
			MethodName: "ReadArtistBioById",
			Handler:    _ArtistService_ReadArtistBioById_Handler,
		},
		{
			MethodName: "ReadArtistFollowerCountById",
			Handler:    _ArtistService_ReadArtistFollowerCountById_Handler,
		},
		{
			MethodName: "ReadArtistLikesCountById",
			Handler:    _ArtistService_ReadArtistLikesCountById_Handler,
		},
		{
			MethodName: "UpdateArtistName",
			Handler:    _ArtistService_UpdateArtistName_Handler,
		},
		{
			MethodName: "UpdateArtistBio",
			Handler:    _ArtistService_UpdateArtistBio_Handler,
		},
		{
			MethodName: "UpdateArtistFollowerCount",
			Handler:    _ArtistService_UpdateArtistFollowerCount_Handler,
		},
		{
			MethodName: "UpdateArtistLikesCount",
			Handler:    _ArtistService_UpdateArtistLikesCount_Handler,
		},
		{
			MethodName: "DeleteArtistById",
			Handler:    _ArtistService_DeleteArtistById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "artist_service.proto",
}
