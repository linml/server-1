package room

import (
	"root/core/log"
	"root/core/log/colorized"
	"root/core/packet"
	"root/core/utils"
)

type (
	game struct {
		*Room
		s         ERoomStatus
		timestamp int64
	}
)

func (self *game) Enter(now int64) {
	self.timestamp = utils.MilliSecondTimeSince1970() + int64(/*config.GetPublicConfig_Int64("HB_SETTLEMENT_TIME")**/1000) // 秒
	log.Infof(colorized.Green("房间:%v game enter duration  water:%v"), self.roomId, RoomMgr.Water_line)
}

func (self *game) Tick(now int64) {
	if utils.MilliSecondTimeSince1970() < self.timestamp {
		return
	}
}

func (self *game) Leave(now int64) {
	log.Debugf(colorized.Green("game leave\n"))
}

func (self *game) Handle(actor int32, msg []byte, session int64) bool {
	pack := packet.NewPacket(msg)
	switch pack.GetMsgID() {
	//case protomsg.Old_MSGID_ENTER_GAME.UInt16(): // 客户端链接进入游戏
	//	self.Old_MSGID_ENTER_GAME(actor, msg, session)
	//case protomsg.Old_MSGID_HONGBAO_ROB_HONGBAO.UInt16(): // 抢红包
	//	self.Old_MSGID_HONGBAO_ROB_HONGBAO(actor, msg, session)
	default:
		//log.Warnf("game 状态 没有处理消息msgId:%v", pack.GetMsgID())
		return false
	}
	return true
}

func (self *game) Old_MSGID_ENTER_GAME(actor int32, msg []byte, session int64) {

}