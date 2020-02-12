// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/fruitmary.proto

package protomsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 网络消息
type FRUITMARYMSG int32

const (
	FRUITMARYMSG_UNKNOW_FRUITMARY            FRUITMARYMSG = 0
	FRUITMARYMSG_CS_ENTER_GAME_FRUITMARY_REQ FRUITMARYMSG = 20001
	FRUITMARYMSG_SC_ENTER_GAME_FRUITMARY_RES FRUITMARYMSG = 20002
	FRUITMARYMSG_CS_LEAVE_GAME_FRUITMARY_REQ FRUITMARYMSG = 20003
	FRUITMARYMSG_SC_LEAVE_GAME_FRUITMARY_RES FRUITMARYMSG = 20004
	FRUITMARYMSG_CS_START_MARY_REQ           FRUITMARYMSG = 20008
	FRUITMARYMSG_SC_START_MARY_RES           FRUITMARYMSG = 20009
	FRUITMARYMSG_SC_UPDATE_MARY_BONUS        FRUITMARYMSG = 20010
	FRUITMARYMSG_CS_START_MARY2_REQ          FRUITMARYMSG = 20011
	FRUITMARYMSG_SC_START_MARY2_RES          FRUITMARYMSG = 20012
	FRUITMARYMSG_CS_NEXT_MARY_RESULT         FRUITMARYMSG = 20013
	FRUITMARYMSG_CS_PLAYERS_LIST_REQ         FRUITMARYMSG = 20015
	FRUITMARYMSG_SC_PLAYERS_LIST_RES         FRUITMARYMSG = 20016
)

var FRUITMARYMSG_name = map[int32]string{
	0:     "UNKNOW_FRUITMARY",
	20001: "CS_ENTER_GAME_FRUITMARY_REQ",
	20002: "SC_ENTER_GAME_FRUITMARY_RES",
	20003: "CS_LEAVE_GAME_FRUITMARY_REQ",
	20004: "SC_LEAVE_GAME_FRUITMARY_RES",
	20008: "CS_START_MARY_REQ",
	20009: "SC_START_MARY_RES",
	20010: "SC_UPDATE_MARY_BONUS",
	20011: "CS_START_MARY2_REQ",
	20012: "SC_START_MARY2_RES",
	20013: "CS_NEXT_MARY_RESULT",
	20015: "CS_PLAYERS_LIST_REQ",
	20016: "SC_PLAYERS_LIST_RES",
}
var FRUITMARYMSG_value = map[string]int32{
	"UNKNOW_FRUITMARY":            0,
	"CS_ENTER_GAME_FRUITMARY_REQ": 20001,
	"SC_ENTER_GAME_FRUITMARY_RES": 20002,
	"CS_LEAVE_GAME_FRUITMARY_REQ": 20003,
	"SC_LEAVE_GAME_FRUITMARY_RES": 20004,
	"CS_START_MARY_REQ":           20008,
	"SC_START_MARY_RES":           20009,
	"SC_UPDATE_MARY_BONUS":        20010,
	"CS_START_MARY2_REQ":          20011,
	"SC_START_MARY2_RES":          20012,
	"CS_NEXT_MARY_RESULT":         20013,
	"CS_PLAYERS_LIST_REQ":         20015,
	"SC_PLAYERS_LIST_RES":         20016,
}

func (x FRUITMARYMSG) String() string {
	return proto.EnumName(FRUITMARYMSG_name, int32(x))
}
func (FRUITMARYMSG) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// 游戏1 图案枚举
type Fruit1ID int32

const (
	Fruit1ID_Fruit1Unknow     Fruit1ID = 0
	Fruit1ID_Fruit1Wild       Fruit1ID = 1
	Fruit1ID_Fruit1Bonus      Fruit1ID = 2
	Fruit1ID_Fruit1Scatter    Fruit1ID = 3
	Fruit1ID_Fruit1Bar        Fruit1ID = 4
	Fruit1ID_Fruit1Cherry     Fruit1ID = 5
	Fruit1ID_Fruit1Bell       Fruit1ID = 6
	Fruit1ID_Fruit1Pineapple  Fruit1ID = 7
	Fruit1ID_Fruit1Grap       Fruit1ID = 8
	Fruit1ID_Fruit1Mango      Fruit1ID = 9
	Fruit1ID_Fruit1Watermelon Fruit1ID = 10
	Fruit1ID_Fruit1Banana     Fruit1ID = 11
)

var Fruit1ID_name = map[int32]string{
	0:  "Fruit1Unknow",
	1:  "Fruit1Wild",
	2:  "Fruit1Bonus",
	3:  "Fruit1Scatter",
	4:  "Fruit1Bar",
	5:  "Fruit1Cherry",
	6:  "Fruit1Bell",
	7:  "Fruit1Pineapple",
	8:  "Fruit1Grap",
	9:  "Fruit1Mango",
	10: "Fruit1Watermelon",
	11: "Fruit1Banana",
}
var Fruit1ID_value = map[string]int32{
	"Fruit1Unknow":     0,
	"Fruit1Wild":       1,
	"Fruit1Bonus":      2,
	"Fruit1Scatter":    3,
	"Fruit1Bar":        4,
	"Fruit1Cherry":     5,
	"Fruit1Bell":       6,
	"Fruit1Pineapple":  7,
	"Fruit1Grap":       8,
	"Fruit1Mango":      9,
	"Fruit1Watermelon": 10,
	"Fruit1Banana":     11,
}

func (x Fruit1ID) String() string {
	return proto.EnumName(Fruit1ID_name, int32(x))
}
func (Fruit1ID) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// 游戏2 图案枚举
type Fruit2ID int32

const (
	Fruit2ID_Fruit2Unknow     Fruit2ID = 0
	Fruit2ID_Fruit2Watermelon Fruit2ID = 1
	Fruit2ID_Fruit2Grap       Fruit2ID = 2
	Fruit2ID_Fruit2Mango      Fruit2ID = 3
	Fruit2ID_Fruit2Cherry     Fruit2ID = 4
	Fruit2ID_Fruit2Banana     Fruit2ID = 5
	Fruit2ID_Fruit2Orange     Fruit2ID = 6
	Fruit2ID_Fruit2Pineapple  Fruit2ID = 7
	Fruit2ID_Fruit2Bomb       Fruit2ID = 8
)

var Fruit2ID_name = map[int32]string{
	0: "Fruit2Unknow",
	1: "Fruit2Watermelon",
	2: "Fruit2Grap",
	3: "Fruit2Mango",
	4: "Fruit2Cherry",
	5: "Fruit2Banana",
	6: "Fruit2Orange",
	7: "Fruit2Pineapple",
	8: "Fruit2Bomb",
}
var Fruit2ID_value = map[string]int32{
	"Fruit2Unknow":     0,
	"Fruit2Watermelon": 1,
	"Fruit2Grap":       2,
	"Fruit2Mango":      3,
	"Fruit2Cherry":     4,
	"Fruit2Banana":     5,
	"Fruit2Orange":     6,
	"Fruit2Pineapple":  7,
	"Fruit2Bomb":       8,
}

func (x Fruit2ID) String() string {
	return proto.EnumName(Fruit2ID_name, int32(x))
}
func (Fruit2ID) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

// 请求进入房间
type ENTER_GAME_FRUITMARY_REQ struct {
	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	RoomID    uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *ENTER_GAME_FRUITMARY_REQ) Reset()                    { *m = ENTER_GAME_FRUITMARY_REQ{} }
func (m *ENTER_GAME_FRUITMARY_REQ) String() string            { return proto.CompactTextString(m) }
func (*ENTER_GAME_FRUITMARY_REQ) ProtoMessage()               {}
func (*ENTER_GAME_FRUITMARY_REQ) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ENTER_GAME_FRUITMARY_REQ) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *ENTER_GAME_FRUITMARY_REQ) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type ENTER_GAME_FRUITMARY_RES struct {
	RoomID       uint32                                         `protobuf:"varint,1,opt,name=RoomID" json:"RoomID,omitempty"`
	Basics       int64                                          `protobuf:"varint,2,opt,name=Basics" json:"Basics,omitempty"`
	Bonus        int64                                          `protobuf:"varint,3,opt,name=Bonus" json:"Bonus,omitempty"`
	LastBet      int64                                          `protobuf:"varint,4,opt,name=LastBet" json:"LastBet,omitempty"`
	Bets         []uint64                                       `protobuf:"varint,5,rep,packed,name=Bets" json:"Bets,omitempty"`
	Ratio        map[int32]*ENTER_GAME_FRUITMARY_RES_FruitRatio `protobuf:"bytes,6,rep,name=Ratio" json:"Ratio,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Mary2_Result *START_MARY2_RES                               `protobuf:"bytes,7,opt,name=Mary2_Result,json=Mary2Result" json:"Mary2_Result,omitempty"`
	FeeCount     int32                                          `protobuf:"varint,8,opt,name=FeeCount" json:"FeeCount,omitempty"`
}

func (m *ENTER_GAME_FRUITMARY_RES) Reset()                    { *m = ENTER_GAME_FRUITMARY_RES{} }
func (m *ENTER_GAME_FRUITMARY_RES) String() string            { return proto.CompactTextString(m) }
func (*ENTER_GAME_FRUITMARY_RES) ProtoMessage()               {}
func (*ENTER_GAME_FRUITMARY_RES) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ENTER_GAME_FRUITMARY_RES) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *ENTER_GAME_FRUITMARY_RES) GetBasics() int64 {
	if m != nil {
		return m.Basics
	}
	return 0
}

func (m *ENTER_GAME_FRUITMARY_RES) GetBonus() int64 {
	if m != nil {
		return m.Bonus
	}
	return 0
}

func (m *ENTER_GAME_FRUITMARY_RES) GetLastBet() int64 {
	if m != nil {
		return m.LastBet
	}
	return 0
}

func (m *ENTER_GAME_FRUITMARY_RES) GetBets() []uint64 {
	if m != nil {
		return m.Bets
	}
	return nil
}

func (m *ENTER_GAME_FRUITMARY_RES) GetRatio() map[int32]*ENTER_GAME_FRUITMARY_RES_FruitRatio {
	if m != nil {
		return m.Ratio
	}
	return nil
}

func (m *ENTER_GAME_FRUITMARY_RES) GetMary2_Result() *START_MARY2_RES {
	if m != nil {
		return m.Mary2_Result
	}
	return nil
}

func (m *ENTER_GAME_FRUITMARY_RES) GetFeeCount() int32 {
	if m != nil {
		return m.FeeCount
	}
	return 0
}

type ENTER_GAME_FRUITMARY_RES_FruitRatio struct {
	ID     Fruit2ID `protobuf:"varint,1,opt,name=ID,enum=protomsg.Fruit2ID" json:"ID,omitempty"`
	Single int32    `protobuf:"varint,2,opt,name=Single" json:"Single,omitempty"`
	Same_2 int32    `protobuf:"varint,3,opt,name=Same_2,json=Same2" json:"Same_2,omitempty"`
	Same_3 int32    `protobuf:"varint,4,opt,name=Same_3,json=Same3" json:"Same_3,omitempty"`
	Same_4 int32    `protobuf:"varint,5,opt,name=Same_4,json=Same4" json:"Same_4,omitempty"`
}

func (m *ENTER_GAME_FRUITMARY_RES_FruitRatio) Reset()         { *m = ENTER_GAME_FRUITMARY_RES_FruitRatio{} }
func (m *ENTER_GAME_FRUITMARY_RES_FruitRatio) String() string { return proto.CompactTextString(m) }
func (*ENTER_GAME_FRUITMARY_RES_FruitRatio) ProtoMessage()    {}
func (*ENTER_GAME_FRUITMARY_RES_FruitRatio) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{1, 0}
}

func (m *ENTER_GAME_FRUITMARY_RES_FruitRatio) GetID() Fruit2ID {
	if m != nil {
		return m.ID
	}
	return Fruit2ID_Fruit2Unknow
}

func (m *ENTER_GAME_FRUITMARY_RES_FruitRatio) GetSingle() int32 {
	if m != nil {
		return m.Single
	}
	return 0
}

func (m *ENTER_GAME_FRUITMARY_RES_FruitRatio) GetSame_2() int32 {
	if m != nil {
		return m.Same_2
	}
	return 0
}

func (m *ENTER_GAME_FRUITMARY_RES_FruitRatio) GetSame_3() int32 {
	if m != nil {
		return m.Same_3
	}
	return 0
}

func (m *ENTER_GAME_FRUITMARY_RES_FruitRatio) GetSame_4() int32 {
	if m != nil {
		return m.Same_4
	}
	return 0
}

// 请求退出房间
type LEAVE_GAME_FRUITMARY_REQ struct {
	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID" json:"AccountID,omitempty"`
	RoomID    uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *LEAVE_GAME_FRUITMARY_REQ) Reset()                    { *m = LEAVE_GAME_FRUITMARY_REQ{} }
func (m *LEAVE_GAME_FRUITMARY_REQ) String() string            { return proto.CompactTextString(m) }
func (*LEAVE_GAME_FRUITMARY_REQ) ProtoMessage()               {}
func (*LEAVE_GAME_FRUITMARY_REQ) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *LEAVE_GAME_FRUITMARY_REQ) GetAccountID() uint32 {
	if m != nil {
		return m.AccountID
	}
	return 0
}

func (m *LEAVE_GAME_FRUITMARY_REQ) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type LEAVE_GAME_FRUITMARY_RES struct {
	Ret    uint32 `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	RoomID uint32 `protobuf:"varint,2,opt,name=RoomID" json:"RoomID,omitempty"`
}

func (m *LEAVE_GAME_FRUITMARY_RES) Reset()                    { *m = LEAVE_GAME_FRUITMARY_RES{} }
func (m *LEAVE_GAME_FRUITMARY_RES) String() string            { return proto.CompactTextString(m) }
func (*LEAVE_GAME_FRUITMARY_RES) ProtoMessage()               {}
func (*LEAVE_GAME_FRUITMARY_RES) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *LEAVE_GAME_FRUITMARY_RES) GetRet() uint32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *LEAVE_GAME_FRUITMARY_RES) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

// //////////////////////////////////////////// 游戏1 /////////////////////////////////////////////
// 请求开始游戏1
type START_MARY_REQ struct {
	Bet uint64 `protobuf:"varint,1,opt,name=Bet" json:"Bet,omitempty"`
}

func (m *START_MARY_REQ) Reset()                    { *m = START_MARY_REQ{} }
func (m *START_MARY_REQ) String() string            { return proto.CompactTextString(m) }
func (*START_MARY_REQ) ProtoMessage()               {}
func (*START_MARY_REQ) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *START_MARY_REQ) GetBet() uint64 {
	if m != nil {
		return m.Bet
	}
	return 0
}

type START_MARY_RES struct {
	Ret          uint64               `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	SumOdds      int64                `protobuf:"varint,2,opt,name=SumOdds" json:"SumOdds,omitempty"`
	Results      []*FRUITMARY_Result  `protobuf:"bytes,3,rep,name=Results" json:"Results,omitempty"`
	PictureList  []int32              `protobuf:"varint,4,rep,packed,name=PictureList" json:"PictureList,omitempty"`
	Bonus        int64                `protobuf:"varint,5,opt,name=Bonus" json:"Bonus,omitempty"`
	Money        int64                `protobuf:"varint,6,opt,name=Money" json:"Money,omitempty"`
	FreeCount    int64                `protobuf:"varint,7,opt,name=FreeCount" json:"FreeCount,omitempty"`
	MaryCount    int64                `protobuf:"varint,8,opt,name=MaryCount" json:"MaryCount,omitempty"`
	FeePositions []*FRUITMARYPosition `protobuf:"bytes,9,rep,name=FeePositions" json:"FeePositions,omitempty"`
}

func (m *START_MARY_RES) Reset()                    { *m = START_MARY_RES{} }
func (m *START_MARY_RES) String() string            { return proto.CompactTextString(m) }
func (*START_MARY_RES) ProtoMessage()               {}
func (*START_MARY_RES) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *START_MARY_RES) GetRet() uint64 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *START_MARY_RES) GetSumOdds() int64 {
	if m != nil {
		return m.SumOdds
	}
	return 0
}

func (m *START_MARY_RES) GetResults() []*FRUITMARY_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *START_MARY_RES) GetPictureList() []int32 {
	if m != nil {
		return m.PictureList
	}
	return nil
}

func (m *START_MARY_RES) GetBonus() int64 {
	if m != nil {
		return m.Bonus
	}
	return 0
}

func (m *START_MARY_RES) GetMoney() int64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *START_MARY_RES) GetFreeCount() int64 {
	if m != nil {
		return m.FreeCount
	}
	return 0
}

func (m *START_MARY_RES) GetMaryCount() int64 {
	if m != nil {
		return m.MaryCount
	}
	return 0
}

func (m *START_MARY_RES) GetFeePositions() []*FRUITMARYPosition {
	if m != nil {
		return m.FeePositions
	}
	return nil
}

type FRUITMARYPosition struct {
	Px int32 `protobuf:"varint,1,opt,name=px" json:"px,omitempty"`
	Py int32 `protobuf:"varint,2,opt,name=py" json:"py,omitempty"`
}

func (m *FRUITMARYPosition) Reset()                    { *m = FRUITMARYPosition{} }
func (m *FRUITMARYPosition) String() string            { return proto.CompactTextString(m) }
func (*FRUITMARYPosition) ProtoMessage()               {}
func (*FRUITMARYPosition) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *FRUITMARYPosition) GetPx() int32 {
	if m != nil {
		return m.Px
	}
	return 0
}

func (m *FRUITMARYPosition) GetPy() int32 {
	if m != nil {
		return m.Py
	}
	return 0
}

type FRUITMARY_Result struct {
	LineId    int32                `protobuf:"varint,1,opt,name=LineId" json:"LineId,omitempty"`
	Count     int32                `protobuf:"varint,2,opt,name=Count" json:"Count,omitempty"`
	Odds      int32                `protobuf:"varint,3,opt,name=Odds" json:"Odds,omitempty"`
	Positions []*FRUITMARYPosition `protobuf:"bytes,4,rep,name=Positions" json:"Positions,omitempty"`
}

func (m *FRUITMARY_Result) Reset()                    { *m = FRUITMARY_Result{} }
func (m *FRUITMARY_Result) String() string            { return proto.CompactTextString(m) }
func (*FRUITMARY_Result) ProtoMessage()               {}
func (*FRUITMARY_Result) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *FRUITMARY_Result) GetLineId() int32 {
	if m != nil {
		return m.LineId
	}
	return 0
}

func (m *FRUITMARY_Result) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *FRUITMARY_Result) GetOdds() int32 {
	if m != nil {
		return m.Odds
	}
	return 0
}

func (m *FRUITMARY_Result) GetPositions() []*FRUITMARYPosition {
	if m != nil {
		return m.Positions
	}
	return nil
}

// 通知更新奖金池
type UPDATE_MARY_BONUS struct {
	Bonus int64 `protobuf:"varint,1,opt,name=Bonus" json:"Bonus,omitempty"`
}

func (m *UPDATE_MARY_BONUS) Reset()                    { *m = UPDATE_MARY_BONUS{} }
func (m *UPDATE_MARY_BONUS) String() string            { return proto.CompactTextString(m) }
func (*UPDATE_MARY_BONUS) ProtoMessage()               {}
func (*UPDATE_MARY_BONUS) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *UPDATE_MARY_BONUS) GetBonus() int64 {
	if m != nil {
		return m.Bonus
	}
	return 0
}

// //////////////////////////////////////////// 游戏2 /////////////////////////////////////////////
// 请求开始游戏2
type START_MARY2_REQ struct {
}

func (m *START_MARY2_REQ) Reset()                    { *m = START_MARY2_REQ{} }
func (m *START_MARY2_REQ) String() string            { return proto.CompactTextString(m) }
func (*START_MARY2_REQ) ProtoMessage()               {}
func (*START_MARY2_REQ) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

type START_MARY2_RES struct {
	Result         []*Mary2_Result `protobuf:"bytes,1,rep,name=Result" json:"Result,omitempty"`
	MarySpareCount int32           `protobuf:"varint,2,opt,name=MarySpareCount" json:"MarySpareCount,omitempty"`
}

func (m *START_MARY2_RES) Reset()                    { *m = START_MARY2_RES{} }
func (m *START_MARY2_RES) String() string            { return proto.CompactTextString(m) }
func (*START_MARY2_RES) ProtoMessage()               {}
func (*START_MARY2_RES) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *START_MARY2_RES) GetResult() []*Mary2_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *START_MARY2_RES) GetMarySpareCount() int32 {
	if m != nil {
		return m.MarySpareCount
	}
	return 0
}

type Mary2_Result struct {
	IndexId int32   `protobuf:"varint,1,opt,name=IndexId" json:"IndexId,omitempty"`
	MaryId  []int32 `protobuf:"varint,2,rep,packed,name=MaryId" json:"MaryId,omitempty"`
	Profit1 int32   `protobuf:"varint,3,opt,name=Profit1" json:"Profit1,omitempty"`
	Profit2 int32   `protobuf:"varint,4,opt,name=Profit2" json:"Profit2,omitempty"`
}

func (m *Mary2_Result) Reset()                    { *m = Mary2_Result{} }
func (m *Mary2_Result) String() string            { return proto.CompactTextString(m) }
func (*Mary2_Result) ProtoMessage()               {}
func (*Mary2_Result) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *Mary2_Result) GetIndexId() int32 {
	if m != nil {
		return m.IndexId
	}
	return 0
}

func (m *Mary2_Result) GetMaryId() []int32 {
	if m != nil {
		return m.MaryId
	}
	return nil
}

func (m *Mary2_Result) GetProfit1() int32 {
	if m != nil {
		return m.Profit1
	}
	return 0
}

func (m *Mary2_Result) GetProfit2() int32 {
	if m != nil {
		return m.Profit2
	}
	return 0
}

type NEXT_MARY_RESULT struct {
}

func (m *NEXT_MARY_RESULT) Reset()                    { *m = NEXT_MARY_RESULT{} }
func (m *NEXT_MARY_RESULT) String() string            { return proto.CompactTextString(m) }
func (*NEXT_MARY_RESULT) ProtoMessage()               {}
func (*NEXT_MARY_RESULT) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

// 请求小玛利玩家列表
type PLAYERS_LIST_RES struct {
	Players []*AccountStorageData `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
}

func (m *PLAYERS_LIST_RES) Reset()                    { *m = PLAYERS_LIST_RES{} }
func (m *PLAYERS_LIST_RES) String() string            { return proto.CompactTextString(m) }
func (*PLAYERS_LIST_RES) ProtoMessage()               {}
func (*PLAYERS_LIST_RES) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func (m *PLAYERS_LIST_RES) GetPlayers() []*AccountStorageData {
	if m != nil {
		return m.Players
	}
	return nil
}

func init() {
	proto.RegisterType((*ENTER_GAME_FRUITMARY_REQ)(nil), "protomsg.ENTER_GAME_FRUITMARY_REQ")
	proto.RegisterType((*ENTER_GAME_FRUITMARY_RES)(nil), "protomsg.ENTER_GAME_FRUITMARY_RES")
	proto.RegisterType((*ENTER_GAME_FRUITMARY_RES_FruitRatio)(nil), "protomsg.ENTER_GAME_FRUITMARY_RES.FruitRatio")
	proto.RegisterType((*LEAVE_GAME_FRUITMARY_REQ)(nil), "protomsg.LEAVE_GAME_FRUITMARY_REQ")
	proto.RegisterType((*LEAVE_GAME_FRUITMARY_RES)(nil), "protomsg.LEAVE_GAME_FRUITMARY_RES")
	proto.RegisterType((*START_MARY_REQ)(nil), "protomsg.START_MARY_REQ")
	proto.RegisterType((*START_MARY_RES)(nil), "protomsg.START_MARY_RES")
	proto.RegisterType((*FRUITMARYPosition)(nil), "protomsg.FRUITMARY_position")
	proto.RegisterType((*FRUITMARY_Result)(nil), "protomsg.FRUITMARY_Result")
	proto.RegisterType((*UPDATE_MARY_BONUS)(nil), "protomsg.UPDATE_MARY_BONUS")
	proto.RegisterType((*START_MARY2_REQ)(nil), "protomsg.START_MARY2_REQ")
	proto.RegisterType((*START_MARY2_RES)(nil), "protomsg.START_MARY2_RES")
	proto.RegisterType((*Mary2_Result)(nil), "protomsg.Mary2_Result")
	proto.RegisterType((*NEXT_MARY_RESULT)(nil), "protomsg.NEXT_MARY_RESULT")
	proto.RegisterType((*PLAYERS_LIST_RES)(nil), "protomsg.PLAYERS_LIST_RES")
	proto.RegisterEnum("protomsg.FRUITMARYMSG", FRUITMARYMSG_name, FRUITMARYMSG_value)
	proto.RegisterEnum("protomsg.Fruit1ID", Fruit1ID_name, Fruit1ID_value)
	proto.RegisterEnum("protomsg.Fruit2ID", Fruit2ID_name, Fruit2ID_value)
}

func init() { proto.RegisterFile("protobuf/fruitmary.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x8e, 0x1b, 0x45,
	0x10, 0xce, 0xd8, 0x1e, 0xdb, 0x5b, 0xde, 0x9f, 0xde, 0xce, 0x12, 0x26, 0x26, 0x07, 0x33, 0x07,
	0x64, 0x22, 0xb1, 0x51, 0x26, 0x2b, 0x84, 0x22, 0x0e, 0xf8, 0x6f, 0x57, 0x06, 0x7b, 0xd7, 0xdb,
	0x6d, 0x13, 0xc2, 0x65, 0xd4, 0xbb, 0xee, 0x35, 0xa3, 0xd8, 0x33, 0xd6, 0x78, 0x0c, 0xeb, 0x33,
	0x67, 0x24, 0x1e, 0xc0, 0x17, 0x7e, 0x84, 0xf8, 0x87, 0x27, 0xe0, 0x55, 0x78, 0x15, 0xd4, 0xdd,
	0x33, 0xee, 0xb1, 0x13, 0x4b, 0x39, 0x70, 0x9a, 0xae, 0xaa, 0xaf, 0xaa, 0xba, 0xaa, 0xbe, 0xae,
	0x01, 0x6b, 0x1a, 0x06, 0x51, 0x70, 0x35, 0xbf, 0x79, 0x74, 0x13, 0xce, 0xbd, 0x68, 0xc2, 0xc2,
	0xc5, 0xb1, 0x54, 0xe1, 0xa2, 0xfc, 0x4c, 0x66, 0xa3, 0xf2, 0xdd, 0x15, 0x66, 0xc8, 0x22, 0xa6,
	0xcc, 0x76, 0x0f, 0xac, 0xd6, 0x79, 0xbf, 0x45, 0xdc, 0xb3, 0x5a, 0xb7, 0xe5, 0x9e, 0x92, 0x41,
	0xbb, 0xdf, 0xad, 0x91, 0xe7, 0x2e, 0x69, 0x5d, 0xe2, 0x07, 0xb0, 0x53, 0xbb, 0xbe, 0x0e, 0xe6,
	0x7e, 0xd4, 0x6e, 0x5a, 0x46, 0xc5, 0xa8, 0xee, 0x11, 0xad, 0xc0, 0xf7, 0x20, 0x4f, 0x82, 0x60,
	0xd2, 0x6e, 0x5a, 0x19, 0x69, 0x8a, 0x25, 0xfb, 0xa7, 0xdc, 0xd6, 0x90, 0x34, 0xe5, 0x64, 0xa4,
	0x9d, 0x84, 0xbe, 0xce, 0x66, 0xde, 0xf5, 0x4c, 0x06, 0xcb, 0x92, 0x58, 0xc2, 0x47, 0x60, 0xd6,
	0x03, 0x7f, 0x3e, 0xb3, 0xb2, 0x52, 0xad, 0x04, 0x6c, 0x41, 0xa1, 0xc3, 0x66, 0x51, 0x9d, 0x47,
	0x56, 0x4e, 0xea, 0x13, 0x11, 0x63, 0xc8, 0xd5, 0x79, 0x34, 0xb3, 0xcc, 0x4a, 0xb6, 0x9a, 0x23,
	0xf2, 0x8c, 0x1b, 0x60, 0x12, 0x16, 0x79, 0x81, 0x95, 0xaf, 0x64, 0xab, 0x25, 0xe7, 0xbd, 0xe3,
	0xa4, 0x23, 0xc7, 0xdb, 0xae, 0x79, 0x2c, 0xf1, 0x2d, 0x3f, 0x0a, 0x17, 0x44, 0xf9, 0xe2, 0x0f,
	0x61, 0xb7, 0xcb, 0xc2, 0x85, 0xe3, 0x12, 0x3e, 0x9b, 0x8f, 0x23, 0xab, 0x50, 0x31, 0xaa, 0x25,
	0xe7, 0xbe, 0x8e, 0x45, 0xfb, 0x35, 0xd2, 0x77, 0x45, 0x04, 0x47, 0x84, 0x20, 0x25, 0x09, 0x57,
	0x68, 0x5c, 0x86, 0xe2, 0x29, 0xe7, 0x0d, 0xd1, 0x39, 0xab, 0x58, 0x31, 0xaa, 0x26, 0x59, 0xc9,
	0xe5, 0x6f, 0x0c, 0x80, 0x53, 0x31, 0x34, 0x95, 0xc8, 0x86, 0x4c, 0xdc, 0x9d, 0x7d, 0x07, 0xeb,
	0xf0, 0x12, 0xe1, 0xb4, 0x9b, 0x24, 0xa3, 0xba, 0x45, 0x3d, 0x7f, 0x34, 0xe6, 0xb2, 0x5b, 0x26,
	0x89, 0x25, 0xfc, 0x06, 0xe4, 0x29, 0x9b, 0x70, 0xd7, 0x91, 0xed, 0x32, 0x89, 0x29, 0x24, 0x67,
	0xa5, 0x7e, 0x22, 0xbb, 0x15, 0xab, 0x9f, 0xac, 0xd4, 0x27, 0x96, 0xa9, 0xd5, 0x27, 0xe5, 0x11,
	0x80, 0x2e, 0x1f, 0x23, 0xc8, 0xbe, 0xe0, 0x0b, 0x79, 0x1f, 0x93, 0x88, 0xa3, 0x68, 0xe7, 0x97,
	0x6c, 0x3c, 0x57, 0xb9, 0x5f, 0xaf, 0x9d, 0xba, 0x3c, 0xa2, 0x7c, 0x9f, 0x66, 0x3e, 0x30, 0x04,
	0xf5, 0x3a, 0xad, 0xda, 0xa7, 0xad, 0xff, 0x8f, 0x7a, 0xcd, 0xad, 0x11, 0xa9, 0x28, 0x84, 0xf0,
	0x28, 0x8e, 0x25, 0x8e, 0x5b, 0xa3, 0xd8, 0xb0, 0xaf, 0x87, 0x29, 0x6f, 0x83, 0x20, 0x5b, 0x8f,
	0x7d, 0x73, 0x44, 0x1c, 0xed, 0x7f, 0x32, 0x1b, 0xa0, 0xb5, 0x04, 0x39, 0x95, 0xc0, 0x82, 0x02,
	0x9d, 0x4f, 0x2e, 0x86, 0xc3, 0x84, 0xd5, 0x89, 0x88, 0x4f, 0xa0, 0xa0, 0x98, 0x21, 0x88, 0x2d,
	0x48, 0x59, 0x4e, 0x4d, 0x5a, 0x5f, 0x5b, 0x42, 0x48, 0x02, 0xc5, 0x15, 0x28, 0xf5, 0xbc, 0xeb,
	0x68, 0x1e, 0xf2, 0x8e, 0x37, 0x13, 0xd4, 0xcf, 0x56, 0x4d, 0x92, 0x56, 0xe9, 0xe7, 0x62, 0xa6,
	0x9f, 0xcb, 0x11, 0x98, 0xdd, 0xc0, 0xe7, 0x0b, 0x2b, 0xaf, 0xb4, 0x52, 0x10, 0x2d, 0x3e, 0x0d,
	0x13, 0x52, 0x16, 0xa4, 0x45, 0x2b, 0x84, 0x55, 0x10, 0x58, 0x53, 0x36, 0x4b, 0xb4, 0x02, 0x7f,
	0x04, 0xbb, 0xa7, 0x9c, 0xf7, 0x82, 0x99, 0x17, 0x79, 0x81, 0x3f, 0xb3, 0x76, 0x64, 0x11, 0x0f,
	0x5e, 0x55, 0xc4, 0x34, 0x06, 0x91, 0x35, 0x0f, 0xfb, 0x04, 0xf0, 0xcb, 0x18, 0xbc, 0x0f, 0x99,
	0xe9, 0x6d, 0x4c, 0xb6, 0xcc, 0xf4, 0x56, 0xca, 0x8b, 0x98, 0xe4, 0x99, 0xe9, 0xc2, 0xfe, 0xd6,
	0x00, 0xb4, 0xd9, 0x1f, 0x31, 0xc7, 0x8e, 0xe7, 0xf3, 0xf6, 0x30, 0x76, 0x8c, 0x25, 0x51, 0xb6,
	0xba, 0xbe, 0xf2, 0x57, 0x82, 0xd8, 0x10, 0x72, 0x22, 0xea, 0x85, 0xc8, 0x33, 0x7e, 0x0a, 0x3b,
	0xba, 0x96, 0xdc, 0x6b, 0xd4, 0xa2, 0xe1, 0xf6, 0xbb, 0x70, 0x38, 0xe8, 0x35, 0x6b, 0xfd, 0x96,
	0x62, 0x42, 0xfd, 0xe2, 0x7c, 0x40, 0xf5, 0x1c, 0x8c, 0xd4, 0x1c, 0xec, 0x43, 0x38, 0x58, 0xdf,
	0x12, 0x97, 0xb6, 0xb7, 0xa9, 0xa2, 0xf8, 0x18, 0xf2, 0xf1, 0x8e, 0x31, 0xe4, 0x4d, 0xee, 0xe9,
	0x9b, 0xa4, 0x37, 0x10, 0x89, 0x51, 0xf8, 0x1d, 0xd8, 0x17, 0x7a, 0x3a, 0x65, 0x21, 0x4f, 0xd7,
	0xbb, 0xa1, 0xb5, 0xa3, 0xf5, 0x0d, 0x26, 0xd8, 0xd9, 0xf6, 0x87, 0xfc, 0x76, 0xd5, 0xb7, 0x44,
	0x14, 0x0d, 0x15, 0xc8, 0xf6, 0xd0, 0xca, 0x48, 0x8a, 0xc5, 0x92, 0xf0, 0xe8, 0x85, 0xc1, 0x8d,
	0x17, 0x3d, 0x8e, 0xbb, 0x97, 0x88, 0xda, 0xe2, 0xc4, 0x2b, 0x26, 0x11, 0x6d, 0x0c, 0xe8, 0xbc,
	0xf5, 0x99, 0x7e, 0x26, 0x83, 0x4e, 0xdf, 0xfe, 0x18, 0x50, 0xaf, 0x53, 0x7b, 0xde, 0x22, 0xd4,
	0xed, 0xb4, 0x69, 0x5f, 0x56, 0xfd, 0x3e, 0x14, 0xa6, 0x63, 0xb6, 0xe0, 0xe1, 0x2c, 0x2e, 0x3b,
	0x35, 0x80, 0xf8, 0xe1, 0xd3, 0x28, 0x08, 0xd9, 0x88, 0x37, 0x59, 0xc4, 0x48, 0x02, 0x7e, 0xf8,
	0x75, 0x16, 0x76, 0x57, 0x03, 0xea, 0xd2, 0x33, 0x7c, 0x04, 0x68, 0x70, 0xfe, 0xc9, 0xf9, 0xc5,
	0x33, 0xfd, 0xfe, 0xd1, 0x1d, 0xfc, 0x36, 0xbc, 0xd5, 0xa0, 0xee, 0xb6, 0x3f, 0x1d, 0xfa, 0x6e,
	0x69, 0x08, 0x08, 0x6d, 0x6c, 0x83, 0x50, 0xf4, 0xbd, 0x82, 0x34, 0xa8, 0xbb, 0x6d, 0x69, 0xa1,
	0x1f, 0x56, 0x51, 0xb6, 0x6d, 0x21, 0xf4, 0xe3, 0xd2, 0xc0, 0x6f, 0xc2, 0x61, 0x83, 0xba, 0xeb,
	0x2b, 0x06, 0xfd, 0xac, 0x0c, 0xb4, 0xb1, 0x6e, 0xa0, 0xe8, 0x97, 0xa5, 0x81, 0xcb, 0x70, 0x44,
	0x1b, 0xee, 0x4b, 0x34, 0x43, 0xbf, 0x2e, 0x0d, 0x6c, 0x01, 0x5e, 0x8b, 0x26, 0x79, 0x85, 0x7e,
	0x53, 0x96, 0xb5, 0x70, 0x92, 0x5e, 0xe8, 0xf7, 0xa5, 0x81, 0xef, 0xc3, 0xdd, 0x06, 0x75, 0x37,
	0xe7, 0x82, 0xfe, 0x58, 0x99, 0x36, 0xc6, 0x73, 0x89, 0xfe, 0x52, 0x26, 0xda, 0xd8, 0x34, 0x51,
	0xf4, 0xf7, 0xd2, 0x78, 0xf8, 0xaf, 0x01, 0x45, 0xb9, 0xe4, 0x1f, 0xb7, 0x9b, 0x18, 0xc1, 0xae,
	0x3a, 0x0f, 0xfc, 0x17, 0x7e, 0xf0, 0x15, 0xba, 0x83, 0xf7, 0xe3, 0x3f, 0xdc, 0xe3, 0x67, 0xde,
	0x78, 0x88, 0x0c, 0x7c, 0x00, 0x25, 0x25, 0xcb, 0x77, 0x81, 0x32, 0xf8, 0x10, 0xf6, 0x94, 0x82,
	0x5e, 0xb3, 0x28, 0xe2, 0x21, 0xca, 0xe2, 0x3d, 0xb1, 0x9e, 0x24, 0x86, 0x85, 0x28, 0xa7, 0x83,
	0x36, 0xbe, 0xe0, 0x61, 0xb8, 0x40, 0xa6, 0x0e, 0x5a, 0xe7, 0xe3, 0x31, 0xca, 0xe3, 0xbb, 0x70,
	0xa0, 0xe4, 0x9e, 0xe7, 0x73, 0x36, 0x9d, 0x8e, 0x39, 0x2a, 0x68, 0xd0, 0x59, 0xc8, 0xa6, 0xa8,
	0xa8, 0x33, 0x77, 0x99, 0x3f, 0x0a, 0xd0, 0x8e, 0xa0, 0x4b, 0x7c, 0x35, 0x16, 0xf1, 0x70, 0xc2,
	0xc7, 0x81, 0x8f, 0x40, 0x67, 0xab, 0x33, 0x9f, 0xf9, 0x0c, 0x95, 0x1e, 0xfe, 0x99, 0x54, 0xe8,
	0xa4, 0x2a, 0x74, 0x56, 0x15, 0x26, 0x61, 0x9c, 0x54, 0x18, 0x63, 0x95, 0xdd, 0x91, 0xd9, 0x33,
	0xab, 0xec, 0x8e, 0xca, 0x9e, 0xd5, 0x81, 0xe2, 0xaa, 0x74, 0x9d, 0x4e, 0x9c, 0xd9, 0xd4, 0x9a,
	0x8b, 0x90, 0xf9, 0x23, 0x9e, 0xaa, 0xd4, 0x79, 0x55, 0xa5, 0x4e, 0x3d, 0x98, 0x5c, 0xa1, 0x62,
	0xfd, 0xe0, 0xf3, 0xbd, 0x30, 0x08, 0xa2, 0x47, 0xc9, 0x2b, 0xba, 0xca, 0xcb, 0xd3, 0x93, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x35, 0x5a, 0xa6, 0xf8, 0x2b, 0x0a, 0x00, 0x00,
}
