package process

import (
	"fmt"
	"go_code/chatroomRebuild/client/model"
	"go_code/chatroomRebuild/common/message"
)

var onlineUsers = make(map[int]*message.User, 500)

// CurUser aa
var CurUser model.CurUser

// updateUserStatus 用户状态改变，改变客户端的onlineUsers
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserID]
	if !ok {
		user = &message.User{
			UserID: notifyUserStatusMes.UserID,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserID] = user
	//更新完之后，显示一下
	showOnlineUsers()
}

//showOnlineUsers 在客户端显示当前的在线用户
func showOnlineUsers() {
	fmt.Println("当前的用户列表为：")
	for id, user := range onlineUsers {
		fmt.Printf("用户id为%d，信息为%v\n", id, user)
	}
}
