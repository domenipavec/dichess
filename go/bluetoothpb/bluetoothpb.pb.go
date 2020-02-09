// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bluetoothpb.proto

package bluetoothpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Request_Type int32

const (
	Request_NOOP            Request_Type = 0
	Request_START_WIFI_SCAN Request_Type = 1
	Request_STOP_WIFI_SCAN  Request_Type = 2
	Request_CONFIGURE_WIFI  Request_Type = 3
	Request_FORGET_WIFI     Request_Type = 4
	Request_CONNECT_WIFI    Request_Type = 5
)

var Request_Type_name = map[int32]string{
	0: "NOOP",
	1: "START_WIFI_SCAN",
	2: "STOP_WIFI_SCAN",
	3: "CONFIGURE_WIFI",
	4: "FORGET_WIFI",
	5: "CONNECT_WIFI",
}

var Request_Type_value = map[string]int32{
	"NOOP":            0,
	"START_WIFI_SCAN": 1,
	"STOP_WIFI_SCAN":  2,
	"CONFIGURE_WIFI":  3,
	"FORGET_WIFI":     4,
	"CONNECT_WIFI":    5,
}

func (x Request_Type) String() string {
	return proto.EnumName(Request_Type_name, int32(x))
}

func (Request_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{1, 0}
}

type Response struct {
	ChessBoard           *Response_ChessBoard    `protobuf:"bytes,1,opt,name=chess_board,json=chessBoard,proto3" json:"chess_board,omitempty"`
	Networks             []*Response_WifiNetwork `protobuf:"bytes,2,rep,name=networks,proto3" json:"networks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{0}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetChessBoard() *Response_ChessBoard {
	if m != nil {
		return m.ChessBoard
	}
	return nil
}

func (m *Response) GetNetworks() []*Response_WifiNetwork {
	if m != nil {
		return m.Networks
	}
	return nil
}

type Response_ChessBoard struct {
	Fen                  string   `protobuf:"bytes,1,opt,name=fen,proto3" json:"fen,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_ChessBoard) Reset()         { *m = Response_ChessBoard{} }
func (m *Response_ChessBoard) String() string { return proto.CompactTextString(m) }
func (*Response_ChessBoard) ProtoMessage()    {}
func (*Response_ChessBoard) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{0, 0}
}

func (m *Response_ChessBoard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_ChessBoard.Unmarshal(m, b)
}
func (m *Response_ChessBoard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_ChessBoard.Marshal(b, m, deterministic)
}
func (m *Response_ChessBoard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_ChessBoard.Merge(m, src)
}
func (m *Response_ChessBoard) XXX_Size() int {
	return xxx_messageInfo_Response_ChessBoard.Size(m)
}
func (m *Response_ChessBoard) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_ChessBoard.DiscardUnknown(m)
}

var xxx_messageInfo_Response_ChessBoard proto.InternalMessageInfo

func (m *Response_ChessBoard) GetFen() string {
	if m != nil {
		return m.Fen
	}
	return ""
}

type Response_WifiNetwork struct {
	Ssid                 string   `protobuf:"bytes,1,opt,name=ssid,proto3" json:"ssid,omitempty"`
	Connected            bool     `protobuf:"varint,2,opt,name=connected,proto3" json:"connected,omitempty"`
	Available            bool     `protobuf:"varint,3,opt,name=available,proto3" json:"available,omitempty"`
	Saved                bool     `protobuf:"varint,4,opt,name=saved,proto3" json:"saved,omitempty"`
	Connecting           bool     `protobuf:"varint,5,opt,name=connecting,proto3" json:"connecting,omitempty"`
	Failed               bool     `protobuf:"varint,6,opt,name=failed,proto3" json:"failed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_WifiNetwork) Reset()         { *m = Response_WifiNetwork{} }
func (m *Response_WifiNetwork) String() string { return proto.CompactTextString(m) }
func (*Response_WifiNetwork) ProtoMessage()    {}
func (*Response_WifiNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{0, 1}
}

func (m *Response_WifiNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_WifiNetwork.Unmarshal(m, b)
}
func (m *Response_WifiNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_WifiNetwork.Marshal(b, m, deterministic)
}
func (m *Response_WifiNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_WifiNetwork.Merge(m, src)
}
func (m *Response_WifiNetwork) XXX_Size() int {
	return xxx_messageInfo_Response_WifiNetwork.Size(m)
}
func (m *Response_WifiNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_WifiNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_Response_WifiNetwork proto.InternalMessageInfo

func (m *Response_WifiNetwork) GetSsid() string {
	if m != nil {
		return m.Ssid
	}
	return ""
}

func (m *Response_WifiNetwork) GetConnected() bool {
	if m != nil {
		return m.Connected
	}
	return false
}

func (m *Response_WifiNetwork) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Response_WifiNetwork) GetSaved() bool {
	if m != nil {
		return m.Saved
	}
	return false
}

func (m *Response_WifiNetwork) GetConnecting() bool {
	if m != nil {
		return m.Connecting
	}
	return false
}

func (m *Response_WifiNetwork) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

type Request struct {
	Type                 Request_Type `protobuf:"varint,1,opt,name=type,proto3,enum=bluetoothpb.Request_Type" json:"type,omitempty"`
	WifiSsid             string       `protobuf:"bytes,2,opt,name=wifi_ssid,json=wifiSsid,proto3" json:"wifi_ssid,omitempty"`
	WifiPsk              string       `protobuf:"bytes,3,opt,name=wifi_psk,json=wifiPsk,proto3" json:"wifi_psk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() Request_Type {
	if m != nil {
		return m.Type
	}
	return Request_NOOP
}

func (m *Request) GetWifiSsid() string {
	if m != nil {
		return m.WifiSsid
	}
	return ""
}

func (m *Request) GetWifiPsk() string {
	if m != nil {
		return m.WifiPsk
	}
	return ""
}

func init() {
	proto.RegisterEnum("bluetoothpb.Request_Type", Request_Type_name, Request_Type_value)
	proto.RegisterType((*Response)(nil), "bluetoothpb.Response")
	proto.RegisterType((*Response_ChessBoard)(nil), "bluetoothpb.Response.ChessBoard")
	proto.RegisterType((*Response_WifiNetwork)(nil), "bluetoothpb.Response.WifiNetwork")
	proto.RegisterType((*Request)(nil), "bluetoothpb.Request")
}

func init() { proto.RegisterFile("bluetoothpb.proto", fileDescriptor_6671346f1aa9b484) }

var fileDescriptor_6671346f1aa9b484 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x6e, 0x9b, 0x40,
	0x14, 0x85, 0xcb, 0x8f, 0x1d, 0xb8, 0x54, 0x09, 0xbd, 0xad, 0x2a, 0x92, 0x56, 0x11, 0xf5, 0xca,
	0x9b, 0xb2, 0x70, 0xd7, 0x5d, 0xb8, 0xc8, 0x8e, 0xbc, 0x01, 0x6b, 0xa0, 0xca, 0x12, 0xf1, 0x33,
	0x34, 0x23, 0x23, 0x86, 0x7a, 0x48, 0xa2, 0x3c, 0x53, 0xdf, 0xa8, 0xcb, 0x3e, 0x49, 0x35, 0x03,
	0x75, 0x50, 0x95, 0xdd, 0xbd, 0xdf, 0x39, 0xf7, 0x68, 0x0e, 0x02, 0xde, 0x14, 0xcd, 0x3d, 0xed,
	0x39, 0xef, 0xef, 0xba, 0x22, 0xe8, 0x8e, 0xbc, 0xe7, 0xe8, 0x4c, 0xd0, 0xe2, 0xb7, 0x0e, 0x16,
	0xa1, 0xa2, 0xe3, 0xad, 0xa0, 0xb8, 0x06, 0xa7, 0xbc, 0xa3, 0x42, 0x64, 0x05, 0xcf, 0x8f, 0x95,
	0xa7, 0xf9, 0xda, 0xd2, 0x59, 0xf9, 0xc1, 0x34, 0xe2, 0x9f, 0x37, 0x08, 0xa5, 0xf1, 0x9b, 0xf4,
	0x11, 0x28, 0x4f, 0x33, 0x7e, 0x05, 0xab, 0xa5, 0xfd, 0x23, 0x3f, 0x1e, 0x84, 0xa7, 0xfb, 0xc6,
	0xd2, 0x59, 0x7d, 0x7a, 0xf9, 0xfe, 0x96, 0xd5, 0x2c, 0x1a, 0x9c, 0xe4, 0x74, 0x72, 0x75, 0x0d,
	0xf0, 0x1c, 0x8c, 0x2e, 0x18, 0x35, 0x6d, 0xd5, 0x3b, 0x6c, 0x22, 0xc7, 0xab, 0x5f, 0x1a, 0x38,
	0x93, 0x4b, 0x44, 0x30, 0x85, 0x60, 0xd5, 0x68, 0x51, 0x33, 0x7e, 0x04, 0xbb, 0xe4, 0x6d, 0x4b,
	0xcb, 0x9e, 0x56, 0x9e, 0xee, 0x6b, 0x4b, 0x8b, 0x3c, 0x03, 0xa9, 0xe6, 0x0f, 0x39, 0x6b, 0xf2,
	0xa2, 0xa1, 0x9e, 0x31, 0xa8, 0x27, 0x80, 0xef, 0x60, 0x26, 0xf2, 0x07, 0x5a, 0x79, 0xa6, 0x52,
	0x86, 0x05, 0xaf, 0x01, 0xc6, 0x00, 0xd6, 0xfe, 0xf0, 0x66, 0x4a, 0x9a, 0x10, 0x7c, 0x0f, 0xf3,
	0x3a, 0x67, 0x0d, 0xad, 0xbc, 0xb9, 0xd2, 0xc6, 0x6d, 0xf1, 0x47, 0x83, 0x33, 0x42, 0x7f, 0xde,
	0x53, 0xd1, 0xe3, 0x67, 0x30, 0xfb, 0xa7, 0x8e, 0xaa, 0x97, 0x9e, 0xaf, 0x2e, 0xff, 0xfb, 0x28,
	0xca, 0x13, 0xa4, 0x4f, 0x1d, 0x25, 0xca, 0x86, 0x1f, 0xc0, 0x7e, 0x64, 0x35, 0xcb, 0x54, 0x3b,
	0x5d, 0xb5, 0xb3, 0x24, 0x48, 0x64, 0xc3, 0x4b, 0x50, 0x73, 0xd6, 0x89, 0x83, 0xaa, 0x60, 0x93,
	0x33, 0xb9, 0xef, 0xc5, 0x61, 0xd1, 0x81, 0x29, 0x53, 0xd0, 0x02, 0x33, 0x8a, 0xe3, 0xbd, 0xfb,
	0x0a, 0xdf, 0xc2, 0x45, 0x92, 0xae, 0x49, 0x9a, 0xdd, 0xee, 0xb6, 0xbb, 0x2c, 0x09, 0xd7, 0x91,
	0xab, 0x21, 0xc2, 0x79, 0x92, 0xc6, 0xfb, 0x09, 0xd3, 0x25, 0x0b, 0xe3, 0x68, 0xbb, 0xbb, 0xf9,
	0x4e, 0x36, 0x4a, 0x70, 0x0d, 0xbc, 0x00, 0x67, 0x1b, 0x93, 0x9b, 0xcd, 0x70, 0xed, 0x9a, 0xe8,
	0xc2, 0xeb, 0x30, 0x8e, 0xa2, 0x4d, 0x38, 0x92, 0x59, 0x31, 0x57, 0x7f, 0xd5, 0x97, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x56, 0x4b, 0x85, 0x34, 0x6a, 0x02, 0x00, 0x00,
}
