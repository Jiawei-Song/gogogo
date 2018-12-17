package message

// User 用户的结构体
type User struct {
	UserID     int    `json:userID`
	UserPWD    string `json:userPWD`
	UserName   string `json:userName`
	UserStatus int    `json:userStatus`
}
