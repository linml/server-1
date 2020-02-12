// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/data.proto

/*
Package protomsg is a generated protocol buffer package.

It is generated from these files:
	protobuf/data.proto
	protobuf/dfdc.proto
	protobuf/fruitmary.proto
	protobuf/hall.proto

It has these top-level messages:
	AccountStorageData
	AccountGameData
	Email
	RoomInfo
	GameInfo
	ENTER_GAME_DFDC_REQ
	ENTER_GAME_DFDC_RES
	LEAVE_GAME_DFDC_REQ
	LEAVE_GAME_DFDC_RES
	START_DFDC_REQ
	START_DFDC_RES
	DFDCPosition
	UPDATE_DFDC_BONUS
	PLAYERS_DFDC_LIST_RES
	ENTER_GAME_FRUITMARY_REQ
	ENTER_GAME_FRUITMARY_RES
	LEAVE_GAME_FRUITMARY_REQ
	LEAVE_GAME_FRUITMARY_RES
	START_MARY_REQ
	START_MARY_RES
	FRUITMARYPosition
	FRUITMARY_Result
	UPDATE_MARY_BONUS
	START_MARY2_REQ
	START_MARY2_RES
	Mary2_Result
	NEXT_MARY_RESULT
	PLAYERS_LIST_RES
	SYNC_SERVER_TIME
	KICK_OUT_HALL
	LOGIN_HALL_REQ
	LOGIN_HALL_RES
	SAFEMONEY_OPERATE_REQ
	SAFEMONEY_OPERATE_RES
	BIND_PHONE_REQ
	BIND_PHONE_RES
	ENTER_ROOM_REQ
	ENTER_ROOM_RES
	EMAILS_REQ
	EMAILS_RES
	EMAIL_READ_REQ
	EMAIL_READ_RES
	EMAIL_REWARD_REQ
	EMAIL_REWARD_RES
	EMAIL_DEL_REQ
	EMAIL_DEL_RES
	UPDATE_MONEY
	EMAIL_NEW
	BROADCAST_MSG
	UPDATE_ROOMLIST
*/
package protomsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 玩家基础信息
type AccountStorageData struct {
	AccountId  uint32 `protobuf:"varint,1,opt,name=AccountId" json:"AccountId,omitempty"`
	UnDevice   string `protobuf:"bytes,2,opt,name=UnDevice" json:"UnDevice,omitempty"`
	Phone      string `protobuf:"bytes,3,opt,name=Phone" json:"Phone,omitempty"`
	WeiXin     string `protobuf:"bytes,4,opt,name=WeiXin" json:"WeiXin,omitempty"`
	Name       string `protobuf:"bytes,5,opt,name=Name" json:"Name,omitempty"`
	HeadURL    string `protobuf:"bytes,6,opt,name=HeadURL" json:"HeadURL,omitempty"`
	Money      uint64 `protobuf:"varint,7,opt,name=Money" json:"Money,omitempty"`
	SafeMoney  uint64 `protobuf:"varint,8,opt,name=SafeMoney" json:"SafeMoney,omitempty"`
	ActiveTime string `protobuf:"bytes,9,opt,name=ActiveTime" json:"ActiveTime,omitempty"`
	ActiveIP   string `protobuf:"bytes,10,opt,name=ActiveIP" json:"ActiveIP,omitempty"`
	FrozenTime uint64 `protobuf:"varint,11,opt,name=FrozenTime" json:"FrozenTime,omitempty"`
	Salesman   int32  `protobuf:"varint,12,opt,name=Salesman" json:"Salesman,omitempty"`
	SalesTime  string `protobuf:"bytes,13,opt,name=SalesTime" json:"SalesTime,omitempty"`
	Special    uint32 `protobuf:"varint,14,opt,name=Special" json:"Special,omitempty"`
	OSType     uint32 `protobuf:"varint,15,opt,name=OSType" json:"OSType,omitempty"`
	LoginTime  int64  `protobuf:"varint,16,opt,name=LoginTime" json:"LoginTime,omitempty"`
	LogoutTime int64  `protobuf:"varint,17,opt,name=LogoutTime" json:"LogoutTime,omitempty"`
	Robot      uint32 `protobuf:"varint,18,opt,name=Robot" json:"Robot,omitempty"`
	Kill       int32  `protobuf:"varint,19,opt,name=Kill" json:"Kill,omitempty"`
}

func (m *AccountStorageData) Reset()                    { *m = AccountStorageData{} }
func (m *AccountStorageData) String() string            { return proto.CompactTextString(m) }
func (*AccountStorageData) ProtoMessage()               {}
func (*AccountStorageData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AccountStorageData) GetAccountId() uint32 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *AccountStorageData) GetUnDevice() string {
	if m != nil {
		return m.UnDevice
	}
	return ""
}

func (m *AccountStorageData) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *AccountStorageData) GetWeiXin() string {
	if m != nil {
		return m.WeiXin
	}
	return ""
}

func (m *AccountStorageData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountStorageData) GetHeadURL() string {
	if m != nil {
		return m.HeadURL
	}
	return ""
}

func (m *AccountStorageData) GetMoney() uint64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *AccountStorageData) GetSafeMoney() uint64 {
	if m != nil {
		return m.SafeMoney
	}
	return 0
}

func (m *AccountStorageData) GetActiveTime() string {
	if m != nil {
		return m.ActiveTime
	}
	return ""
}

func (m *AccountStorageData) GetActiveIP() string {
	if m != nil {
		return m.ActiveIP
	}
	return ""
}

func (m *AccountStorageData) GetFrozenTime() uint64 {
	if m != nil {
		return m.FrozenTime
	}
	return 0
}

func (m *AccountStorageData) GetSalesman() int32 {
	if m != nil {
		return m.Salesman
	}
	return 0
}

func (m *AccountStorageData) GetSalesTime() string {
	if m != nil {
		return m.SalesTime
	}
	return ""
}

func (m *AccountStorageData) GetSpecial() uint32 {
	if m != nil {
		return m.Special
	}
	return 0
}

func (m *AccountStorageData) GetOSType() uint32 {
	if m != nil {
		return m.OSType
	}
	return 0
}

func (m *AccountStorageData) GetLoginTime() int64 {
	if m != nil {
		return m.LoginTime
	}
	return 0
}

func (m *AccountStorageData) GetLogoutTime() int64 {
	if m != nil {
		return m.LogoutTime
	}
	return 0
}

func (m *AccountStorageData) GetRobot() uint32 {
	if m != nil {
		return m.Robot
	}
	return 0
}

func (m *AccountStorageData) GetKill() int32 {
	if m != nil {
		return m.Kill
	}
	return 0
}

// 玩家游戏信息
type AccountGameData struct {
	RoomID uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *AccountGameData) Reset()                    { *m = AccountGameData{} }
func (m *AccountGameData) String() string            { return proto.CompactTextString(m) }
func (*AccountGameData) ProtoMessage()               {}
func (*AccountGameData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AccountGameData) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type Email struct {
	EmailID   uint32 `protobuf:"varint,1,opt,name=EmailID" json:"EmailID,omitempty"`
	EmailType uint32 `protobuf:"varint,2,opt,name=EmailType" json:"EmailType,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=Content" json:"Content,omitempty"`
	Money     uint64 `protobuf:"varint,4,opt,name=Money" json:"Money,omitempty"`
	SendTime  int64  `protobuf:"varint,5,opt,name=SendTime" json:"SendTime,omitempty"`
	IsRead    uint32 `protobuf:"varint,6,opt,name=IsRead" json:"IsRead,omitempty"`
}

func (m *Email) Reset()                    { *m = Email{} }
func (m *Email) String() string            { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()               {}
func (*Email) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Email) GetEmailID() uint32 {
	if m != nil {
		return m.EmailID
	}
	return 0
}

func (m *Email) GetEmailType() uint32 {
	if m != nil {
		return m.EmailType
	}
	return 0
}

func (m *Email) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Email) GetMoney() uint64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *Email) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func (m *Email) GetIsRead() uint32 {
	if m != nil {
		return m.IsRead
	}
	return 0
}

type RoomInfo struct {
	RoomID   uint32   `protobuf:"varint,1,opt,name=RoomID" json:"RoomID,omitempty"`
	MinMoney uint64   `protobuf:"varint,2,opt,name=MinMoney" json:"MinMoney,omitempty"`
	Bets     []uint64 `protobuf:"varint,3,rep,packed,name=Bets" json:"Bets,omitempty"`
	Type     uint32   `protobuf:"varint,4,opt,name=type" json:"type,omitempty"`
	Order    uint32   `protobuf:"varint,5,opt,name=Order" json:"Order,omitempty"`
}

func (m *RoomInfo) Reset()                    { *m = RoomInfo{} }
func (m *RoomInfo) String() string            { return proto.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()               {}
func (*RoomInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RoomInfo) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *RoomInfo) GetMinMoney() uint64 {
	if m != nil {
		return m.MinMoney
	}
	return 0
}

func (m *RoomInfo) GetBets() []uint64 {
	if m != nil {
		return m.Bets
	}
	return nil
}

func (m *RoomInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *RoomInfo) GetOrder() uint32 {
	if m != nil {
		return m.Order
	}
	return 0
}

type GameInfo struct {
	GameType uint32      `protobuf:"varint,1,opt,name=GameType" json:"GameType,omitempty"`
	Rooms    []*RoomInfo `protobuf:"bytes,2,rep,name=rooms" json:"rooms,omitempty"`
}

func (m *GameInfo) Reset()                    { *m = GameInfo{} }
func (m *GameInfo) String() string            { return proto.CompactTextString(m) }
func (*GameInfo) ProtoMessage()               {}
func (*GameInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GameInfo) GetGameType() uint32 {
	if m != nil {
		return m.GameType
	}
	return 0
}

func (m *GameInfo) GetRooms() []*RoomInfo {
	if m != nil {
		return m.Rooms
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountStorageData)(nil), "protomsg.AccountStorageData")
	proto.RegisterType((*AccountGameData)(nil), "protomsg.AccountGameData")
	proto.RegisterType((*Email)(nil), "protomsg.Email")
	proto.RegisterType((*RoomInfo)(nil), "protomsg.RoomInfo")
	proto.RegisterType((*GameInfo)(nil), "protomsg.GameInfo")
}

func init() { proto.RegisterFile("protobuf/data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x93, 0x6d, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x95, 0x35, 0xe9, 0x5a, 0x8f, 0x50, 0xf0, 0xd0, 0x64, 0x4d, 0x08, 0x45, 0x7d, 0x15,
	0xde, 0x74, 0x12, 0x7c, 0x82, 0x8d, 0xf2, 0x50, 0xd1, 0xb1, 0xca, 0xdd, 0x04, 0xe2, 0x9d, 0xdb,
	0x5c, 0x8b, 0xa5, 0xc4, 0xae, 0x12, 0x77, 0xd2, 0x10, 0x9f, 0x81, 0x2f, 0xc1, 0x17, 0x45, 0x77,
	0xce, 0xd3, 0xab, 0xde, 0xef, 0x7f, 0xb9, 0xf3, 0xff, 0xce, 0x2e, 0x3b, 0x3f, 0x94, 0xd6, 0xd9,
	0xcd, 0x71, 0x77, 0x95, 0x29, 0xa7, 0x66, 0x44, 0x7c, 0x44, 0x3f, 0x45, 0xb5, 0x9f, 0xfe, 0x0d,
	0x19, 0xbf, 0xde, 0x6e, 0xed, 0xd1, 0xb8, 0xb5, 0xb3, 0xa5, 0xda, 0xc3, 0x5c, 0x39, 0xc5, 0x5f,
	0xb3, 0x71, 0xad, 0x2e, 0x32, 0x11, 0x24, 0x41, 0x1a, 0xcb, 0x4e, 0xe0, 0x97, 0x6c, 0xf4, 0x60,
	0xe6, 0xf0, 0xa8, 0xb7, 0x20, 0x4e, 0x92, 0x20, 0x1d, 0xcb, 0x96, 0xf9, 0x2b, 0x16, 0xad, 0x7e,
	0x59, 0x03, 0x62, 0x40, 0x09, 0x0f, 0xfc, 0x82, 0x0d, 0xbf, 0x83, 0xfe, 0xa1, 0x8d, 0x08, 0x49,
	0xae, 0x89, 0x73, 0x16, 0x7e, 0x53, 0x05, 0x88, 0x88, 0x54, 0x8a, 0xb9, 0x60, 0xa7, 0x5f, 0x40,
	0x65, 0x0f, 0x72, 0x29, 0x86, 0x24, 0x37, 0x88, 0xbd, 0x6f, 0xad, 0x81, 0x27, 0x71, 0x9a, 0x04,
	0x69, 0x28, 0x3d, 0xa0, 0xd7, 0xb5, 0xda, 0x81, 0xcf, 0x8c, 0x28, 0xd3, 0x09, 0xfc, 0x0d, 0x63,
	0xd7, 0x5b, 0xa7, 0x1f, 0xe1, 0x5e, 0x17, 0x20, 0xc6, 0xd4, 0xb0, 0xa7, 0xe0, 0x2c, 0x9e, 0x16,
	0x2b, 0xc1, 0xfc, 0x2c, 0x0d, 0x63, 0xed, 0xa7, 0xd2, 0xfe, 0x06, 0x43, 0xb5, 0x67, 0xd4, 0xba,
	0xa7, 0x60, 0xed, 0x5a, 0xe5, 0x50, 0x15, 0xca, 0x88, 0x67, 0x49, 0x90, 0x46, 0xb2, 0x65, 0xef,
	0x2a, 0x87, 0x8a, 0x4a, 0x63, 0x6a, 0xdc, 0x09, 0x38, 0xe3, 0xfa, 0x00, 0x5b, 0xad, 0x72, 0xf1,
	0x9c, 0xb6, 0xdb, 0x20, 0x6e, 0xea, 0x6e, 0x7d, 0xff, 0x74, 0x00, 0x31, 0xa1, 0x44, 0x4d, 0xd8,
	0x6f, 0x69, 0xf7, 0xda, 0x5b, 0x79, 0x91, 0x04, 0xe9, 0x40, 0x76, 0x02, 0x3a, 0x5d, 0xda, 0xbd,
	0x3d, 0x3a, 0x4a, 0xbf, 0xa4, 0x74, 0x4f, 0xc1, 0xcd, 0x49, 0xbb, 0xb1, 0x4e, 0x70, 0x6a, 0xea,
	0x01, 0xb7, 0xff, 0x55, 0xe7, 0xb9, 0x38, 0x27, 0xef, 0x14, 0x4f, 0xdf, 0xb2, 0x49, 0x7d, 0xd1,
	0x9f, 0x55, 0xe1, 0x1f, 0xc3, 0x05, 0x1b, 0x4a, 0x6b, 0x8b, 0xc5, 0x9c, 0x2e, 0x3b, 0x96, 0x35,
	0x4d, 0xff, 0x05, 0x2c, 0xfa, 0x58, 0x28, 0x9d, 0xe3, 0x38, 0x14, 0x2c, 0xe6, 0xf5, 0x63, 0x69,
	0x10, 0x6d, 0x53, 0x48, 0x13, 0xf9, 0xf2, 0x4e, 0xc0, 0xba, 0x0f, 0xd6, 0x38, 0x30, 0xae, 0x7e,
	0x2e, 0x0d, 0x76, 0x57, 0x1d, 0xf6, 0xaf, 0x1a, 0x17, 0x0e, 0x26, 0xa3, 0x21, 0x23, 0x1a, 0xb2,
	0x65, 0x74, 0xb9, 0xa8, 0x24, 0xa8, 0x8c, 0x5e, 0x4d, 0x2c, 0x6b, 0x9a, 0xfe, 0x61, 0x23, 0xf2,
	0x6b, 0x76, 0xb6, 0x37, 0x49, 0xd0, 0x9f, 0x04, 0xfb, 0xde, 0x6a, 0xe3, 0x0f, 0x3c, 0xa1, 0x03,
	0x5b, 0xc6, 0x25, 0xdd, 0x80, 0xab, 0xc4, 0x20, 0x19, 0xa4, 0xa1, 0xa4, 0x18, 0x35, 0x87, 0x03,
	0x85, 0xd4, 0x85, 0x62, 0x74, 0x7c, 0x57, 0x66, 0x50, 0x92, 0xb1, 0x58, 0x7a, 0x98, 0xae, 0xd8,
	0x08, 0xf7, 0x48, 0xa7, 0x5f, 0xfa, 0x98, 0x56, 0xe1, 0xcf, 0x6f, 0x99, 0xa7, 0x2c, 0x2a, 0xad,
	0x2d, 0x2a, 0x71, 0x92, 0x0c, 0xd2, 0xb3, 0x77, 0x7c, 0xd6, 0xfc, 0x43, 0x67, 0x8d, 0x79, 0xe9,
	0x3f, 0xb8, 0x99, 0xfc, 0x8c, 0x4b, 0x6b, 0xdd, 0x55, 0xf3, 0xc1, 0x66, 0x48, 0xd1, 0xfb, 0xff,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x05, 0x99, 0x43, 0xf9, 0xea, 0x03, 0x00, 0x00,
}
