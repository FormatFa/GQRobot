package main

import (
	"fmt"
	"testing"
)

// 测试大模型prompt
func TestChatApi(t *testing.T) {

	res := reqChatApi("有，但是不多 --- 用英文合理的回答和解析语法： ")
	fmt.Println(res)
}
