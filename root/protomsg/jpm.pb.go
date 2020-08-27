// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: protobuf/jpm.proto

package protomsg

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 网络消息
type JPMMSG int32

const (
	JPMMSG_UNKNOW_JPM              JPMMSG = 0
	JPMMSG_CS_ENTER_GAME_JPM_REQ   JPMMSG = 14001 // 玩家请求进入房间
	JPMMSG_SC_ENTER_GAME_JPM_RES   JPMMSG = 14002
	JPMMSG_CS_LEAVE_GAME_JPM_REQ   JPMMSG = 14003 // 玩家请求退出房间
	JPMMSG_SC_LEAVE_GAME_JPM_RES   JPMMSG = 14004
	JPMMSG_CS_START_JPM_REQ        JPMMSG = 14008 // 请求开始游戏1
	JPMMSG_SC_START_JPM_RES        JPMMSG = 14009
	JPMMSG_SC_UPDATE_JPM_BONUS     JPMMSG = 14010 // 通知更新奖金池
	JPMMSG_CS_PLAYERS_JPM_LIST_REQ JPMMSG = 14015 // 请求在线玩家列表
	JPMMSG_SC_PLAYERS_JPM_LIST_RES JPMMSG = 14016
)

// Enum value maps for JPMMSG.
var (
	JPMMSG_name = map[int32]string{
		0:     "UNKNOW_JPM",
		14001: "CS_ENTER_GAME_JPM_REQ",
		14002: "SC_ENTER_GAME_JPM_RES",
		14003: "CS_LEAVE_GAME_JPM_REQ",
		14004: "SC_LEAVE_GAME_JPM_RES",
		14008: "CS_START_JPM_REQ",
		14009: "SC_START_JPM_RES",
		14010: "SC_UPDATE_JPM_BONUS",
		14015: "CS_PLAYERS_JPM_LIST_REQ",
		14016: "SC_PLAYERS_JPM_LIST_RES",
	}
	JPMMSG_value = map[string]int32{
		"UNKNOW_JPM":              0,
		"CS_ENTER_GAME_JPM_REQ":   14001,
		"SC_ENTER_GAME_JPM_RES":   14002,
		"CS_LEAVE_GAME_JPM_REQ":   14003,
		"SC_LEAVE_GAME_JPM_RES":   14004,
		"CS_START_JPM_REQ":        14008,
		"SC_START_JPM_RES":        14009,
		"SC_UPDATE_JPM_BONUS":     14010,
		"CS_PLAYERS_JPM_LIST_REQ": 14015,
		"SC_PLAYERS_JPM_LIST_RES": 14016,
	}
)

func (x JPMMSG) Enum() *JPMMSG {
	p := new(JPMMSG)
	*p = x
	return p
}

func (x JPMMSG) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JPMMSG) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_jpm_proto_enumTypes[0].Descriptor()
}

func (JPMMSG) Type() protoreflect.EnumType {
	return &file_protobuf_jpm_proto_enumTypes[0]
}

func (x JPMMSG) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JPMMSG.Descriptor instead.
func (JPMMSG) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_jpm_proto_rawDescGZIP(), []int{0}
}

// 游戏 图案枚举
type JPMID int32

const (
	JPMID_JPMUnknow JPMID = 0
	JPMID_JPM1      JPMID = 1  // 肚兜Wild
	JPMID_JPM2      JPMID = 2  // 庞春梅Scatter
	JPMID_JPM3      JPMID = 3  // 潘金莲Jackpot
	JPMID_JPM4      JPMID = 4  // 李瓶儿
	JPMID_JPM5      JPMID = 5  // 西门庆
	JPMID_JPM6      JPMID = 6  // 武大郎
	JPMID_JPM7      JPMID = 7  // 玉势
	JPMID_JPM8      JPMID = 8  // 缅玲
	JPMID_JPM9      JPMID = 9  // 木马
	JPMID_JPM10     JPMID = 10 // 春宫瓷器
	JPMID_JPM11     JPMID = 11 // 银托子
)

// Enum value maps for JPMID.
var (
	JPMID_name = map[int32]string{
		0:  "JPMUnknow",
		1:  "JPM1",
		2:  "JPM2",
		3:  "JPM3",
		4:  "JPM4",
		5:  "JPM5",
		6:  "JPM6",
		7:  "JPM7",
		8:  "JPM8",
		9:  "JPM9",
		10: "JPM10",
		11: "JPM11",
	}
	JPMID_value = map[string]int32{
		"JPMUnknow": 0,
		"JPM1":      1,
		"JPM2":      2,
		"JPM3":      3,
		"JPM4":      4,
		"JPM5":      5,
		"JPM6":      6,
		"JPM7":      7,
		"JPM8":      8,
		"JPM9":      9,
		"JPM10":     10,
		"JPM11":     11,
	}
)

func (x JPMID) Enum() *JPMID {
	p := new(JPMID)
	*p = x
	return p
}

func (x JPMID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JPMID) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_jpm_proto_enumTypes[1].Descriptor()
}

func (JPMID) Type() protoreflect.EnumType {
	return &file_protobuf_jpm_proto_enumTypes[1]
}

func (x JPMID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JPMID.Descriptor instead.
func (JPMID) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_jpm_proto_rawDescGZIP(), []int{1}
}

// 请求进入房间
type ENTER_GAME_JPM_REQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"` //
	RoomID    uint32 `protobuf:"varint,2,opt,name=RoomID,proto3" json:"RoomID,omitempty"`       // 房间ID
}

func (x *ENTER_GAME_JPM_REQ) Reset() {
	*x = ENTER_GAME_JPM_REQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jpm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ENTER_GAME_JPM_REQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ENTER_GAME_JPM_REQ) ProtoMessage() {}

func (x *ENTER_GAME_JPM_REQ) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jpm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ENTER_GAME_JPM_REQ.ProtoReflect.Descriptor instead.
func (*ENTER_GAME_JPM_REQ) Descriptor() ([]byte, []int) {
	return file_protobuf_jpm_proto_rawDescGZIP(), []int{0}
}

func (x *ENTER_GAME_JPM_REQ) GetAccountID() uint32 {
	if x != nil {
		return x.AccountID
	}
	return 0
}

func (x *ENTER_GAME_JPM_REQ) GetRoomID() uint32 {
	if x != nil {
		return x.RoomID
	}
	return 0
}

type ENTER_GAME_JPM_RES struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID   uint32   `protobuf:"varint,1,opt,name=RoomID,proto3" json:"RoomID,omitempty"`     // 房间ID
	Basics   int64    `protobuf:"varint,2,opt,name=Basics,proto3" json:"Basics,omitempty"`     // 基础金额
	Bonus    int64    `protobuf:"varint,3,opt,name=Bonus,proto3" json:"Bonus,omitempty"`       // 奖金池
	LastBet  int64    `protobuf:"varint,4,opt,name=LastBet,proto3" json:"LastBet,omitempty"`   // 最后一次压住
	Bets     []uint64 `protobuf:"varint,5,rep,packed,name=Bets,proto3" json:"Bets,omitempty"`  // 可以选择的押注金额
	FeeCount int32    `protobuf:"varint,8,opt,name=FeeCount,proto3" json:"FeeCount,omitempty"` // 免费次数
}

func (x *ENTER_GAME_JPM_RES) Reset() {
	*x = ENTER_GAME_JPM_RES{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jpm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ENTER_GAME_JPM_RES) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ENTER_GAME_JPM_RES) ProtoMessage() {}

func (x *ENTER_GAME_JPM_RES) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jpm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ENTER_GAME_JPM_RES.ProtoReflect.Descriptor instead.
func (*ENTER_GAME_JPM_RES) Descriptor() ([]byte, []int) {
	return file_protobuf_jpm_proto_rawDescGZIP(), []int{1}
}

func (x *ENTER_GAME_JPM_RES) GetRoomID() uint32 {
	if x != nil {
		return x.RoomID
	}
	return 0
}

func (x *ENTER_GAME_JPM_RES) GetBasics() int64 {
	if x != nil {
		return x.Basics
	}
	return 0
}

func (x *ENTER_GAME_JPM_RES) GetBonus() int64 {
	if x != nil {
		return x.Bonus
	}
	return 0
}

func (x *ENTER_GAME_JPM_RES) GetLastBet() int64 {
	if x != nil {
		return x.LastBet
	}
	return 0
}

func (x *ENTER_GAME_JPM_RES) GetBets() []uint64 {
	if x != nil {
		return x.Bets
	}
	return nil
}

func (x *ENTER_GAME_JPM_RES) GetFeeCount() int32 {
	if x != nil {
		return x.FeeCount
	}
	return 0
}

// 请求退出房间
type LEAVE_GAME_JPM_REQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID uint32 `protobuf:"varint,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"` //
	RoomID    uint32 `protobuf:"varint,2,opt,name=RoomID,proto3" json:"RoomID,omitempty"`       // 房间ID
}

func (x *LEAVE_GAME_JPM_REQ) Reset() {
	*x = LEAVE_GAME_JPM_REQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jpm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LEAVE_GAME_JPM_REQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LEAVE_GAME_JPM_REQ) ProtoMessage() {}

func (x *LEAVE_GAME_JPM_REQ) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jpm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LEAVE_GAME_JPM_REQ.ProtoReflect.Descriptor instead.
func (*LEAVE_GAME_JPM_REQ) Descriptor() ([]byte, []int) {
	return file_protobuf_jpm_proto_rawDescGZIP(), []int{2}
}

func (x *LEAVE_GAME_JPM_REQ) GetAccountID() uint32 {
	if x != nil {
		return x.AccountID
	}
	return 0
}

func (x *LEAVE_GAME_JPM_REQ) GetRoomID() uint32 {
	if x != nil {
		return x.RoomID
	}
	return 0
}

type LEAVE_GAME_JPM_RES struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret    uint32 `protobuf:"varint,1,opt,name=Ret,proto3" json:"Ret,omitempty"`       // 0.可以退出  1.不能退出房间
	RoomID uint32 `protobuf:"varint,2,opt,name=RoomID,proto3" json:"RoomID,omitempty"` // 房间ID
}

func (x *LEAVE_GAME_JPM_RES) Reset() {
	*x = LEAVE_GAME_JPM_RES{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jpm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LEAVE_GAME_JPM_RES) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LEAVE_GAME_JPM_RES) ProtoMessage() {}

func (x *LEAVE_GAME_JPM_RES) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jpm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LEAVE_GAME_JPM_RES.ProtoReflect.Descriptor instead.
func (*LEAVE_GAME_JPM_RES) Descriptor() ([]byte, []int) {
	return file_protobuf_jpm_proto_rawDescGZIP(), []int{3}
}

func (x *LEAVE_GAME_JPM_RES) GetRet() uint32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *LEAVE_GAME_JPM_RES) GetRoomID() uint32 {
	if x != nil {
		return x.RoomID
	}
	return 0
}

////////////////////////////////////////////// 游戏1 /////////////////////////////////////////////
// 请求开始游戏
type START_JPM_REQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bet uint64 `protobuf:"varint,1,opt,name=Bet,proto3" json:"Bet,omitempty"`
}

func (x *START_JPM_REQ) Reset() {
	*x = START_JPM_REQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jpm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *START_JPM_REQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*START_JPM_REQ) ProtoMessage() {}

func (x *START_JPM_REQ) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jpm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use START_JPM_REQ.ProtoReflect.Descriptor instead.
func (*START_JPM_REQ) Descriptor() ([]byte, []int) {
	return file_protobuf_jpm_proto_rawDescGZIP(), []int{4}
}

func (x *START_JPM_REQ) GetBet() uint64 {
	if x != nil {
		return x.Bet
	}
	return 0
}

type START_JPM_RES struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret          uint64         `protobuf:"varint,1,opt,name=Ret,proto3" json:"Ret,omitempty"`                        // 0成功
	SumOdds      int64          `protobuf:"varint,2,opt,name=SumOdds,proto3" json:"SumOdds,omitempty"`                // 9条线总赔率
	Results      []*JPM_Result  `protobuf:"bytes,3,rep,name=Results,proto3" json:"Results,omitempty"`                 // 中奖线列表
	PictureList  []int32        `protobuf:"varint,4,rep,packed,name=PictureList,proto3" json:"PictureList,omitempty"` // 图形列表( [0-14] 即15个图形)
	Bonus        int64          `protobuf:"varint,5,opt,name=Bonus,proto3" json:"Bonus,omitempty"`                    // 赢得水池里的金币
	Money        int64          `protobuf:"varint,6,opt,name=Money,proto3" json:"Money,omitempty"`                    // 身上的钱
	FreeCount    int64          `protobuf:"varint,7,opt,name=FreeCount,proto3" json:"FreeCount,omitempty"`            // 获得免费次数
	FeePositions []*JPMPosition `protobuf:"bytes,9,rep,name=FeePositions,proto3" json:"FeePositions,omitempty"`       //中免费的坐标
}

func (x *START_JPM_RES) Reset() {
	*x = START_JPM_RES{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jpm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *START_JPM_RES) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*START_JPM_RES) ProtoMessage() {}

func (x *START_JPM_RES) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jpm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use START_JPM_RES.ProtoReflect.Descriptor instead.
func (*START_JPM_RES) Descriptor() ([]byte, []int) {
	return file_protobuf_jpm_proto_rawDescGZIP(), []int{5}
}

func (x *START_JPM_RES) GetRet() uint64 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *START_JPM_RES) GetSumOdds() int64 {
	if x != nil {
		return x.SumOdds
	}
	return 0
}

func (x *START_JPM_RES) GetResults() []*JPM_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *START_JPM_RES) GetPictureList() []int32 {
	if x != nil {
		return x.PictureList
	}
	return nil
}

func (x *START_JPM_RES) GetBonus() int64 {
	if x != nil {
		return x.Bonus
	}
	return 0
}

func (x *START_JPM_RES) GetMoney() int64 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *START_JPM_RES) GetFreeCount() int64 {
	if x != nil {
		return x.FreeCount
	}
	return 0
}

func (x *START_JPM_RES) GetFeePositions() []*JPMPosition {
	if x != nil {
		return x.FeePositions
	}
	return nil
}

type JPMPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Px int32 `protobuf:"varint,1,opt,name=px,proto3" json:"px,omitempty"`
	Py int32 `protobuf:"varint,2,opt,name=py,proto3" json:"py,omitempty"`
}

func (x *JPMPosition) Reset() {
	*x = JPMPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jpm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JPMPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JPMPosition) ProtoMessage() {}

func (x *JPMPosition) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jpm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JPMPosition.ProtoReflect.Descriptor instead.
func (*JPMPosition) Descriptor() ([]byte, []int) {
	return file_protobuf_jpm_proto_rawDescGZIP(), []int{6}
}

func (x *JPMPosition) GetPx() int32 {
	if x != nil {
		return x.Px
	}
	return 0
}

func (x *JPMPosition) GetPy() int32 {
	if x != nil {
		return x.Py
	}
	return 0
}

type JPM_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LineId    int32          `protobuf:"varint,1,opt,name=LineId,proto3" json:"LineId,omitempty"` // 表示第几条线
	Count     int32          `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`   // 图形项连的个数
	Odds      int32          `protobuf:"varint,3,opt,name=Odds,proto3" json:"Odds,omitempty"`     // 该条线的赔率
	Positions []*JPMPosition `protobuf:"bytes,4,rep,name=Positions,proto3" json:"Positions,omitempty"`
}

func (x *JPM_Result) Reset() {
	*x = JPM_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jpm_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JPM_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JPM_Result) ProtoMessage() {}

func (x *JPM_Result) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jpm_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JPM_Result.ProtoReflect.Descriptor instead.
func (*JPM_Result) Descriptor() ([]byte, []int) {
	return file_protobuf_jpm_proto_rawDescGZIP(), []int{7}
}

func (x *JPM_Result) GetLineId() int32 {
	if x != nil {
		return x.LineId
	}
	return 0
}

func (x *JPM_Result) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *JPM_Result) GetOdds() int32 {
	if x != nil {
		return x.Odds
	}
	return 0
}

func (x *JPM_Result) GetPositions() []*JPMPosition {
	if x != nil {
		return x.Positions
	}
	return nil
}

// 通知更新奖金池
type UPDATE_JPM_BONUS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bonus int64 `protobuf:"varint,1,opt,name=Bonus,proto3" json:"Bonus,omitempty"`
}

func (x *UPDATE_JPM_BONUS) Reset() {
	*x = UPDATE_JPM_BONUS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jpm_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UPDATE_JPM_BONUS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UPDATE_JPM_BONUS) ProtoMessage() {}

func (x *UPDATE_JPM_BONUS) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jpm_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UPDATE_JPM_BONUS.ProtoReflect.Descriptor instead.
func (*UPDATE_JPM_BONUS) Descriptor() ([]byte, []int) {
	return file_protobuf_jpm_proto_rawDescGZIP(), []int{8}
}

func (x *UPDATE_JPM_BONUS) GetBonus() int64 {
	if x != nil {
		return x.Bonus
	}
	return 0
}

// 请求jpm玩家列表
type PLAYERS_JPM_LIST_RES struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*AccountStorageData `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"` // 房间内所有的玩家
}

func (x *PLAYERS_JPM_LIST_RES) Reset() {
	*x = PLAYERS_JPM_LIST_RES{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_jpm_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PLAYERS_JPM_LIST_RES) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PLAYERS_JPM_LIST_RES) ProtoMessage() {}

func (x *PLAYERS_JPM_LIST_RES) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_jpm_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PLAYERS_JPM_LIST_RES.ProtoReflect.Descriptor instead.
func (*PLAYERS_JPM_LIST_RES) Descriptor() ([]byte, []int) {
	return file_protobuf_jpm_proto_rawDescGZIP(), []int{9}
}

func (x *PLAYERS_JPM_LIST_RES) GetPlayers() []*AccountStorageData {
	if x != nil {
		return x.Players
	}
	return nil
}

var File_protobuf_jpm_proto protoreflect.FileDescriptor

var file_protobuf_jpm_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6a, 0x70, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x73, 0x67, 0x1a, 0x13,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x12, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x47, 0x41, 0x4d,
	0x45, 0x5f, 0x4a, 0x50, 0x4d, 0x5f, 0x52, 0x45, 0x51, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x22,
	0xa4, 0x01, 0x0a, 0x12, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4a,
	0x50, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x42, 0x61, 0x73, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x4c, 0x61, 0x73, 0x74, 0x42, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4c,
	0x61, 0x73, 0x74, 0x42, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x65, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x04, 0x42, 0x65, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x65,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x65,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x12, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x5f,
	0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4a, 0x50, 0x4d, 0x5f, 0x52, 0x45, 0x51, 0x12, 0x1c, 0x0a, 0x09,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f,
	0x6f, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d,
	0x49, 0x44, 0x22, 0x3e, 0x0a, 0x12, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x5f, 0x47, 0x41, 0x4d, 0x45,
	0x5f, 0x4a, 0x50, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x52, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f,
	0x6f, 0x6d, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d,
	0x49, 0x44, 0x22, 0x21, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x4a, 0x50, 0x4d, 0x5f,
	0x52, 0x45, 0x51, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x42, 0x65, 0x74, 0x22, 0x93, 0x02, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f,
	0x4a, 0x50, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x52, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x6d,
	0x4f, 0x64, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x75, 0x6d, 0x4f,
	0x64, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x73, 0x67, 0x2e,
	0x4a, 0x50, 0x4d, 0x5f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x72, 0x65, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x46, 0x72, 0x65, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3a, 0x0a, 0x0c, 0x46, 0x65, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x73, 0x67,
	0x2e, 0x4a, 0x50, 0x4d, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x46,
	0x65, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x2e, 0x0a, 0x0c, 0x4a,
	0x50, 0x4d, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x70,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x70, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x70,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x70, 0x79, 0x22, 0x84, 0x01, 0x0a, 0x0a,
	0x4a, 0x50, 0x4d, 0x5f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4c, 0x69, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x64, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x4f, 0x64, 0x64, 0x73, 0x12, 0x34, 0x0a, 0x09,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x73, 0x67, 0x2e, 0x4a, 0x50, 0x4d, 0x5f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x28, 0x0a, 0x10, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4a, 0x50, 0x4d,
	0x5f, 0x42, 0x4f, 0x4e, 0x55, 0x53, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x22, 0x4e, 0x0a, 0x14,
	0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x53, 0x5f, 0x4a, 0x50, 0x4d, 0x5f, 0x4c, 0x49, 0x53, 0x54,
	0x5f, 0x52, 0x45, 0x53, 0x12, 0x36, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x73, 0x67,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2a, 0x8c, 0x02, 0x0a,
	0x06, 0x4a, 0x50, 0x4d, 0x4d, 0x53, 0x47, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x5f, 0x4a, 0x50, 0x4d, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x15, 0x43, 0x53, 0x5f, 0x45, 0x4e,
	0x54, 0x45, 0x52, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4a, 0x50, 0x4d, 0x5f, 0x52, 0x45, 0x51,
	0x10, 0xb1, 0x6d, 0x12, 0x1a, 0x0a, 0x15, 0x53, 0x43, 0x5f, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f,
	0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4a, 0x50, 0x4d, 0x5f, 0x52, 0x45, 0x53, 0x10, 0xb2, 0x6d, 0x12,
	0x1a, 0x0a, 0x15, 0x43, 0x53, 0x5f, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x5f, 0x47, 0x41, 0x4d, 0x45,
	0x5f, 0x4a, 0x50, 0x4d, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xb3, 0x6d, 0x12, 0x1a, 0x0a, 0x15, 0x53,
	0x43, 0x5f, 0x4c, 0x45, 0x41, 0x56, 0x45, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4a, 0x50, 0x4d,
	0x5f, 0x52, 0x45, 0x53, 0x10, 0xb4, 0x6d, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x53, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x5f, 0x4a, 0x50, 0x4d, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xb8, 0x6d, 0x12, 0x15,
	0x0a, 0x10, 0x53, 0x43, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x4a, 0x50, 0x4d, 0x5f, 0x52,
	0x45, 0x53, 0x10, 0xb9, 0x6d, 0x12, 0x18, 0x0a, 0x13, 0x53, 0x43, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x4a, 0x50, 0x4d, 0x5f, 0x42, 0x4f, 0x4e, 0x55, 0x53, 0x10, 0xba, 0x6d, 0x12,
	0x1c, 0x0a, 0x17, 0x43, 0x53, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x53, 0x5f, 0x4a, 0x50,
	0x4d, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x10, 0xbf, 0x6d, 0x12, 0x1c, 0x0a,
	0x17, 0x53, 0x43, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x53, 0x5f, 0x4a, 0x50, 0x4d, 0x5f,
	0x4c, 0x49, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x10, 0xc0, 0x6d, 0x2a, 0x86, 0x01, 0x0a, 0x05,
	0x4a, 0x50, 0x4d, 0x49, 0x44, 0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x50, 0x4d, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x50, 0x4d, 0x31, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x4a, 0x50, 0x4d, 0x32, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x50, 0x4d, 0x33,
	0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x50, 0x4d, 0x34, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04,
	0x4a, 0x50, 0x4d, 0x35, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x50, 0x4d, 0x36, 0x10, 0x06,
	0x12, 0x08, 0x0a, 0x04, 0x4a, 0x50, 0x4d, 0x37, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x50,
	0x4d, 0x38, 0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x50, 0x4d, 0x39, 0x10, 0x09, 0x12, 0x09,
	0x0a, 0x05, 0x4a, 0x50, 0x4d, 0x31, 0x30, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x4a, 0x50, 0x4d,
	0x31, 0x31, 0x10, 0x0b, 0x42, 0x0f, 0x5a, 0x0d, 0x72, 0x6f, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_jpm_proto_rawDescOnce sync.Once
	file_protobuf_jpm_proto_rawDescData = file_protobuf_jpm_proto_rawDesc
)

func file_protobuf_jpm_proto_rawDescGZIP() []byte {
	file_protobuf_jpm_proto_rawDescOnce.Do(func() {
		file_protobuf_jpm_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_jpm_proto_rawDescData)
	})
	return file_protobuf_jpm_proto_rawDescData
}

var file_protobuf_jpm_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protobuf_jpm_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protobuf_jpm_proto_goTypes = []interface{}{
	(JPMMSG)(0),                  // 0: protomsg.JPMMSG
	(JPMID)(0),                   // 1: protomsg.JPMID
	(*ENTER_GAME_JPM_REQ)(nil),   // 2: protomsg.ENTER_GAME_JPM_REQ
	(*ENTER_GAME_JPM_RES)(nil),   // 3: protomsg.ENTER_GAME_JPM_RES
	(*LEAVE_GAME_JPM_REQ)(nil),   // 4: protomsg.LEAVE_GAME_JPM_REQ
	(*LEAVE_GAME_JPM_RES)(nil),   // 5: protomsg.LEAVE_GAME_JPM_RES
	(*START_JPM_REQ)(nil),        // 6: protomsg.START_JPM_REQ
	(*START_JPM_RES)(nil),        // 7: protomsg.START_JPM_RES
	(*JPMPosition)(nil),          // 8: protomsg.JPM_position
	(*JPM_Result)(nil),           // 9: protomsg.JPM_Result
	(*UPDATE_JPM_BONUS)(nil),     // 10: protomsg.UPDATE_JPM_BONUS
	(*PLAYERS_JPM_LIST_RES)(nil), // 11: protomsg.PLAYERS_JPM_LIST_RES
	(*AccountStorageData)(nil),   // 12: protomsg.AccountStorageData
}
var file_protobuf_jpm_proto_depIdxs = []int32{
	9,  // 0: protomsg.START_JPM_RES.Results:type_name -> protomsg.JPM_Result
	8,  // 1: protomsg.START_JPM_RES.FeePositions:type_name -> protomsg.JPM_position
	8,  // 2: protomsg.JPM_Result.Positions:type_name -> protomsg.JPM_position
	12, // 3: protomsg.PLAYERS_JPM_LIST_RES.players:type_name -> protomsg.AccountStorageData
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_protobuf_jpm_proto_init() }
func file_protobuf_jpm_proto_init() {
	if File_protobuf_jpm_proto != nil {
		return
	}
	file_protobuf_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protobuf_jpm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ENTER_GAME_JPM_REQ); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_jpm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ENTER_GAME_JPM_RES); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_jpm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LEAVE_GAME_JPM_REQ); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_jpm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LEAVE_GAME_JPM_RES); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_jpm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*START_JPM_REQ); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_jpm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*START_JPM_RES); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_jpm_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JPMPosition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_jpm_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JPM_Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_jpm_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UPDATE_JPM_BONUS); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protobuf_jpm_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PLAYERS_JPM_LIST_RES); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_jpm_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_jpm_proto_goTypes,
		DependencyIndexes: file_protobuf_jpm_proto_depIdxs,
		EnumInfos:         file_protobuf_jpm_proto_enumTypes,
		MessageInfos:      file_protobuf_jpm_proto_msgTypes,
	}.Build()
	File_protobuf_jpm_proto = out.File
	file_protobuf_jpm_proto_rawDesc = nil
	file_protobuf_jpm_proto_goTypes = nil
	file_protobuf_jpm_proto_depIdxs = nil
}
