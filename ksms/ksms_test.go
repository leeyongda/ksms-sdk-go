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
		TplId:     "0",
		TplParams: `{"number":"123457"}`,
	})
	if re, err := cl.SendSms("1234567890"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(re))
	}
}
