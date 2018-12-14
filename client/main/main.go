package main

import (
	"fmt"
	"go_code/chatroomRebuild/client/process"
)

var userID int
var userPWD string
var userName string

func main() {
	var key int
	// var loop = true
	for {
		fmt.Println("--------欢迎使用多人聊天系统-----------------")
		fmt.Println("            1.登陆聊天室")
		fmt.Println("            2.注册用户")
		fmt.Println("            3.退   出")
		fmt.Println("请选择（1-3）")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户ID")
			fmt.Scanln(&userID)
			fmt.Println("请输入用户密码")
			fmt.Scanln(&userPWD)
			up := &process.UserProcess{}
			up.Login(userID, userPWD)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户ID")
			fmt.Scanln(&userID)
			fmt.Println("请输入用户密码")
			fmt.Scanln(&userPWD)
			fmt.Println("请输入用户名字")
			fmt.Scanln(&userName)
			up := &process.UserProcess{}
			up.Register(userID, userPWD, userName)
		case 3:
			fmt.Println("退   出")
			// loop = false
		default:
			fmt.Println("输入的信息有误，请重新输入")
		}
	}
}
