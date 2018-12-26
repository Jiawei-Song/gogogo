package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroomRebuild/common/message"
	"go_code/chatroomRebuild/server/utils"
	"net"
)

// SmsProcess 处理消息的结构体
type SmsProcess struct {
}

// SendGroupMes 服务端处理消息的函数
func (smsProcess *SmsProcess) SendGroupMes(mes *message.Message) (err error) {
	// 遍历在线用户，把消息发给他们
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("服务端 SendGroupMes json.Unmarshal fail, err = ", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("服务端 SendGroupMes json.Marshal(mes) fail, err = ", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserID {
			continue
		}
		err = smsProcess.SendMesToEveryOnlineUser(data, up.Conn)
		if err != nil {
			return err
		}
	}
	return
}

// SendMesToEveryOnlineUser 给在线用户发送消息
func (smsProcess *SmsProcess) SendMesToEveryOnlineUser(data []byte, conn net.Conn) (err error) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("conn.Write data 失败", err)
		return
	}
	return
}
