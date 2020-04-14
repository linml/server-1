// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/hongbao.proto

package protomsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 网络消息
type HBMSG int32

const (
	HBMSG_UNKNOW_HB              HBMSG = 0
	HBMSG_CS_ENTER_GAME_HB_REQ   HBMSG = 19001
	HBMSG_SC_ENTER_GAME_HB_RES   HBMSG = 19002
	HBMSG_CS_LEAVE_GAME_HB_REQ   HBMSG = 19003
	HBMSG_SC_LEAVE_GAME_HB_RES   HBMSG = 19004
	HBMSG_CS_ASSIGN_HB_REQ       HBMSG = 19008
	HBMSG_SC_ASSIGN_HB_RES       HBMSG = 19009
	HBMSG_SC_BROADCAST_NEW_HB    HBMSG = 19010
	HBMSG_CS_GRAB_HB_REQ         HBMSG = 19015
	HBMSG_SC_GRAB_HB_RES         HBMSG = 19016
	HBMSG_SC_BROADCAST_UPDATE_HB HBMSG = 19017
	HBMSG_CS_PLAYERS_HB_LIST_REQ HBMSG = 19020
	HBMSG_SC_PLAYERS_HB_LIST_RES HBMSG = 19021
)

var HBMSG_name = map[int32]string{
	0:     "UNKNOW_HB",
	19001: "CS_ENTER_GAME_HB_REQ",
	19002: "SC_ENTER_GAME_HB_RES",
	19003: "CS_LEAVE_GAME_HB_REQ",
	19004: "SC_LEAVE_GAME_HB_RES",
	19008: "CS_ASSIGN_HB_REQ",
	19009: "SC_ASSIGN_HB_RES",
	19010: "SC_BROADCAST_NEW_HB",
	19015: "CS_GRAB_HB_REQ",
	19016: "SC_GRAB_HB_RES",
	19017: "SC_BROADCAST_UPDATE_HB",
	19020: "CS_PLAYERS_HB_LIST_REQ",
	19021: "SC_PLAYERS_HB_LIST_RES",
}
var HBMSG_value = map[string]int32{
	"UNKNOW_HB":              0,
	"CS_ENTER_GAME_HB_REQ":   19001,
	"SC_ENTER_GAME_HB_RES":   19002,
	"CS_LEAVE_GAME_HB_REQ":   19003,
	"SC_LEAVE_GAME_HB_RES":   19004,
	"CS_ASSIGN_HB_REQ":       19008,
	"SC_ASSIGN_HB_RES":       19009,
	"SC_BROADCAST_NEW_HB":    19010,
	"CS_GRAB_HB_REQ":         19015,
	"SC_GRAB_HB_RES":         19016,
	"SC_BROADCAST_UPDATE_HB": 19017,
	"CS_PLAYERS_HB_LIST_REQ": 19020,
	"SC_PLAYERS_HB_LIST_RES": 19021,
}

func (x HBMSG) String() string {
	return proto.EnumName(HBMSG_name, int32(x))
}
func (HBMSG) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

// 请求进入房间
type ENTER_GAME_HB_REQ struct {
	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	RoomID    uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *ENTER_GAME_HB_REQ) Reset()                    { *m = ENTER_GAME_HB_REQ{} }
func (m *ENTER_GAME_HB_REQ) String() string            { return proto.CompactTextString(m) }
func (*ENTER_GAME_HB_REQ) ProtoMessage()               {}
func (*ENTER_GAME_HB_REQ) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *ENTER_GAME_HB_REQ) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *ENTER_GAME_HB_REQ) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

// 红包结构
type HONGBAO struct {
	ID            uint32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	AssignerAccID uint32 `protobuf:"varint,2,opt,name=AssignerAccID" json:"AssignerAccID,omitempty"`
	AssignerName  uint32 `protobuf:"varint,3,opt,name=AssignerName" json:"AssignerName,omitempty"`
	Value         uint64 `protobuf:"varint,4,opt,name=Value" json:"Value,omitempty"`
	Count         uint64 `protobuf:"varint,5,opt,name=Count" json:"Count,omitempty"`
	Spare         uint64 `protobuf:"varint,6,opt,name=Spare" json:"Spare,omitempty"`
	BombNumber    uint64 `protobuf:"varint,7,opt,name=BombNumber" json:"BombNumber,omitempty"`
}

func (m *HONGBAO) Reset()                    { *m = HONGBAO{} }
func (m *HONGBAO) String() string            { return proto.CompactTextString(m) }
func (*HONGBAO) ProtoMessage()               {}
func (*HONGBAO) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *HONGBAO) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *HONGBAO) GetAssignerAccID() uint32 {
	if m != nil {
		return m.AssignerAccID
	}
	return 0
}

func (m *HONGBAO) GetAssignerName() uint32 {
	if m != nil {
		return m.AssignerName
	}
	return 0
}

func (m *HONGBAO) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *HONGBAO) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *HONGBAO) GetSpare() uint64 {
	if m != nil {
		return m.Spare
	}
	return 0
}

func (m *HONGBAO) GetBombNumber() uint64 {
	if m != nil {
		return m.BombNumber
	}
	return 0
}

type ENTER_GAME_HB_RES struct {
	RoomID         uint32                `protobuf:"varint,1,opt,name=RoomID" json:"RoomID,omitempty"`
	HongBaoList    []*HONGBAO            `protobuf:"bytes,2,rep,name=HongBaoList" json:"HongBaoList,omitempty"`
	LuckPlayer     *AccountStorageData   `protobuf:"bytes,3,opt,name=luckPlayer" json:"luckPlayer,omitempty"`
	BigWealth      *AccountStorageData   `protobuf:"bytes,4,opt,name=bigWealth" json:"bigWealth,omitempty"`
	RankPlayers    []*AccountStorageData `protobuf:"bytes,5,rep,name=RankPlayers" json:"RankPlayers,omitempty"`
	Conf_MinValue  uint64                `protobuf:"varint,6,opt,name=Conf_MinValue,json=ConfMinValue" json:"Conf_MinValue,omitempty"`
	Conf_MaxValue  uint64                `protobuf:"varint,7,opt,name=Conf_MaxValue,json=ConfMaxValue" json:"Conf_MaxValue,omitempty"`
	Conf_OnceCount uint32                `protobuf:"varint,8,opt,name=Conf_OnceCount,json=ConfOnceCount" json:"Conf_OnceCount,omitempty"`
	Conf_Pump      uint32                `protobuf:"varint,9,opt,name=Conf_Pump,json=ConfPump" json:"Conf_Pump,omitempty"`
}

func (m *ENTER_GAME_HB_RES) Reset()                    { *m = ENTER_GAME_HB_RES{} }
func (m *ENTER_GAME_HB_RES) String() string            { return proto.CompactTextString(m) }
func (*ENTER_GAME_HB_RES) ProtoMessage()               {}
func (*ENTER_GAME_HB_RES) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *ENTER_GAME_HB_RES) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *ENTER_GAME_HB_RES) GetHongBaoList() []*HONGBAO {
	if m != nil {
		return m.HongBaoList
	}
	return nil
}

func (m *ENTER_GAME_HB_RES) GetLuckPlayer() *AccountStorageData {
	if m != nil {
		return m.LuckPlayer
	}
	return nil
}

func (m *ENTER_GAME_HB_RES) GetBigWealth() *AccountStorageData {
	if m != nil {
		return m.BigWealth
	}
	return nil
}

func (m *ENTER_GAME_HB_RES) GetRankPlayers() []*AccountStorageData {
	if m != nil {
		return m.RankPlayers
	}
	return nil
}

func (m *ENTER_GAME_HB_RES) GetConf_MinValue() uint64 {
	if m != nil {
		return m.Conf_MinValue
	}
	return 0
}

func (m *ENTER_GAME_HB_RES) GetConf_MaxValue() uint64 {
	if m != nil {
		return m.Conf_MaxValue
	}
	return 0
}

func (m *ENTER_GAME_HB_RES) GetConf_OnceCount() uint32 {
	if m != nil {
		return m.Conf_OnceCount
	}
	return 0
}

func (m *ENTER_GAME_HB_RES) GetConf_Pump() uint32 {
	if m != nil {
		return m.Conf_Pump
	}
	return 0
}

// 请求退出房间
type LEAVE_GAME_HB_REQ struct {
	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	RoomID    uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *LEAVE_GAME_HB_REQ) Reset()                    { *m = LEAVE_GAME_HB_REQ{} }
func (m *LEAVE_GAME_HB_REQ) String() string            { return proto.CompactTextString(m) }
func (*LEAVE_GAME_HB_REQ) ProtoMessage()               {}
func (*LEAVE_GAME_HB_REQ) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *LEAVE_GAME_HB_REQ) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *LEAVE_GAME_HB_REQ) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type LEAVE_GAME_HB_RES struct {
	Ret    uint32 `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	RoomID uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *LEAVE_GAME_HB_RES) Reset()                    { *m = LEAVE_GAME_HB_RES{} }
func (m *LEAVE_GAME_HB_RES) String() string            { return proto.CompactTextString(m) }
func (*LEAVE_GAME_HB_RES) ProtoMessage()               {}
func (*LEAVE_GAME_HB_RES) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *LEAVE_GAME_HB_RES) GetRet() uint32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *LEAVE_GAME_HB_RES) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

// 发红包
type ASSIGN_HB_REQ struct {
	Value      uint64 `protobuf:"varint,1,opt,name=Value" json:"Value,omitempty"`
	Count      uint32 `protobuf:"varint,2,opt,name=Count" json:"Count,omitempty"`
	BombNumber uint32 `protobuf:"varint,3,opt,name=BombNumber" json:"BombNumber,omitempty"`
	Num        uint32 `protobuf:"varint,4,opt,name=Num" json:"Num,omitempty"`
}

func (m *ASSIGN_HB_REQ) Reset()                    { *m = ASSIGN_HB_REQ{} }
func (m *ASSIGN_HB_REQ) String() string            { return proto.CompactTextString(m) }
func (*ASSIGN_HB_REQ) ProtoMessage()               {}
func (*ASSIGN_HB_REQ) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *ASSIGN_HB_REQ) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *ASSIGN_HB_REQ) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ASSIGN_HB_REQ) GetBombNumber() uint32 {
	if m != nil {
		return m.BombNumber
	}
	return 0
}

func (m *ASSIGN_HB_REQ) GetNum() uint32 {
	if m != nil {
		return m.Num
	}
	return 0
}

// 返回发红包结果
type ASSIGN_HB_RES struct {
	Ret uint64 `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
}

func (m *ASSIGN_HB_RES) Reset()                    { *m = ASSIGN_HB_RES{} }
func (m *ASSIGN_HB_RES) String() string            { return proto.CompactTextString(m) }
func (*ASSIGN_HB_RES) ProtoMessage()               {}
func (*ASSIGN_HB_RES) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *ASSIGN_HB_RES) GetRet() uint64 {
	if m != nil {
		return m.Ret
	}
	return 0
}

// 服务器广播添加一个新红包
type BROADCAST_NEW_HB struct {
	New *HONGBAO `protobuf:"bytes,1,opt,name=new" json:"new,omitempty"`
}

func (m *BROADCAST_NEW_HB) Reset()                    { *m = BROADCAST_NEW_HB{} }
func (m *BROADCAST_NEW_HB) String() string            { return proto.CompactTextString(m) }
func (*BROADCAST_NEW_HB) ProtoMessage()               {}
func (*BROADCAST_NEW_HB) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *BROADCAST_NEW_HB) GetNew() *HONGBAO {
	if m != nil {
		return m.New
	}
	return nil
}

// 抢红包
type GRAB_HB_REQ struct {
	ID uint32 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *GRAB_HB_REQ) Reset()                    { *m = GRAB_HB_REQ{} }
func (m *GRAB_HB_REQ) String() string            { return proto.CompactTextString(m) }
func (*GRAB_HB_REQ) ProtoMessage()               {}
func (*GRAB_HB_REQ) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *GRAB_HB_REQ) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

// 返回抢红包结果
type GRAB_HB_RES struct {
	Ret uint64 `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
}

func (m *GRAB_HB_RES) Reset()                    { *m = GRAB_HB_RES{} }
func (m *GRAB_HB_RES) String() string            { return proto.CompactTextString(m) }
func (*GRAB_HB_RES) ProtoMessage()               {}
func (*GRAB_HB_RES) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *GRAB_HB_RES) GetRet() uint64 {
	if m != nil {
		return m.Ret
	}
	return 0
}

// 服务器广播更新红包包数信息
type BROADCAST_UPDATE_HB struct {
	UD    uint32 `protobuf:"varint,1,opt,name=UD" json:"UD,omitempty"`
	Spare uint32 `protobuf:"varint,2,opt,name=Spare" json:"Spare,omitempty"`
}

func (m *BROADCAST_UPDATE_HB) Reset()                    { *m = BROADCAST_UPDATE_HB{} }
func (m *BROADCAST_UPDATE_HB) String() string            { return proto.CompactTextString(m) }
func (*BROADCAST_UPDATE_HB) ProtoMessage()               {}
func (*BROADCAST_UPDATE_HB) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

func (m *BROADCAST_UPDATE_HB) GetUD() uint32 {
	if m != nil {
		return m.UD
	}
	return 0
}

func (m *BROADCAST_UPDATE_HB) GetSpare() uint32 {
	if m != nil {
		return m.Spare
	}
	return 0
}

// 请求HB玩家列表
type PLAYERS_HB_LIST_RES struct {
	Players []*AccountStorageData `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
}

func (m *PLAYERS_HB_LIST_RES) Reset()                    { *m = PLAYERS_HB_LIST_RES{} }
func (m *PLAYERS_HB_LIST_RES) String() string            { return proto.CompactTextString(m) }
func (*PLAYERS_HB_LIST_RES) ProtoMessage()               {}
func (*PLAYERS_HB_LIST_RES) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{11} }

func (m *PLAYERS_HB_LIST_RES) GetPlayers() []*AccountStorageData {
	if m != nil {
		return m.Players
	}
	return nil
}

func init() {
	proto.RegisterType((*ENTER_GAME_HB_REQ)(nil), "protomsg.ENTER_GAME_HB_REQ")
	proto.RegisterType((*HONGBAO)(nil), "protomsg.HONGBAO")
	proto.RegisterType((*ENTER_GAME_HB_RES)(nil), "protomsg.ENTER_GAME_HB_RES")
	proto.RegisterType((*LEAVE_GAME_HB_REQ)(nil), "protomsg.LEAVE_GAME_HB_REQ")
	proto.RegisterType((*LEAVE_GAME_HB_RES)(nil), "protomsg.LEAVE_GAME_HB_RES")
	proto.RegisterType((*ASSIGN_HB_REQ)(nil), "protomsg.ASSIGN_HB_REQ")
	proto.RegisterType((*ASSIGN_HB_RES)(nil), "protomsg.ASSIGN_HB_RES")
	proto.RegisterType((*BROADCAST_NEW_HB)(nil), "protomsg.BROADCAST_NEW_HB")
	proto.RegisterType((*GRAB_HB_REQ)(nil), "protomsg.GRAB_HB_REQ")
	proto.RegisterType((*GRAB_HB_RES)(nil), "protomsg.GRAB_HB_RES")
	proto.RegisterType((*BROADCAST_UPDATE_HB)(nil), "protomsg.BROADCAST_UPDATE_HB")
	proto.RegisterType((*PLAYERS_HB_LIST_RES)(nil), "protomsg.PLAYERS_HB_LIST_RES")
	proto.RegisterEnum("protomsg.HBMSG", HBMSG_name, HBMSG_value)
}

func init() { proto.RegisterFile("protobuf/hongbao.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xda, 0x4a,
	0x10, 0x3e, 0x86, 0x10, 0xc2, 0x10, 0x72, 0x36, 0x4b, 0x84, 0x7c, 0x72, 0x72, 0x4e, 0x53, 0xa7,
	0x95, 0xa2, 0x5e, 0x24, 0x52, 0x22, 0xb5, 0x52, 0xff, 0x24, 0xdb, 0x58, 0x80, 0x0a, 0x86, 0x78,
	0x21, 0x51, 0x7b, 0x63, 0x19, 0xea, 0x10, 0x54, 0x6c, 0x23, 0x30, 0x6a, 0xfb, 0x1e, 0x3c, 0x44,
	0x5f, 0xa1, 0x3f, 0x17, 0x6d, 0xa5, 0xaa, 0xad, 0xd4, 0x5e, 0xf7, 0x75, 0xaa, 0xdd, 0x35, 0xc6,
	0x8e, 0xe9, 0xcf, 0x45, 0xaf, 0xd8, 0xf9, 0xe6, 0x9b, 0xd9, 0xd9, 0x6f, 0x3e, 0x03, 0xa5, 0xd1,
	0xd8, 0xf3, 0xbd, 0xee, 0xf4, 0xfc, 0xf0, 0xc2, 0x73, 0xfb, 0x5d, 0xcb, 0x3b, 0x60, 0x00, 0x5e,
	0x63, 0x3f, 0xce, 0xa4, 0xbf, 0x5d, 0x0c, 0x19, 0x8f, 0x2d, 0xdf, 0xe2, 0x69, 0xa9, 0x06, 0x9b,
	0x9a, 0xde, 0xd6, 0x0c, 0xb3, 0x22, 0x37, 0x34, 0xb3, 0xaa, 0x98, 0x86, 0x76, 0x82, 0x77, 0x20,
	0x27, 0xf7, 0x7a, 0xde, 0xd4, 0xf5, 0x6b, 0x65, 0x51, 0xd8, 0x15, 0xf6, 0x0b, 0xc6, 0x02, 0xc0,
	0x25, 0x58, 0x35, 0x3c, 0xcf, 0xa9, 0x95, 0xc5, 0x14, 0x4b, 0x05, 0x91, 0xf4, 0x41, 0x80, 0x6c,
	0xb5, 0xa9, 0x57, 0x14, 0xb9, 0x89, 0x37, 0x20, 0x15, 0x96, 0xa6, 0x6a, 0x65, 0x7c, 0x0d, 0x0a,
	0xf2, 0x64, 0x32, 0xe8, 0xbb, 0xf6, 0x58, 0xee, 0xf5, 0xc2, 0xd2, 0x38, 0x88, 0x25, 0x58, 0x9f,
	0x03, 0xba, 0xe5, 0xd8, 0x62, 0x9a, 0x91, 0x62, 0x18, 0xde, 0x82, 0xcc, 0xa9, 0x35, 0x9c, 0xda,
	0xe2, 0xca, 0xae, 0xb0, 0xbf, 0x62, 0xf0, 0x80, 0xa2, 0x2a, 0x1d, 0x4f, 0xcc, 0x70, 0x94, 0x05,
	0x14, 0x25, 0x23, 0x6b, 0x6c, 0x8b, 0xab, 0x1c, 0x65, 0x01, 0xfe, 0x1f, 0x40, 0xf1, 0x9c, 0xae,
	0x3e, 0x75, 0xba, 0xf6, 0x58, 0xcc, 0xb2, 0x54, 0x04, 0x91, 0x5e, 0xa4, 0x93, 0x9a, 0x90, 0xc8,
	0xab, 0x85, 0xe8, 0xab, 0xf1, 0x31, 0xe4, 0xab, 0x9e, 0xdb, 0x57, 0x2c, 0xaf, 0x3e, 0x98, 0xf8,
	0x62, 0x6a, 0x37, 0xbd, 0x9f, 0x3f, 0xda, 0x3c, 0x98, 0xab, 0x7e, 0x10, 0x28, 0x62, 0x44, 0x59,
	0xf8, 0x2e, 0xc0, 0x70, 0xda, 0x7b, 0xd2, 0x1a, 0x5a, 0xcf, 0xed, 0x31, 0x7b, 0x66, 0xfe, 0x68,
	0x67, 0x51, 0x13, 0x68, 0x4d, 0x7c, 0x6f, 0x6c, 0xf5, 0xed, 0xb2, 0xe5, 0x5b, 0x46, 0x84, 0x8f,
	0x6f, 0x43, 0xae, 0x3b, 0xe8, 0x9f, 0xd9, 0xd6, 0xd0, 0xbf, 0x60, 0x32, 0xfc, 0xaa, 0x78, 0x41,
	0xc7, 0xf7, 0x21, 0x6f, 0x58, 0x6e, 0xd0, 0x69, 0x22, 0x66, 0xd8, 0xb8, 0x3f, 0xaf, 0x8e, 0x16,
	0xe0, 0x3d, 0x28, 0xa8, 0x9e, 0x7b, 0x6e, 0x36, 0x06, 0x2e, 0x5f, 0x03, 0x97, 0x76, 0x9d, 0x82,
	0x73, 0x6c, 0x41, 0xb2, 0x9e, 0x71, 0x52, 0x36, 0x42, 0x0a, 0x30, 0x7c, 0x1d, 0x36, 0x18, 0xa9,
	0xe9, 0xf6, 0x6c, 0xbe, 0xbb, 0x35, 0xee, 0x09, 0x8a, 0x86, 0x20, 0xfe, 0x17, 0x72, 0x8c, 0xd6,
	0x9a, 0x3a, 0x23, 0x31, 0xc7, 0x18, 0x6b, 0x14, 0xa0, 0x31, 0x75, 0x6f, 0x5d, 0x93, 0x4f, 0xb5,
	0x3f, 0xe0, 0xde, 0x7b, 0xc9, 0x56, 0x04, 0x23, 0x48, 0x1b, 0xb6, 0x1f, 0x34, 0xa1, 0xc7, 0x1f,
	0x96, 0x3b, 0x50, 0x90, 0x09, 0xa9, 0x55, 0xf4, 0xf9, 0x14, 0xa1, 0x4f, 0x85, 0xa5, 0x3e, 0xe5,
	0xd5, 0x81, 0x4f, 0xe3, 0x8e, 0xe4, 0xae, 0x8f, 0x20, 0x74, 0x0c, 0x7d, 0xea, 0xb0, 0x55, 0x17,
	0x0c, 0x7a, 0x94, 0xae, 0xc6, 0xaf, 0x8b, 0x4d, 0xba, 0xc2, 0x26, 0x95, 0x6e, 0x01, 0x52, 0x8c,
	0xa6, 0x5c, 0x56, 0x65, 0xd2, 0x36, 0x75, 0xed, 0xcc, 0xac, 0x2a, 0x78, 0x0f, 0xd2, 0xae, 0xfd,
	0x94, 0xb1, 0x96, 0x9a, 0x94, 0x66, 0xa5, 0xff, 0x20, 0x5f, 0x31, 0x64, 0x65, 0xfe, 0x90, 0x4b,
	0x9f, 0xb2, 0x74, 0x25, 0x9a, 0x5e, 0x76, 0xf1, 0x1d, 0x28, 0x2e, 0x2e, 0xee, 0xb4, 0xca, 0x72,
	0x9b, 0xea, 0x49, 0xfb, 0x74, 0xc2, 0x3e, 0x9d, 0xf2, 0xe2, 0xe3, 0x0c, 0xa4, 0x60, 0x81, 0xd4,
	0x80, 0x62, 0xab, 0x2e, 0x3f, 0xd4, 0x0c, 0x42, 0x2f, 0xa8, 0xd7, 0x48, 0x9b, 0xdd, 0x72, 0x13,
	0xb2, 0xa3, 0xc0, 0xb2, 0xc2, 0x6f, 0x58, 0x76, 0x4e, 0xbe, 0xf1, 0x2d, 0x05, 0x99, 0xaa, 0xd2,
	0x20, 0x15, 0x5c, 0x80, 0x5c, 0x47, 0x7f, 0xa0, 0x37, 0xa9, 0x0e, 0xe8, 0x2f, 0xbc, 0x0d, 0x5b,
	0x2a, 0x31, 0x13, 0x7f, 0x7d, 0xe8, 0xe5, 0x4c, 0xa0, 0x39, 0xa2, 0x26, 0x72, 0x04, 0xbd, 0xe2,
	0x39, 0x95, 0x98, 0x09, 0xd3, 0xa1, 0xd7, 0x61, 0x5d, 0xc2, 0x45, 0xe8, 0xcd, 0x4c, 0xc0, 0x25,
	0x40, 0x2a, 0x31, 0x63, 0x16, 0x41, 0x6f, 0x39, 0x4e, 0xd4, 0x18, 0x4e, 0xd0, 0xbb, 0x99, 0x80,
	0xff, 0x81, 0x22, 0x51, 0xcd, 0xcb, 0x0b, 0x44, 0xef, 0x67, 0x02, 0xde, 0x82, 0x0d, 0x95, 0x98,
	0x91, 0x15, 0xa1, 0x8f, 0x1c, 0x25, 0x6a, 0x04, 0x25, 0xe8, 0xd3, 0x4c, 0xc0, 0x3b, 0x50, 0x8a,
	0xb5, 0x09, 0xd7, 0x81, 0x3e, 0xf3, 0xac, 0x4a, 0xcc, 0xa4, 0xde, 0x27, 0xe8, 0x4b, 0x58, 0xbb,
	0x64, 0x1b, 0xe8, 0xeb, 0x4c, 0x50, 0xfe, 0x7e, 0x54, 0x18, 0x7b, 0x9e, 0x7f, 0x38, 0x5f, 0x43,
	0x77, 0x95, 0x9d, 0x8e, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x06, 0x4f, 0xaf, 0x68, 0x89, 0x06,
	0x00, 0x00,
}
