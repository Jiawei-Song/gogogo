package process

import (
	"fmt"
	"go_code/chatroomRebuild/common/message"
)

// outputGroupMes 输出收到的mes
func outputGroupMes(smsMes *message.SmsMes) {
	fmt.Printf("用户id为%d\t 对大家说：\n %s\n", smsMes.UserID, smsMes.Content)
}
