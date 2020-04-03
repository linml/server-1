// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/luckfruit.proto

package protomsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 网络消息
type LUCKFRUITMSG int32

const (
	LUCKFRUITMSG_UNKNOW_LUCKFRUIT              LUCKFRUITMSG = 0
	LUCKFRUITMSG_CS_ENTER_GAME_LUCKFRUIT_REQ   LUCKFRUITMSG = 15001
	LUCKFRUITMSG_SC_ENTER_GAME_LUCKFRUIT_RES   LUCKFRUITMSG = 15002
	LUCKFRUITMSG_CS_LEAVE_GAME_LUCKFRUIT_REQ   LUCKFRUITMSG = 15003
	LUCKFRUITMSG_SC_LEAVE_GAME_LUCKFRUIT_RES   LUCKFRUITMSG = 15004
	LUCKFRUITMSG_CS_START_LUCKFRUIT_REQ        LUCKFRUITMSG = 15008
	LUCKFRUITMSG_SC_START_LUCKFRUIT_RES        LUCKFRUITMSG = 15009
	LUCKFRUITMSG_SC_UPDATE_LUCKFRUIT_BONUS     LUCKFRUITMSG = 15010
	LUCKFRUITMSG_CS_PLAYERS_LUCKFRUIT_LIST_REQ LUCKFRUITMSG = 15015
	LUCKFRUITMSG_SC_PLAYERS_LUCKFRUIT_LIST_RES LUCKFRUITMSG = 15016
)

var LUCKFRUITMSG_name = map[int32]string{
	0:     "UNKNOW_LUCKFRUIT",
	15001: "CS_ENTER_GAME_LUCKFRUIT_REQ",
	15002: "SC_ENTER_GAME_LUCKFRUIT_RES",
	15003: "CS_LEAVE_GAME_LUCKFRUIT_REQ",
	15004: "SC_LEAVE_GAME_LUCKFRUIT_RES",
	15008: "CS_START_LUCKFRUIT_REQ",
	15009: "SC_START_LUCKFRUIT_RES",
	15010: "SC_UPDATE_LUCKFRUIT_BONUS",
	15015: "CS_PLAYERS_LUCKFRUIT_LIST_REQ",
	15016: "SC_PLAYERS_LUCKFRUIT_LIST_RES",
}
var LUCKFRUITMSG_value = map[string]int32{
	"UNKNOW_LUCKFRUIT":              0,
	"CS_ENTER_GAME_LUCKFRUIT_REQ":   15001,
	"SC_ENTER_GAME_LUCKFRUIT_RES":   15002,
	"CS_LEAVE_GAME_LUCKFRUIT_REQ":   15003,
	"SC_LEAVE_GAME_LUCKFRUIT_RES":   15004,
	"CS_START_LUCKFRUIT_REQ":        15008,
	"SC_START_LUCKFRUIT_RES":        15009,
	"SC_UPDATE_LUCKFRUIT_BONUS":     15010,
	"CS_PLAYERS_LUCKFRUIT_LIST_REQ": 15015,
	"SC_PLAYERS_LUCKFRUIT_LIST_RES": 15016,
}

func (x LUCKFRUITMSG) String() string {
	return proto.EnumName(LUCKFRUITMSG_name, int32(x))
}
func (LUCKFRUITMSG) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

// 游戏1 图案枚举
type LUCKFRUITID int32

const (
	LUCKFRUITID_LUCKFRUITUnknow LUCKFRUITID = 0
	LUCKFRUITID_LUCKFRUIT1      LUCKFRUITID = 1
	LUCKFRUITID_LUCKFRUIT2      LUCKFRUITID = 2
	LUCKFRUITID_LUCKFRUIT3      LUCKFRUITID = 3
	LUCKFRUITID_LUCKFRUIT4      LUCKFRUITID = 4
	LUCKFRUITID_LUCKFRUIT5      LUCKFRUITID = 5
	LUCKFRUITID_LUCKFRUIT6      LUCKFRUITID = 6
	LUCKFRUITID_LUCKFRUIT7      LUCKFRUITID = 7
	LUCKFRUITID_LUCKFRUIT8      LUCKFRUITID = 8
	LUCKFRUITID_LUCKFRUIT9      LUCKFRUITID = 9
	LUCKFRUITID_LUCKFRUIT10     LUCKFRUITID = 10
	LUCKFRUITID_LUCKFRUIT11     LUCKFRUITID = 11
)

var LUCKFRUITID_name = map[int32]string{
	0:  "LUCKFRUITUnknow",
	1:  "LUCKFRUIT1",
	2:  "LUCKFRUIT2",
	3:  "LUCKFRUIT3",
	4:  "LUCKFRUIT4",
	5:  "LUCKFRUIT5",
	6:  "LUCKFRUIT6",
	7:  "LUCKFRUIT7",
	8:  "LUCKFRUIT8",
	9:  "LUCKFRUIT9",
	10: "LUCKFRUIT10",
	11: "LUCKFRUIT11",
}
var LUCKFRUITID_value = map[string]int32{
	"LUCKFRUITUnknow": 0,
	"LUCKFRUIT1":      1,
	"LUCKFRUIT2":      2,
	"LUCKFRUIT3":      3,
	"LUCKFRUIT4":      4,
	"LUCKFRUIT5":      5,
	"LUCKFRUIT6":      6,
	"LUCKFRUIT7":      7,
	"LUCKFRUIT8":      8,
	"LUCKFRUIT9":      9,
	"LUCKFRUIT10":     10,
	"LUCKFRUIT11":     11,
}

func (x LUCKFRUITID) String() string {
	return proto.EnumName(LUCKFRUITID_name, int32(x))
}
func (LUCKFRUITID) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

// 请求进入房间
type ENTER_GAME_LUCKFRUIT_REQ struct {
	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	RoomID    uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *ENTER_GAME_LUCKFRUIT_REQ) Reset()                    { *m = ENTER_GAME_LUCKFRUIT_REQ{} }
func (m *ENTER_GAME_LUCKFRUIT_REQ) String() string            { return proto.CompactTextString(m) }
func (*ENTER_GAME_LUCKFRUIT_REQ) ProtoMessage()               {}
func (*ENTER_GAME_LUCKFRUIT_REQ) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *ENTER_GAME_LUCKFRUIT_REQ) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *ENTER_GAME_LUCKFRUIT_REQ) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type ENTER_GAME_LUCKFRUIT_RES struct {
	RoomID    uint32   `protobuf:"varint,1,opt,name=RoomID" json:"RoomID,omitempty"`
	Basics    int64    `protobuf:"varint,2,opt,name=Basics" json:"Basics,omitempty"`
	Bonus     int64    `protobuf:"varint,3,opt,name=Bonus" json:"Bonus,omitempty"`
	LastBet   int64    `protobuf:"varint,4,opt,name=LastBet" json:"LastBet,omitempty"`
	Bets      []uint64 `protobuf:"varint,5,rep,packed,name=Bets" json:"Bets,omitempty"`
	FeeCount  int32    `protobuf:"varint,8,opt,name=FeeCount" json:"FeeCount,omitempty"`
	FeeProfit int64    `protobuf:"varint,9,opt,name=FeeProfit" json:"FeeProfit,omitempty"`
}

func (m *ENTER_GAME_LUCKFRUIT_RES) Reset()                    { *m = ENTER_GAME_LUCKFRUIT_RES{} }
func (m *ENTER_GAME_LUCKFRUIT_RES) String() string            { return proto.CompactTextString(m) }
func (*ENTER_GAME_LUCKFRUIT_RES) ProtoMessage()               {}
func (*ENTER_GAME_LUCKFRUIT_RES) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *ENTER_GAME_LUCKFRUIT_RES) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *ENTER_GAME_LUCKFRUIT_RES) GetBasics() int64 {
	if m != nil {
		return m.Basics
	}
	return 0
}

func (m *ENTER_GAME_LUCKFRUIT_RES) GetBonus() int64 {
	if m != nil {
		return m.Bonus
	}
	return 0
}

func (m *ENTER_GAME_LUCKFRUIT_RES) GetLastBet() int64 {
	if m != nil {
		return m.LastBet
	}
	return 0
}

func (m *ENTER_GAME_LUCKFRUIT_RES) GetBets() []uint64 {
	if m != nil {
		return m.Bets
	}
	return nil
}

func (m *ENTER_GAME_LUCKFRUIT_RES) GetFeeCount() int32 {
	if m != nil {
		return m.FeeCount
	}
	return 0
}

func (m *ENTER_GAME_LUCKFRUIT_RES) GetFeeProfit() int64 {
	if m != nil {
		return m.FeeProfit
	}
	return 0
}

// 请求退出房间
type LEAVE_GAME_LUCKFRUIT_REQ struct {
	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	RoomID    uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *LEAVE_GAME_LUCKFRUIT_REQ) Reset()                    { *m = LEAVE_GAME_LUCKFRUIT_REQ{} }
func (m *LEAVE_GAME_LUCKFRUIT_REQ) String() string            { return proto.CompactTextString(m) }
func (*LEAVE_GAME_LUCKFRUIT_REQ) ProtoMessage()               {}
func (*LEAVE_GAME_LUCKFRUIT_REQ) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *LEAVE_GAME_LUCKFRUIT_REQ) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *LEAVE_GAME_LUCKFRUIT_REQ) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type LEAVE_GAME_LUCKFRUIT_RES struct {
	Ret    uint32 `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	RoomID uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *LEAVE_GAME_LUCKFRUIT_RES) Reset()                    { *m = LEAVE_GAME_LUCKFRUIT_RES{} }
func (m *LEAVE_GAME_LUCKFRUIT_RES) String() string            { return proto.CompactTextString(m) }
func (*LEAVE_GAME_LUCKFRUIT_RES) ProtoMessage()               {}
func (*LEAVE_GAME_LUCKFRUIT_RES) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *LEAVE_GAME_LUCKFRUIT_RES) GetRet() uint32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *LEAVE_GAME_LUCKFRUIT_RES) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

// //////////////////////////////////////////// 游戏 /////////////////////////////////////////////
// 请求开始游戏
type START_LUCKFRUIT_REQ struct {
	Bet uint64 `protobuf:"varint,1,opt,name=Bet" json:"Bet,omitempty"`
}

func (m *START_LUCKFRUIT_REQ) Reset()                    { *m = START_LUCKFRUIT_REQ{} }
func (m *START_LUCKFRUIT_REQ) String() string            { return proto.CompactTextString(m) }
func (*START_LUCKFRUIT_REQ) ProtoMessage()               {}
func (*START_LUCKFRUIT_REQ) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *START_LUCKFRUIT_REQ) GetBet() uint64 {
	if m != nil {
		return m.Bet
	}
	return 0
}

type START_LUCKFRUIT_RES struct {
	Ret          uint64               `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	SumOdds      int64                `protobuf:"varint,2,opt,name=SumOdds" json:"SumOdds,omitempty"`
	Results      []*LUCKFRUIT_Result  `protobuf:"bytes,3,rep,name=Results" json:"Results,omitempty"`
	PictureList  []int32              `protobuf:"varint,4,rep,packed,name=PictureList" json:"PictureList,omitempty"`
	Bonus        int64                `protobuf:"varint,5,opt,name=Bonus" json:"Bonus,omitempty"`
	Money        int64                `protobuf:"varint,6,opt,name=Money" json:"Money,omitempty"`
	FreeCount    int64                `protobuf:"varint,7,opt,name=FreeCount" json:"FreeCount,omitempty"`
	FeePositions []*LUCKFRUITPosition `protobuf:"bytes,8,rep,name=FeePositions" json:"FeePositions,omitempty"`
	FeeProfit    int64                `protobuf:"varint,9,opt,name=FeeProfit" json:"FeeProfit,omitempty"`
}

func (m *START_LUCKFRUIT_RES) Reset()                    { *m = START_LUCKFRUIT_RES{} }
func (m *START_LUCKFRUIT_RES) String() string            { return proto.CompactTextString(m) }
func (*START_LUCKFRUIT_RES) ProtoMessage()               {}
func (*START_LUCKFRUIT_RES) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *START_LUCKFRUIT_RES) GetRet() uint64 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *START_LUCKFRUIT_RES) GetSumOdds() int64 {
	if m != nil {
		return m.SumOdds
	}
	return 0
}

func (m *START_LUCKFRUIT_RES) GetResults() []*LUCKFRUIT_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *START_LUCKFRUIT_RES) GetPictureList() []int32 {
	if m != nil {
		return m.PictureList
	}
	return nil
}

func (m *START_LUCKFRUIT_RES) GetBonus() int64 {
	if m != nil {
		return m.Bonus
	}
	return 0
}

func (m *START_LUCKFRUIT_RES) GetMoney() int64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *START_LUCKFRUIT_RES) GetFreeCount() int64 {
	if m != nil {
		return m.FreeCount
	}
	return 0
}

func (m *START_LUCKFRUIT_RES) GetFeePositions() []*LUCKFRUITPosition {
	if m != nil {
		return m.FeePositions
	}
	return nil
}

func (m *START_LUCKFRUIT_RES) GetFeeProfit() int64 {
	if m != nil {
		return m.FeeProfit
	}
	return 0
}

type LUCKFRUITPosition struct {
	Px int32 `protobuf:"varint,1,opt,name=px" json:"px,omitempty"`
	Py int32 `protobuf:"varint,2,opt,name=py" json:"py,omitempty"`
}

func (m *LUCKFRUITPosition) Reset()                    { *m = LUCKFRUITPosition{} }
func (m *LUCKFRUITPosition) String() string            { return proto.CompactTextString(m) }
func (*LUCKFRUITPosition) ProtoMessage()               {}
func (*LUCKFRUITPosition) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *LUCKFRUITPosition) GetPx() int32 {
	if m != nil {
		return m.Px
	}
	return 0
}

func (m *LUCKFRUITPosition) GetPy() int32 {
	if m != nil {
		return m.Py
	}
	return 0
}

type LUCKFRUIT_Result struct {
	LineId    int32                `protobuf:"varint,1,opt,name=LineId" json:"LineId,omitempty"`
	Count     int32                `protobuf:"varint,2,opt,name=Count" json:"Count,omitempty"`
	Odds      int32                `protobuf:"varint,3,opt,name=Odds" json:"Odds,omitempty"`
	Positions []*LUCKFRUITPosition `protobuf:"bytes,4,rep,name=Positions" json:"Positions,omitempty"`
}

func (m *LUCKFRUIT_Result) Reset()                    { *m = LUCKFRUIT_Result{} }
func (m *LUCKFRUIT_Result) String() string            { return proto.CompactTextString(m) }
func (*LUCKFRUIT_Result) ProtoMessage()               {}
func (*LUCKFRUIT_Result) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *LUCKFRUIT_Result) GetLineId() int32 {
	if m != nil {
		return m.LineId
	}
	return 0
}

func (m *LUCKFRUIT_Result) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *LUCKFRUIT_Result) GetOdds() int32 {
	if m != nil {
		return m.Odds
	}
	return 0
}

func (m *LUCKFRUIT_Result) GetPositions() []*LUCKFRUITPosition {
	if m != nil {
		return m.Positions
	}
	return nil
}

// 通知更新奖金池
type UPDATE_LUCKFRUIT_BONUS struct {
	Bonus int64 `protobuf:"varint,1,opt,name=Bonus" json:"Bonus,omitempty"`
}

func (m *UPDATE_LUCKFRUIT_BONUS) Reset()                    { *m = UPDATE_LUCKFRUIT_BONUS{} }
func (m *UPDATE_LUCKFRUIT_BONUS) String() string            { return proto.CompactTextString(m) }
func (*UPDATE_LUCKFRUIT_BONUS) ProtoMessage()               {}
func (*UPDATE_LUCKFRUIT_BONUS) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

func (m *UPDATE_LUCKFRUIT_BONUS) GetBonus() int64 {
	if m != nil {
		return m.Bonus
	}
	return 0
}

// 请求LUCKFRUIT玩家列表
type PLAYERS_LUCKFRUIT_LIST_RES struct {
	Players []*AccountStorageData `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
}

func (m *PLAYERS_LUCKFRUIT_LIST_RES) Reset()                    { *m = PLAYERS_LUCKFRUIT_LIST_RES{} }
func (m *PLAYERS_LUCKFRUIT_LIST_RES) String() string            { return proto.CompactTextString(m) }
func (*PLAYERS_LUCKFRUIT_LIST_RES) ProtoMessage()               {}
func (*PLAYERS_LUCKFRUIT_LIST_RES) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *PLAYERS_LUCKFRUIT_LIST_RES) GetPlayers() []*AccountStorageData {
	if m != nil {
		return m.Players
	}
	return nil
}

func init() {
	proto.RegisterType((*ENTER_GAME_LUCKFRUIT_REQ)(nil), "protomsg.ENTER_GAME_LUCKFRUIT_REQ")
	proto.RegisterType((*ENTER_GAME_LUCKFRUIT_RES)(nil), "protomsg.ENTER_GAME_LUCKFRUIT_RES")
	proto.RegisterType((*LEAVE_GAME_LUCKFRUIT_REQ)(nil), "protomsg.LEAVE_GAME_LUCKFRUIT_REQ")
	proto.RegisterType((*LEAVE_GAME_LUCKFRUIT_RES)(nil), "protomsg.LEAVE_GAME_LUCKFRUIT_RES")
	proto.RegisterType((*START_LUCKFRUIT_REQ)(nil), "protomsg.START_LUCKFRUIT_REQ")
	proto.RegisterType((*START_LUCKFRUIT_RES)(nil), "protomsg.START_LUCKFRUIT_RES")
	proto.RegisterType((*LUCKFRUITPosition)(nil), "protomsg.LUCKFRUIT_position")
	proto.RegisterType((*LUCKFRUIT_Result)(nil), "protomsg.LUCKFRUIT_Result")
	proto.RegisterType((*UPDATE_LUCKFRUIT_BONUS)(nil), "protomsg.UPDATE_LUCKFRUIT_BONUS")
	proto.RegisterType((*PLAYERS_LUCKFRUIT_LIST_RES)(nil), "protomsg.PLAYERS_LUCKFRUIT_LIST_RES")
	proto.RegisterEnum("protomsg.LUCKFRUITMSG", LUCKFRUITMSG_name, LUCKFRUITMSG_value)
	proto.RegisterEnum("protomsg.LUCKFRUITID", LUCKFRUITID_name, LUCKFRUITID_value)
}

func init() { proto.RegisterFile("protobuf/luckfruit.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4b, 0x6f, 0xda, 0x4c,
	0x14, 0x8d, 0x31, 0xe6, 0x71, 0xc9, 0x63, 0x34, 0x89, 0x22, 0x7f, 0x24, 0x5f, 0x85, 0xd8, 0x14,
	0x65, 0x41, 0x9a, 0x47, 0xd3, 0xc7, 0xaa, 0x60, 0x48, 0x84, 0x42, 0x08, 0x99, 0x81, 0x56, 0xed,
	0x06, 0x39, 0xe0, 0x44, 0x56, 0x12, 0x8c, 0xec, 0xb1, 0x1a, 0xfe, 0x45, 0xb7, 0x7d, 0x2c, 0xda,
	0x6e, 0xda, 0xff, 0xd1, 0x4d, 0x57, 0x5d, 0xf5, 0x07, 0x55, 0x33, 0xb6, 0xb1, 0x9d, 0x62, 0xb5,
	0x8b, 0xae, 0x98, 0x73, 0xe6, 0xdc, 0xe3, 0x99, 0x7b, 0xee, 0x00, 0xea, 0xc4, 0xb6, 0x98, 0x75,
	0xee, 0x5e, 0x6c, 0x5f, 0xbb, 0xc3, 0xab, 0x0b, 0xdb, 0x35, 0x59, 0x55, 0x50, 0x38, 0x27, 0x7e,
	0x6e, 0x9c, 0xcb, 0xe2, 0xea, 0x4c, 0x33, 0xd2, 0x99, 0xee, 0x6d, 0x97, 0xbb, 0xa0, 0x36, 0x3b,
	0xbd, 0x26, 0x19, 0x1c, 0xd5, 0x4e, 0x9a, 0x83, 0x76, 0x5f, 0x3b, 0x3e, 0x24, 0xfd, 0x56, 0x6f,
	0x40, 0x9a, 0x67, 0x78, 0x13, 0xf2, 0xb5, 0xe1, 0xd0, 0x72, 0xc7, 0xac, 0xd5, 0x50, 0xa5, 0x92,
	0x54, 0x59, 0x22, 0x21, 0x81, 0xd7, 0x21, 0x43, 0x2c, 0xeb, 0xa6, 0xd5, 0x50, 0x53, 0x62, 0xcb,
	0x47, 0xe5, 0xef, 0x52, 0xa2, 0x25, 0x8d, 0x14, 0x49, 0xd1, 0x22, 0xce, 0xd7, 0x75, 0xc7, 0x1c,
	0x3a, 0xc2, 0x4c, 0x26, 0x3e, 0xc2, 0x6b, 0xa0, 0xd4, 0xad, 0xb1, 0xeb, 0xa8, 0xb2, 0xa0, 0x3d,
	0x80, 0x55, 0xc8, 0xb6, 0x75, 0x87, 0xd5, 0x0d, 0xa6, 0xa6, 0x05, 0x1f, 0x40, 0x8c, 0x21, 0x5d,
	0x37, 0x98, 0xa3, 0x2a, 0x25, 0xb9, 0x92, 0x26, 0x62, 0x8d, 0x8b, 0x90, 0x3b, 0x34, 0x0c, 0x8d,
	0x1f, 0x5b, 0xcd, 0x95, 0xa4, 0x8a, 0x42, 0x66, 0x98, 0x5f, 0xf1, 0xd0, 0x30, 0xba, 0xb6, 0x75,
	0x61, 0x32, 0x35, 0x2f, 0xbc, 0x42, 0x82, 0x37, 0xa7, 0xdd, 0xac, 0x3d, 0x6f, 0xfe, 0xbb, 0xe6,
	0x34, 0x12, 0x1d, 0x29, 0x46, 0x20, 0x13, 0x83, 0xf9, 0x5e, 0x7c, 0x99, 0xe8, 0x72, 0x1f, 0x56,
	0x69, 0xaf, 0x46, 0x7a, 0x77, 0x8e, 0x84, 0x40, 0xae, 0xfb, 0x06, 0x69, 0xc2, 0x97, 0xe5, 0x6f,
	0xa9, 0x79, 0xca, 0xd8, 0xa7, 0xd2, 0xde, 0xa7, 0x54, 0xc8, 0x52, 0xf7, 0xe6, 0x74, 0x34, 0x0a,
	0x12, 0x08, 0x20, 0xde, 0x87, 0x2c, 0x31, 0x1c, 0xf7, 0x9a, 0xf1, 0x10, 0xe4, 0x4a, 0x61, 0xb7,
	0x58, 0x0d, 0x46, 0xaa, 0x1a, 0x71, 0x15, 0x12, 0x12, 0x48, 0x71, 0x09, 0x0a, 0x5d, 0x73, 0xc8,
	0x5c, 0xdb, 0x68, 0x9b, 0x0e, 0x8f, 0x49, 0xae, 0x28, 0x24, 0x4a, 0x85, 0xd1, 0x2a, 0xd1, 0x68,
	0xd7, 0x40, 0x39, 0xb1, 0xc6, 0xc6, 0x54, 0xcd, 0x78, 0xac, 0x00, 0x22, 0x26, 0x3b, 0xc8, 0x30,
	0xeb, 0xc7, 0x14, 0x10, 0xf8, 0x19, 0x2c, 0xf2, 0xcc, 0x2c, 0xc7, 0x64, 0xa6, 0x35, 0x76, 0xd4,
	0x9c, 0x38, 0xe6, 0xe6, 0xbc, 0x63, 0x4e, 0x7c, 0x11, 0x89, 0x55, 0xfc, 0x61, 0x0c, 0xf6, 0x01,
	0xff, 0xee, 0x80, 0x97, 0x21, 0x35, 0xb9, 0x15, 0x2d, 0x54, 0x48, 0x6a, 0x72, 0x2b, 0xf0, 0x54,
	0x34, 0x8f, 0xe3, 0x69, 0xf9, 0x8d, 0x04, 0xe8, 0x6e, 0x7f, 0x78, 0xa2, 0x6d, 0x73, 0x6c, 0xb4,
	0x46, 0x7e, 0xa1, 0x8f, 0xf8, 0xb5, 0xbd, 0xcb, 0x79, 0xf5, 0x1e, 0xe0, 0xd3, 0x2c, 0x12, 0x91,
	0x05, 0x29, 0xd6, 0xf8, 0x29, 0xe4, 0xc3, 0x9b, 0xa6, 0xff, 0xe2, 0xa6, 0xa1, 0xbc, 0x5c, 0x85,
	0xf5, 0x7e, 0xb7, 0x51, 0xeb, 0x45, 0x27, 0xaf, 0x7e, 0xda, 0xe9, 0xd3, 0x30, 0x0c, 0x29, 0x12,
	0x46, 0xb9, 0x07, 0xc5, 0x6e, 0xbb, 0xf6, 0xb2, 0x49, 0x68, 0xa4, 0xa0, 0xdd, 0xa2, 0xde, 0x10,
	0x1d, 0x40, 0x76, 0x72, 0xad, 0x4f, 0x0d, 0x9b, 0x57, 0xdd, 0x39, 0x87, 0xff, 0x12, 0x28, 0xb3,
	0x6c, 0xfd, 0xd2, 0x68, 0xe8, 0x4c, 0x27, 0x81, 0x78, 0xeb, 0x47, 0x0a, 0x16, 0x67, 0x76, 0x27,
	0xf4, 0x08, 0xaf, 0x01, 0xea, 0x77, 0x8e, 0x3b, 0xa7, 0x2f, 0xc2, 0xaf, 0xa0, 0x05, 0x5c, 0x82,
	0x0d, 0x8d, 0x0e, 0x92, 0xfe, 0x9c, 0xd0, 0x5b, 0x97, 0x2b, 0xa8, 0x96, 0xa4, 0xa0, 0xe8, 0x9d,
	0xeb, 0x7b, 0x24, 0xbd, 0x61, 0xf4, 0x3e, 0xf0, 0x48, 0x7a, 0x93, 0xe8, 0x83, 0x8b, 0x37, 0x60,
	0x5d, 0xa3, 0x83, 0x39, 0xef, 0x0d, 0x7d, 0x14, 0x9b, 0x54, 0x9b, 0xb3, 0x49, 0xd1, 0x27, 0x17,
	0xdf, 0x83, 0xff, 0xa8, 0x36, 0x98, 0xdf, 0x71, 0xf4, 0xd9, 0xc5, 0x65, 0xf8, 0x5f, 0xa3, 0x83,
	0xc4, 0x0e, 0x9f, 0xa1, 0x2f, 0x42, 0x43, 0xb5, 0x64, 0x0d, 0x45, 0x5f, 0xdd, 0xad, 0x9f, 0x12,
	0x14, 0x66, 0x3b, 0xad, 0x06, 0x5e, 0x85, 0x95, 0x19, 0xec, 0x8f, 0xaf, 0xc6, 0xd6, 0x6b, 0xb4,
	0x80, 0x97, 0x01, 0x66, 0xe4, 0x0e, 0x92, 0x62, 0x78, 0x17, 0xa5, 0x62, 0x78, 0x0f, 0xc9, 0x31,
	0xbc, 0x8f, 0xd2, 0x31, 0xfc, 0x10, 0x29, 0x31, 0x7c, 0x80, 0x32, 0x31, 0xfc, 0x08, 0x65, 0x63,
	0xf8, 0x31, 0xca, 0xc5, 0xf0, 0x13, 0x94, 0xc7, 0x2b, 0x91, 0x33, 0xef, 0x3c, 0x40, 0x10, 0x27,
	0x76, 0x50, 0xa1, 0xbe, 0xf2, 0x6a, 0xc9, 0xb6, 0x2c, 0xb6, 0x1d, 0x0c, 0xd5, 0x79, 0x46, 0xac,
	0xf6, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x60, 0x19, 0x89, 0xed, 0x06, 0x00, 0x00,
}
