package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroomRebuild/client/utils"
	"go_code/chatroomRebuild/common/message"
)

// SmsProcess 处理消息的结构体
type SmsProcess struct{}

// SendGroupMes 客户端发送群聊消息的函数
func (smsProcess *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserID = CurUser.UserID
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("smsMes 序列化失败， err =", err)
		return
	}
	mes.Data = string(data)
	// 将mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes序列化失败， err =", err)
		return
	}
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("conn.Write data 失败", err)
		return
	}
	return
}
