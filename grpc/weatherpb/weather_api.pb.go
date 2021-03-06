// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gprc/weatherpb/weather_api.proto

package weatherpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetWeatherRequest struct {
	City                 string   `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWeatherRequest) Reset()         { *m = GetWeatherRequest{} }
func (m *GetWeatherRequest) String() string { return proto.CompactTextString(m) }
func (*GetWeatherRequest) ProtoMessage()    {}
func (*GetWeatherRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_80db36882793d3f8, []int{0}
}

func (m *GetWeatherRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWeatherRequest.Unmarshal(m, b)
}
func (m *GetWeatherRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWeatherRequest.Marshal(b, m, deterministic)
}
func (m *GetWeatherRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWeatherRequest.Merge(m, src)
}
func (m *GetWeatherRequest) XXX_Size() int {
	return xxx_messageInfo_GetWeatherRequest.Size(m)
}
func (m *GetWeatherRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWeatherRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWeatherRequest proto.InternalMessageInfo

func (m *GetWeatherRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

type GetWeatherResponse struct {
	Temperature          string   `protobuf:"bytes,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWeatherResponse) Reset()         { *m = GetWeatherResponse{} }
func (m *GetWeatherResponse) String() string { return proto.CompactTextString(m) }
func (*GetWeatherResponse) ProtoMessage()    {}
func (*GetWeatherResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_80db36882793d3f8, []int{1}
}

func (m *GetWeatherResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWeatherResponse.Unmarshal(m, b)
}
func (m *GetWeatherResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWeatherResponse.Marshal(b, m, deterministic)
}
func (m *GetWeatherResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWeatherResponse.Merge(m, src)
}
func (m *GetWeatherResponse) XXX_Size() int {
	return xxx_messageInfo_GetWeatherResponse.Size(m)
}
func (m *GetWeatherResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWeatherResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetWeatherResponse proto.InternalMessageInfo

func (m *GetWeatherResponse) GetTemperature() string {
	if m != nil {
		return m.Temperature
	}
	return ""
}

func init() {
	proto.RegisterType((*GetWeatherRequest)(nil), "weatherpb.GetWeatherRequest")
	proto.RegisterType((*GetWeatherResponse)(nil), "weatherpb.GetWeatherResponse")
}

func init() { proto.RegisterFile("gprc/weatherpb/weather_api.proto", fileDescriptor_80db36882793d3f8) }

var fileDescriptor_80db36882793d3f8 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0x2f, 0x28, 0x4a,
	0xd6, 0x2f, 0x4f, 0x4d, 0x2c, 0xc9, 0x48, 0x2d, 0x2a, 0x48, 0x82, 0xb1, 0xe2, 0x13, 0x0b, 0x32,
	0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0xe1, 0x92, 0x4a, 0xea, 0x5c, 0x82, 0xee, 0xa9,
	0x25, 0xe1, 0x10, 0x7e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x72,
	0x66, 0x49, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0x64, 0xc6, 0x25, 0x84,
	0xac, 0xb0, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x81, 0x8b, 0xbb, 0x24, 0x35, 0xb7, 0x20,
	0xb5, 0x28, 0xb1, 0xa4, 0xb4, 0x28, 0x15, 0xaa, 0x01, 0x59, 0xc8, 0x28, 0x9c, 0x8b, 0x0b, 0xaa,
	0xc9, 0x31, 0xc0, 0x53, 0xc8, 0x93, 0x8b, 0x0b, 0x61, 0x8a, 0x90, 0x8c, 0x1e, 0xdc, 0x21, 0x7a,
	0x18, 0xae, 0x90, 0x92, 0xc5, 0x21, 0x0b, 0xb1, 0x3a, 0x89, 0x0d, 0xec, 0x17, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa1, 0x12, 0x48, 0x36, 0xef, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WeatherAPIClient is the client API for WeatherAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WeatherAPIClient interface {
	GetWeather(ctx context.Context, in *GetWeatherRequest, opts ...grpc.CallOption) (*GetWeatherResponse, error)
}

type weatherAPIClient struct {
	cc *grpc.ClientConn
}

func NewWeatherAPIClient(cc *grpc.ClientConn) WeatherAPIClient {
	return &weatherAPIClient{cc}
}

func (c *weatherAPIClient) GetWeather(ctx context.Context, in *GetWeatherRequest, opts ...grpc.CallOption) (*GetWeatherResponse, error) {
	out := new(GetWeatherResponse)
	err := c.cc.Invoke(ctx, "/weatherpb.WeatherAPI/GetWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherAPIServer is the server API for WeatherAPI service.
type WeatherAPIServer interface {
	GetWeather(context.Context, *GetWeatherRequest) (*GetWeatherResponse, error)
}

// UnimplementedWeatherAPIServer can be embedded to have forward compatible implementations.
type UnimplementedWeatherAPIServer struct {
}

func (*UnimplementedWeatherAPIServer) GetWeather(ctx context.Context, req *GetWeatherRequest) (*GetWeatherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeather not implemented")
}

func RegisterWeatherAPIServer(s *grpc.Server, srv WeatherAPIServer) {
	s.RegisterService(&_WeatherAPI_serviceDesc, srv)
}

func _WeatherAPI_GetWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherAPIServer).GetWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weatherpb.WeatherAPI/GetWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherAPIServer).GetWeather(ctx, req.(*GetWeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WeatherAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "weatherpb.WeatherAPI",
	HandlerType: (*WeatherAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWeather",
			Handler:    _WeatherAPI_GetWeather_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gprc/weatherpb/weather_api.proto",
}
