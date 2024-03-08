// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0--rc3
// source: playlist_service.proto

package grpcapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Playlist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId int32  `protobuf:"varint,1,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
	Title      string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CreatorId  int32  `protobuf:"varint,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Playlist) Reset() {
	*x = Playlist{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Playlist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Playlist) ProtoMessage() {}

func (x *Playlist) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Playlist.ProtoReflect.Descriptor instead.
func (*Playlist) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{0}
}

func (x *Playlist) GetPlaylistId() int32 {
	if x != nil {
		return x.PlaylistId
	}
	return 0
}

func (x *Playlist) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Playlist) GetCreatorId() int32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Playlist) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ViewSongsByPlaylistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId int32 `protobuf:"varint,1,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
}

func (x *ViewSongsByPlaylistRequest) Reset() {
	*x = ViewSongsByPlaylistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewSongsByPlaylistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewSongsByPlaylistRequest) ProtoMessage() {}

func (x *ViewSongsByPlaylistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewSongsByPlaylistRequest.ProtoReflect.Descriptor instead.
func (*ViewSongsByPlaylistRequest) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{1}
}

func (x *ViewSongsByPlaylistRequest) GetPlaylistId() int32 {
	if x != nil {
		return x.PlaylistId
	}
	return 0
}

type ViewSongsByPlaylistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Songs []*Song `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"`
}

func (x *ViewSongsByPlaylistResponse) Reset() {
	*x = ViewSongsByPlaylistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewSongsByPlaylistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewSongsByPlaylistResponse) ProtoMessage() {}

func (x *ViewSongsByPlaylistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewSongsByPlaylistResponse.ProtoReflect.Descriptor instead.
func (*ViewSongsByPlaylistResponse) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{2}
}

func (x *ViewSongsByPlaylistResponse) GetSongs() []*Song {
	if x != nil {
		return x.Songs
	}
	return nil
}

type GetPlaylistByTitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GetPlaylistByTitleRequest) Reset() {
	*x = GetPlaylistByTitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlaylistByTitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlaylistByTitleRequest) ProtoMessage() {}

func (x *GetPlaylistByTitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlaylistByTitleRequest.ProtoReflect.Descriptor instead.
func (*GetPlaylistByTitleRequest) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetPlaylistByTitleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetPlaylistByTitleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Playlists []*Playlist `protobuf:"bytes,1,rep,name=playlists,proto3" json:"playlists,omitempty"`
}

func (x *GetPlaylistByTitleResponse) Reset() {
	*x = GetPlaylistByTitleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlaylistByTitleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlaylistByTitleResponse) ProtoMessage() {}

func (x *GetPlaylistByTitleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlaylistByTitleResponse.ProtoReflect.Descriptor instead.
func (*GetPlaylistByTitleResponse) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlaylistByTitleResponse) GetPlaylists() []*Playlist {
	if x != nil {
		return x.Playlists
	}
	return nil
}

type GetPlaylistByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId int32 `protobuf:"varint,1,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
}

func (x *GetPlaylistByIdRequest) Reset() {
	*x = GetPlaylistByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlaylistByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlaylistByIdRequest) ProtoMessage() {}

func (x *GetPlaylistByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlaylistByIdRequest.ProtoReflect.Descriptor instead.
func (*GetPlaylistByIdRequest) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetPlaylistByIdRequest) GetPlaylistId() int32 {
	if x != nil {
		return x.PlaylistId
	}
	return 0
}

type GetPlaylistByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Playlist *Playlist `protobuf:"bytes,1,opt,name=playlist,proto3" json:"playlist,omitempty"`
}

func (x *GetPlaylistByIdResponse) Reset() {
	*x = GetPlaylistByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlaylistByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlaylistByIdResponse) ProtoMessage() {}

func (x *GetPlaylistByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlaylistByIdResponse.ProtoReflect.Descriptor instead.
func (*GetPlaylistByIdResponse) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetPlaylistByIdResponse) GetPlaylist() *Playlist {
	if x != nil {
		return x.Playlist
	}
	return nil
}

type GetPlaylistCreationTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistId int32 `protobuf:"varint,1,opt,name=playlist_id,json=playlistId,proto3" json:"playlist_id,omitempty"`
}

func (x *GetPlaylistCreationTimeRequest) Reset() {
	*x = GetPlaylistCreationTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlaylistCreationTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlaylistCreationTimeRequest) ProtoMessage() {}

func (x *GetPlaylistCreationTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlaylistCreationTimeRequest.ProtoReflect.Descriptor instead.
func (*GetPlaylistCreationTimeRequest) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetPlaylistCreationTimeRequest) GetPlaylistId() int32 {
	if x != nil {
		return x.PlaylistId
	}
	return 0
}

type GetPlaylistCreationTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreationTime string `protobuf:"bytes,1,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"` // Consider using a timestamp type
}

func (x *GetPlaylistCreationTimeResponse) Reset() {
	*x = GetPlaylistCreationTimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlaylistCreationTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlaylistCreationTimeResponse) ProtoMessage() {}

func (x *GetPlaylistCreationTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlaylistCreationTimeResponse.ProtoReflect.Descriptor instead.
func (*GetPlaylistCreationTimeResponse) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetPlaylistCreationTimeResponse) GetCreationTime() string {
	if x != nil {
		return x.CreationTime
	}
	return ""
}

var File_playlist_service_proto protoreflect.FileDescriptor

var file_playlist_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x70, 0x64, 0x62, 0x1a, 0x13,
	0x73, 0x6f, 0x6e, 0x67, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x3d, 0x0a, 0x1a, 0x56, 0x69, 0x65, 0x77, 0x53, 0x6f, 0x6e, 0x67,
	0x73, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x1b, 0x56, 0x69, 0x65, 0x77, 0x53, 0x6f, 0x6e, 0x67, 0x73,
	0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x70, 0x64, 0x62, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x05, 0x73,
	0x6f, 0x6e, 0x67, 0x73, 0x22, 0x31, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x4a, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x70, 0x64, 0x62, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x73, 0x22, 0x39, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x22, 0x45,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x70,
	0x64, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79,
	0x6c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x32, 0xfe, 0x02, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x13, 0x56, 0x69, 0x65, 0x77, 0x53, 0x6f, 0x6e, 0x67,
	0x73, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x6d, 0x70,
	0x64, 0x62, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x42, 0x79, 0x50, 0x6c,
	0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x6d, 0x70, 0x64, 0x62, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x42, 0x79,
	0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x57, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x42,
	0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x6d, 0x70, 0x64, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x70, 0x64, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x6d,
	0x70, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x70, 0x64,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x2e, 0x6d, 0x70, 0x64, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x70, 0x64,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x73, 0x68, 0x69, 0x62, 0x6c, 0x69, 0x32, 0x30, 0x30, 0x32, 0x2f, 0x6d, 0x62, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_playlist_service_proto_rawDescOnce sync.Once
	file_playlist_service_proto_rawDescData = file_playlist_service_proto_rawDesc
)

func file_playlist_service_proto_rawDescGZIP() []byte {
	file_playlist_service_proto_rawDescOnce.Do(func() {
		file_playlist_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_playlist_service_proto_rawDescData)
	})
	return file_playlist_service_proto_rawDescData
}

var file_playlist_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_playlist_service_proto_goTypes = []interface{}{
	(*Playlist)(nil),                        // 0: mpdb.Playlist
	(*ViewSongsByPlaylistRequest)(nil),      // 1: mpdb.ViewSongsByPlaylistRequest
	(*ViewSongsByPlaylistResponse)(nil),     // 2: mpdb.ViewSongsByPlaylistResponse
	(*GetPlaylistByTitleRequest)(nil),       // 3: mpdb.GetPlaylistByTitleRequest
	(*GetPlaylistByTitleResponse)(nil),      // 4: mpdb.GetPlaylistByTitleResponse
	(*GetPlaylistByIdRequest)(nil),          // 5: mpdb.GetPlaylistByIdRequest
	(*GetPlaylistByIdResponse)(nil),         // 6: mpdb.GetPlaylistByIdResponse
	(*GetPlaylistCreationTimeRequest)(nil),  // 7: mpdb.GetPlaylistCreationTimeRequest
	(*GetPlaylistCreationTimeResponse)(nil), // 8: mpdb.GetPlaylistCreationTimeResponse
	(*Song)(nil),                            // 9: mpdb.Song
}
var file_playlist_service_proto_depIdxs = []int32{
	9, // 0: mpdb.ViewSongsByPlaylistResponse.songs:type_name -> mpdb.Song
	0, // 1: mpdb.GetPlaylistByTitleResponse.playlists:type_name -> mpdb.Playlist
	0, // 2: mpdb.GetPlaylistByIdResponse.playlist:type_name -> mpdb.Playlist
	1, // 3: mpdb.PlaylistService.ViewSongsByPlaylist:input_type -> mpdb.ViewSongsByPlaylistRequest
	3, // 4: mpdb.PlaylistService.GetPlaylistByTitle:input_type -> mpdb.GetPlaylistByTitleRequest
	5, // 5: mpdb.PlaylistService.GetPlaylistById:input_type -> mpdb.GetPlaylistByIdRequest
	7, // 6: mpdb.PlaylistService.GetPlaylistCreationTime:input_type -> mpdb.GetPlaylistCreationTimeRequest
	2, // 7: mpdb.PlaylistService.ViewSongsByPlaylist:output_type -> mpdb.ViewSongsByPlaylistResponse
	4, // 8: mpdb.PlaylistService.GetPlaylistByTitle:output_type -> mpdb.GetPlaylistByTitleResponse
	6, // 9: mpdb.PlaylistService.GetPlaylistById:output_type -> mpdb.GetPlaylistByIdResponse
	8, // 10: mpdb.PlaylistService.GetPlaylistCreationTime:output_type -> mpdb.GetPlaylistCreationTimeResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_playlist_service_proto_init() }
func file_playlist_service_proto_init() {
	if File_playlist_service_proto != nil {
		return
	}
	file_songs_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_playlist_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Playlist); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewSongsByPlaylistRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewSongsByPlaylistResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlaylistByTitleRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlaylistByTitleResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlaylistByIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlaylistByIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlaylistCreationTimeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlaylistCreationTimeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_playlist_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_playlist_service_proto_goTypes,
		DependencyIndexes: file_playlist_service_proto_depIdxs,
		MessageInfos:      file_playlist_service_proto_msgTypes,
	}.Build()
	File_playlist_service_proto = out.File
	file_playlist_service_proto_rawDesc = nil
	file_playlist_service_proto_goTypes = nil
	file_playlist_service_proto_depIdxs = nil
}
