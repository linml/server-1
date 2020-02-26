package logic

import (
	"root/common"
	"root/common/config"
	"root/common/tools"
	"root/core"
	"root/core/log"
	"root/core/log/colorized"
	"root/core/network"
	"root/core/packet"
	"root/core/utils"
	"fmt"
	"root/protomsg"
	"root/server/red2black/account"
	"root/server/red2black/room"
	"root/server/red2black/send_tools"
	"strconv"
)

type (
	Red2Black struct {
		owner     *core.Actor
		con_count int
	}
)

func NewRed2Black() *Red2Black {
	return &Red2Black{}
}

func (self *Red2Black) Init(actor *core.Actor) bool {
	// 先处理脚本
	if core.ScriptDir != "" {
		core.InitScript(core.ScriptDir)
	}

	self.owner = actor

	// 连接hall
	hallIP := core.CoreAppConfString("connectHall")
	connectHall_actor := network.NewTCPClient(self.owner, func() string {
		return core.CoreAppConfString("connectHall")
	}, self.registerHall)
	child := core.NewActor(common.EActorType_CONNECT_HALL.Int32(), connectHall_actor, make(chan core.IMessage, 1000))
	core.CoreRegisteActor(child)
	log.Infof(colorized.Yellow("连接hall:[%v]"), hallIP)
	return true
}

// 向hall注册
func (self *Red2Black) registerHall() {
	// 发送注册消息登记自身信息
	sid, _ := strconv.Atoi(core.Appname)
	count := uint32(room.RoomMgr.RoomCount())

	GAME_TO_HALL_MAP_KEY := config.GetPublicConfig_String("GAME_TO_HALL_MAP_KEY")
	strSign := fmt.Sprintf("%v%v%v", sid, count, GAME_TO_HALL_MAP_KEY)
	strSign = tools.MD5(strSign)

	// 组装消息
	Databytes := packet.NewPacket(nil)
	Databytes.WriteUInt16(uint16(sid))
	Databytes.WriteUInt32(count)
	Databytes.WriteString(strSign)
	Databytes.SetMsgID(protomsg.Old_MSGID_SS_MAPING.UInt16())
	core.CoreSend(self.owner.Id, common.EActorType_CONNECT_HALL.Int32(), Databytes.GetData(), 0)
	log.Infof("连接大厅成功  sid:%v", sid)

	room.RoomMgr.InitRoomMgr()

	self.StartService()
}

func (self *Red2Black) StartService() {
	// 监听端口，客户端连接用
	var customer []*core.Actor
	customer = append(customer, self.owner)
	listen_actor := network.NewTCPServer(customer, beego.AppConfig.DefaultString(core.Appname+"::listen", ""))
	child := core.NewActor(common.EActorType_SERVER.Int32(), listen_actor, make(chan core.IMessage, 10000))
	core.CoreRegisteActor(child)
}

func (self *Red2Black) Stop() {

}

func (self *Red2Black) Old_MSGID_BACKSTAGE_SET_WATER_LINE(actor int32, msg []byte, session int64) {
	pack := packet.NewPacket(msg)
	nBet := pack.ReadUInt32()
	iWaterLine := pack.ReadInt32()

	if session != 0 {
		log.Warnf("Error, 异常session:%v 处理消息编号:%v", session, pack.GetMsgID())
		return
	}

	iOldWaterLine := room.RoomMgr.Water_line
	room.RoomMgr.Water_line = int64(iWaterLine)

	log.Infof("后台设置水位线, ServerID:%v, OldWaterLine:%v, NewWaterLine:%v", core.Appname, iOldWaterLine, room.RoomMgr.Water_line)
	room.RoomMgr.SaveWaterLine()

	strLog := fmt.Sprintf("INSERT INTO log_water_line (log_ServerID, log_OldValue, log_NewValue, log_Bet, log_Time) VALUES (%v, %v, %v, %v, '%v');", core.Appname, iOldWaterLine, iWaterLine, nBet, utils.DateString())
	tSendToHall := packet.NewPacket(nil)
	tSendToHall.SetMsgID(protomsg.MSGID_SAVE_LOG.UInt16())
	tSendToHall.WriteString(strLog)
	tSendToHall.WriteUInt16(1) // 日志类型, 大厅可将多条日志组装成一个消息回存
	send_tools.Send2Hall(tSendToHall.GetData())
}

func (self *Red2Black) HandlerInitWaterLine(actor int32, msg []byte, session int64) {
	pack := packet.NewPacket(msg)
	strWaterLine := pack.ReadString()
	iWaterLine, e := strconv.ParseInt(strWaterLine, 10, 64)
	if e == nil {
		room.RoomMgr.Water_line = iWaterLine
	} else {
		room.RoomMgr.Water_line = 0
	}
	log.Infof("================= OnCompleteLoadWaterLine:%v", room.RoomMgr.Water_line)
}

func (self *Red2Black) HandleMessage(actor int32, msg []byte, session int64) bool {
	pack := packet.NewPacket(msg)
	switch pack.GetMsgID() {
	case protomsg.Old_MSGID_MAINTENANCE_NOTICE.UInt16(): // 关服通知
		self.Old_MSGID_MAINTENANCE_NOTICE(actor, msg, session)
	case protomsg.Old_MSGID_CREATE_ROOM.UInt16(): // 大厅请求创建房间
		self.Old_MSGID_CREATE_ROOM(actor, msg, session)
	case protomsg.Old_MSGID_RECV_ACCOUNT_INFO.UInt16(): // 同步账号数据
		self.Old_MSGID_RECV_ACCOUNT_INFO(actor, msg, session)
	case protomsg.Old_MSGID_ENTER_GAME.UInt16(), protomsg.Old_MSGID_R2B_GAME_ENTER_GAME.UInt16(): // 客户端链接进入游戏
		self.Old_MSGID_R2B_ENTER_GAME(actor, msg, session)
	case protomsg.MSGID_GET_ONE_WATERLINE.UInt16(): // 初始化水位线
		self.HandlerInitWaterLine(actor, msg, session)
	case protomsg.Old_MSGID_BACKSTAGE_SET_WATER_LINE.UInt16(): // 后台设置水位线
		self.Old_MSGID_BACKSTAGE_SET_WATER_LINE(actor, msg, session)
	case protomsg.Old_MSGID_CHANGE_PLAYER_INFO.UInt16(): // 修改玩家信息
		self.Old_MSGID_CHANGE_PLAYER_INFO(actor, msg, session)
	case protomsg.Old_MSGID_CLIENT_KEEPALIVE.UInt16():
		send_tools.Send2Account(msg, session)
	case protomsg.Old_MSGID_SS_TEST_NETWORK.UInt16():
		log.Infof("收到测试网络消息 SessionID:%v", session)
	case protomsg.Old_MSGID_HG_TEST_NETWORK.UInt16():
		log.Infof("收到来自大厅的测试网络消息 SessionID:%v", session)
		req := packet.NewPacket(nil)
		req.SetMsgID(protomsg.Old_MSGID_HG_TEST_NETWORK.UInt16())
		send_tools.Send2Hall(req.GetData())
	default: // 客户端游戏消息，统一发送给房间处理
		acc := account.AccountMgr.GetAccountBySessionID(session)
		if acc == nil {
			log.Warnf("找不到session 关联的玩家 session:%v msgId：%v", session, pack.GetMsgID())
			return false
		}
		core.CoreSend(self.owner.Id, int32(acc.RoomID), msg, session)
		break
	}
	return true
}

func (self *Red2Black) Old_MSGID_MAINTENANCE_NOTICE(actor int32, msg []byte, session int64) {
	pack := packet.NewPacket(msg)
	if session != 0 {
		log.Warnf("Error, 异常session:%v 处理消息编号:%v", session, pack.GetMsgID())
		return
	}
	room.Close(nil)
}
