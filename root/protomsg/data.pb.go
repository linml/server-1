// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/data.proto

/*
Package protomsg is a generated protocol buffer package.

It is generated from these files:
	protobuf/data.proto
	protobuf/fruitmary.proto
	protobuf/hall.proto

It has these top-level messages:
	AccountStorageData
	AccountGameData
	Email
	RoomInfo
	GameInfo
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
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe5, 0xd8, 0x4e, 0x93, 0x2d, 0x26, 0xb0, 0xa0, 0x6a, 0x55, 0x21, 0x64, 0xe5, 0x64,
	0x2e, 0xa9, 0x04, 0x4f, 0xd0, 0x12, 0x3e, 0x22, 0xa5, 0x34, 0x5a, 0xb7, 0x02, 0x71, 0xdb, 0xd8,
	0x93, 0x60, 0x29, 0xde, 0x8d, 0xec, 0x4d, 0xa5, 0x22, 0xde, 0x86, 0xb7, 0xe4, 0x84, 0x66, 0xd6,
	0x5f, 0xa7, 0xcc, 0xef, 0x3f, 0x9e, 0xd9, 0xff, 0xcc, 0x6e, 0xd8, 0xab, 0x63, 0x65, 0xac, 0xd9,
	0x9e, 0x76, 0x57, 0xb9, 0xb2, 0x6a, 0x41, 0xc4, 0x27, 0xf4, 0x53, 0xd6, 0xfb, 0xf9, 0x3f, 0x9f,
	0xf1, 0xeb, 0x2c, 0x33, 0x27, 0x6d, 0x53, 0x6b, 0x2a, 0xb5, 0x87, 0xa5, 0xb2, 0x8a, 0xbf, 0x61,
	0xd3, 0x46, 0x5d, 0xe5, 0xc2, 0x8b, 0xbd, 0x24, 0x92, 0xbd, 0xc0, 0x2f, 0xd9, 0xe4, 0x41, 0x2f,
	0xe1, 0xb1, 0xc8, 0x40, 0x8c, 0x62, 0x2f, 0x99, 0xca, 0x8e, 0xf9, 0x6b, 0x16, 0x6e, 0x7e, 0x19,
	0x0d, 0xc2, 0xa7, 0x84, 0x03, 0x7e, 0xc1, 0xc6, 0xdf, 0xa1, 0xf8, 0x51, 0x68, 0x11, 0x90, 0xdc,
	0x10, 0xe7, 0x2c, 0xf8, 0xa6, 0x4a, 0x10, 0x21, 0xa9, 0x14, 0x73, 0xc1, 0xce, 0xbe, 0x82, 0xca,
	0x1f, 0xe4, 0x5a, 0x8c, 0x49, 0x6e, 0x11, 0x7b, 0xdf, 0x1a, 0x0d, 0x4f, 0xe2, 0x2c, 0xf6, 0x92,
	0x40, 0x3a, 0x40, 0xaf, 0xa9, 0xda, 0x81, 0xcb, 0x4c, 0x28, 0xd3, 0x0b, 0xfc, 0x2d, 0x63, 0xd7,
	0x99, 0x2d, 0x1e, 0xe1, 0xbe, 0x28, 0x41, 0x4c, 0xa9, 0xe1, 0x40, 0xc1, 0x59, 0x1c, 0xad, 0x36,
	0x82, 0xb9, 0x59, 0x5a, 0xc6, 0xda, 0xcf, 0x95, 0xf9, 0x0d, 0x9a, 0x6a, 0xcf, 0xa9, 0xf5, 0x40,
	0xc1, 0xda, 0x54, 0x1d, 0xa0, 0x2e, 0x95, 0x16, 0xcf, 0x62, 0x2f, 0x09, 0x65, 0xc7, 0xce, 0xd5,
	0x01, 0x6a, 0x2a, 0x8d, 0xa8, 0x71, 0x2f, 0xe0, 0x8c, 0xe9, 0x11, 0xb2, 0x42, 0x1d, 0xc4, 0x73,
	0xda, 0x6e, 0x8b, 0xb8, 0xa9, 0xbb, 0xf4, 0xfe, 0xe9, 0x08, 0x62, 0x46, 0x89, 0x86, 0xb0, 0xdf,
	0xda, 0xec, 0x0b, 0x67, 0xe5, 0x45, 0xec, 0x25, 0xbe, 0xec, 0x05, 0x74, 0xba, 0x36, 0x7b, 0x73,
	0xb2, 0x94, 0x7e, 0x49, 0xe9, 0x81, 0x82, 0x9b, 0x93, 0x66, 0x6b, 0xac, 0xe0, 0xd4, 0xd4, 0xc1,
	0xfc, 0x1d, 0x9b, 0x35, 0x97, 0xfa, 0x45, 0x95, 0xee, 0xe2, 0x2f, 0xd8, 0x58, 0x1a, 0x53, 0xae,
	0x96, 0x74, 0xb1, 0x91, 0x6c, 0x68, 0xfe, 0xd7, 0x63, 0xe1, 0xa7, 0x52, 0x15, 0x07, 0xb4, 0x4e,
	0xc1, 0x6a, 0xd9, 0x3c, 0x8c, 0x16, 0xd1, 0x22, 0x85, 0xe4, 0xde, 0x95, 0xf7, 0x02, 0xd6, 0x7d,
	0x34, 0xda, 0x82, 0xb6, 0xcd, 0xd3, 0x68, 0xb1, 0xbf, 0xd6, 0x60, 0x78, 0xad, 0xb8, 0x5c, 0xd0,
	0x39, 0x0d, 0x14, 0xd2, 0x40, 0x1d, 0xa3, 0xcb, 0x55, 0x2d, 0x41, 0xe5, 0xf4, 0x42, 0x22, 0xd9,
	0xd0, 0xfc, 0x0f, 0x9b, 0x90, 0x5f, 0xbd, 0x33, 0x83, 0x49, 0xbc, 0xe1, 0x24, 0xd8, 0xf7, 0xb6,
	0xd0, 0xee, 0xc0, 0x11, 0x1d, 0xd8, 0x31, 0x3e, 0xc7, 0x1b, 0xb0, 0xb5, 0xf0, 0x63, 0x3f, 0x09,
	0x24, 0xc5, 0xa8, 0x59, 0x1c, 0x28, 0xa0, 0x2e, 0x14, 0xa3, 0xe3, 0xbb, 0x2a, 0x87, 0x8a, 0x8c,
	0x45, 0xd2, 0xc1, 0x7c, 0xc3, 0x26, 0xb8, 0x47, 0x3a, 0xfd, 0xd2, 0xc5, 0xb4, 0x0a, 0x77, 0x7e,
	0xc7, 0x3c, 0x61, 0x61, 0x65, 0x4c, 0x59, 0x8b, 0x51, 0xec, 0x27, 0xe7, 0xef, 0xf9, 0xa2, 0xfd,
	0x37, 0x2e, 0x5a, 0xf3, 0xd2, 0x7d, 0x70, 0x33, 0xfb, 0x19, 0x55, 0xc6, 0xd8, 0xab, 0xf6, 0x83,
	0xed, 0x98, 0xa2, 0x0f, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x58, 0x7e, 0x3c, 0xd6, 0x03,
	0x00, 0x00,
}
