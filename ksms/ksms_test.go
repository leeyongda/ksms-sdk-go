package ksms

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	cl := NewSmsClient(Client{
		AccessKey: "",
		SecretKey: "",
		SignName:  "签名",
		TplId:     "1",
	})
	// 发送单条短信, 必须先设置模版内容，才能发送短信
	if res, err := cl.SetTplParams("number", "123458").
		SendSms("12345678901"); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(res))
	}
	// 发送多条短信, 必须先设置模版内容，才能发送短信
	if res, err := cl.SetTplParams("number", "123458").
		SendBatchSms([]string{"12345678901"}); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(res))
	}
}
