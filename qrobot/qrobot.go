package qrobot

import (
	"fmt"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
)

type Robot interface {
	getLatestChat(groupName string) []string
}

var qqWindowsName = "QQ"

func isMainQ() bool {
	title := robotgo.GetTitle()
	fmt.Println("title@@@ ", title)
	return title == "QQ"
}
func ActiveQ() {
	ids, _ := robotgo.FindIds("QQ")
	fmt.Println("pids", ids)
	id := -1
	for _, value := range ids {
		name, _ := robotgo.FindName(value)
		fmt.Println("name;" + name)
		if name == qqWindowsName {
			id = value
			break
		}
	}
	if id == -1 {
		fmt.Println("qq not found ")
	}

	// 显示出来
	robotgo.ActivePid(id)
	robotgo.MilliSleep(200)
	title := robotgo.GetTitle()
	if title != "QQ" {
		fmt.Println("打开QQ失败")
	} else {
		fmt.Println("打开QQ成功")
	}

}

// 发送消息
func SendMsg(msg string) {

	clipboard.WriteAll(msg)
	robotgo.MilliSleep(1000)
	robotgo.MoveSmooth(30, 825)
	robotgo.MilliSleep(300)
	robotgo.KeyToggle("command")
	robotgo.MilliSleep(100)
	robotgo.KeyTap("v")
	robotgo.KeyToggle("command", "up")
	robotgo.MilliSleep(500)
	robotgo.KeyTap("enter")
	clipboard.WriteAll("")
	robotgo.MilliSleep(500)
}
func GetChatText(groupName string) []string {

	fmt.Println("start get chat text:")

	var chats []string

	title := robotgo.GetTitle()
	fmt.Println("title@@@ ", title)
	rect := robotgo.GetDisplayRect(0)

	robotgo.MouseSleep = 20

	title = robotgo.GetTitle()
	fmt.Println("title@@@ ", title)
	if title != groupName {
		fmt.Println("windows name change.")
		return chats
	}
	robotgo.MilliSleep(1000)
	// 移动鼠标到历史按钮
	robotgo.Move(1480, 800)
	robotgo.MilliSleep(1000)
	// 1480,800
	// 点击
	robotgo.Click("left")

	robotgo.MilliSleep(1000)
	//
	startX := 70
	startY := 235
	step := 10
	fmt.Println("window h:", rect.H)
	lastText := ""

	// 遍历屏幕高度，这里/2表示这查到一半就行
	for i := startY; i < rect.H/2; i += step {

		clipboard.WriteAll("")
		title = robotgo.GetTitle()
		if title != groupName {
			fmt.Println("name change .stop")
			return chats
		}
		// fmt.Println("move to:", startX, i)
		robotgo.MoveSmooth(startX, i)
		robotgo.MilliSleep(100)
		color := robotgo.GetLocationColor()

		// fmt.Println(color)
		// 需要修改的地方，用来判断当前的像素是不是背景来的（不然会误点视频，小程序那些，qq黑暗模式和白天模式主题会不一样)
		if color != "111111" {
			// fmt.Println("skip white link")
			continue
		}
		// if color != "f2f2f2" {
		// 	// fmt.Println("skip white link")
		// 	continue
		// }
		// fmt.Println(robotgo.GetLocationColor(0))
		robotgo.Click()
		robotgo.MilliSleep(100)
		robotgo.KeyTap("command")
		robotgo.KeyTap("a")
		robotgo.KeyTap("a", "up")

		robotgo.KeyTap("c")
		robotgo.KeyTap("c", "up")

		robotgo.MilliSleep(300)
		copyText, _ := clipboard.ReadAll()

		robotgo.MilliSleep(1000)

		// fmt.Println("last text=" + lastText + " copy:" + copyText)
		if copyText == lastText {

			continue
		}
		if strings.HasPrefix(copyText, "emm") {
			continue
		}
		// if len(copyText) > 30 {
		// 	continue
		// }
		// 出现过的对话跳过
		// if InSlice(historyChats, copyText) {
		// 	continue
		// }
		// hasHan := false
		// // 跳过中文
		// for _, r := range copyText {
		// 	if unicode.Is(unicode.Scripts["Han"], r) {
		// 		hasHan = true
		// 		continue
		// 	}
		// }
		// if hasHan {
		// 	continue
		// }
		fmt.Println(copyText)

		lastText = copyText
		chats = append(chats, copyText)
		// historyChats = append(historyChats, copyText)
	}
	robotgo.CloseWindow()

	// robotgo.KeyToggle("command", "up")
	// robotgo.MilliSleep(100)
	// fmt.Println("command.")
	// robotgo.KeyToggle("command")
	// robotgo.KeyTap("w")
	// robotgo.MilliSleep(100)
	// robotgo.KeyTap("w", "up")
	// robotgo.KeyToggle("command", "up")
	robotgo.CloseWindow()

	robotgo.Sleep(1)
	title = robotgo.GetTitle()
	fmt.Println("close history，title=", title)

	// if title != "QQ" {
	// 	fmt.Println("close his")
	// } else {
	// 	fmt.Println("打开QQ成功")
	// }

	fmt.Println(chats)
	robotgo.MilliSleep(2000)
	return chats
}
