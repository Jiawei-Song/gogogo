package process

import (
	"fmt"
	"go_code/chatroomRebuild/client/utils"
	"net"
	"os"
)

//ShowMenu 登陆成功后显示聊天室的界面
func ShowMenu() {
	var key int
	// var loop = true
	for {
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
		case 2:
			fmt.Println("发送消息")
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

func serverProcessMes(conn net.Conn) {
	fmt.Println("进入聊天室后启动了一个协程，监听服务端返回的消息")
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("正在读取服务端返回的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("serverProcessMes 里的 tf.ReadPkg 出错， err = ", err)
			return
		}
		fmt.Println(mes)
	}
}
