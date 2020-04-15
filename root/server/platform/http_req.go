package platform

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"root/core"
	"root/core/log"
	"root/core/utils"
	"strconv"
	"time"
)

func Asyn_addMoney(trycount int, addr_url, unique string, num int64, roomID int32, gamestring, desc string, back func(backunique string, backmoney int64, bwType int32), errback func()) {
	if trycount == 0 {
		errback()
		return
	}
	desc += " time:" + utils.DateString()
	go func() {
		utils.Try(func() {
			send := url.Values{"channelId": {"DDHYLC"},
				"gameId": {gamestring},
				"userId": {unique},
				"num":    {strconv.Itoa(int(num))},
				"desc":   {desc},
			}
			log.Infof("http-request try:%v :%v addr:%v ", trycount, send, addr_url)
			resp, err := http.PostForm(addr_url,
				send)

			if err != nil {
				if resp != nil && resp.Body != nil {
					resp.Body.Close()
				}
				log.Errorf("三方平台，http 请求错误:%v", err.Error())
				for i := 0; i < 10; i++ {
					time.Sleep(1 * time.Second)
					resp, err = http.PostForm(addr_url,
						send)
					if err == nil {
						break
					} else {
						if resp != nil && resp.Body != nil {
							resp.Body.Close()
						}
						log.Errorf("三方平台，http %v 请求错误:%v send:%v", i+1, err.Error(), send)
					}
				}

				// 10次请求以后，如果还有错，就直接返回了
				if err != nil {
					if errback != nil {
						core.LocalCoreSend(0, roomID, func() {
							errback()
						})
					}
					return
				}
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Errorf("三方平台，read 错误:%v", err.Error())
				return
			}
			log.Infof("http-response try:%v :%v", trycount, string(body))
			var jsonstr map[string]interface{}
			e := json.Unmarshal(body, &jsonstr)
			if e != nil {
				log.Errorf("json 解析错误:%v ", e.Error())
				return
			}
			if err, e := jsonstr["status"]; e && int(err.(float64)) != 0 {
				log.Errorf("平台返回错误码:%v ", int(err.(float64)))
				Asyn_addMoney(trycount-1, addr_url, unique, num, roomID, gamestring, desc, back, errback)
				return
			} else {
				data := jsonstr["data"].(map[string]interface{})
				gold := data["gold"].(float64)
				bwType := data["bwType"].(float64)
				if back != nil {
					core.LocalCoreSend(0, roomID, func() {
						back(unique, int64(gold), int32(bwType))
					})
				}
			}
		})
	}()
}
