// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchain_security/ccv/provider/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	types1 "github.com/cosmos/ibc-go/v4/modules/light-clients/07-tendermint/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgAssignConsumerKey struct {
	// The chain id of the consumer chain to assign a consensus public key to
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// The validator address on the provider
	ProviderAddr string `protobuf:"bytes,2,opt,name=provider_addr,json=providerAddr,proto3" json:"provider_addr,omitempty" yaml:"address"`
	// The consensus public key to use on the consumer
	ConsumerKey *types.Any `protobuf:"bytes,3,opt,name=consumer_key,json=consumerKey,proto3" json:"consumer_key,omitempty"`
}

func (m *MsgAssignConsumerKey) Reset()         { *m = MsgAssignConsumerKey{} }
func (m *MsgAssignConsumerKey) String() string { return proto.CompactTextString(m) }
func (*MsgAssignConsumerKey) ProtoMessage()    {}
func (*MsgAssignConsumerKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_43221a4391e9fbf4, []int{0}
}
func (m *MsgAssignConsumerKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAssignConsumerKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAssignConsumerKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAssignConsumerKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAssignConsumerKey.Merge(m, src)
}
func (m *MsgAssignConsumerKey) XXX_Size() int {
	return m.Size()
}
func (m *MsgAssignConsumerKey) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAssignConsumerKey.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAssignConsumerKey proto.InternalMessageInfo

type MsgAssignConsumerKeyResponse struct {
}

func (m *MsgAssignConsumerKeyResponse) Reset()         { *m = MsgAssignConsumerKeyResponse{} }
func (m *MsgAssignConsumerKeyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAssignConsumerKeyResponse) ProtoMessage()    {}
func (*MsgAssignConsumerKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43221a4391e9fbf4, []int{1}
}
func (m *MsgAssignConsumerKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAssignConsumerKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAssignConsumerKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAssignConsumerKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAssignConsumerKeyResponse.Merge(m, src)
}
func (m *MsgAssignConsumerKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAssignConsumerKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAssignConsumerKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAssignConsumerKeyResponse proto.InternalMessageInfo

// MsgSubmitConsumerMisbehaviour defines a message that reports a misbehaviour
// observed on a consumer chain
type MsgSubmitConsumerMisbehaviour struct {
	Submitter string `protobuf:"bytes,1,opt,name=submitter,proto3" json:"submitter,omitempty"`
	// The Misbehaviour of the consumer chain wrapping
	// two conflicting IBC headers
	Misbehaviour *types1.Misbehaviour `protobuf:"bytes,2,opt,name=misbehaviour,proto3" json:"misbehaviour,omitempty"`
}

func (m *MsgSubmitConsumerMisbehaviour) Reset()         { *m = MsgSubmitConsumerMisbehaviour{} }
func (m *MsgSubmitConsumerMisbehaviour) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitConsumerMisbehaviour) ProtoMessage()    {}
func (*MsgSubmitConsumerMisbehaviour) Descriptor() ([]byte, []int) {
	return fileDescriptor_43221a4391e9fbf4, []int{2}
}
func (m *MsgSubmitConsumerMisbehaviour) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitConsumerMisbehaviour) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitConsumerMisbehaviour.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitConsumerMisbehaviour) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitConsumerMisbehaviour.Merge(m, src)
}
func (m *MsgSubmitConsumerMisbehaviour) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitConsumerMisbehaviour) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitConsumerMisbehaviour.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitConsumerMisbehaviour proto.InternalMessageInfo

type MsgSubmitConsumerMisbehaviourResponse struct {
}

func (m *MsgSubmitConsumerMisbehaviourResponse) Reset()         { *m = MsgSubmitConsumerMisbehaviourResponse{} }
func (m *MsgSubmitConsumerMisbehaviourResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitConsumerMisbehaviourResponse) ProtoMessage()    {}
func (*MsgSubmitConsumerMisbehaviourResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43221a4391e9fbf4, []int{3}
}
func (m *MsgSubmitConsumerMisbehaviourResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitConsumerMisbehaviourResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitConsumerMisbehaviourResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitConsumerMisbehaviourResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitConsumerMisbehaviourResponse.Merge(m, src)
}
func (m *MsgSubmitConsumerMisbehaviourResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitConsumerMisbehaviourResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitConsumerMisbehaviourResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitConsumerMisbehaviourResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAssignConsumerKey)(nil), "interchain_security.ccv.provider.v1.MsgAssignConsumerKey")
	proto.RegisterType((*MsgAssignConsumerKeyResponse)(nil), "interchain_security.ccv.provider.v1.MsgAssignConsumerKeyResponse")
	proto.RegisterType((*MsgSubmitConsumerMisbehaviour)(nil), "interchain_security.ccv.provider.v1.MsgSubmitConsumerMisbehaviour")
	proto.RegisterType((*MsgSubmitConsumerMisbehaviourResponse)(nil), "interchain_security.ccv.provider.v1.MsgSubmitConsumerMisbehaviourResponse")
}

func init() {
	proto.RegisterFile("interchain_security/ccv/provider/v1/tx.proto", fileDescriptor_43221a4391e9fbf4)
}

var fileDescriptor_43221a4391e9fbf4 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x3f, 0x6f, 0x13, 0x3f,
	0x18, 0x8e, 0x5b, 0xe9, 0xf7, 0x6b, 0x9d, 0x80, 0xc4, 0x29, 0x43, 0x7a, 0x0a, 0x97, 0xea, 0x10,
	0xa2, 0x43, 0x6b, 0x2b, 0x61, 0x40, 0x64, 0x4b, 0x98, 0xa0, 0x8a, 0x54, 0x0e, 0x26, 0x96, 0xe8,
	0xce, 0x67, 0x1c, 0x8b, 0x9c, 0x7d, 0xb2, 0x7d, 0xa7, 0xde, 0x37, 0x60, 0x84, 0x91, 0xad, 0x1f,
	0x80, 0x91, 0xef, 0x40, 0xc5, 0xd4, 0x91, 0x09, 0xa1, 0x64, 0x61, 0xe6, 0x13, 0xa0, 0xdc, 0x9f,
	0xe6, 0x2a, 0x42, 0x55, 0xc1, 0xe6, 0xd7, 0xcf, 0xe3, 0xe7, 0x7d, 0x9e, 0xf7, 0x7c, 0x86, 0x87,
	0x5c, 0x18, 0xaa, 0xc8, 0xcc, 0xe7, 0x62, 0xaa, 0x29, 0x49, 0x14, 0x37, 0x19, 0x26, 0x24, 0xc5,
	0xb1, 0x92, 0x29, 0x0f, 0xa9, 0xc2, 0x69, 0x1f, 0x9b, 0x53, 0x14, 0x2b, 0x69, 0xa4, 0x75, 0x6f,
	0x03, 0x1b, 0x11, 0x92, 0xa2, 0x8a, 0x8d, 0xd2, 0xbe, 0xdd, 0x65, 0x52, 0xb2, 0x39, 0xc5, 0x7e,
	0xcc, 0xb1, 0x2f, 0x84, 0x34, 0xbe, 0xe1, 0x52, 0xe8, 0x42, 0xc2, 0x6e, 0x33, 0xc9, 0x64, 0xbe,
	0xc4, 0xab, 0x55, 0xb9, 0xbb, 0x47, 0xa4, 0x8e, 0xa4, 0x9e, 0x16, 0x40, 0x51, 0x54, 0x50, 0x29,
	0x97, 0x57, 0x41, 0xf2, 0x1a, 0xfb, 0x22, 0x2b, 0x21, 0xcc, 0x03, 0x82, 0xe7, 0x9c, 0xcd, 0x0c,
	0x99, 0x73, 0x2a, 0x8c, 0xc6, 0x86, 0x8a, 0x90, 0xaa, 0x88, 0x0b, 0x93, 0xfb, 0xbe, 0xac, 0x8a,
	0x03, 0xee, 0x67, 0x00, 0xdb, 0x13, 0xcd, 0x46, 0x5a, 0x73, 0x26, 0x9e, 0x48, 0xa1, 0x93, 0x88,
	0xaa, 0x63, 0x9a, 0x59, 0x7b, 0x70, 0xa7, 0x48, 0xc5, 0xc3, 0x0e, 0xd8, 0x07, 0x07, 0xbb, 0xde,
	0xff, 0x79, 0xfd, 0x34, 0xb4, 0x1e, 0xc1, 0x5b, 0x55, 0xba, 0xa9, 0x1f, 0x86, 0xaa, 0xb3, 0xb5,
	0xc2, 0xc7, 0xd6, 0xcf, 0x6f, 0xbd, 0xdb, 0x99, 0x1f, 0xcd, 0x87, 0xee, 0x6a, 0x97, 0x6a, 0xed,
	0x7a, 0xad, 0x8a, 0x38, 0x0a, 0x43, 0x65, 0x3d, 0x87, 0x2d, 0x52, 0xb6, 0x98, 0xbe, 0xa1, 0x59,
	0x67, 0x7b, 0x1f, 0x1c, 0x34, 0x07, 0x6d, 0x54, 0xe4, 0x41, 0x55, 0x1e, 0x34, 0x12, 0xd9, 0xb8,
	0xf3, 0xe5, 0xd3, 0x51, 0xbb, 0x8c, 0x4d, 0x54, 0x16, 0x1b, 0x89, 0x4e, 0x92, 0xe0, 0x98, 0x66,
	0x5e, 0x93, 0xac, 0x6d, 0x0e, 0x77, 0xde, 0x9e, 0xf5, 0x1a, 0x3f, 0xce, 0x7a, 0x0d, 0xd7, 0x81,
	0xdd, 0x4d, 0x41, 0x3c, 0xaa, 0x63, 0x29, 0x34, 0x75, 0x3f, 0x00, 0x78, 0x77, 0xa2, 0xd9, 0x8b,
	0x24, 0x88, 0xb8, 0xa9, 0x08, 0x13, 0xae, 0x03, 0x3a, 0xf3, 0x53, 0x2e, 0x13, 0x65, 0x75, 0xe1,
	0xae, 0xce, 0x51, 0x43, 0x55, 0x99, 0x79, 0xbd, 0x61, 0x9d, 0xc0, 0x56, 0x54, 0x63, 0xe7, 0xa1,
	0x9b, 0x83, 0x43, 0xc4, 0x03, 0x82, 0xea, 0x13, 0x47, 0xb5, 0x19, 0xa7, 0x7d, 0x54, 0xef, 0xe0,
	0x5d, 0x51, 0xa8, 0x79, 0x7f, 0x00, 0xef, 0x5f, 0x6b, 0xad, 0x0a, 0x31, 0x38, 0xdf, 0x82, 0xdb,
	0x13, 0xcd, 0xac, 0xf7, 0x00, 0xde, 0xf9, 0xfd, 0x9b, 0x3d, 0x46, 0x37, 0xb8, 0x8d, 0x68, 0xd3,
	0x94, 0xec, 0xd1, 0x5f, 0x1f, 0xad, 0xbc, 0x59, 0x1f, 0x01, 0xb4, 0xaf, 0x99, 0xee, 0xf8, 0xa6,
	0x1d, 0xfe, 0xac, 0x61, 0x3f, 0xfb, 0x77, 0x8d, 0xca, 0xee, 0xf8, 0xe5, 0xf9, 0xc2, 0x01, 0x17,
	0x0b, 0x07, 0x7c, 0x5f, 0x38, 0xe0, 0xdd, 0xd2, 0x69, 0x5c, 0x2c, 0x9d, 0xc6, 0xd7, 0xa5, 0xd3,
	0x78, 0x35, 0x64, 0xdc, 0xcc, 0x92, 0x00, 0x11, 0x19, 0x95, 0x3f, 0x1e, 0x5e, 0xb7, 0x3d, 0xba,
	0x7c, 0x13, 0x4e, 0xaf, 0xbe, 0x0a, 0x26, 0x8b, 0xa9, 0x0e, 0xfe, 0xcb, 0x2f, 0xf1, 0xc3, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x41, 0x76, 0xf4, 0x46, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	AssignConsumerKey(ctx context.Context, in *MsgAssignConsumerKey, opts ...grpc.CallOption) (*MsgAssignConsumerKeyResponse, error)
	SubmitConsumerMisbehaviour(ctx context.Context, in *MsgSubmitConsumerMisbehaviour, opts ...grpc.CallOption) (*MsgSubmitConsumerMisbehaviourResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AssignConsumerKey(ctx context.Context, in *MsgAssignConsumerKey, opts ...grpc.CallOption) (*MsgAssignConsumerKeyResponse, error) {
	out := new(MsgAssignConsumerKeyResponse)
	err := c.cc.Invoke(ctx, "/interchain_security.ccv.provider.v1.Msg/AssignConsumerKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitConsumerMisbehaviour(ctx context.Context, in *MsgSubmitConsumerMisbehaviour, opts ...grpc.CallOption) (*MsgSubmitConsumerMisbehaviourResponse, error) {
	out := new(MsgSubmitConsumerMisbehaviourResponse)
	err := c.cc.Invoke(ctx, "/interchain_security.ccv.provider.v1.Msg/SubmitConsumerMisbehaviour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AssignConsumerKey(context.Context, *MsgAssignConsumerKey) (*MsgAssignConsumerKeyResponse, error)
	SubmitConsumerMisbehaviour(context.Context, *MsgSubmitConsumerMisbehaviour) (*MsgSubmitConsumerMisbehaviourResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AssignConsumerKey(ctx context.Context, req *MsgAssignConsumerKey) (*MsgAssignConsumerKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignConsumerKey not implemented")
}
func (*UnimplementedMsgServer) SubmitConsumerMisbehaviour(ctx context.Context, req *MsgSubmitConsumerMisbehaviour) (*MsgSubmitConsumerMisbehaviourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitConsumerMisbehaviour not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AssignConsumerKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAssignConsumerKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AssignConsumerKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interchain_security.ccv.provider.v1.Msg/AssignConsumerKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AssignConsumerKey(ctx, req.(*MsgAssignConsumerKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitConsumerMisbehaviour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitConsumerMisbehaviour)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitConsumerMisbehaviour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interchain_security.ccv.provider.v1.Msg/SubmitConsumerMisbehaviour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitConsumerMisbehaviour(ctx, req.(*MsgSubmitConsumerMisbehaviour))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interchain_security.ccv.provider.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignConsumerKey",
			Handler:    _Msg_AssignConsumerKey_Handler,
		},
		{
			MethodName: "SubmitConsumerMisbehaviour",
			Handler:    _Msg_SubmitConsumerMisbehaviour_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interchain_security/ccv/provider/v1/tx.proto",
}

func (m *MsgAssignConsumerKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAssignConsumerKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAssignConsumerKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ConsumerKey != nil {
		{
			size, err := m.ConsumerKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ProviderAddr) > 0 {
		i -= len(m.ProviderAddr)
		copy(dAtA[i:], m.ProviderAddr)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ProviderAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAssignConsumerKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAssignConsumerKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAssignConsumerKeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSubmitConsumerMisbehaviour) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitConsumerMisbehaviour) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitConsumerMisbehaviour) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Misbehaviour != nil {
		{
			size, err := m.Misbehaviour.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Submitter) > 0 {
		i -= len(m.Submitter)
		copy(dAtA[i:], m.Submitter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Submitter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitConsumerMisbehaviourResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitConsumerMisbehaviourResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitConsumerMisbehaviourResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAssignConsumerKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ProviderAddr)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ConsumerKey != nil {
		l = m.ConsumerKey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAssignConsumerKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSubmitConsumerMisbehaviour) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Submitter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Misbehaviour != nil {
		l = m.Misbehaviour.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitConsumerMisbehaviourResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAssignConsumerKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAssignConsumerKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAssignConsumerKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsumerKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConsumerKey == nil {
				m.ConsumerKey = &types.Any{}
			}
			if err := m.ConsumerKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAssignConsumerKeyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAssignConsumerKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAssignConsumerKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitConsumerMisbehaviour) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitConsumerMisbehaviour: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitConsumerMisbehaviour: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Submitter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Submitter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Misbehaviour", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Misbehaviour == nil {
				m.Misbehaviour = &types1.Misbehaviour{}
			}
			if err := m.Misbehaviour.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitConsumerMisbehaviourResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitConsumerMisbehaviourResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitConsumerMisbehaviourResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
