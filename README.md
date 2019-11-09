## 金山云短信服务 Go-SDK

show me code(golang)
```Golang
cl := NewSmsClient(Client{
		AccessKey: "",
		SecretKey: "",
		SignName:  "签名",
		TplId:     "0",
		TplParams: `{"number":"123456"}`,
	})
	if res, err := cl.SendSms("12345678901"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(res))
	}
```


