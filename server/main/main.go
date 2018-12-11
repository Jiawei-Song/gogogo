package main

import (
	"fmt"
	"io"
	"net"
)

// func readPkg(conn net.Conn) (mes message.Message, err error) {
// 	buf := make([]byte, 8096)
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		if err == io.EOF {
// 			return
// 		}
// 		fmt.Println("conn.Read, err", err)
// 		return
// 	}
// 	fmt.Println("读到的buf", buf[:4])

// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[0:4])
// 	n, err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Read fail, err =", err)
// 	}

// 	err = json.Unmarshal(buf[:pkgLen], &mes)
// 	if err != nil {
// 		fmt.Println("序列化失败, err = ", err)
// 	}
// 	return
// }

// func writePkg(conn net.Conn, data []byte) (err error) {
// 	var pkglen uint32
// 	pkglen = uint32(len(data))
// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkglen)
// 	n, err := conn.Write(buf[:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write 失败", err)
// 		return
// 	}
// 	n, err = conn.Write(data)
// 	if n != int(pkglen) || err != nil {
// 		fmt.Println("conn.Write 失败", err)
// 		return
// 	}

// 	fmt.Printf("客户端，发送的消息长度为%d\n", pkglen)
// 	return
// }

// // serverLoginMes 服务端处理登陆消息的函数
// func serverLoginMes(conn net.Conn, mes *message.Message) (err error) {
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal fail , err = ", err)
// 		return
// 	}
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType

// 	var loginResMes message.LoginResMes

// 	// 如果用户 id 为 100 ，密码为 123456， 就合法
// 	if loginMes.UserID == 100 && loginMes.UserPWD == "123456" {
// 		loginResMes.Code = 200
// 	} else {
// 		loginResMes.Code = 500
// 		loginResMes.Error = "用户不存在"
// 	}

// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail , err = ", err)
// 	}
// 	resMes.Data = string(data)

// 	// resMes 构建完成，序列化之后发送
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal fail , err = ", err)
// 	}
// 	err = writePkg(conn, data)

// 	return
// }

// // 根据消息的种类，判断调用那个函数来进行处理
// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
// 	switch mes.Type {
// 	case message.LoginMesType:
// 		// 处理登陆的逻辑
// 		err = serverLoginMes(conn, mes)
// 	case message.RegisterMesType:
// 		// 处理注册的逻辑
// 	default:
// 		fmt.Println("消息类型不存在，无法处理")

// 	}
// 	return
// }

func process(conn net.Conn) {
	defer conn.Close()
	for {
		// 将读取数据包进行一个封装，readPkg(),返回一个Message, Err
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出")
				return
			}
			fmt.Println("readPkg 出错, err = ", err)
			return
		}
		fmt.Println(mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			fmt.Println()
			return
		}
	}
}

func main() {
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
		fmt.Println(conn)
		go process(conn)
	}
}
