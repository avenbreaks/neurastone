// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neura/coinomics/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type QueryMaxSupplyRequest struct {
}

func (m *QueryMaxSupplyRequest) Reset()         { *m = QueryMaxSupplyRequest{} }
func (m *QueryMaxSupplyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryMaxSupplyRequest) ProtoMessage()    {}
func (*QueryMaxSupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bccdbe1913fce8f0, []int{0}
}
func (m *QueryMaxSupplyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMaxSupplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMaxSupplyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMaxSupplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMaxSupplyRequest.Merge(m, src)
}
func (m *QueryMaxSupplyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryMaxSupplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMaxSupplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMaxSupplyRequest proto.InternalMessageInfo

type QueryMaxSupplyResponse struct {
	MaxSupply github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,1,opt,name=max_supply,json=maxSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"max_supply"`
}

func (m *QueryMaxSupplyResponse) Reset()         { *m = QueryMaxSupplyResponse{} }
func (m *QueryMaxSupplyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMaxSupplyResponse) ProtoMessage()    {}
func (*QueryMaxSupplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bccdbe1913fce8f0, []int{1}
}
func (m *QueryMaxSupplyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMaxSupplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMaxSupplyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMaxSupplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMaxSupplyResponse.Merge(m, src)
}
func (m *QueryMaxSupplyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMaxSupplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMaxSupplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMaxSupplyResponse proto.InternalMessageInfo

type QueryRewardCoefficientRequest struct {
}

func (m *QueryRewardCoefficientRequest) Reset()         { *m = QueryRewardCoefficientRequest{} }
func (m *QueryRewardCoefficientRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRewardCoefficientRequest) ProtoMessage()    {}
func (*QueryRewardCoefficientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bccdbe1913fce8f0, []int{2}
}
func (m *QueryRewardCoefficientRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRewardCoefficientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRewardCoefficientRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRewardCoefficientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRewardCoefficientRequest.Merge(m, src)
}
func (m *QueryRewardCoefficientRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRewardCoefficientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRewardCoefficientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRewardCoefficientRequest proto.InternalMessageInfo

type QueryRewardCoefficientResponse struct {
	// rate by which the total supply increases within one era
	RewardCoefficient github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=reward_coefficient,json=rewardCoefficient,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_coefficient"`
}

func (m *QueryRewardCoefficientResponse) Reset()         { *m = QueryRewardCoefficientResponse{} }
func (m *QueryRewardCoefficientResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRewardCoefficientResponse) ProtoMessage()    {}
func (*QueryRewardCoefficientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bccdbe1913fce8f0, []int{3}
}
func (m *QueryRewardCoefficientResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRewardCoefficientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRewardCoefficientResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRewardCoefficientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRewardCoefficientResponse.Merge(m, src)
}
func (m *QueryRewardCoefficientResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRewardCoefficientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRewardCoefficientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRewardCoefficientResponse proto.InternalMessageInfo

type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bccdbe1913fce8f0, []int{4}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bccdbe1913fce8f0, []int{5}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryMaxSupplyRequest)(nil), "neura.coinomics.v1.QueryMaxSupplyRequest")
	proto.RegisterType((*QueryMaxSupplyResponse)(nil), "neura.coinomics.v1.QueryMaxSupplyResponse")
	proto.RegisterType((*QueryRewardCoefficientRequest)(nil), "neura.coinomics.v1.QueryRewardCoefficientRequest")
	proto.RegisterType((*QueryRewardCoefficientResponse)(nil), "neura.coinomics.v1.QueryRewardCoefficientResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "neura.coinomics.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "neura.coinomics.v1.QueryParamsResponse")
}

func init() { proto.RegisterFile("neura/coinomics/v1/query.proto", fileDescriptor_bccdbe1913fce8f0) }

var fileDescriptor_bccdbe1913fce8f0 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x63, 0x2e, 0x91, 0x32, 0xac, 0x32, 0x94, 0x4b, 0x0d, 0x75, 0x8a, 0xa5, 0xa6, 0x65,
	0x91, 0x19, 0x5c, 0x16, 0xec, 0x53, 0x16, 0x6c, 0x40, 0x60, 0x76, 0x48, 0xa8, 0x9a, 0x38, 0x53,
	0x77, 0xd4, 0x7a, 0x8e, 0xe3, 0x99, 0xa4, 0x09, 0x1b, 0x24, 0xd8, 0xb0, 0x44, 0xe2, 0x21, 0x78,
	0x05, 0x1e, 0xa1, 0xcb, 0x4a, 0x6c, 0x10, 0x8b, 0x0a, 0x25, 0x3c, 0x08, 0xf2, 0x78, 0x92, 0x36,
	0x75, 0x22, 0xda, 0x95, 0xad, 0xf3, 0x9f, 0xcb, 0xa7, 0xf3, 0x9f, 0x41, 0x6b, 0xfb, 0xac, 0xd7,
	0xa3, 0x11, 0x08, 0x09, 0x89, 0x88, 0x14, 0x1d, 0x04, 0xb4, 0xd7, 0xe7, 0xd9, 0x88, 0xa4, 0x19,
	0x68, 0xc0, 0xf5, 0x5c, 0x26, 0x33, 0x99, 0x0c, 0x02, 0x77, 0x25, 0x86, 0x18, 0x8c, 0x4a, 0xf3,
	0xbf, 0x22, 0xd1, 0x7d, 0x18, 0x03, 0xc4, 0x87, 0x9c, 0xb2, 0x54, 0x50, 0x26, 0x25, 0x68, 0xa6,
	0x05, 0x48, 0x65, 0x55, 0x2f, 0x02, 0x95, 0x80, 0xa2, 0x1d, 0xa6, 0x38, 0x1d, 0x04, 0x1d, 0xae,
	0x59, 0x60, 0x86, 0x5a, 0xbd, 0x51, 0xa6, 0x88, 0xb9, 0xe4, 0x4a, 0xd8, 0x06, 0xfe, 0x3d, 0x74,
	0xe7, 0x4d, 0x8e, 0xf5, 0x92, 0x0d, 0xdf, 0xf6, 0xd3, 0xf4, 0x70, 0x14, 0xf2, 0x5e, 0x9f, 0x2b,
	0xed, 0x7f, 0x76, 0xd0, 0xdd, 0x8b, 0x8a, 0x4a, 0x41, 0x2a, 0x8e, 0x05, 0x42, 0x09, 0x1b, 0xee,
	0x2a, 0x13, 0xbd, 0xef, 0xac, 0x3b, 0x5b, 0xb7, 0xb6, 0x57, 0x49, 0x41, 0x42, 0x72, 0x12, 0x62,
	0x49, 0xc8, 0x0e, 0x08, 0xd9, 0xa6, 0xc7, 0xa7, 0x8d, 0xca, 0xef, 0xd3, 0xc6, 0x66, 0x2c, 0xf4,
	0x7e, 0xbf, 0x43, 0x22, 0x48, 0xa8, 0xc5, 0x2e, 0x3e, 0x2d, 0xd5, 0x3d, 0xa0, 0x7a, 0x94, 0x72,
	0x65, 0x0a, 0xc2, 0x5a, 0x32, 0x1d, 0xe9, 0x37, 0xd0, 0x9a, 0x81, 0x08, 0xf9, 0x11, 0xcb, 0xba,
	0x3b, 0xc0, 0xf7, 0xf6, 0x44, 0x24, 0xb8, 0xd4, 0x53, 0xcc, 0x8f, 0xc8, 0x5b, 0x96, 0x60, 0x69,
	0xdf, 0x23, 0x9c, 0x19, 0x71, 0x37, 0x3a, 0x53, 0x0d, 0x75, 0xad, 0x4d, 0x2c, 0x5a, 0xf3, 0x12,
	0x68, 0xcf, 0x79, 0x14, 0xd6, 0xb3, 0x8b, 0x63, 0xfc, 0x15, 0x84, 0x0d, 0xc0, 0x6b, 0x96, 0xb1,
	0x44, 0x4d, 0xb1, 0x5e, 0xa1, 0xdb, 0x73, 0x51, 0xcb, 0xf2, 0x0c, 0x55, 0x53, 0x13, 0x99, 0x6d,
	0xad, 0x74, 0x06, 0xa4, 0x28, 0x69, 0xdf, 0xc8, 0xd1, 0x42, 0x9b, 0xbe, 0xfd, 0xe3, 0x3a, 0xba,
	0x69, 0x1a, 0xe2, 0x2f, 0x0e, 0xaa, 0xcd, 0x2c, 0xc1, 0x5b, 0x0b, 0x1a, 0x2c, 0xf4, 0xd3, 0x7d,
	0x7c, 0x89, 0xcc, 0x82, 0xd2, 0x6f, 0x7e, 0xfa, 0xf9, 0xf7, 0xdb, 0xb5, 0x75, 0xec, 0xd1, 0xbc,
	0xa4, 0x3b, 0x7f, 0x3e, 0x67, 0xce, 0xe3, 0xef, 0x0e, 0xaa, 0x97, 0xf6, 0x8e, 0x9f, 0x2c, 0x1b,
	0xb4, 0xcc, 0x43, 0x37, 0xb8, 0x42, 0x85, 0x45, 0x6c, 0x19, 0xc4, 0x4d, 0xbc, 0x41, 0xcb, 0x07,
	0x5e, 0x76, 0x1b, 0x7f, 0x40, 0xd5, 0x62, 0xad, 0x78, 0x63, 0xd9, 0xac, 0x39, 0xff, 0xdc, 0xe6,
	0xff, 0xd2, 0x2c, 0xc7, 0x23, 0xc3, 0xf1, 0x00, 0xaf, 0x2e, 0xe0, 0x28, 0xac, 0x6b, 0xbf, 0x38,
	0x1e, 0x7b, 0xce, 0xc9, 0xd8, 0x73, 0xfe, 0x8c, 0x3d, 0xe7, 0xeb, 0xc4, 0xab, 0x9c, 0x4c, 0xbc,
	0xca, 0xaf, 0x89, 0x57, 0x79, 0x47, 0xce, 0x5d, 0x5d, 0x5e, 0xde, 0x92, 0x5c, 0x1f, 0x41, 0x76,
	0x50, 0xf4, 0x1a, 0x9e, 0xeb, 0x66, 0x2e, 0xb0, 0x53, 0x35, 0x4f, 0xf6, 0xe9, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xfc, 0x25, 0x93, 0xcb, 0x5b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// MaxSupply retrieves total coins of all eras and when mint ended.
	MaxSupply(ctx context.Context, in *QueryMaxSupplyRequest, opts ...grpc.CallOption) (*QueryMaxSupplyResponse, error)
	// InflationRewardCoefficientRate APY rate for staking rewards
	RewardCoefficient(ctx context.Context, in *QueryRewardCoefficientRequest, opts ...grpc.CallOption) (*QueryRewardCoefficientResponse, error)
	// Params retrieves coinomics moudle params.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) MaxSupply(ctx context.Context, in *QueryMaxSupplyRequest, opts ...grpc.CallOption) (*QueryMaxSupplyResponse, error) {
	out := new(QueryMaxSupplyResponse)
	err := c.cc.Invoke(ctx, "/neura.coinomics.v1.Query/MaxSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RewardCoefficient(ctx context.Context, in *QueryRewardCoefficientRequest, opts ...grpc.CallOption) (*QueryRewardCoefficientResponse, error) {
	out := new(QueryRewardCoefficientResponse)
	err := c.cc.Invoke(ctx, "/neura.coinomics.v1.Query/RewardCoefficient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/neura.coinomics.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// MaxSupply retrieves total coins of all eras and when mint ended.
	MaxSupply(context.Context, *QueryMaxSupplyRequest) (*QueryMaxSupplyResponse, error)
	// InflationRewardCoefficientRate APY rate for staking rewards
	RewardCoefficient(context.Context, *QueryRewardCoefficientRequest) (*QueryRewardCoefficientResponse, error)
	// Params retrieves coinomics moudle params.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) MaxSupply(ctx context.Context, req *QueryMaxSupplyRequest) (*QueryMaxSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaxSupply not implemented")
}
func (*UnimplementedQueryServer) RewardCoefficient(ctx context.Context, req *QueryRewardCoefficientRequest) (*QueryRewardCoefficientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RewardCoefficient not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_MaxSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMaxSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MaxSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neura.coinomics.v1.Query/MaxSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MaxSupply(ctx, req.(*QueryMaxSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RewardCoefficient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRewardCoefficientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RewardCoefficient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neura.coinomics.v1.Query/RewardCoefficient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RewardCoefficient(ctx, req.(*QueryRewardCoefficientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neura.coinomics.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "neura.coinomics.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MaxSupply",
			Handler:    _Query_MaxSupply_Handler,
		},
		{
			MethodName: "RewardCoefficient",
			Handler:    _Query_RewardCoefficient_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "neura/coinomics/v1/query.proto",
}

func (m *QueryMaxSupplyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMaxSupplyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMaxSupplyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryMaxSupplyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMaxSupplyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMaxSupplyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxSupply.Size()
		i -= size
		if _, err := m.MaxSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryRewardCoefficientRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRewardCoefficientRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRewardCoefficientRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryRewardCoefficientResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRewardCoefficientResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRewardCoefficientResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RewardCoefficient.Size()
		i -= size
		if _, err := m.RewardCoefficient.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryMaxSupplyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryMaxSupplyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MaxSupply.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryRewardCoefficientRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryRewardCoefficientResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.RewardCoefficient.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryMaxSupplyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryMaxSupplyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMaxSupplyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryMaxSupplyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryMaxSupplyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMaxSupplyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryRewardCoefficientRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRewardCoefficientRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRewardCoefficientRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryRewardCoefficientResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRewardCoefficientResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRewardCoefficientResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardCoefficient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardCoefficient.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
