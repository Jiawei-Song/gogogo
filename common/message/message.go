package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
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
	Code  int    `json:code` // 返回的状态码
	Error string `json:err`
}

// RegisterMes 注册消息
type RegisterMes struct {
	UserID   int    `json:user_id`
	UserPWD  string `json:user_password`
	UserName string `json:user_name`
}
