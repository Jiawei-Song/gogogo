package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroomRebuild/client/utils"
	"go_code/chatroomRebuild/common/message"
	"net"
	"os"
)

//ShowMenu 登陆成功后显示聊天室的界面
func ShowMenu() {
	var key int
	// var loop = true
	for {
		smsProcess := &SmsProcess{}
		fmt.Println("-------------恭喜XXX登陆成功------------")
		fmt.Println("            1.显示在线用户列表")
		fmt.Println("            2.发送消息")
		fmt.Println("            3.消息列表")
		fmt.Println("            4.退   出")
		fmt.Println("请选择（1-4）")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("显示在线用户列表")
			showOnlineUsers()
		case 2:
			fmt.Println("请输入消息")
			var content string
			fmt.Scanln(&content)
			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("消息列表")
		case 4:
			fmt.Println("退   出")
			os.Exit(0)
		default:
			fmt.Println("输入的信息有误，请重新输入")
		}
		// fmt.Println(loop)
	}
}

// serverProcessMes 客户端跑的一个协程，用于监听服务端返回的消息
func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("serverProcessMes 里的 tf.ReadPkg 出错， err = ", err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			// 上下线的消息 一顿处理。。。 先反序列化，再调用写好的函数
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("serverProcessMes 函数中NotifyUserStatusMes反序列化失败")
				return
			}
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			// 聊天的消息 一顿处理。。。 先反序列化，再调用写好的函数
			var smsMes message.SmsMes
			err = json.Unmarshal([]byte(mes.Data), &smsMes)
			if err != nil {
				fmt.Println("serverProcessMes 函数中SmsMes反序列化失败")
				return
			}
			outputGroupMes(&smsMes)
		default:
			fmt.Println("服务端返回的消息类型处理不了")
		}
		fmt.Println(mes)
	}
}
