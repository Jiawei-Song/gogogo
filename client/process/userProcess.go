package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroomRebuild/client/utils"
	"go_code/chatroomRebuild/common/message"
	"net"
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
	fmt.Println(string(data))

	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write 失败", err)
		return
	}
	fmt.Printf("客户端，发送的消息长度为%d\n", len(data))

	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write data 失败", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("tf.ReadPkg() fail, err = ", err)
		return
	}

	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
		ShowMenu()
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
