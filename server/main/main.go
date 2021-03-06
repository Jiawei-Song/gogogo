package main

import (
	"fmt"
	"go_code/chatroomRebuild/server/model"
	"io"
	"net"
)

func goprocess(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.LoopProcess()
	if err != nil && err != io.EOF {
		fmt.Println("客户端和服务端协程发生了错误", err)
		return
	}
	fmt.Println("一次连接结束")
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	// 初始化一个redis的连接池
	initPool("127.0.0.1:6379", 8, 100)
	initUserDao()
	fmt.Println("服务器在8889监听。。。。")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	for {
		fmt.Println("等待客户端来链接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() err = ", err)
			// return
		}
		go goprocess(conn)
	}
}
