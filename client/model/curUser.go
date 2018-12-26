package model

import (
	"go_code/chatroomRebuild/common/message"
	"net"
)

// CurUser aa
type CurUser struct {
	Conn net.Conn
	message.User
}
