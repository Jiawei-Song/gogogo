package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

// Message 消息结构体
type Message struct {
	Type string `json:type` // 消息类型
	Data string `json:data`
}

// LoginMes 登陆消息
type LoginMes struct {
	UserID   int    `json:user_id`
	UserPWD  string `json:user_password`
	UserName string `json:user_name`
}

// LoginResMes 登陆返回消息
type LoginResMes struct {
	Code          int `json:code` // 返回的状态码
	OnlineUserIDs []int
	Error         string `json:err`
}

// RegisterMes 注册消息
type RegisterMes struct {
	User User `json:user`
}

// RegisterResMes 注册返回的消息体
type RegisterResMes struct {
	Code  int    `json:code`
	Error string `json:err`
}

// NotifyUserStatusMes 用户状态变化时用到的结构体
type NotifyUserStatusMes struct {
	UserID int `json:userID`
	Status int `json:status`
}
