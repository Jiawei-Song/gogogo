package main

import (
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

// 根据消息的种类，判断调用那个函数来进行处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登陆的逻辑
		err = serverLoginMes(conn, mes)
	case message.RegisterMesType:
		// 处理注册的逻辑
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return
}
