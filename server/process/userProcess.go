package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroomRebuild/common/message"
	"go_code/chatroomRebuild/server/model"
	"go_code/chatroomRebuild/server/utils"
	"net"
)

// UserProcess 处理用户相关的,含用户的连接Conn和UserID
type UserProcess struct {
	Conn   net.Conn
	UserID int
}

// NotifyOthersOnlineUser 向其他人推送这个人上线的消息
func (userMgr *UserMgr) NotifyOthersOnlineUser(userID int) {
	for id, up := range userMgr.onlineUsers {
		if id == userID {
			continue
		}
		// 获取到在线用户的连接后，逐个通知新的上线用户
		up.NoticeOnline(userID)
	}
}

// NoticeOnline 在线用户通知新用户上线
func (userProcess *UserProcess) NoticeOnline(userID int) {
	//开始组装NotifyUserStatusMes消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserID = userID
	notifyUserStatusMes.Status = message.UserOnline
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("NoticeOnline jsonMarshal fail, err =", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("NoticeOnline jsonMarshal fail, err =", err)
		return
	}
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) fail, err =", err)
		return
	}
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
		// 登陆成功过后的逻辑
		userProcess.UserID = loginMes.UserID
		// userMgr是全局的，先维护全局的userMgr.onlineUsers，然后再添加到登陆返回消息loginResMes.OnlineUserIDs
		userMgr.AddOnlineUser(userProcess)
		for id := range userMgr.onlineUsers {
			loginResMes.OnlineUserIDs = append(loginResMes.OnlineUserIDs, id)
		}
		userMgr.NotifyOthersOnlineUser(loginMes.UserID)
		fmt.Println("loginResMes.OnlineUserIDs = ", loginResMes.OnlineUserIDs)
		fmt.Println(user, "服务端登陆成功函数走完")
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

// ServerRegisterMes 服务端处理登陆消息的函数
func (userProcess *UserProcess) ServerRegisterMes(mes *message.Message) (err error) {
	// 接收到的registerMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail , err = ", err)
		return
	}
	// 构建一个注册返回的消息体RegisterResMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	var registerResMes message.RegisterResMes
	// 服务端进行redis校验
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		fmt.Println("--------------------", err)
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 308
			// "用户已存在"
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
			fmt.Println("注册发生未知错误 = ", err)
		}
	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功")
	}

	data, err := json.Marshal(registerResMes)
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
