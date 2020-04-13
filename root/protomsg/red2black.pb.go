// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/red2black.proto

package protomsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// //////////////////////////////////////////////////// 网络消息 /////////////////////////////////////////////////////////////
type RED2BLACKMSG int32

const (
	RED2BLACKMSG_UNKNOW_RED2BLACK                RED2BLACKMSG = 0
	RED2BLACKMSG_CS_ENTER_GAME_RED2BLACK_REQ     RED2BLACKMSG = 16001
	RED2BLACKMSG_SC_ENTER_GAME_RED2BLACK_RES     RED2BLACKMSG = 16002
	RED2BLACKMSG_CS_LEAVE_GAME_RED2BLACK_REQ     RED2BLACKMSG = 16003
	RED2BLACKMSG_SC_LEAVE_GAME_RED2BLACK_RES     RED2BLACKMSG = 16004
	RED2BLACKMSG_CS_BET_RED2BLACK_REQ            RED2BLACKMSG = 16008
	RED2BLACKMSG_SC_BET_RED2BLACK_RES            RED2BLACKMSG = 16009
	RED2BLACKMSG_CS_CLEAN_BET_RED2BLACK_REQ      RED2BLACKMSG = 16010
	RED2BLACKMSG_SC_CLEAN_BET_RED2BLACK_RES      RED2BLACKMSG = 16011
	RED2BLACKMSG_CS_PLAYERS_RED2BLACK_LIST_REQ   RED2BLACKMSG = 16015
	RED2BLACKMSG_SC_PLAYERS_RED2BLACK_LIST_RES   RED2BLACKMSG = 16016
	RED2BLACKMSG_SC_SWITCH_GAME_STATUS_BROADCAST RED2BLACKMSG = 16020
)

var RED2BLACKMSG_name = map[int32]string{
	0:     "UNKNOW_RED2BLACK",
	16001: "CS_ENTER_GAME_RED2BLACK_REQ",
	16002: "SC_ENTER_GAME_RED2BLACK_RES",
	16003: "CS_LEAVE_GAME_RED2BLACK_REQ",
	16004: "SC_LEAVE_GAME_RED2BLACK_RES",
	16008: "CS_BET_RED2BLACK_REQ",
	16009: "SC_BET_RED2BLACK_RES",
	16010: "CS_CLEAN_BET_RED2BLACK_REQ",
	16011: "SC_CLEAN_BET_RED2BLACK_RES",
	16015: "CS_PLAYERS_RED2BLACK_LIST_REQ",
	16016: "SC_PLAYERS_RED2BLACK_LIST_RES",
	16020: "SC_SWITCH_GAME_STATUS_BROADCAST",
}
var RED2BLACKMSG_value = map[string]int32{
	"UNKNOW_RED2BLACK":                0,
	"CS_ENTER_GAME_RED2BLACK_REQ":     16001,
	"SC_ENTER_GAME_RED2BLACK_RES":     16002,
	"CS_LEAVE_GAME_RED2BLACK_REQ":     16003,
	"SC_LEAVE_GAME_RED2BLACK_RES":     16004,
	"CS_BET_RED2BLACK_REQ":            16008,
	"SC_BET_RED2BLACK_RES":            16009,
	"CS_CLEAN_BET_RED2BLACK_REQ":      16010,
	"SC_CLEAN_BET_RED2BLACK_RES":      16011,
	"CS_PLAYERS_RED2BLACK_LIST_REQ":   16015,
	"SC_PLAYERS_RED2BLACK_LIST_RES":   16016,
	"SC_SWITCH_GAME_STATUS_BROADCAST": 16020,
}

func (x RED2BLACKMSG) String() string {
	return proto.EnumName(RED2BLACKMSG_name, int32(x))
}
func (RED2BLACKMSG) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

// //////////////////////////////////////////////////// 游戏枚举 /////////////////////////////////////////////////////////////
// 下注区域
type RED2BLACKAREA int32

const (
	RED2BLACKAREA_RED2BLACK_AREA_Unknow RED2BLACKAREA = 0
	RED2BLACKAREA_RED2BLACK_AREA_RED    RED2BLACKAREA = 1
	RED2BLACKAREA_RED2BLACK_AREA_BLACK  RED2BLACKAREA = 2
	RED2BLACKAREA_RED2BLACK_AREA_LUCK   RED2BLACKAREA = 3
)

var RED2BLACKAREA_name = map[int32]string{
	0: "RED2BLACK_AREA_Unknow",
	1: "RED2BLACK_AREA_RED",
	2: "RED2BLACK_AREA_BLACK",
	3: "RED2BLACK_AREA_LUCK",
}
var RED2BLACKAREA_value = map[string]int32{
	"RED2BLACK_AREA_Unknow": 0,
	"RED2BLACK_AREA_RED":    1,
	"RED2BLACK_AREA_BLACK":  2,
	"RED2BLACK_AREA_LUCK":   3,
}

func (x RED2BLACKAREA) String() string {
	return proto.EnumName(RED2BLACKAREA_name, int32(x))
}
func (RED2BLACKAREA) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

// 牌型
type RED2BLACKCARDTYPE int32

const (
	RED2BLACKCARDTYPE_RED2BLACK_CARDTYPE_Unknow RED2BLACKCARDTYPE = 0
	RED2BLACKCARDTYPE_RED2BLACK_CARDTYPE_1      RED2BLACKCARDTYPE = 1
	RED2BLACKCARDTYPE_RED2BLACK_CARDTYPE_2      RED2BLACKCARDTYPE = 2
	RED2BLACKCARDTYPE_RED2BLACK_CARDTYPE_3      RED2BLACKCARDTYPE = 3
	RED2BLACKCARDTYPE_RED2BLACK_CARDTYPE_4      RED2BLACKCARDTYPE = 4
	RED2BLACKCARDTYPE_RED2BLACK_CARDTYPE_5      RED2BLACKCARDTYPE = 5
	RED2BLACKCARDTYPE_RED2BLACK_CARDTYPE_6      RED2BLACKCARDTYPE = 6
	RED2BLACKCARDTYPE_RED2BLACK_CARDTYPE_7      RED2BLACKCARDTYPE = 7
)

var RED2BLACKCARDTYPE_name = map[int32]string{
	0: "RED2BLACK_CARDTYPE_Unknow",
	1: "RED2BLACK_CARDTYPE_1",
	2: "RED2BLACK_CARDTYPE_2",
	3: "RED2BLACK_CARDTYPE_3",
	4: "RED2BLACK_CARDTYPE_4",
	5: "RED2BLACK_CARDTYPE_5",
	6: "RED2BLACK_CARDTYPE_6",
	7: "RED2BLACK_CARDTYPE_7",
}
var RED2BLACKCARDTYPE_value = map[string]int32{
	"RED2BLACK_CARDTYPE_Unknow": 0,
	"RED2BLACK_CARDTYPE_1":      1,
	"RED2BLACK_CARDTYPE_2":      2,
	"RED2BLACK_CARDTYPE_3":      3,
	"RED2BLACK_CARDTYPE_4":      4,
	"RED2BLACK_CARDTYPE_5":      5,
	"RED2BLACK_CARDTYPE_6":      6,
	"RED2BLACK_CARDTYPE_7":      7,
}

func (x RED2BLACKCARDTYPE) String() string {
	return proto.EnumName(RED2BLACKCARDTYPE_name, int32(x))
}
func (RED2BLACKCARDTYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

// 游戏状态
type RED2BLACKGAMESTATUS int32

const (
	RED2BLACKGAMESTATUS_RED2BLACK_GAMESTATUS_Unknow RED2BLACKGAMESTATUS = 0
	RED2BLACKGAMESTATUS_RED2BLACK_GAMESTATUS_WAIT   RED2BLACKGAMESTATUS = 1
	RED2BLACKGAMESTATUS_RED2BLACK_GAMESTATUS_BET    RED2BLACKGAMESTATUS = 2
	RED2BLACKGAMESTATUS_RED2BLACK_GAMESTATUS_STOP   RED2BLACKGAMESTATUS = 3
	RED2BLACKGAMESTATUS_RED2BLACK_GAMESTATUS_SETTLE RED2BLACKGAMESTATUS = 4
)

var RED2BLACKGAMESTATUS_name = map[int32]string{
	0: "RED2BLACK_GAMESTATUS_Unknow",
	1: "RED2BLACK_GAMESTATUS_WAIT",
	2: "RED2BLACK_GAMESTATUS_BET",
	3: "RED2BLACK_GAMESTATUS_STOP",
	4: "RED2BLACK_GAMESTATUS_SETTLE",
}
var RED2BLACKGAMESTATUS_value = map[string]int32{
	"RED2BLACK_GAMESTATUS_Unknow": 0,
	"RED2BLACK_GAMESTATUS_WAIT":   1,
	"RED2BLACK_GAMESTATUS_BET":    2,
	"RED2BLACK_GAMESTATUS_STOP":   3,
	"RED2BLACK_GAMESTATUS_SETTLE": 4,
}

func (x RED2BLACKGAMESTATUS) String() string {
	return proto.EnumName(RED2BLACKGAMESTATUS_name, int32(x))
}
func (RED2BLACKGAMESTATUS) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

// //////////////////////////////////////////////////// 消息数据结构 /////////////////////////////////////////////////////////////
// 请求进入房间
type ENTER_GAME_RED2BLACK_REQ struct {
	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	RoomID    uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *ENTER_GAME_RED2BLACK_REQ) Reset()                    { *m = ENTER_GAME_RED2BLACK_REQ{} }
func (m *ENTER_GAME_RED2BLACK_REQ) String() string            { return proto.CompactTextString(m) }
func (*ENTER_GAME_RED2BLACK_REQ) ProtoMessage()               {}
func (*ENTER_GAME_RED2BLACK_REQ) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *ENTER_GAME_RED2BLACK_REQ) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *ENTER_GAME_RED2BLACK_REQ) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

// 进入房间，服务器返回房间、游戏状态相关数据
type ENTER_GAME_RED2BLACK_RES struct {
	RoomID         uint32                             `protobuf:"varint,1,opt,name=RoomID" json:"RoomID,omitempty"`
	HistoryWinners []*ENTER_GAME_RED2BLACK_RES_Winner `protobuf:"bytes,2,rep,name=HistoryWinners" json:"HistoryWinners,omitempty"`
	Bets           []int64                            `protobuf:"varint,3,rep,packed,name=bets" json:"bets,omitempty"`
	ShowNum        uint32                             `protobuf:"varint,4,opt,name=ShowNum" json:"ShowNum,omitempty"`
	BetLimit       uint64                             `protobuf:"varint,5,opt,name=BetLimit" json:"BetLimit,omitempty"`
	Status         *StatusMsg                         `protobuf:"bytes,6,opt,name=Status" json:"Status,omitempty"`
}

func (m *ENTER_GAME_RED2BLACK_RES) Reset()                    { *m = ENTER_GAME_RED2BLACK_RES{} }
func (m *ENTER_GAME_RED2BLACK_RES) String() string            { return proto.CompactTextString(m) }
func (*ENTER_GAME_RED2BLACK_RES) ProtoMessage()               {}
func (*ENTER_GAME_RED2BLACK_RES) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *ENTER_GAME_RED2BLACK_RES) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *ENTER_GAME_RED2BLACK_RES) GetHistoryWinners() []*ENTER_GAME_RED2BLACK_RES_Winner {
	if m != nil {
		return m.HistoryWinners
	}
	return nil
}

func (m *ENTER_GAME_RED2BLACK_RES) GetBets() []int64 {
	if m != nil {
		return m.Bets
	}
	return nil
}

func (m *ENTER_GAME_RED2BLACK_RES) GetShowNum() uint32 {
	if m != nil {
		return m.ShowNum
	}
	return 0
}

func (m *ENTER_GAME_RED2BLACK_RES) GetBetLimit() uint64 {
	if m != nil {
		return m.BetLimit
	}
	return 0
}

func (m *ENTER_GAME_RED2BLACK_RES) GetStatus() *StatusMsg {
	if m != nil {
		return m.Status
	}
	return nil
}

type ENTER_GAME_RED2BLACK_RES_Winner struct {
	WinArea     RED2BLACKAREA     `protobuf:"varint,1,opt,name=WinArea,enum=protomsg.RED2BLACKAREA" json:"WinArea,omitempty"`
	WinCardType RED2BLACKCARDTYPE `protobuf:"varint,2,opt,name=WinCardType,enum=protomsg.RED2BLACKCARDTYPE" json:"WinCardType,omitempty"`
}

func (m *ENTER_GAME_RED2BLACK_RES_Winner) Reset()         { *m = ENTER_GAME_RED2BLACK_RES_Winner{} }
func (m *ENTER_GAME_RED2BLACK_RES_Winner) String() string { return proto.CompactTextString(m) }
func (*ENTER_GAME_RED2BLACK_RES_Winner) ProtoMessage()    {}
func (*ENTER_GAME_RED2BLACK_RES_Winner) Descriptor() ([]byte, []int) {
	return fileDescriptor9, []int{1, 0}
}

func (m *ENTER_GAME_RED2BLACK_RES_Winner) GetWinArea() RED2BLACKAREA {
	if m != nil {
		return m.WinArea
	}
	return RED2BLACKAREA_RED2BLACK_AREA_Unknow
}

func (m *ENTER_GAME_RED2BLACK_RES_Winner) GetWinCardType() RED2BLACKCARDTYPE {
	if m != nil {
		return m.WinCardType
	}
	return RED2BLACKCARDTYPE_RED2BLACK_CARDTYPE_Unknow
}

// 请求退出房间
type LEAVE_GAME_RED2BLACK_REQ struct {
	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	RoomID    uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *LEAVE_GAME_RED2BLACK_REQ) Reset()                    { *m = LEAVE_GAME_RED2BLACK_REQ{} }
func (m *LEAVE_GAME_RED2BLACK_REQ) String() string            { return proto.CompactTextString(m) }
func (*LEAVE_GAME_RED2BLACK_REQ) ProtoMessage()               {}
func (*LEAVE_GAME_RED2BLACK_REQ) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *LEAVE_GAME_RED2BLACK_REQ) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *LEAVE_GAME_RED2BLACK_REQ) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type LEAVE_GAME_RED2BLACK_RES struct {
	Ret    uint32 `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	RoomID uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *LEAVE_GAME_RED2BLACK_RES) Reset()                    { *m = LEAVE_GAME_RED2BLACK_RES{} }
func (m *LEAVE_GAME_RED2BLACK_RES) String() string            { return proto.CompactTextString(m) }
func (*LEAVE_GAME_RED2BLACK_RES) ProtoMessage()               {}
func (*LEAVE_GAME_RED2BLACK_RES) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *LEAVE_GAME_RED2BLACK_RES) GetRet() uint32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *LEAVE_GAME_RED2BLACK_RES) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

// //////////////////////////////////////////// 游戏 /////////////////////////////////////////////
// 服务器广播切换状态
type SWITCH_GAME_STATUS_BROADCAST struct {
	NextStatus *StatusMsg `protobuf:"bytes,1,opt,name=NextStatus" json:"NextStatus,omitempty"`
}

func (m *SWITCH_GAME_STATUS_BROADCAST) Reset()                    { *m = SWITCH_GAME_STATUS_BROADCAST{} }
func (m *SWITCH_GAME_STATUS_BROADCAST) String() string            { return proto.CompactTextString(m) }
func (*SWITCH_GAME_STATUS_BROADCAST) ProtoMessage()               {}
func (*SWITCH_GAME_STATUS_BROADCAST) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *SWITCH_GAME_STATUS_BROADCAST) GetNextStatus() *StatusMsg {
	if m != nil {
		return m.NextStatus
	}
	return nil
}

// 状态结构
type StatusMsg struct {
	Status           RED2BLACKGAMESTATUS `protobuf:"varint,1,opt,name=Status,enum=protomsg.RED2BLACKGAMESTATUS" json:"Status,omitempty"`
	Status_StartTime uint64              `protobuf:"varint,2,opt,name=Status_StartTime,json=StatusStartTime" json:"Status_StartTime,omitempty"`
	Status_EndTime   uint64              `protobuf:"varint,3,opt,name=Status_EndTime,json=StatusEndTime" json:"Status_EndTime,omitempty"`
	RedCards         []*Card             `protobuf:"bytes,4,rep,name=RedCards" json:"RedCards,omitempty"`
	BlackCards       []*Card             `protobuf:"bytes,5,rep,name=BlackCards" json:"BlackCards,omitempty"`
	AreaBetVal       map[int32]int64     `protobuf:"bytes,6,rep,name=AreaBetVal" json:"AreaBetVal,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	AreaBetVal_Own   map[int32]int64     `protobuf:"bytes,7,rep,name=AreaBetVal_Own,json=AreaBetValOwn" json:"AreaBetVal_Own,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Status_Data      []byte              `protobuf:"bytes,8,opt,name=Status_Data,json=StatusData,proto3" json:"Status_Data,omitempty"`
}

func (m *StatusMsg) Reset()                    { *m = StatusMsg{} }
func (m *StatusMsg) String() string            { return proto.CompactTextString(m) }
func (*StatusMsg) ProtoMessage()               {}
func (*StatusMsg) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

func (m *StatusMsg) GetStatus() RED2BLACKGAMESTATUS {
	if m != nil {
		return m.Status
	}
	return RED2BLACKGAMESTATUS_RED2BLACK_GAMESTATUS_Unknow
}

func (m *StatusMsg) GetStatus_StartTime() uint64 {
	if m != nil {
		return m.Status_StartTime
	}
	return 0
}

func (m *StatusMsg) GetStatus_EndTime() uint64 {
	if m != nil {
		return m.Status_EndTime
	}
	return 0
}

func (m *StatusMsg) GetRedCards() []*Card {
	if m != nil {
		return m.RedCards
	}
	return nil
}

func (m *StatusMsg) GetBlackCards() []*Card {
	if m != nil {
		return m.BlackCards
	}
	return nil
}

func (m *StatusMsg) GetAreaBetVal() map[int32]int64 {
	if m != nil {
		return m.AreaBetVal
	}
	return nil
}

func (m *StatusMsg) GetAreaBetVal_Own() map[int32]int64 {
	if m != nil {
		return m.AreaBetVal_Own
	}
	return nil
}

func (m *StatusMsg) GetStatus_Data() []byte {
	if m != nil {
		return m.Status_Data
	}
	return nil
}

// 1.等待
type Status_Wait struct {
}

func (m *Status_Wait) Reset()                    { *m = Status_Wait{} }
func (m *Status_Wait) String() string            { return proto.CompactTextString(m) }
func (*Status_Wait) ProtoMessage()               {}
func (*Status_Wait) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{6} }

// 2.押注
type Status_Bet struct {
}

func (m *Status_Bet) Reset()                    { *m = Status_Bet{} }
func (m *Status_Bet) String() string            { return proto.CompactTextString(m) }
func (*Status_Bet) ProtoMessage()               {}
func (*Status_Bet) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{7} }

// 3.停止下注
type Status_Stop struct {
}

func (m *Status_Stop) Reset()                    { *m = Status_Stop{} }
func (m *Status_Stop) String() string            { return proto.CompactTextString(m) }
func (*Status_Stop) ProtoMessage()               {}
func (*Status_Stop) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{8} }

// 4.结算
type Status_Settle struct {
	WinArea      RED2BLACKAREA     `protobuf:"varint,1,opt,name=WinArea,enum=protomsg.RED2BLACKAREA" json:"WinArea,omitempty"`
	WinCardType  RED2BLACKCARDTYPE `protobuf:"varint,2,opt,name=WinCardType,enum=protomsg.RED2BLACKCARDTYPE" json:"WinCardType,omitempty"`
	LossCardType RED2BLACKCARDTYPE `protobuf:"varint,3,opt,name=LossCardType,enum=protomsg.RED2BLACKCARDTYPE" json:"LossCardType,omitempty"`
	WinOdds      uint64            `protobuf:"varint,4,opt,name=winOdds" json:"winOdds,omitempty"`
	Players      map[int32]int64   `protobuf:"bytes,5,rep,name=Players" json:"Players,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *Status_Settle) Reset()                    { *m = Status_Settle{} }
func (m *Status_Settle) String() string            { return proto.CompactTextString(m) }
func (*Status_Settle) ProtoMessage()               {}
func (*Status_Settle) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{9} }

func (m *Status_Settle) GetWinArea() RED2BLACKAREA {
	if m != nil {
		return m.WinArea
	}
	return RED2BLACKAREA_RED2BLACK_AREA_Unknow
}

func (m *Status_Settle) GetWinCardType() RED2BLACKCARDTYPE {
	if m != nil {
		return m.WinCardType
	}
	return RED2BLACKCARDTYPE_RED2BLACK_CARDTYPE_Unknow
}

func (m *Status_Settle) GetLossCardType() RED2BLACKCARDTYPE {
	if m != nil {
		return m.LossCardType
	}
	return RED2BLACKCARDTYPE_RED2BLACK_CARDTYPE_Unknow
}

func (m *Status_Settle) GetWinOdds() uint64 {
	if m != nil {
		return m.WinOdds
	}
	return 0
}

func (m *Status_Settle) GetPlayers() map[int32]int64 {
	if m != nil {
		return m.Players
	}
	return nil
}

// 请求下注
type BET_RED2BLACK_REQ struct {
	AccountID uint32        `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	Area      RED2BLACKAREA `protobuf:"varint,2,opt,name=Area,enum=protomsg.RED2BLACKAREA" json:"Area,omitempty"`
	Bet       uint64        `protobuf:"varint,3,opt,name=Bet" json:"Bet,omitempty"`
	BetType   uint32        `protobuf:"varint,4,opt,name=BetType" json:"BetType,omitempty"`
}

func (m *BET_RED2BLACK_REQ) Reset()                    { *m = BET_RED2BLACK_REQ{} }
func (m *BET_RED2BLACK_REQ) String() string            { return proto.CompactTextString(m) }
func (*BET_RED2BLACK_REQ) ProtoMessage()               {}
func (*BET_RED2BLACK_REQ) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{10} }

func (m *BET_RED2BLACK_REQ) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *BET_RED2BLACK_REQ) GetArea() RED2BLACKAREA {
	if m != nil {
		return m.Area
	}
	return RED2BLACKAREA_RED2BLACK_AREA_Unknow
}

func (m *BET_RED2BLACK_REQ) GetBet() uint64 {
	if m != nil {
		return m.Bet
	}
	return 0
}

func (m *BET_RED2BLACK_REQ) GetBetType() uint32 {
	if m != nil {
		return m.BetType
	}
	return 0
}

// 间隔 200 毫秒一次 广播玩家下注
type BET_RED2BLACK_RES struct {
	Players    []*BET_RED2BLACK_RES_BetPlayer `protobuf:"bytes,1,rep,name=Players" json:"Players,omitempty"`
	AreaBetVal map[int32]int64                `protobuf:"bytes,2,rep,name=AreaBetVal" json:"AreaBetVal,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *BET_RED2BLACK_RES) Reset()                    { *m = BET_RED2BLACK_RES{} }
func (m *BET_RED2BLACK_RES) String() string            { return proto.CompactTextString(m) }
func (*BET_RED2BLACK_RES) ProtoMessage()               {}
func (*BET_RED2BLACK_RES) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{11} }

func (m *BET_RED2BLACK_RES) GetPlayers() []*BET_RED2BLACK_RES_BetPlayer {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *BET_RED2BLACK_RES) GetAreaBetVal() map[int32]int64 {
	if m != nil {
		return m.AreaBetVal
	}
	return nil
}

type BET_RED2BLACK_RES_BetPlayer struct {
	AccountID uint32        `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	Area      RED2BLACKAREA `protobuf:"varint,2,opt,name=Area,enum=protomsg.RED2BLACKAREA" json:"Area,omitempty"`
	Bet       uint64        `protobuf:"varint,3,opt,name=Bet" json:"Bet,omitempty"`
}

func (m *BET_RED2BLACK_RES_BetPlayer) Reset()                    { *m = BET_RED2BLACK_RES_BetPlayer{} }
func (m *BET_RED2BLACK_RES_BetPlayer) String() string            { return proto.CompactTextString(m) }
func (*BET_RED2BLACK_RES_BetPlayer) ProtoMessage()               {}
func (*BET_RED2BLACK_RES_BetPlayer) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{11, 0} }

func (m *BET_RED2BLACK_RES_BetPlayer) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *BET_RED2BLACK_RES_BetPlayer) GetArea() RED2BLACKAREA {
	if m != nil {
		return m.Area
	}
	return RED2BLACKAREA_RED2BLACK_AREA_Unknow
}

func (m *BET_RED2BLACK_RES_BetPlayer) GetBet() uint64 {
	if m != nil {
		return m.Bet
	}
	return 0
}

// 请求清除下注
type CLEAN_BET_RED2BLACK_REQ struct {
	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
}

func (m *CLEAN_BET_RED2BLACK_REQ) Reset()                    { *m = CLEAN_BET_RED2BLACK_REQ{} }
func (m *CLEAN_BET_RED2BLACK_REQ) String() string            { return proto.CompactTextString(m) }
func (*CLEAN_BET_RED2BLACK_REQ) ProtoMessage()               {}
func (*CLEAN_BET_RED2BLACK_REQ) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{12} }

func (m *CLEAN_BET_RED2BLACK_REQ) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

// 服务器广播清除下注
type CLEAN_BET_RED2BLACK_RES struct {
	AccountID        uint32          `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	PlayerAreaBetVal map[int32]int64 `protobuf:"bytes,2,rep,name=PlayerAreaBetVal" json:"PlayerAreaBetVal,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	AreaBetVal       map[int32]int64 `protobuf:"bytes,3,rep,name=AreaBetVal" json:"AreaBetVal,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *CLEAN_BET_RED2BLACK_RES) Reset()                    { *m = CLEAN_BET_RED2BLACK_RES{} }
func (m *CLEAN_BET_RED2BLACK_RES) String() string            { return proto.CompactTextString(m) }
func (*CLEAN_BET_RED2BLACK_RES) ProtoMessage()               {}
func (*CLEAN_BET_RED2BLACK_RES) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{13} }

func (m *CLEAN_BET_RED2BLACK_RES) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *CLEAN_BET_RED2BLACK_RES) GetPlayerAreaBetVal() map[int32]int64 {
	if m != nil {
		return m.PlayerAreaBetVal
	}
	return nil
}

func (m *CLEAN_BET_RED2BLACK_RES) GetAreaBetVal() map[int32]int64 {
	if m != nil {
		return m.AreaBetVal
	}
	return nil
}

// 请求RED2BLACK玩家列表
type PLAYERS_RED2BLACK_LIST_RES struct {
	Players []*AccountStorageData `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
}

func (m *PLAYERS_RED2BLACK_LIST_RES) Reset()                    { *m = PLAYERS_RED2BLACK_LIST_RES{} }
func (m *PLAYERS_RED2BLACK_LIST_RES) String() string            { return proto.CompactTextString(m) }
func (*PLAYERS_RED2BLACK_LIST_RES) ProtoMessage()               {}
func (*PLAYERS_RED2BLACK_LIST_RES) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{14} }

func (m *PLAYERS_RED2BLACK_LIST_RES) GetPlayers() []*AccountStorageData {
	if m != nil {
		return m.Players
	}
	return nil
}

func init() {
	proto.RegisterType((*ENTER_GAME_RED2BLACK_REQ)(nil), "protomsg.ENTER_GAME_RED2BLACK_REQ")
	proto.RegisterType((*ENTER_GAME_RED2BLACK_RES)(nil), "protomsg.ENTER_GAME_RED2BLACK_RES")
	proto.RegisterType((*ENTER_GAME_RED2BLACK_RES_Winner)(nil), "protomsg.ENTER_GAME_RED2BLACK_RES.Winner")
	proto.RegisterType((*LEAVE_GAME_RED2BLACK_REQ)(nil), "protomsg.LEAVE_GAME_RED2BLACK_REQ")
	proto.RegisterType((*LEAVE_GAME_RED2BLACK_RES)(nil), "protomsg.LEAVE_GAME_RED2BLACK_RES")
	proto.RegisterType((*SWITCH_GAME_STATUS_BROADCAST)(nil), "protomsg.SWITCH_GAME_STATUS_BROADCAST")
	proto.RegisterType((*StatusMsg)(nil), "protomsg.StatusMsg")
	proto.RegisterType((*Status_Wait)(nil), "protomsg.Status_Wait")
	proto.RegisterType((*Status_Bet)(nil), "protomsg.Status_Bet")
	proto.RegisterType((*Status_Stop)(nil), "protomsg.Status_Stop")
	proto.RegisterType((*Status_Settle)(nil), "protomsg.Status_Settle")
	proto.RegisterType((*BET_RED2BLACK_REQ)(nil), "protomsg.BET_RED2BLACK_REQ")
	proto.RegisterType((*BET_RED2BLACK_RES)(nil), "protomsg.BET_RED2BLACK_RES")
	proto.RegisterType((*BET_RED2BLACK_RES_BetPlayer)(nil), "protomsg.BET_RED2BLACK_RES.BetPlayer")
	proto.RegisterType((*CLEAN_BET_RED2BLACK_REQ)(nil), "protomsg.CLEAN_BET_RED2BLACK_REQ")
	proto.RegisterType((*CLEAN_BET_RED2BLACK_RES)(nil), "protomsg.CLEAN_BET_RED2BLACK_RES")
	proto.RegisterType((*PLAYERS_RED2BLACK_LIST_RES)(nil), "protomsg.PLAYERS_RED2BLACK_LIST_RES")
	proto.RegisterEnum("protomsg.RED2BLACKMSG", RED2BLACKMSG_name, RED2BLACKMSG_value)
	proto.RegisterEnum("protomsg.RED2BLACKAREA", RED2BLACKAREA_name, RED2BLACKAREA_value)
	proto.RegisterEnum("protomsg.RED2BLACKCARDTYPE", RED2BLACKCARDTYPE_name, RED2BLACKCARDTYPE_value)
	proto.RegisterEnum("protomsg.RED2BLACKGAMESTATUS", RED2BLACKGAMESTATUS_name, RED2BLACKGAMESTATUS_value)
}

func init() { proto.RegisterFile("protobuf/red2black.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 1171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdf, 0x4e, 0xe3, 0xc6,
	0x17, 0x5e, 0xc7, 0x21, 0x81, 0x03, 0x09, 0xb3, 0x03, 0xbb, 0x98, 0x00, 0x22, 0xca, 0x6f, 0xf7,
	0xa7, 0x2c, 0x48, 0x59, 0x11, 0xba, 0x4b, 0xb5, 0xd2, 0x76, 0x6b, 0x3b, 0xd6, 0x2e, 0x22, 0x24,
	0x30, 0x63, 0x36, 0xda, 0xde, 0x58, 0x86, 0xb8, 0x6c, 0x04, 0xd8, 0x28, 0x99, 0x94, 0x52, 0x29,
	0x17, 0xfd, 0x73, 0xd1, 0x6e, 0x2f, 0x5a, 0x55, 0x7d, 0x86, 0xbe, 0x40, 0xdf, 0xa2, 0x4f, 0xd0,
	0x47, 0xe8, 0x7d, 0x6f, 0x7a, 0x59, 0x79, 0xec, 0xd8, 0xce, 0x1f, 0x07, 0xa1, 0x6d, 0x7b, 0x95,
	0x99, 0x73, 0xbe, 0xf3, 0x9d, 0x33, 0xdf, 0x99, 0x39, 0x31, 0x48, 0x97, 0x6d, 0x87, 0x39, 0xc7,
	0xdd, 0x4f, 0x1f, 0xb7, 0xad, 0x66, 0xf9, 0xf8, 0xdc, 0x3c, 0x39, 0x2b, 0x71, 0x13, 0x9e, 0xe6,
	0x3f, 0x17, 0x9d, 0xd3, 0xdc, 0x42, 0x80, 0x69, 0x9a, 0xcc, 0xf4, 0xdc, 0x85, 0x03, 0x90, 0xb4,
	0x9a, 0xae, 0x11, 0xe3, 0xa5, 0xbc, 0xaf, 0x19, 0x44, 0xab, 0x94, 0x95, 0xaa, 0xac, 0xee, 0x19,
	0x44, 0x3b, 0xc4, 0xab, 0x30, 0x23, 0x9f, 0x9c, 0x38, 0x5d, 0x9b, 0xed, 0x56, 0x24, 0x21, 0x2f,
	0x14, 0x33, 0x24, 0x34, 0xe0, 0xfb, 0x90, 0x22, 0x8e, 0x73, 0xb1, 0x5b, 0x91, 0x12, 0xdc, 0xe5,
	0xef, 0x0a, 0x7f, 0x25, 0x62, 0x29, 0x69, 0x24, 0x48, 0x88, 0x06, 0xe1, 0x43, 0xc8, 0xbe, 0x6a,
	0x75, 0x98, 0xd3, 0xbe, 0x6e, 0xb4, 0x6c, 0xdb, 0x6a, 0x77, 0xa4, 0x44, 0x5e, 0x2c, 0xce, 0x96,
	0x1f, 0x95, 0xfa, 0xe5, 0x97, 0xe2, 0x38, 0x4b, 0x5e, 0x04, 0x19, 0x22, 0xc0, 0x18, 0x92, 0xc7,
	0x16, 0xeb, 0x48, 0x62, 0x5e, 0x2c, 0x8a, 0x84, 0xaf, 0xb1, 0x04, 0x69, 0xfa, 0xd6, 0xb9, 0xaa,
	0x75, 0x2f, 0xa4, 0x24, 0xcf, 0xdf, 0xdf, 0xe2, 0x1c, 0x4c, 0x2b, 0x16, 0xab, 0xb6, 0x2e, 0x5a,
	0x4c, 0x9a, 0xca, 0x0b, 0xc5, 0x24, 0x09, 0xf6, 0x78, 0x13, 0x52, 0x94, 0x99, 0xac, 0xdb, 0x91,
	0x52, 0x79, 0xa1, 0x38, 0x5b, 0x5e, 0x08, 0x8b, 0xf2, 0xec, 0xfb, 0x9d, 0x53, 0xe2, 0x43, 0x72,
	0x5f, 0x40, 0xca, 0xab, 0x00, 0x6f, 0x41, 0xba, 0xd1, 0xb2, 0xe5, 0xb6, 0x65, 0xf2, 0xc3, 0x66,
	0xcb, 0x4b, 0x61, 0x5c, 0x70, 0x02, 0x99, 0x68, 0x32, 0xe9, 0xe3, 0xf0, 0x73, 0x98, 0x6d, 0xb4,
	0x6c, 0xd5, 0x6c, 0x37, 0xf5, 0xeb, 0x4b, 0x8b, 0x0b, 0x9b, 0x2d, 0xaf, 0x8c, 0x09, 0x53, 0x65,
	0x52, 0xd1, 0xdf, 0x1c, 0x68, 0x24, 0x8a, 0x77, 0x9b, 0x59, 0xd5, 0xe4, 0xd7, 0xda, 0x3f, 0xd7,
	0xcc, 0x4a, 0x2c, 0x23, 0xc5, 0x08, 0x44, 0x62, 0x31, 0x9f, 0xcb, 0x5d, 0xc6, 0xb2, 0x50, 0x58,
	0xa5, 0x8d, 0x5d, 0x5d, 0x7d, 0xe5, 0xd1, 0x50, 0x5d, 0xd6, 0x8f, 0xa8, 0xa1, 0x90, 0xba, 0x5c,
	0x51, 0x65, 0xaa, 0xe3, 0x6d, 0x80, 0x9a, 0xf5, 0x39, 0xf3, 0x45, 0x16, 0xe2, 0x45, 0x8e, 0xc0,
	0x0a, 0xbf, 0x24, 0x61, 0x26, 0xf0, 0xe0, 0x27, 0x41, 0x8f, 0x3c, 0xad, 0xd7, 0xc6, 0x88, 0xe6,
	0x26, 0xf7, 0x72, 0xf7, 0xbb, 0x85, 0x1f, 0x01, 0xf2, 0x56, 0x06, 0x65, 0x66, 0x9b, 0xe9, 0xad,
	0x0b, 0x4f, 0xf5, 0x24, 0x99, 0xf7, 0xec, 0x81, 0x19, 0x3f, 0x84, 0xac, 0x0f, 0xd5, 0xec, 0x26,
	0x07, 0x8a, 0x1c, 0x98, 0xf1, 0xac, 0xbe, 0x11, 0x6f, 0xc0, 0x34, 0xb1, 0x9a, 0x6e, 0x4b, 0x3a,
	0x52, 0x92, 0xdf, 0xe1, 0x6c, 0x58, 0x8a, 0x6b, 0x26, 0x81, 0x1f, 0x97, 0x00, 0x14, 0xf7, 0xa9,
	0x7a, 0xe8, 0xa9, 0xb1, 0xe8, 0x08, 0x02, 0xab, 0x00, 0xee, 0x35, 0x51, 0x2c, 0xf6, 0xda, 0x3c,
	0x97, 0x52, 0x1c, 0xff, 0xbf, 0x31, 0x3a, 0x95, 0x42, 0x94, 0x66, 0xb3, 0xf6, 0x35, 0x89, 0x84,
	0xe1, 0x7d, 0xc8, 0x86, 0x3b, 0xa3, 0x7e, 0x65, 0x4b, 0x69, 0x4e, 0xf4, 0xff, 0xc9, 0x44, 0xf5,
	0x2b, 0xdb, 0xe3, 0xca, 0x0c, 0xd8, 0xf0, 0x3a, 0xcc, 0xfa, 0xb2, 0x54, 0x4c, 0x66, 0x4a, 0xd3,
	0x79, 0xa1, 0x38, 0x47, 0xc0, 0x33, 0xb9, 0x96, 0xdc, 0x73, 0x98, 0x1f, 0x2a, 0xc7, 0xbd, 0x39,
	0x67, 0xd6, 0x35, 0xef, 0xd4, 0x14, 0x71, 0x97, 0x78, 0x11, 0xa6, 0x3e, 0x33, 0xcf, 0xbb, 0x9e,
	0xf8, 0x22, 0xf1, 0x36, 0xcf, 0x12, 0x1f, 0x0a, 0xb9, 0x8f, 0x01, 0x8f, 0x16, 0x71, 0x1b, 0x86,
	0x42, 0x26, 0xa8, 0xb0, 0x61, 0xb6, 0x58, 0x61, 0x0e, 0xfc, 0xea, 0x0c, 0xc5, 0x62, 0x11, 0x27,
	0x65, 0xce, 0x65, 0xe1, 0xf7, 0x04, 0x64, 0xfa, 0x7b, 0x8b, 0xb1, 0x73, 0xeb, 0xbf, 0x7f, 0xc5,
	0xf8, 0x05, 0xcc, 0x55, 0x9d, 0x4e, 0x27, 0x88, 0x17, 0x6f, 0x8e, 0x1f, 0x08, 0x70, 0xa7, 0xdc,
	0x55, 0xcb, 0xae, 0x37, 0xf9, 0x0d, 0x74, 0xaf, 0x68, 0x7f, 0x8b, 0x3f, 0x82, 0xf4, 0xc1, 0xb9,
	0x79, 0xed, 0xce, 0x57, 0xef, 0xb6, 0x3d, 0x18, 0x6e, 0xba, 0x7f, 0xec, 0x92, 0x0f, 0xf3, 0x5a,
	0xde, 0x0f, 0xca, 0x3d, 0x83, 0xb9, 0xa8, 0xe3, 0x56, 0x6d, 0x78, 0x27, 0xc0, 0x5d, 0x45, 0xd3,
	0x6f, 0x35, 0x96, 0x36, 0x21, 0xc9, 0x95, 0x4f, 0x4c, 0x56, 0x9e, 0x83, 0xdc, 0x62, 0x14, 0x8b,
	0xf9, 0xaf, 0xd2, 0x5d, 0xba, 0x42, 0x28, 0x16, 0xe3, 0x22, 0xfa, 0xe3, 0xde, 0xdf, 0x16, 0x7e,
	0x4b, 0x8c, 0x16, 0x43, 0xf1, 0x8b, 0x50, 0x1e, 0x81, 0xcb, 0xf3, 0x30, 0xcc, 0x38, 0x82, 0x2e,
	0x29, 0x16, 0xf3, 0xd0, 0x81, 0x3e, 0x78, 0x6f, 0xe0, 0x81, 0x7a, 0x7f, 0x61, 0x9b, 0x93, 0x38,
	0x26, 0x3c, 0xd4, 0xdc, 0x5b, 0x98, 0x09, 0x52, 0xfc, 0xab, 0x3a, 0xbd, 0xe7, 0x13, 0x2d, 0xec,
	0xc0, 0x92, 0x5a, 0xd5, 0xe4, 0x9a, 0x71, 0xcb, 0xf6, 0x16, 0xfe, 0x4c, 0xc4, 0x45, 0xd2, 0x1b,
	0x0e, 0x7c, 0x02, 0xc8, 0x13, 0x66, 0x44, 0xee, 0x9d, 0xc8, 0xfc, 0x1c, 0x4f, 0x5d, 0x1a, 0x8e,
	0xf4, 0xa4, 0x1f, 0x21, 0xc4, 0x87, 0x03, 0xdd, 0x14, 0x39, 0xfd, 0xd6, 0xcd, 0xf4, 0x93, 0x7a,
	0xaa, 0xc2, 0xbd, 0xb1, 0xd9, 0x6f, 0x35, 0x12, 0xdf, 0xb3, 0x5d, 0x3a, 0xe4, 0x0e, 0xaa, 0xf2,
	0x1b, 0x8d, 0xd0, 0x48, 0xe1, 0xd5, 0x5d, 0xaa, 0x73, 0xdd, 0x9f, 0x42, 0xfa, 0x72, 0xe0, 0x0d,
	0xac, 0x86, 0x27, 0xf6, 0xf5, 0xa7, 0xcc, 0x69, 0x9b, 0xa7, 0x96, 0x3b, 0xdd, 0x49, 0x1f, 0xbc,
	0xf1, 0x93, 0x08, 0x73, 0x01, 0xdd, 0x3e, 0x7d, 0x89, 0x17, 0x01, 0x1d, 0xd5, 0xf6, 0x6a, 0xf5,
	0x46, 0x98, 0x05, 0xdd, 0xc1, 0x79, 0x58, 0x51, 0xa9, 0x11, 0xf7, 0xc9, 0x89, 0xbe, 0xec, 0xb9,
	0x08, 0xaa, 0xc6, 0x21, 0x28, 0xfa, 0xaa, 0xe7, 0x73, 0xc4, 0x7d, 0xe9, 0xa0, 0xaf, 0xfb, 0x1c,
	0x71, 0x5f, 0x2e, 0xe8, 0x9b, 0x1e, 0x5e, 0x86, 0x45, 0x95, 0x8e, 0x5e, 0x58, 0xf4, 0x2d, 0x77,
	0x51, 0x75, 0xb4, 0xaf, 0xe8, 0xbb, 0x1e, 0x5e, 0x87, 0x9c, 0x4a, 0x8d, 0x98, 0xcb, 0x8e, 0xde,
	0x71, 0x00, 0x55, 0x63, 0x00, 0x14, 0x7d, 0xdf, 0xc3, 0x05, 0x58, 0x53, 0xa9, 0x11, 0xab, 0xff,
	0x21, 0xfa, 0x81, 0x63, 0xa8, 0x1a, 0x8f, 0xa1, 0xe8, 0xc7, 0x1e, 0x7e, 0x00, 0xeb, 0x54, 0x35,
	0x26, 0x7d, 0x55, 0xa1, 0x9f, 0x7b, 0x1b, 0x5d, 0xc8, 0x0c, 0x4c, 0x00, 0xbc, 0x0c, 0xf7, 0x42,
	0x3e, 0xd7, 0x62, 0x1c, 0xd9, 0x67, 0xb6, 0x73, 0x85, 0xee, 0xe0, 0xfb, 0x80, 0x87, 0x5c, 0x44,
	0xab, 0x20, 0x01, 0x4b, 0xb0, 0x38, 0x64, 0xf7, 0x7a, 0x99, 0xc0, 0x4b, 0xb0, 0x30, 0xe4, 0xa9,
	0x1e, 0xa9, 0x7b, 0x48, 0xdc, 0xf8, 0x43, 0x80, 0xbb, 0x23, 0x7f, 0x52, 0x78, 0x0d, 0x96, 0x43,
	0x78, 0xdf, 0x1a, 0xe6, 0x1f, 0xc8, 0x13, 0xb8, 0xb7, 0x86, 0x2b, 0x08, 0x3c, 0x65, 0x94, 0x88,
	0xf1, 0x6c, 0x23, 0x31, 0xc6, 0xf3, 0x01, 0x4a, 0xc6, 0x78, 0x9e, 0xa0, 0xa9, 0x18, 0xcf, 0x53,
	0x94, 0x8a, 0xf1, 0xec, 0xa0, 0xf4, 0xc6, 0xaf, 0x42, 0x44, 0x84, 0xf0, 0x03, 0x13, 0xaf, 0xc3,
	0x4a, 0x18, 0x11, 0xda, 0xc3, 0xe3, 0x0e, 0xa8, 0x11, 0x01, 0x34, 0xe4, 0x5d, 0x1d, 0x09, 0x78,
	0x15, 0xa4, 0xb1, 0x6e, 0x45, 0xd3, 0x51, 0x22, 0x36, 0x98, 0xea, 0xf5, 0x03, 0x24, 0xc6, 0x26,
	0xa7, 0x9a, 0xae, 0x57, 0x35, 0x94, 0x54, 0xe6, 0x3f, 0xc9, 0xb4, 0x1d, 0x87, 0x3d, 0xee, 0xbf,
	0xec, 0xe3, 0x14, 0x5f, 0x6d, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x42, 0xf9, 0x77, 0x92, 0x48,
	0x0e, 0x00, 0x00,
}
