package main

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"root/core"
	"root/core/log"
	"root/core/log/colorized"
	"root/core/packet"
	"root/protomsg"
)

type (
	Game struct {
		owner  *core.Actor
		init   bool // 重新建立连接是否需要拉取所有数据
		roomID uint32
		feec   int
	}
)

var count = 0
var fee = 0
var total = 0
var retmap = make(map[int]int)

func NewGame() *Game {
	return &Game{}
}

var game_GLobal *Client

func (self *Game) Init(actor *core.Actor) bool {
	self.owner = actor

	game_GLobal = NewWebsocketClient(addr+":41901", "/connect")
	game_GLobal.connect()
	if game_GLobal.ws == nil {
		log.Printf("connect faild \r\n")
		return false
	}
	fmt.Println("connected success :", game_GLobal.ws.RemoteAddr())
	go func() {
		for {
			recv := make([]byte, 65535)
			n, err := game_GLobal.ws.Read(recv)
			if err != nil {
				log.Warnf("err:%v", err.Error())
				continue
			}
			//log.Infof("读出%v个字节",n)
			recv = recv[0:n]
			buffer := new(bytes.Buffer)
			buffer.Write(recv)
			_, content, errcode := self.decode(buffer)
			if errcode != 0 {
				log.Warnf("错误:%v", errcode)
			}
			self.HandleMessage(0, content, 0)
		}
	}()

	Send2Game(protomsg.HBMSG_CS_ENTER_GAME_HB_REQ.UInt16(), &protomsg.ENTER_GAME_HB_REQ{
		AccountID: AccountID,
		RoomID:    self.roomID,
	})

	self.owner.AddTimer(30000, -1, func(dt int64) {
		Send2Game(protomsg.MSG_CLIENT_KEEPALIVE.UInt16(), nil)
	})
	return true
}

// 解密&处理粘包
func (self *Game) decode(buffer *bytes.Buffer) (uint16, []byte, int) {
	buff_len := uint16(buffer.Len())
	if buff_len < packet.PACKET_HEAD_LEN { // 长度不够
		return 0, nil, 0
	}

	origin := buffer.Bytes()
	// 获取长度
	var msg_len uint16 = 0
	msg_len = uint16(origin[4])
	msg_len |= uint16(origin[5]) << 8
	if msg_len < packet.PACKET_HEAD_LEN || msg_len > packet.PACKET_BUFFER_LEN || msg_len > buff_len {
		return 0, nil, 0
	}

	content := buffer.Next(int(msg_len))

	newBuf := make([]byte, 0, msg_len)
	newBuf = append(newBuf, content...)

	// 解析msgid
	var msgid uint16 = 0
	msgid = uint16(newBuf[2])
	msgid |= uint16(newBuf[3]) << 8
	return msgid, newBuf, 0
}

func (self *Game) registerHall() {

}

func (self *Game) Stop() {

}

func (self *Game) HandleMessage(actor int32, msg []byte, session int64) bool {
	pack := packet.NewPacket(msg)
	switch pack.GetMsgID() {
	case protomsg.HBMSG_SC_ENTER_GAME_HB_RES.UInt16():
		pb := packet.PBUnmarshal(pack.ReadBytes(), &protomsg.ENTER_GAME_HB_RES{}).(*protomsg.ENTER_GAME_HB_RES)
		log.Infof(colorized.Blue("进入游戏成功：%+v"), pb)

	case protomsg.HBMSG_SC_ASSIGN_HB_RES.UInt16(): // 发红包结果
		pb := packet.PBUnmarshal(pack.ReadBytes(), &protomsg.ASSIGN_HB_RES{}).(*protomsg.ASSIGN_HB_RES)
		log.Infof("%v ", pb.String())

	case protomsg.HBMSG_SC_BROADCAST_NEW_HB.UInt16(): // 新红包
		pb := packet.PBUnmarshal(pack.ReadBytes(), &protomsg.BROADCAST_NEW_HB{}).(*protomsg.BROADCAST_NEW_HB)
		log.Infof("%v ", pb.String())

	case protomsg.HBMSG_SC_GRAB_HB_RES.UInt16(): // 抢红包结果
		pb := packet.PBUnmarshal(pack.ReadBytes(), &protomsg.GRAB_HB_RES{}).(*protomsg.GRAB_HB_RES)
		log.Infof("%v ", pb.String())

	case protomsg.HBMSG_SC_BROADCAST_UPDATE_HB.UInt16():
		pb := packet.PBUnmarshal(pack.ReadBytes(), &protomsg.BROADCAST_UPDATE_HB{}).(*protomsg.BROADCAST_UPDATE_HB)
		log.Infof("更新红包数量:%v ", pb.String())

	case protomsg.HBMSG_SC_BROADCAST_UPDATE_GRAB.UInt16():
		pb := packet.PBUnmarshal(pack.ReadBytes(), &protomsg.BROADCAST_UPDATE_GRAB{}).(*protomsg.BROADCAST_UPDATE_GRAB)
		log.Infof("抢到红包：%v ", pb.String())

	case protomsg.HBMSG_SC_BROADCAST_UPDATE_BOMB.UInt16():
		pb := packet.PBUnmarshal(pack.ReadBytes(), &protomsg.BROADCAST_UPDATE_BOMB{}).(*protomsg.BROADCAST_UPDATE_BOMB)
		log.Infof("中雷：%v ", pb.String())
	}

	return true
}

func Send2Game(msgId uint16, pb proto.Message) {
	var bytes []byte
	if pb == nil {
		bytes = []byte{}
	} else {
		data, error := proto.Marshal(pb)
		if error != nil {
			log.Errorf("发送数据出错 :%v", error.Error())
			return
		}
		bytes = data
	}
	req := packet.NewPacket(nil)
	req.SetMsgID(msgId)
	req.WriteBytes(bytes)
	game_GLobal.SendMessage(req.GetData())
}
