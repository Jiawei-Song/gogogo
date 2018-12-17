package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroomRebuild/client/utils"
	"go_code/chatroomRebuild/common/message"
	"io"
	"net"
	"os"
)

// UserProcess 处理用户的
type UserProcess struct {
}

// Login client登陆的函数
func (userProcess *UserProcess) Login(userID int, userPWD string) (err error) {
	// 连接
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial, err = ", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMesType

	// 构建一个loginmes结构体
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPWD = userPWD
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginmes 序列化失败， err =", err)
		return
	}
	mes.Data = string(data)
	// 将mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes 序列化失败， err =", err)
		return
	}
	fmt.Println("登陆序列化的data =", string(data))

	// 消息构建完成之后，创建一个tf实例，用于通讯
	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.WritePkg(data)
	if err != nil {
		fmt.Println("conn.Write data 失败", err)
		return
	}
	mes, err = tf.ReadPkg()
	if err != nil && err != io.EOF {
		fmt.Println("tf.ReadPkg() fail, err = ", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
		// 处理在线用户相关
		for _, v := range loginResMes.OnlineUserIDs {
			if v == userID {
				continue
			}
			fmt.Println("在线用户id有：", v)
			// onlineUsers[v]
		}
		// 启动一个协程，监听服务端发送给客户端的数据，如果有，就显示出来
		go serverProcessMes(conn)
		ShowMenu()

	} else if loginResMes.Code != 200 {
		fmt.Println(loginResMes.Code, loginResMes.Error)
	}
	return
}

// Register client用户注册的函数
func (userProcess *UserProcess) Register(userID int, userPWD string, userName string) (err error) {
	// 连接
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial, err = ", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 构建一个register结构体
	var register message.RegisterMes
	register.User.UserID = userID
	register.User.UserPWD = userPWD
	register.User.UserName = userName
	data, err := json.Marshal(register)
	if err != nil {
		fmt.Println("register 序列化失败， err =", err)
		return
	}
	mes.Data = string(data)

	// 将mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes 序列化失败， err =", err)
		return
	}
	fmt.Println("注册序列化的data =", string(data))

	// 消息构建完成之后，创建一个tf实例，用于通讯
	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.WritePkg(data)
	if err != nil {
		fmt.Println("conn.Write data 失败", err)
		return
	}
	mes, err = tf.ReadPkg()
	if err != nil && err != io.EOF {
		fmt.Println("tf.ReadPkg() fail, err = ", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功")
		os.Exit(0)
		// // 启动一个协程，监听服务端发送给客户端的数据，如果有，就显示出来
		// go serverProcessMes(conn)
		// ShowMenu()
	} else if registerResMes.Code != 200 {
		fmt.Println(registerResMes.Code, registerResMes.Error)
		os.Exit(0)
	}
	return
}
