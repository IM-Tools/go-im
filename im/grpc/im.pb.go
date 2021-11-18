// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: im/grpc/proto/im.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	FromId      int32  `protobuf:"varint,2,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	Msg         string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	ToId        int32  `protobuf:"varint,4,opt,name=to_id,json=toId,proto3" json:"to_id,omitempty"`
	Status      int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	MsgType     int32  `protobuf:"varint,6,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	ChannelType int32  `protobuf:"varint,7,opt,name=channel_type,json=channelType,proto3" json:"channel_type,omitempty"`
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_grpc_proto_im_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_im_grpc_proto_im_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_im_grpc_proto_im_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRequest) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MessageRequest) GetFromId() int32 {
	if x != nil {
		return x.FromId
	}
	return 0
}

func (x *MessageRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *MessageRequest) GetToId() int32 {
	if x != nil {
		return x.ToId
	}
	return 0
}

func (x *MessageRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MessageRequest) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *MessageRequest) GetChannelType() int32 {
	if x != nil {
		return x.ChannelType
	}
	return 0
}

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_grpc_proto_im_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_im_grpc_proto_im_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_im_grpc_proto_im_proto_rawDescGZIP(), []int{1}
}

func (x *MessageResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_im_grpc_proto_im_proto protoreflect.FileDescriptor

var file_im_grpc_proto_im_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0xba, 0x01, 0x0a,
	0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x13,
	0x0a, 0x05, 0x74, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x6f, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d,
	0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0x68, 0x0a, 0x0c, 0x49, 0x6d, 0x52, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x58, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x22, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f,
	0x69, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_im_grpc_proto_im_proto_rawDescOnce sync.Once
	file_im_grpc_proto_im_proto_rawDescData = file_im_grpc_proto_im_proto_rawDesc
)

func file_im_grpc_proto_im_proto_rawDescGZIP() []byte {
	file_im_grpc_proto_im_proto_rawDescOnce.Do(func() {
		file_im_grpc_proto_im_proto_rawDescData = protoimpl.X.CompressGZIP(file_im_grpc_proto_im_proto_rawDescData)
	})
	return file_im_grpc_proto_im_proto_rawDescData
}

var file_im_grpc_proto_im_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_im_grpc_proto_im_proto_goTypes = []interface{}{
	(*MessageRequest)(nil),  // 0: Message.SearchNode.MessageRequest
	(*MessageResponse)(nil), // 1: Message.SearchNode.MessageResponse
}
var file_im_grpc_proto_im_proto_depIdxs = []int32{
	0, // 0: Message.SearchNode.ImRpcService.SendMessage:input_type -> Message.SearchNode.MessageRequest
	1, // 1: Message.SearchNode.ImRpcService.SendMessage:output_type -> Message.SearchNode.MessageResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_im_grpc_proto_im_proto_init() }
func file_im_grpc_proto_im_proto_init() {
	if File_im_grpc_proto_im_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_im_grpc_proto_im_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRequest); i {
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
		file_im_grpc_proto_im_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageResponse); i {
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
			RawDescriptor: file_im_grpc_proto_im_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_im_grpc_proto_im_proto_goTypes,
		DependencyIndexes: file_im_grpc_proto_im_proto_depIdxs,
		MessageInfos:      file_im_grpc_proto_im_proto_msgTypes,
	}.Build()
	File_im_grpc_proto_im_proto = out.File
	file_im_grpc_proto_im_proto_rawDesc = nil
	file_im_grpc_proto_im_proto_goTypes = nil
	file_im_grpc_proto_im_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ImRpcServiceClient is the client API for ImRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImRpcServiceClient interface {
	SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type imRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImRpcServiceClient(cc grpc.ClientConnInterface) ImRpcServiceClient {
	return &imRpcServiceClient{cc}
}

func (c *imRpcServiceClient) SendMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/Message.SearchNode.ImRpcService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImRpcServiceServer is the server API for ImRpcService service.
type ImRpcServiceServer interface {
	SendMessage(context.Context, *MessageRequest) (*MessageResponse, error)
}

// UnimplementedImRpcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedImRpcServiceServer struct {
}

func (*UnimplementedImRpcServiceServer) SendMessage(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterImRpcServiceServer(s *grpc.Server, srv ImRpcServiceServer) {
	s.RegisterService(&_ImRpcService_serviceDesc, srv)
}

func _ImRpcService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImRpcServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Message.SearchNode.ImRpcService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImRpcServiceServer).SendMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImRpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Message.SearchNode.ImRpcService",
	HandlerType: (*ImRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ImRpcService_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im/grpc/proto/im.proto",
}