// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/dfdc.proto

package protomsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 网络消息
type DFDCMSG int32

const (
	DFDCMSG_UNKNOW_DFDC              DFDCMSG = 0
	DFDCMSG_CS_ENTER_GAME_DFDC_REQ   DFDCMSG = 13001
	DFDCMSG_SC_ENTER_GAME_DFDC_RES   DFDCMSG = 13002
	DFDCMSG_CS_LEAVE_GAME_DFDC_REQ   DFDCMSG = 13003
	DFDCMSG_SC_LEAVE_GAME_DFDC_RES   DFDCMSG = 13004
	DFDCMSG_CS_START_DFDC_REQ        DFDCMSG = 13008
	DFDCMSG_SC_START_DFDC_RES        DFDCMSG = 13009
	DFDCMSG_SC_UPDATE_DFDC_BONUS     DFDCMSG = 13010
	DFDCMSG_CS_PLAYERS_DFDC_LIST_REQ DFDCMSG = 13015
	DFDCMSG_SC_PLAYERS_DFDC_LIST_RES DFDCMSG = 13016
)

var DFDCMSG_name = map[int32]string{
	0:     "UNKNOW_DFDC",
	13001: "CS_ENTER_GAME_DFDC_REQ",
	13002: "SC_ENTER_GAME_DFDC_RES",
	13003: "CS_LEAVE_GAME_DFDC_REQ",
	13004: "SC_LEAVE_GAME_DFDC_RES",
	13008: "CS_START_DFDC_REQ",
	13009: "SC_START_DFDC_RES",
	13010: "SC_UPDATE_DFDC_BONUS",
	13015: "CS_PLAYERS_DFDC_LIST_REQ",
	13016: "SC_PLAYERS_DFDC_LIST_RES",
}
var DFDCMSG_value = map[string]int32{
	"UNKNOW_DFDC":              0,
	"CS_ENTER_GAME_DFDC_REQ":   13001,
	"SC_ENTER_GAME_DFDC_RES":   13002,
	"CS_LEAVE_GAME_DFDC_REQ":   13003,
	"SC_LEAVE_GAME_DFDC_RES":   13004,
	"CS_START_DFDC_REQ":        13008,
	"SC_START_DFDC_RES":        13009,
	"SC_UPDATE_DFDC_BONUS":     13010,
	"CS_PLAYERS_DFDC_LIST_REQ": 13015,
	"SC_PLAYERS_DFDC_LIST_RES": 13016,
}

func (x DFDCMSG) String() string {
	return proto.EnumName(DFDCMSG_name, int32(x))
}
func (DFDCMSG) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// 游戏1 图案枚举
type DFDCID int32

const (
	DFDCID_DFDCUnknow DFDCID = 0
	DFDCID_DFDC1      DFDCID = 1
	DFDCID_DFDC2      DFDCID = 2
	DFDCID_DFDC3      DFDCID = 3
	DFDCID_DFDC4      DFDCID = 4
	DFDCID_DFDC5      DFDCID = 5
	DFDCID_DFDC6      DFDCID = 6
	DFDCID_DFDC7      DFDCID = 7
	DFDCID_DFDC8      DFDCID = 8
	DFDCID_DFDC9      DFDCID = 9
	DFDCID_DFDC10     DFDCID = 10
	DFDCID_DFDC11     DFDCID = 11
	DFDCID_DFDC12     DFDCID = 12
	DFDCID_DFDC13     DFDCID = 13
)

var DFDCID_name = map[int32]string{
	0:  "DFDCUnknow",
	1:  "DFDC1",
	2:  "DFDC2",
	3:  "DFDC3",
	4:  "DFDC4",
	5:  "DFDC5",
	6:  "DFDC6",
	7:  "DFDC7",
	8:  "DFDC8",
	9:  "DFDC9",
	10: "DFDC10",
	11: "DFDC11",
	12: "DFDC12",
	13: "DFDC13",
}
var DFDCID_value = map[string]int32{
	"DFDCUnknow": 0,
	"DFDC1":      1,
	"DFDC2":      2,
	"DFDC3":      3,
	"DFDC4":      4,
	"DFDC5":      5,
	"DFDC6":      6,
	"DFDC7":      7,
	"DFDC8":      8,
	"DFDC9":      9,
	"DFDC10":     10,
	"DFDC11":     11,
	"DFDC12":     12,
	"DFDC13":     13,
}

func (x DFDCID) String() string {
	return proto.EnumName(DFDCID_name, int32(x))
}
func (DFDCID) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// 请求进入房间
type ENTER_GAME_DFDC_REQ struct {
	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	RoomID    uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *ENTER_GAME_DFDC_REQ) Reset()                    { *m = ENTER_GAME_DFDC_REQ{} }
func (m *ENTER_GAME_DFDC_REQ) String() string            { return proto.CompactTextString(m) }
func (*ENTER_GAME_DFDC_REQ) ProtoMessage()               {}
func (*ENTER_GAME_DFDC_REQ) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ENTER_GAME_DFDC_REQ) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *ENTER_GAME_DFDC_REQ) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type ENTER_GAME_DFDC_RES struct {
	RoomID   uint32                                   `protobuf:"varint,1,opt,name=RoomID" json:"RoomID,omitempty"`
	Basics   int64                                    `protobuf:"varint,2,opt,name=Basics" json:"Basics,omitempty"`
	Bonus    map[int32]int64                          `protobuf:"bytes,3,rep,name=Bonus" json:"Bonus,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	LastBet  int64                                    `protobuf:"varint,4,opt,name=LastBet" json:"LastBet,omitempty"`
	Bets     []uint64                                 `protobuf:"varint,5,rep,packed,name=Bets" json:"Bets,omitempty"`
	Ratio    map[int32]*ENTER_GAME_DFDC_RES_DfdcRatio `protobuf:"bytes,6,rep,name=Ratio" json:"Ratio,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FeeCount int32                                    `protobuf:"varint,8,opt,name=FeeCount" json:"FeeCount,omitempty"`
}

func (m *ENTER_GAME_DFDC_RES) Reset()                    { *m = ENTER_GAME_DFDC_RES{} }
func (m *ENTER_GAME_DFDC_RES) String() string            { return proto.CompactTextString(m) }
func (*ENTER_GAME_DFDC_RES) ProtoMessage()               {}
func (*ENTER_GAME_DFDC_RES) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ENTER_GAME_DFDC_RES) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *ENTER_GAME_DFDC_RES) GetBasics() int64 {
	if m != nil {
		return m.Basics
	}
	return 0
}

func (m *ENTER_GAME_DFDC_RES) GetBonus() map[int32]int64 {
	if m != nil {
		return m.Bonus
	}
	return nil
}

func (m *ENTER_GAME_DFDC_RES) GetLastBet() int64 {
	if m != nil {
		return m.LastBet
	}
	return 0
}

func (m *ENTER_GAME_DFDC_RES) GetBets() []uint64 {
	if m != nil {
		return m.Bets
	}
	return nil
}

func (m *ENTER_GAME_DFDC_RES) GetRatio() map[int32]*ENTER_GAME_DFDC_RES_DfdcRatio {
	if m != nil {
		return m.Ratio
	}
	return nil
}

func (m *ENTER_GAME_DFDC_RES) GetFeeCount() int32 {
	if m != nil {
		return m.FeeCount
	}
	return 0
}

type ENTER_GAME_DFDC_RES_DfdcRatio struct {
	ID     DFDCID `protobuf:"varint,1,opt,name=ID,enum=protomsg.DFDCID" json:"ID,omitempty"`
	Same_3 int32  `protobuf:"varint,4,opt,name=Same_3,json=Same3" json:"Same_3,omitempty"`
	Same_4 int32  `protobuf:"varint,5,opt,name=Same_4,json=Same4" json:"Same_4,omitempty"`
	Same_5 int32  `protobuf:"varint,6,opt,name=Same_5,json=Same5" json:"Same_5,omitempty"`
}

func (m *ENTER_GAME_DFDC_RES_DfdcRatio) Reset()         { *m = ENTER_GAME_DFDC_RES_DfdcRatio{} }
func (m *ENTER_GAME_DFDC_RES_DfdcRatio) String() string { return proto.CompactTextString(m) }
func (*ENTER_GAME_DFDC_RES_DfdcRatio) ProtoMessage()    {}
func (*ENTER_GAME_DFDC_RES_DfdcRatio) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 0}
}

func (m *ENTER_GAME_DFDC_RES_DfdcRatio) GetID() DFDCID {
	if m != nil {
		return m.ID
	}
	return DFDCID_DFDCUnknow
}

func (m *ENTER_GAME_DFDC_RES_DfdcRatio) GetSame_3() int32 {
	if m != nil {
		return m.Same_3
	}
	return 0
}

func (m *ENTER_GAME_DFDC_RES_DfdcRatio) GetSame_4() int32 {
	if m != nil {
		return m.Same_4
	}
	return 0
}

func (m *ENTER_GAME_DFDC_RES_DfdcRatio) GetSame_5() int32 {
	if m != nil {
		return m.Same_5
	}
	return 0
}

// 请求退出房间
type LEAVE_GAME_DFDC_REQ struct {
	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	RoomID    uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *LEAVE_GAME_DFDC_REQ) Reset()                    { *m = LEAVE_GAME_DFDC_REQ{} }
func (m *LEAVE_GAME_DFDC_REQ) String() string            { return proto.CompactTextString(m) }
func (*LEAVE_GAME_DFDC_REQ) ProtoMessage()               {}
func (*LEAVE_GAME_DFDC_REQ) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *LEAVE_GAME_DFDC_REQ) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *LEAVE_GAME_DFDC_REQ) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type LEAVE_GAME_DFDC_RES struct {
	Ret    uint32 `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	RoomID uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *LEAVE_GAME_DFDC_RES) Reset()                    { *m = LEAVE_GAME_DFDC_RES{} }
func (m *LEAVE_GAME_DFDC_RES) String() string            { return proto.CompactTextString(m) }
func (*LEAVE_GAME_DFDC_RES) ProtoMessage()               {}
func (*LEAVE_GAME_DFDC_RES) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *LEAVE_GAME_DFDC_RES) GetRet() uint32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *LEAVE_GAME_DFDC_RES) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

// //////////////////////////////////////////// 游戏 /////////////////////////////////////////////
// 请求开始游戏1
type START_DFDC_REQ struct {
	Bet uint64 `protobuf:"varint,1,opt,name=Bet" json:"Bet,omitempty"`
}

func (m *START_DFDC_REQ) Reset()                    { *m = START_DFDC_REQ{} }
func (m *START_DFDC_REQ) String() string            { return proto.CompactTextString(m) }
func (*START_DFDC_REQ) ProtoMessage()               {}
func (*START_DFDC_REQ) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *START_DFDC_REQ) GetBet() uint64 {
	if m != nil {
		return m.Bet
	}
	return 0
}

type START_DFDC_RES struct {
	Ret         uint64          `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	PictureList []int32         `protobuf:"varint,2,rep,packed,name=PictureList" json:"PictureList,omitempty"`
	Bonus       map[int32]int64 `protobuf:"bytes,3,rep,name=Bonus" json:"Bonus,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Money       int64           `protobuf:"varint,4,opt,name=Money" json:"Money,omitempty"`
	TotalOdds   int64           `protobuf:"varint,5,opt,name=TotalOdds" json:"TotalOdds,omitempty"`
	FreeCount   int64           `protobuf:"varint,6,opt,name=FreeCount" json:"FreeCount,omitempty"`
	Shows       []*DFDCPosition `protobuf:"bytes,7,rep,name=Shows" json:"Shows,omitempty"`
}

func (m *START_DFDC_RES) Reset()                    { *m = START_DFDC_RES{} }
func (m *START_DFDC_RES) String() string            { return proto.CompactTextString(m) }
func (*START_DFDC_RES) ProtoMessage()               {}
func (*START_DFDC_RES) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *START_DFDC_RES) GetRet() uint64 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *START_DFDC_RES) GetPictureList() []int32 {
	if m != nil {
		return m.PictureList
	}
	return nil
}

func (m *START_DFDC_RES) GetBonus() map[int32]int64 {
	if m != nil {
		return m.Bonus
	}
	return nil
}

func (m *START_DFDC_RES) GetMoney() int64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *START_DFDC_RES) GetTotalOdds() int64 {
	if m != nil {
		return m.TotalOdds
	}
	return 0
}

func (m *START_DFDC_RES) GetFreeCount() int64 {
	if m != nil {
		return m.FreeCount
	}
	return 0
}

func (m *START_DFDC_RES) GetShows() []*DFDCPosition {
	if m != nil {
		return m.Shows
	}
	return nil
}

type DFDCPosition struct {
	Px int32 `protobuf:"varint,1,opt,name=px" json:"px,omitempty"`
	Py int32 `protobuf:"varint,2,opt,name=py" json:"py,omitempty"`
}

func (m *DFDCPosition) Reset()                    { *m = DFDCPosition{} }
func (m *DFDCPosition) String() string            { return proto.CompactTextString(m) }
func (*DFDCPosition) ProtoMessage()               {}
func (*DFDCPosition) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *DFDCPosition) GetPx() int32 {
	if m != nil {
		return m.Px
	}
	return 0
}

func (m *DFDCPosition) GetPy() int32 {
	if m != nil {
		return m.Py
	}
	return 0
}

// 通知更新奖金池
type UPDATE_DFDC_BONUS struct {
	Bonus map[int32]int64 `protobuf:"bytes,1,rep,name=Bonus" json:"Bonus,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *UPDATE_DFDC_BONUS) Reset()                    { *m = UPDATE_DFDC_BONUS{} }
func (m *UPDATE_DFDC_BONUS) String() string            { return proto.CompactTextString(m) }
func (*UPDATE_DFDC_BONUS) ProtoMessage()               {}
func (*UPDATE_DFDC_BONUS) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *UPDATE_DFDC_BONUS) GetBonus() map[int32]int64 {
	if m != nil {
		return m.Bonus
	}
	return nil
}

// 请求dfdc玩家列表
type PLAYERS_DFDC_LIST_RES struct {
	Players []*AccountStorageData `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
}

func (m *PLAYERS_DFDC_LIST_RES) Reset()                    { *m = PLAYERS_DFDC_LIST_RES{} }
func (m *PLAYERS_DFDC_LIST_RES) String() string            { return proto.CompactTextString(m) }
func (*PLAYERS_DFDC_LIST_RES) ProtoMessage()               {}
func (*PLAYERS_DFDC_LIST_RES) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *PLAYERS_DFDC_LIST_RES) GetPlayers() []*AccountStorageData {
	if m != nil {
		return m.Players
	}
	return nil
}

func init() {
	proto.RegisterType((*ENTER_GAME_DFDC_REQ)(nil), "protomsg.ENTER_GAME_DFDC_REQ")
	proto.RegisterType((*ENTER_GAME_DFDC_RES)(nil), "protomsg.ENTER_GAME_DFDC_RES")
	proto.RegisterType((*ENTER_GAME_DFDC_RES_DfdcRatio)(nil), "protomsg.ENTER_GAME_DFDC_RES.DfdcRatio")
	proto.RegisterType((*LEAVE_GAME_DFDC_REQ)(nil), "protomsg.LEAVE_GAME_DFDC_REQ")
	proto.RegisterType((*LEAVE_GAME_DFDC_RES)(nil), "protomsg.LEAVE_GAME_DFDC_RES")
	proto.RegisterType((*START_DFDC_REQ)(nil), "protomsg.START_DFDC_REQ")
	proto.RegisterType((*START_DFDC_RES)(nil), "protomsg.START_DFDC_RES")
	proto.RegisterType((*DFDCPosition)(nil), "protomsg.DFDC_position")
	proto.RegisterType((*UPDATE_DFDC_BONUS)(nil), "protomsg.UPDATE_DFDC_BONUS")
	proto.RegisterType((*PLAYERS_DFDC_LIST_RES)(nil), "protomsg.PLAYERS_DFDC_LIST_RES")
	proto.RegisterEnum("protomsg.DFDCMSG", DFDCMSG_name, DFDCMSG_value)
	proto.RegisterEnum("protomsg.DFDCID", DFDCID_name, DFDCID_value)
}

func init() { proto.RegisterFile("protobuf/dfdc.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdb, 0x6e, 0xe3, 0x44,
	0x18, 0x5e, 0xdb, 0xb1, 0xd3, 0xfe, 0xa1, 0xdd, 0xd9, 0xd9, 0x6e, 0x31, 0x61, 0x91, 0xa2, 0x20,
	0x41, 0xb4, 0x12, 0x2d, 0x4d, 0xda, 0xa5, 0x8b, 0x38, 0x28, 0xb1, 0xdd, 0x55, 0xb4, 0xe9, 0x61,
	0x67, 0x52, 0x10, 0xdc, 0x58, 0xde, 0x64, 0xba, 0x44, 0xdb, 0x66, 0x22, 0x7b, 0xc2, 0x36, 0xcf,
	0xc0, 0x03, 0x70, 0xcb, 0x3d, 0x6f, 0xc0, 0x1b, 0x70, 0xb8, 0x00, 0x6e, 0xe0, 0x71, 0xd0, 0x8c,
	0x3d, 0x76, 0x42, 0x0d, 0x42, 0xc0, 0x55, 0xbe, 0xff, 0xf4, 0xfd, 0xbf, 0xbf, 0x7c, 0x36, 0xdc,
	0x9d, 0xc5, 0x5c, 0xf0, 0x67, 0xf3, 0x8b, 0xdd, 0xf1, 0xc5, 0x78, 0xb4, 0xa3, 0x22, 0xbc, 0xa6,
	0x7e, 0xae, 0x92, 0xe7, 0xf5, 0xa5, 0x72, 0x24, 0xa2, 0xb4, 0xdc, 0x7c, 0x02, 0x77, 0x83, 0x93,
	0x61, 0x40, 0xc2, 0xc7, 0xdd, 0xe3, 0x20, 0xf4, 0x8f, 0x7c, 0x2f, 0x24, 0xc1, 0x53, 0x7c, 0x1f,
	0xd6, 0xbb, 0xa3, 0x11, 0x9f, 0x4f, 0x45, 0xdf, 0x77, 0x8d, 0x86, 0xd1, 0xda, 0x20, 0x45, 0x02,
	0x6f, 0x83, 0x43, 0x38, 0xbf, 0xea, 0xfb, 0xae, 0xa9, 0x4a, 0x59, 0xd4, 0xfc, 0xa6, 0x52, 0xc6,
	0x46, 0x97, 0xfa, 0x8d, 0xe5, 0x7e, 0x99, 0xef, 0x45, 0xc9, 0x64, 0x94, 0x28, 0x1e, 0x8b, 0x64,
	0x11, 0xfe, 0x08, 0xec, 0x1e, 0x9f, 0xce, 0x13, 0xd7, 0x6a, 0x58, 0xad, 0x5a, 0xbb, 0xb5, 0xa3,
	0x9f, 0x61, 0xa7, 0x84, 0x7d, 0x47, 0xb5, 0x06, 0x53, 0x11, 0x2f, 0x48, 0x3a, 0x86, 0x5d, 0xa8,
	0x0e, 0xa2, 0x44, 0xf4, 0x98, 0x70, 0x2b, 0x8a, 0x58, 0x87, 0x18, 0x43, 0xa5, 0xc7, 0x44, 0xe2,
	0xda, 0x0d, 0xab, 0x55, 0x21, 0x0a, 0xcb, 0x6d, 0x24, 0x12, 0x13, 0xee, 0x3a, 0xff, 0x64, 0x9b,
	0x6a, 0xcd, 0xb6, 0x29, 0x8c, 0xeb, 0xb0, 0x76, 0xc4, 0x98, 0x27, 0xb5, 0x71, 0xd7, 0x1a, 0x46,
	0xcb, 0x26, 0x79, 0x5c, 0x8f, 0x61, 0xdd, 0xbf, 0x18, 0x8f, 0xd2, 0xc6, 0x06, 0x98, 0x99, 0x04,
	0x9b, 0x6d, 0x54, 0x6c, 0x91, 0xd4, 0x7d, 0x9f, 0x98, 0x7d, 0x1f, 0xdf, 0x03, 0x87, 0x46, 0x57,
	0x2c, 0xec, 0xa8, 0xbb, 0x6d, 0x62, 0xcb, 0xa8, 0x93, 0xa7, 0xf7, 0x5d, 0xbb, 0x48, 0xef, 0xe7,
	0xe9, 0x03, 0xd7, 0x29, 0xd2, 0x07, 0xf5, 0x43, 0x80, 0x42, 0x12, 0x8c, 0xc0, 0x7a, 0xc1, 0x16,
	0x6a, 0xab, 0x4d, 0x24, 0xc4, 0x5b, 0x60, 0x7f, 0x19, 0x5d, 0xce, 0x59, 0x26, 0x7a, 0x1a, 0xbc,
	0x6f, 0x1e, 0x1a, 0xf5, 0x08, 0xa0, 0x78, 0xbc, 0x92, 0xc9, 0x0f, 0x97, 0x27, 0x6b, 0xed, 0xb7,
	0xff, 0x5e, 0xa9, 0xfc, 0xc1, 0x97, 0x56, 0x48, 0xbf, 0x0d, 0x82, 0xee, 0x27, 0xc1, 0xff, 0xe2,
	0xb7, 0x8f, 0xcb, 0xc8, 0xa8, 0x3c, 0x9c, 0x30, 0x91, 0xd1, 0x48, 0xf8, 0x97, 0x04, 0x4d, 0xd8,
	0xa4, 0xc3, 0x2e, 0x19, 0x16, 0x87, 0x20, 0xb0, 0x7a, 0xd9, 0x6c, 0x85, 0x48, 0xd8, 0xfc, 0xce,
	0xfc, 0x53, 0xd3, 0xca, 0x82, 0x4a, 0xba, 0xa0, 0x01, 0xb5, 0xb3, 0xc9, 0x48, 0xcc, 0x63, 0x36,
	0x98, 0x24, 0xc2, 0x35, 0x1b, 0x56, 0xcb, 0x26, 0xcb, 0x29, 0xfc, 0x68, 0xd5, 0xd3, 0x6f, 0x16,
	0xda, 0xad, 0x92, 0x97, 0xd8, 0x79, 0x0b, 0xec, 0x63, 0x3e, 0x65, 0x8b, 0xcc, 0xcc, 0x69, 0x20,
	0x25, 0x1b, 0x72, 0x11, 0x5d, 0x9e, 0x8e, 0xc7, 0x89, 0xf2, 0x85, 0x45, 0x8a, 0x84, 0xac, 0x1e,
	0xc5, 0xda, 0x95, 0x4e, 0x5a, 0xcd, 0x13, 0xf8, 0x1d, 0xb0, 0xe9, 0x17, 0xfc, 0x65, 0xe2, 0x56,
	0xd5, 0x31, 0xaf, 0xae, 0x9a, 0x31, 0x9c, 0xf1, 0x64, 0x22, 0x26, 0x7c, 0x4a, 0xd2, 0xae, 0x7f,
	0xef, 0xa8, 0xe6, 0x2e, 0x6c, 0xac, 0x30, 0xe2, 0x4d, 0x30, 0x67, 0xd7, 0xd9, 0xac, 0x39, 0xbb,
	0x56, 0xf1, 0x42, 0xcd, 0xc9, 0x78, 0xd1, 0xfc, 0xca, 0x80, 0x3b, 0xe7, 0x67, 0x7e, 0x77, 0x98,
	0xfd, 0x9f, 0xbd, 0xd3, 0x93, 0x73, 0x8a, 0x3f, 0xd0, 0xe2, 0x19, 0xea, 0xde, 0xb7, 0x8a, 0x7b,
	0x6f, 0xf4, 0xde, 0xd4, 0xef, 0x3f, 0x9c, 0x7f, 0x0a, 0xf7, 0xce, 0x06, 0xdd, 0xcf, 0x02, 0x42,
	0xd3, 0x0d, 0x83, 0x3e, 0x1d, 0x2a, 0x07, 0x3c, 0x84, 0xea, 0xec, 0x32, 0x5a, 0xb0, 0x58, 0x9f,
	0x74, 0xbf, 0x38, 0x29, 0xf3, 0x2d, 0x15, 0x3c, 0x8e, 0x9e, 0x33, 0x3f, 0x12, 0x11, 0xd1, 0xcd,
	0x0f, 0xbe, 0x36, 0xa1, 0x2a, 0x99, 0x8e, 0xe9, 0x63, 0x7c, 0x1b, 0x6a, 0xe7, 0x27, 0x4f, 0x4e,
	0x4e, 0x3f, 0x55, 0xdc, 0xe8, 0x16, 0x7e, 0x1d, 0xb6, 0x3d, 0x1a, 0x96, 0x7c, 0x8e, 0xd1, 0xf7,
	0x4c, 0x16, 0xa9, 0x57, 0x52, 0xa4, 0xe8, 0x07, 0x96, 0x4d, 0x96, 0xbc, 0x58, 0xe8, 0x47, 0x3d,
	0x59, 0xf2, 0xa2, 0xa0, 0x9f, 0x18, 0xde, 0x86, 0x3b, 0x1e, 0x0d, 0x57, 0x5f, 0x02, 0xf4, 0xb3,
	0xca, 0x53, 0x6f, 0x35, 0x4f, 0xd1, 0x2f, 0x0c, 0xbf, 0x06, 0x5b, 0xd4, 0x0b, 0x6f, 0xa8, 0x8e,
	0x7e, 0x65, 0xf8, 0x0d, 0x70, 0x3d, 0x1a, 0x96, 0xe9, 0xf5, 0x14, 0xfd, 0xa6, 0xca, 0xd4, 0x2b,
	0x2d, 0x53, 0xf4, 0x3b, 0x7b, 0xf0, 0xad, 0x01, 0x4e, 0xfa, 0x25, 0xc4, 0x9b, 0x00, 0x12, 0x9d,
	0x4f, 0x5f, 0x4c, 0xf9, 0x4b, 0x74, 0x0b, 0xaf, 0x83, 0x2d, 0xe3, 0x3d, 0x64, 0x68, 0xd8, 0x46,
	0xa6, 0x86, 0x1d, 0x64, 0x69, 0xb8, 0x8f, 0x2a, 0x1a, 0x1e, 0x20, 0x5b, 0xc3, 0x87, 0xc8, 0xd1,
	0xf0, 0x3d, 0x54, 0xd5, 0xf0, 0x10, 0xad, 0x69, 0xf8, 0x08, 0xad, 0x63, 0x48, 0x97, 0xef, 0xbd,
	0x8b, 0x20, 0xc7, 0x7b, 0xa8, 0x96, 0xe3, 0x36, 0x7a, 0x25, 0xc7, 0x1d, 0xb4, 0xd1, 0xbb, 0xfd,
	0xf9, 0x46, 0xcc, 0xb9, 0xd8, 0xd5, 0x7f, 0xfa, 0x33, 0x47, 0xa1, 0xce, 0x1f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x6b, 0xc3, 0xb4, 0xea, 0x84, 0x07, 0x00, 0x00,
}
