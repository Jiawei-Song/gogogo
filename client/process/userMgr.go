package process

import (
	"go_code/chatroomRebuild/common/message"
)

var onlineUsers = make(map[int]*message.User, 500)
