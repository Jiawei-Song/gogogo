package main

import (
	"fmt"
	"go_code/chatroomRebuild/common/message"
	"go_code/chatroomRebuild/server/process"
	"go_code/chatroomRebuild/server/utils"
	"io"
	"net"
)

// Processor 处理的结构体
type Processor struct {
	Conn net.Conn
}

// ServerProcessMes 根据消息的种类，判断调用那个函数来进行处理
func (processor *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登陆的逻辑   创建一个用户登陆处理的实例
		up := &process.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerLoginMes(mes)
	case message.RegisterMesType:
		// 处理注册的逻辑
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return
}

//LoopProcess 循环处理的主函数
func (processor *Processor) LoopProcess() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: processor.Conn,
		}
		// 将读取数据包进行一个封装，readPkg(),返回一个Message, Err
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("读完了")
				return err
			}
			fmt.Println("readPkg 出错, err = ", err)
			return err
		}
		fmt.Println(mes)
		err = processor.serverProcessMes(&mes)
		if err != nil {
			fmt.Println()
			return err
		}
	}
	return
}
