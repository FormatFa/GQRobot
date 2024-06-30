package qrobot

import (
	"testing"
)

func TestGetChat(t *testing.T) {
	// activeQ()
	// robotgo.Sleep(3)
	// texts := getChatText("英语学习交流")
	// fmt.Println(texts)
	// // pids, _ := robotgo.FindIds("qq")
	// fmt.Println(pids)
	// for _, value := range pids {
	// 	fmt.Println("open :", value)
	// 	robotgo.ActivePid(value)
	// 	robotgo.Sleep(2)
	// }

}

// 激活QQ 测试
func TestActiveQQ(t *testing.T) {
	ActiveQ()
}
