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

type Settings_Language int32

const (
	Settings_ENGLISH   Settings_Language = 0
	Settings_SLOVENIAN Settings_Language = 1
)

var Settings_Language_name = map[int32]string{
	0: "ENGLISH",
	1: "SLOVENIAN",
}

var Settings_Language_value = map[string]int32{
	"ENGLISH":   0,
	"SLOVENIAN": 1,
}

func (x Settings_Language) String() string {
	return proto.EnumName(Settings_Language_name, int32(x))
}

func (Settings_Language) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{0, 0}
}

type Settings_PlayerType int32

const (
	Settings_HUMAN    Settings_PlayerType = 0
	Settings_COMPUTER Settings_PlayerType = 1
)

var Settings_PlayerType_name = map[int32]string{
	0: "HUMAN",
	1: "COMPUTER",
}

var Settings_PlayerType_value = map[string]int32{
	"HUMAN":    0,
	"COMPUTER": 1,
}

func (x Settings_PlayerType) String() string {
	return proto.EnumName(Settings_PlayerType_name, int32(x))
}

func (Settings_PlayerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{0, 1}
}

type Response_Type int32

const (
	Response_NOOP         Response_Type = 0
	Response_GAME_UPDATE  Response_Type = 1
	Response_WIFI_UPDATE  Response_Type = 2
	Response_STATE_UPDATE Response_Type = 3
)

var Response_Type_name = map[int32]string{
	0: "NOOP",
	1: "GAME_UPDATE",
	2: "WIFI_UPDATE",
	3: "STATE_UPDATE",
}

var Response_Type_value = map[string]int32{
	"NOOP":         0,
	"GAME_UPDATE":  1,
	"WIFI_UPDATE":  2,
	"STATE_UPDATE": 3,
}

func (x Response_Type) String() string {
	return proto.EnumName(Response_Type_name, int32(x))
}

func (Response_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{1, 0}
}

type Request_Type int32

const (
	Request_NOOP            Request_Type = 0
	Request_START_WIFI_SCAN Request_Type = 1
	Request_STOP_WIFI_SCAN  Request_Type = 2
	Request_CONFIGURE_WIFI  Request_Type = 3
	Request_FORGET_WIFI     Request_Type = 4
	Request_CONNECT_WIFI    Request_Type = 5
	Request_UPDATE_SETTINGS Request_Type = 6
	Request_UNDO_MOVE       Request_Type = 7
	Request_MOVE            Request_Type = 8
	Request_NEW_GAME        Request_Type = 9
	Request_GET_SETTINGS    Request_Type = 10
)

var Request_Type_name = map[int32]string{
	0:  "NOOP",
	1:  "START_WIFI_SCAN",
	2:  "STOP_WIFI_SCAN",
	3:  "CONFIGURE_WIFI",
	4:  "FORGET_WIFI",
	5:  "CONNECT_WIFI",
	6:  "UPDATE_SETTINGS",
	7:  "UNDO_MOVE",
	8:  "MOVE",
	9:  "NEW_GAME",
	10: "GET_SETTINGS",
}

var Request_Type_value = map[string]int32{
	"NOOP":            0,
	"START_WIFI_SCAN": 1,
	"STOP_WIFI_SCAN":  2,
	"CONFIGURE_WIFI":  3,
	"FORGET_WIFI":     4,
	"CONNECT_WIFI":    5,
	"UPDATE_SETTINGS": 6,
	"UNDO_MOVE":       7,
	"MOVE":            8,
	"NEW_GAME":        9,
	"GET_SETTINGS":    10,
}

func (x Request_Type) String() string {
	return proto.EnumName(Request_Type_name, int32(x))
}

func (Request_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{2, 0}
}

type Settings struct {
	Sound                bool                       `protobuf:"varint,1,opt,name=sound,proto3" json:"sound,omitempty"`
	Language             Settings_Language          `protobuf:"varint,2,opt,name=language,proto3,enum=bluetoothpb.Settings_Language" json:"language,omitempty"`
	VoiceRecognition     bool                       `protobuf:"varint,3,opt,name=voice_recognition,json=voiceRecognition,proto3" json:"voice_recognition,omitempty"`
	AutoMove             bool                       `protobuf:"varint,4,opt,name=auto_move,json=autoMove,proto3" json:"auto_move,omitempty"`
	RandomBw             bool                       `protobuf:"varint,5,opt,name=random_bw,json=randomBw,proto3" json:"random_bw,omitempty"`
	Player1              Settings_PlayerType        `protobuf:"varint,6,opt,name=player1,proto3,enum=bluetoothpb.Settings_PlayerType" json:"player1,omitempty"`
	Player2              Settings_PlayerType        `protobuf:"varint,7,opt,name=player2,proto3,enum=bluetoothpb.Settings_PlayerType" json:"player2,omitempty"`
	ComputerSettings     *Settings_ComputerSettings `protobuf:"bytes,8,opt,name=computerSettings,proto3" json:"computerSettings,omitempty"`
	Intro                bool                       `protobuf:"varint,9,opt,name=intro,proto3" json:"intro,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{0}
}

func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetSound() bool {
	if m != nil {
		return m.Sound
	}
	return false
}

func (m *Settings) GetLanguage() Settings_Language {
	if m != nil {
		return m.Language
	}
	return Settings_ENGLISH
}

func (m *Settings) GetVoiceRecognition() bool {
	if m != nil {
		return m.VoiceRecognition
	}
	return false
}

func (m *Settings) GetAutoMove() bool {
	if m != nil {
		return m.AutoMove
	}
	return false
}

func (m *Settings) GetRandomBw() bool {
	if m != nil {
		return m.RandomBw
	}
	return false
}

func (m *Settings) GetPlayer1() Settings_PlayerType {
	if m != nil {
		return m.Player1
	}
	return Settings_HUMAN
}

func (m *Settings) GetPlayer2() Settings_PlayerType {
	if m != nil {
		return m.Player2
	}
	return Settings_HUMAN
}

func (m *Settings) GetComputerSettings() *Settings_ComputerSettings {
	if m != nil {
		return m.ComputerSettings
	}
	return nil
}

func (m *Settings) GetIntro() bool {
	if m != nil {
		return m.Intro
	}
	return false
}

type Settings_ComputerSettings struct {
	TimeLimitMs int32 `protobuf:"varint,1,opt,name=time_limit_ms,json=timeLimitMs,proto3" json:"time_limit_ms,omitempty"`
	// from 0 to 20
	SkillLevel int32 `protobuf:"varint,2,opt,name=skill_level,json=skillLevel,proto3" json:"skill_level,omitempty"`
	// enables elo over skill level
	LimitStrength bool `protobuf:"varint,3,opt,name=limit_strength,json=limitStrength,proto3" json:"limit_strength,omitempty"`
	// from 1350 to 2850
	Elo                  int32    `protobuf:"varint,4,opt,name=elo,proto3" json:"elo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_ComputerSettings) Reset()         { *m = Settings_ComputerSettings{} }
func (m *Settings_ComputerSettings) String() string { return proto.CompactTextString(m) }
func (*Settings_ComputerSettings) ProtoMessage()    {}
func (*Settings_ComputerSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{0, 0}
}

func (m *Settings_ComputerSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_ComputerSettings.Unmarshal(m, b)
}
func (m *Settings_ComputerSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_ComputerSettings.Marshal(b, m, deterministic)
}
func (m *Settings_ComputerSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_ComputerSettings.Merge(m, src)
}
func (m *Settings_ComputerSettings) XXX_Size() int {
	return xxx_messageInfo_Settings_ComputerSettings.Size(m)
}
func (m *Settings_ComputerSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_ComputerSettings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_ComputerSettings proto.InternalMessageInfo

func (m *Settings_ComputerSettings) GetTimeLimitMs() int32 {
	if m != nil {
		return m.TimeLimitMs
	}
	return 0
}

func (m *Settings_ComputerSettings) GetSkillLevel() int32 {
	if m != nil {
		return m.SkillLevel
	}
	return 0
}

func (m *Settings_ComputerSettings) GetLimitStrength() bool {
	if m != nil {
		return m.LimitStrength
	}
	return false
}

func (m *Settings_ComputerSettings) GetElo() int32 {
	if m != nil {
		return m.Elo
	}
	return 0
}

type Response struct {
	Type           Response_Type           `protobuf:"varint,1,opt,name=type,proto3,enum=bluetoothpb.Response_Type" json:"type,omitempty"`
	Networks       []*Response_WifiNetwork `protobuf:"bytes,2,rep,name=networks,proto3" json:"networks,omitempty"`
	Settings       *Settings               `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	GameInProgress bool                    `protobuf:"varint,4,opt,name=gameInProgress,proto3" json:"gameInProgress,omitempty"`
	Moves          []string                `protobuf:"bytes,5,rep,name=moves,proto3" json:"moves,omitempty"`
	// not used for now
	WhiteTurn            bool                 `protobuf:"varint,6,opt,name=whiteTurn,proto3" json:"whiteTurn,omitempty"`
	State                string               `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	ChessBoard           *Response_ChessBoard `protobuf:"bytes,8,opt,name=chess_board,json=chessBoard,proto3" json:"chess_board,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{1}
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

func (m *Response) GetType() Response_Type {
	if m != nil {
		return m.Type
	}
	return Response_NOOP
}

func (m *Response) GetNetworks() []*Response_WifiNetwork {
	if m != nil {
		return m.Networks
	}
	return nil
}

func (m *Response) GetSettings() *Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *Response) GetGameInProgress() bool {
	if m != nil {
		return m.GameInProgress
	}
	return false
}

func (m *Response) GetMoves() []string {
	if m != nil {
		return m.Moves
	}
	return nil
}

func (m *Response) GetWhiteTurn() bool {
	if m != nil {
		return m.WhiteTurn
	}
	return false
}

func (m *Response) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Response) GetChessBoard() *Response_ChessBoard {
	if m != nil {
		return m.ChessBoard
	}
	return nil
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
	return fileDescriptor_6671346f1aa9b484, []int{1, 0}
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

type Response_ChessBoard struct {
	Fen                  string   `protobuf:"bytes,1,opt,name=fen,proto3" json:"fen,omitempty"`
	Rotate               bool     `protobuf:"varint,2,opt,name=rotate,proto3" json:"rotate,omitempty"`
	CanMakeMove          bool     `protobuf:"varint,3,opt,name=canMakeMove,proto3" json:"canMakeMove,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_ChessBoard) Reset()         { *m = Response_ChessBoard{} }
func (m *Response_ChessBoard) String() string { return proto.CompactTextString(m) }
func (*Response_ChessBoard) ProtoMessage()    {}
func (*Response_ChessBoard) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{1, 1}
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

func (m *Response_ChessBoard) GetRotate() bool {
	if m != nil {
		return m.Rotate
	}
	return false
}

func (m *Response_ChessBoard) GetCanMakeMove() bool {
	if m != nil {
		return m.CanMakeMove
	}
	return false
}

type Request struct {
	Type                 Request_Type `protobuf:"varint,1,opt,name=type,proto3,enum=bluetoothpb.Request_Type" json:"type,omitempty"`
	WifiSsid             string       `protobuf:"bytes,2,opt,name=wifi_ssid,json=wifiSsid,proto3" json:"wifi_ssid,omitempty"`
	WifiPsk              string       `protobuf:"bytes,3,opt,name=wifi_psk,json=wifiPsk,proto3" json:"wifi_psk,omitempty"`
	Settings             *Settings    `protobuf:"bytes,4,opt,name=settings,proto3" json:"settings,omitempty"`
	Move                 string       `protobuf:"bytes,5,opt,name=move,proto3" json:"move,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6671346f1aa9b484, []int{2}
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

func (m *Request) GetSettings() *Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *Request) GetMove() string {
	if m != nil {
		return m.Move
	}
	return ""
}

func init() {
	proto.RegisterEnum("bluetoothpb.Settings_Language", Settings_Language_name, Settings_Language_value)
	proto.RegisterEnum("bluetoothpb.Settings_PlayerType", Settings_PlayerType_name, Settings_PlayerType_value)
	proto.RegisterEnum("bluetoothpb.Response_Type", Response_Type_name, Response_Type_value)
	proto.RegisterEnum("bluetoothpb.Request_Type", Request_Type_name, Request_Type_value)
	proto.RegisterType((*Settings)(nil), "bluetoothpb.Settings")
	proto.RegisterType((*Settings_ComputerSettings)(nil), "bluetoothpb.Settings.ComputerSettings")
	proto.RegisterType((*Response)(nil), "bluetoothpb.Response")
	proto.RegisterType((*Response_WifiNetwork)(nil), "bluetoothpb.Response.WifiNetwork")
	proto.RegisterType((*Response_ChessBoard)(nil), "bluetoothpb.Response.ChessBoard")
	proto.RegisterType((*Request)(nil), "bluetoothpb.Request")
}

func init() { proto.RegisterFile("bluetoothpb.proto", fileDescriptor_6671346f1aa9b484) }

var fileDescriptor_6671346f1aa9b484 = []byte{
	// 892 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x8e, 0xe2, 0x46,
	0x10, 0x1e, 0x03, 0x1e, 0xec, 0xf2, 0x0e, 0xeb, 0xed, 0xfc, 0xc8, 0x3b, 0x1b, 0x6d, 0x08, 0xd2,
	0x8e, 0x90, 0xa2, 0x20, 0x2d, 0xb9, 0xad, 0x94, 0x03, 0xcb, 0x78, 0x58, 0x24, 0xb0, 0x51, 0xdb,
	0xec, 0xe4, 0x66, 0x19, 0xe8, 0x61, 0x5a, 0x18, 0x37, 0x71, 0x37, 0xa0, 0x79, 0x82, 0x5c, 0xf3,
	0x1e, 0x39, 0xe4, 0x0d, 0xf2, 0x5c, 0x39, 0x46, 0xdd, 0x36, 0x3f, 0xc3, 0x72, 0x48, 0x6e, 0x55,
	0x5f, 0x7f, 0x55, 0xaa, 0xae, 0xef, 0x73, 0x1b, 0x5e, 0x4d, 0x92, 0x35, 0x11, 0x8c, 0x89, 0xc7,
	0xd5, 0xa4, 0xb5, 0xca, 0x98, 0x60, 0xc8, 0x3a, 0x82, 0x1a, 0xbf, 0xeb, 0x60, 0x04, 0x44, 0x08,
	0x9a, 0xce, 0x39, 0xfa, 0x1a, 0x74, 0xce, 0xd6, 0xe9, 0xcc, 0xd1, 0xea, 0x5a, 0xd3, 0xc0, 0x79,
	0x82, 0x3e, 0x80, 0x91, 0xc4, 0xe9, 0x7c, 0x1d, 0xcf, 0x89, 0x53, 0xaa, 0x6b, 0xcd, 0x5a, 0xfb,
	0x6d, 0xeb, 0xb8, 0xeb, 0xae, 0xbc, 0x35, 0x28, 0x58, 0x78, 0xcf, 0x47, 0x3f, 0xc2, 0xab, 0x0d,
	0xa3, 0x53, 0x12, 0x65, 0x64, 0xca, 0xe6, 0x29, 0x15, 0x94, 0xa5, 0x4e, 0x59, 0x75, 0xb7, 0xd5,
	0x01, 0x3e, 0xe0, 0xe8, 0x0d, 0x98, 0xf1, 0x5a, 0xb0, 0x68, 0xc9, 0x36, 0xc4, 0xa9, 0x28, 0x92,
	0x21, 0x81, 0x21, 0xdb, 0x10, 0x79, 0x98, 0xc5, 0xe9, 0x8c, 0x2d, 0xa3, 0xc9, 0xd6, 0xd1, 0xf3,
	0xc3, 0x1c, 0xf8, 0xb8, 0x45, 0x1f, 0xa0, 0xba, 0x4a, 0xe2, 0x27, 0x92, 0xbd, 0x77, 0x2e, 0xd5,
	0x84, 0xf5, 0xf3, 0x13, 0x8e, 0x14, 0x29, 0x7c, 0x5a, 0x11, 0xbc, 0x2b, 0x38, 0xd4, 0xb6, 0x9d,
	0xea, 0xff, 0xab, 0x6d, 0x23, 0x0c, 0xf6, 0x94, 0x2d, 0x57, 0x6b, 0x41, 0xb2, 0x1d, 0xcf, 0x31,
	0xea, 0x5a, 0xd3, 0x6a, 0xdf, 0x9c, 0x6f, 0xd2, 0x3d, 0x61, 0xe3, 0x2f, 0xea, 0xa5, 0x08, 0x34,
	0x15, 0x19, 0x73, 0xcc, 0x5c, 0x04, 0x95, 0x5c, 0xff, 0xa1, 0x81, 0x7d, 0x5a, 0x8c, 0x1a, 0x70,
	0x25, 0xe8, 0x92, 0x44, 0x09, 0x5d, 0x52, 0x11, 0x2d, 0xb9, 0xd2, 0x4d, 0xc7, 0x96, 0x04, 0x07,
	0x12, 0x1b, 0x72, 0xf4, 0x3d, 0x58, 0x7c, 0x41, 0x93, 0x24, 0x4a, 0xc8, 0x86, 0x24, 0x4a, 0x40,
	0x1d, 0x83, 0x82, 0x06, 0x12, 0x41, 0xef, 0xa0, 0x96, 0xd7, 0x73, 0x91, 0x91, 0x74, 0x2e, 0x1e,
	0x0b, 0x7d, 0xae, 0x14, 0x1a, 0x14, 0x20, 0xb2, 0xa1, 0x4c, 0x12, 0xa6, 0x64, 0xd1, 0xb1, 0x0c,
	0x1b, 0x37, 0x60, 0xec, 0x14, 0x47, 0x16, 0x54, 0x5d, 0xaf, 0x37, 0xe8, 0x07, 0x9f, 0xec, 0x0b,
	0x74, 0x05, 0x66, 0x30, 0xf0, 0x3f, 0xbb, 0x5e, 0xbf, 0xe3, 0xd9, 0x5a, 0xe3, 0x1d, 0xc0, 0x61,
	0x77, 0xc8, 0x04, 0xfd, 0xd3, 0x78, 0xd8, 0xf1, 0xec, 0x0b, 0xf4, 0x02, 0x8c, 0xae, 0x3f, 0x1c,
	0x8d, 0x43, 0x17, 0xdb, 0x5a, 0xe3, 0x2f, 0x1d, 0x0c, 0x4c, 0xf8, 0x8a, 0xa5, 0x9c, 0xa0, 0x16,
	0x54, 0xc4, 0xd3, 0x8a, 0xa8, 0x0b, 0xd5, 0xda, 0xd7, 0xcf, 0x96, 0xb9, 0x23, 0xb5, 0x94, 0x16,
	0x8a, 0x87, 0x7e, 0x01, 0x23, 0x25, 0x62, 0xcb, 0xb2, 0x05, 0x77, 0x4a, 0xf5, 0x72, 0xd3, 0x6a,
	0xff, 0x70, 0xbe, 0xe6, 0x9e, 0x3e, 0x50, 0x2f, 0x67, 0xe2, 0x7d, 0x09, 0x7a, 0x0f, 0x06, 0xdf,
	0xe9, 0x57, 0x56, 0xfa, 0x7d, 0x73, 0x56, 0x3f, 0xbc, 0xa7, 0xa1, 0x1b, 0xa8, 0xcd, 0xe3, 0x25,
	0xe9, 0xa7, 0xa3, 0x8c, 0xcd, 0x33, 0xc2, 0x79, 0xe1, 0xd8, 0x13, 0x54, 0xca, 0x29, 0xfd, 0xcc,
	0x1d, 0xbd, 0x5e, 0x6e, 0x9a, 0x38, 0x4f, 0xd0, 0x77, 0x60, 0x6e, 0x1f, 0xa9, 0x20, 0xe1, 0x3a,
	0x4b, 0x95, 0x65, 0x0d, 0x7c, 0x00, 0xd4, 0x77, 0x28, 0x62, 0x41, 0x94, 0x21, 0x4d, 0x9c, 0x27,
	0xa8, 0x03, 0xd6, 0xf4, 0x91, 0x70, 0x1e, 0x4d, 0x58, 0x9c, 0xcd, 0x0a, 0x9f, 0xd5, 0xcf, 0x5f,
	0xb3, 0x2b, 0x89, 0x1f, 0x25, 0x0f, 0xc3, 0x74, 0x1f, 0x5f, 0xff, 0xa9, 0x81, 0x75, 0xb4, 0x01,
	0x84, 0xa0, 0xc2, 0x39, 0xcd, 0xbf, 0x77, 0x13, 0xab, 0x58, 0x8e, 0x36, 0x65, 0x69, 0x4a, 0xa6,
	0x82, 0xcc, 0x94, 0x5d, 0x0c, 0x7c, 0x00, 0xe4, 0x69, 0xbc, 0x89, 0x69, 0x12, 0x4f, 0x12, 0x52,
	0x18, 0xe5, 0x00, 0xa8, 0xc1, 0xe3, 0x0d, 0x99, 0x15, 0xbb, 0xc8, 0x13, 0xf4, 0x16, 0xa0, 0x68,
	0x40, 0xd3, 0x79, 0xf1, 0xed, 0x1e, 0x21, 0xe8, 0x5b, 0xb8, 0x7c, 0x88, 0x69, 0x42, 0x66, 0xc5,
	0x26, 0x8a, 0xec, 0xfa, 0x57, 0x80, 0xc3, 0x3d, 0xa4, 0x01, 0x1f, 0x48, 0x5a, 0x8c, 0x2a, 0x43,
	0x59, 0x97, 0x31, 0xb5, 0xa7, 0x7c, 0xcc, 0x22, 0x43, 0x75, 0xb0, 0xa6, 0x71, 0x3a, 0x8c, 0x17,
	0x44, 0xbe, 0x1c, 0xc5, 0x94, 0xc7, 0x50, 0xe3, 0x16, 0x2a, 0xca, 0x8c, 0x06, 0x54, 0x3c, 0xdf,
	0x1f, 0xd9, 0x17, 0xe8, 0x25, 0x58, 0xbd, 0xce, 0xd0, 0x8d, 0xc6, 0xa3, 0xdb, 0x4e, 0xe8, 0xda,
	0x9a, 0x04, 0xee, 0xfb, 0x77, 0xfd, 0x1d, 0x50, 0x42, 0x36, 0xbc, 0x08, 0xc2, 0x4e, 0xb8, 0xa7,
	0x94, 0x1b, 0xff, 0x94, 0xa0, 0x8a, 0xc9, 0x6f, 0x6b, 0xc2, 0x05, 0xfa, 0xe9, 0x99, 0x61, 0x5f,
	0x9f, 0xa8, 0xa2, 0x38, 0xc7, 0x7e, 0x7d, 0x03, 0xe6, 0x96, 0x3e, 0xd0, 0x48, 0x6d, 0xbf, 0xa4,
	0xae, 0x64, 0x48, 0x20, 0x90, 0x0a, 0xbc, 0x06, 0x15, 0x47, 0x2b, 0xbe, 0x50, 0xc3, 0x9b, 0xb8,
	0x2a, 0xf3, 0x11, 0x5f, 0x3c, 0x33, 0x6a, 0xe5, 0xbf, 0x19, 0x15, 0x41, 0x45, 0x3d, 0xa8, 0x7a,
	0xae, 0xb1, 0x8c, 0x1b, 0x7f, 0x6b, 0x5f, 0x2c, 0xe0, 0x2b, 0x78, 0x19, 0x84, 0x1d, 0x1c, 0x46,
	0xea, 0xd6, 0x41, 0x57, 0x7e, 0xba, 0x08, 0x41, 0x2d, 0x08, 0xfd, 0xd1, 0x11, 0x56, 0x92, 0x58,
	0xd7, 0xf7, 0xee, 0xfa, 0xbd, 0x31, 0x76, 0xd5, 0x81, 0x5d, 0x96, 0xcb, 0xba, 0xf3, 0x71, 0xcf,
	0xcd, 0xab, 0xed, 0x8a, 0x5c, 0x56, 0xd7, 0xf7, 0x3c, 0xb7, 0x5b, 0x20, 0xba, 0xec, 0x9f, 0x2f,
	0x2e, 0x0a, 0xdc, 0x30, 0xec, 0x7b, 0xbd, 0xc0, 0xbe, 0x94, 0x2f, 0xc5, 0xd8, 0xbb, 0xf5, 0xa3,
	0xa1, 0xff, 0xd9, 0xb5, 0xab, 0x72, 0x1a, 0x15, 0x19, 0xf2, 0x69, 0xf0, 0xdc, 0xfb, 0x48, 0x4a,
	0x62, 0x9b, 0xb2, 0x9b, 0xec, 0xbd, 0x2f, 0x84, 0xc9, 0xa5, 0xfa, 0x95, 0xfd, 0xfc, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x4f, 0xf6, 0x49, 0xfb, 0xdf, 0x06, 0x00, 0x00,
}
