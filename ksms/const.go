package ksms

import (
	"net/url"
	"time"
)

const (
	accesskey             = "Accesskey"
	secretKey             = "SecretKey"
	signNameKey           = "SignName"
	tplIdKey              = "TplId"
	tplParamsKey          = "TplParams"
	mobileKey             = "Mobile"
	signatureKey          = "Signature"
	actionKey             = "Action"
	actionBatchSendSmsKey = "BatchSendSms"
)

var (
	postUrl          = "https://ksms.api.ksyun.com"
	defaultCommentUv = url.Values{
		"Service":          []string{"ksms"},
		"Action":           []string{"SendSms"},
		"Version":          []string{"2019-05-01"},
		"Timestamp":        []string{time.Now().UTC().Format("2006-01-02T15:04:05Z")},
		"SignatureVersion": []string{"1.0"},
		"SignatureMethod":  []string{"HMAC-SHA256"},
	}
)
