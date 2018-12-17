package process

import "fmt"

// UserMgr 用户在线相关
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

var userMgr *UserMgr

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
	fmt.Println(userMgr)
}

//AddOnlineUser 添加在线用户
func (userMgr *UserMgr) AddOnlineUser(up *UserProcess) {
	fmt.Println("进入了AddOnlineUser函数, up=", up)
	userMgr.onlineUsers[up.UserID] = up
}

//DelOnlineUser 添加在线用户
func (userMgr *UserMgr) DelOnlineUser(up *UserProcess) {
	delete(userMgr.onlineUsers, up.UserID)
}

// GetAllOnlineUsers 获取在线用户的map
func (userMgr *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return userMgr.onlineUsers
}

// GetUserByID 根据ID获取用户连接的Conn
func (userMgr *UserMgr) GetUserByID(userID int) (up *UserProcess, err error) {
	up, ok := userMgr.onlineUsers[userID]
	if !ok {
		err = fmt.Errorf("userID为%d的用户不存在", userID)
		return
	}
	return
}
