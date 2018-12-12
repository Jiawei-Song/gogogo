package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroomRebuild/common/message"
	"go_code/chatroomRebuild/server/utils"
	"net"
)

// UserProcess 处理用户相关的结构体ß
type UserProcess struct {
	Conn net.Conn
}

// ServerLoginMes 服务端处理登陆消息的函数
func (userProcess *UserProcess) ServerLoginMes(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail , err = ", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	// 如果用户 id 为 100 ，密码为 123456， 就合法
	if loginMes.UserID == 100 && loginMes.UserPWD == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在"
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail , err = ", err)
	}
	resMes.Data = string(data)

	// resMes 构建完成，序列化之后发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail , err = ", err)
	}
	fmt.Println("这个是服务端返回的数据， data =", data)

	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg(data)

	return
}
