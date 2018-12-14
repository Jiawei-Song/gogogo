package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroomRebuild/common/message"
	"go_code/chatroomRebuild/server/model"
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

	// 服务端进行redis校验
	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPWD)
	if err != nil {
		fmt.Println("MyUserDao.Login 返回的错误", err)
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 404
			loginResMes.Error = "用户不存在"
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 304
			loginResMes.Error = "密码不正确"
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "未知错误"
		}

	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登陆成功")
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
	fmt.Println("这个是服务端返回的数据， data =", string(data))

	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg(data)

	return
}
