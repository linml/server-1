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
	time2 "time"
)

type (
	Game struct {
		owner          *core.Actor
		init           bool // 重新建立连接是否需要拉取所有数据
		roomID         uint32
		feec		int
	}
)

var count = 0
var fee = 0
var total = 0
func NewGame() *Game {
	return &Game{}
}
var game_GLobal *Client
func (self *Game) Init(actor *core.Actor) bool {
	self.owner = actor
	
	//game_GLobal = NewWebsocketClient("47.108.87.29:41201","/connect")
	game_GLobal = NewWebsocketClient("192.168.2.100:41401","/connect")
	game_GLobal.connect()
	fmt.Println("connected success :", game_GLobal.ws.RemoteAddr())
	go func() {
		for {
			recv := make([]byte,65535)
			n,err := game_GLobal.ws.Read(recv)
			if err != nil{
				log.Warnf("err:%v",err.Error())
				continue
			}
			//log.Infof("读出%v个字节",n)
			recv = recv[0:n]
			buffer := new(bytes.Buffer)
			buffer.Write(recv)
			_, content, errcode := self.decode(buffer)
			if errcode != 0{
				log.Warnf("错误:%v",errcode)
			}
			self.HandleMessage(0, content,0)
		}
	}()

	Send2Game(protomsg.JPMMSG_CS_ENTER_GAME_JPM_REQ.UInt16(),&protomsg.ENTER_GAME_JPM_REQ{
		AccountID:AccountID,
		RoomID:self.roomID,
	})

	self.owner.AddTimer(30000,-1, func(dt int64) {
		Send2Game(protomsg.MSG_CLIENT_KEEPALIVE.UInt16(),nil)
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
	case protomsg.JPMMSG_SC_ENTER_GAME_JPM_RES.UInt16():
		pb := packet.PBUnmarshal(pack.ReadBytes(),&protomsg.ENTER_GAME_JPM_RES{}).(*protomsg.ENTER_GAME_JPM_RES)
		log.Infof(colorized.Blue("进入游戏成功：%+v"),pb)

	case protomsg.JPMMSG_SC_START_JPM_RES.UInt16():
		pb := packet.PBUnmarshal(pack.ReadBytes(),&protomsg.START_JPM_RES{}).(*protomsg.START_JPM_RES)
		log.Infof(colorized.Blue("开始游戏：%+v"),pb)
		//time2.Sleep(time2.Duration(rand.Int63n(10000))*time2.Microsecond)
		self.feec = int(pb.GetFreeCount())
		if self.feec > 0{
			Send2Game(protomsg.JPMMSG_CS_START_JPM_REQ.UInt16(),&protomsg.START_JPM_REQ{Bet:uint64(1000)})
			return true
		}

		count--
		if count > 0 {
			Send2Game(protomsg.JPMMSG_CS_START_JPM_REQ.UInt16(),&protomsg.START_JPM_REQ{Bet:uint64(1000)})
		}else{
			log.Infof("身上的钱--:%v", pb.GetMoney())
		}

		if count % 100000 == 0{
			log.Infof("sleep start")
			time2.Sleep(1346*time2.Millisecond)
			log.Infof("sleep end")
		}
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