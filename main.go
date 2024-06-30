package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/go-vgo/robotgo"
	"online.indigo6a.gorobot/qrobot"
)

var historyChats []string

func main() {
	qrobot.ActiveQ()
	// 不断循环获取最新消息回复
	for k := 0; k < 240; k += 1 {
		work()
		robotgo.Sleep(10 * 1)
	}
}
func work() {

	chats := qrobot.GetChatText("UnityStudio")
	for _, value := range chats {
		if slices.Contains(historyChats, value) {
			continue
		}
		historyChats = append(historyChats, value)

		// 自定义过滤
		// 包含指定开头时才回复
		// 特定开头时,
		// if !strings.Contains(value, "@indigo6a") {
		// 	continue
		// }
		// 包含中文时
		// hasHan := false
		// // 跳过中文
		// for _, r := range value {
		// 	if unicode.Is(unicode.Scripts["Han"], r) {
		// 		hasHan = true
		// 		continue
		// 	}
		// }

		// if !hasHan {
		// 	fmt.Println("skip plain english:", value)
		// 	continue
		// }
		if strings.HasPrefix(value, "[表情]") {
			continue
		}

		// 忽略机器人说的话
		if strings.HasPrefix(value, "emm") {
			continue
		}
		fmt.Println("process chat:", value)
		value = strings.ReplaceAll(value, "@indigo6a", "")
		// res := reqChatApi("请用英文回答并解析英文的语法：" + value)
		res := reqChatApi(value + " --- 用英文合理的回答和解析语法 ")

		fmt.Println("ai response:", res)
		qrobot.SendMsg("emm," + res)
	}

}

// 请求大模型接口，得到响应
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content,string"`
}
type Payload struct {
	Messages []Message `json:"messages"`
}

// 请求大模型接口，可自定义
func reqChatApi(question string) string {
	targetUrl := "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant?access_token=24.0eaaefe1fa779a5a1b429d4ac3068d27.2592000.1722123944.282335"

	// fmt.Println(question)
	var messages []Message
	msg := Message{
		Role:    "user",
		Content: question,
	}
	// messages = append(messages, msg)
	// for _, item := range chats {
	// 	msg := Message{
	// 		Role:    "user",
	// 		Content: item,
	// 	}
	// 	messages = append(messages, msg)
	// }
	messages = append(messages, msg)
	data := Payload{
		Messages: messages,
	}
	b, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	question = string(b)
	fmt.Println(question)
	payload := strings.NewReader(question)

	req, _ := http.NewRequest("POST", targetUrl, payload)

	req.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err == nil {
		res, err := io.ReadAll(response.Body)
		if err == nil {
			resJson := string(res)
			fmt.Println(resJson)
			var data map[string]interface{}
			_ = json.Unmarshal(res, &data)
			result := data["result"]
			fmt.Println(result)
			return result.(string)

		}
	} else {
		fmt.Println("http fail", err)

	}
	return ""

}
