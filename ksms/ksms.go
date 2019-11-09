package ksms

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	AccessKey  string
	SecretKey  string
	SignName   string
	TplId      string
	tplParams  map[string]interface{}
	httpClient http.Client
	dopts      clientOptions
	params     params
	signature  string
}

type clientOptions struct {
	service          string
	action           string
	version          string
	signatureVersion string
	signatureMethod  string
	timestamp        string
}

type ClientOption interface {
	apply(*clientOptions)
}
type funcClientOption struct {
	f func(*clientOptions)
}

func newFuncClientOption(f func(*clientOptions)) *funcClientOption {
	return &funcClientOption{
		f: f,
	}
}

func (fdo *funcClientOption) apply(do *clientOptions) {
	fdo.f(do)
}

func WithService(service string) ClientOption {
	return newFuncClientOption(func(o *clientOptions) {
		o.service = service
	})
}

func WithAction(action string) ClientOption {
	return newFuncClientOption(func(o *clientOptions) {
		o.action = action
	})
}

func WithVersion(version string) ClientOption {
	return newFuncClientOption(func(o *clientOptions) {
		o.version = version
	})
}

func WithSignatureVersion(signatureVersion string) ClientOption {
	return newFuncClientOption(func(o *clientOptions) {
		o.signatureVersion = signatureVersion
	})
}

func WithSignatureMethod(signatureMethod string) ClientOption {
	return newFuncClientOption(func(o *clientOptions) {
		o.signatureMethod = signatureMethod
	})
}

func NewSmsClient(cl Client, opts ...ClientOption) *Client {
	co := clientOptions{
		service:          "ksms",
		action:           "SendSms",
		version:          "2019-05-01",
		timestamp:        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		signatureMethod:  "HMAC-SHA256",
		signatureVersion: "1.0",
	}
	for _, opt := range opts {
		opt.apply(&co)
	}
	cl.tplParams = make(map[string]interface{})
	cl.params = params{uv: url.Values{}}
	cl.processParams()
	cl.httpClient = http.Client{Timeout: 10 * time.Second}
	cl.dopts = co
	return &cl
}

func (c *Client) genSha256(rawStr, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(rawStr))
	hashCode := fmt.Sprintf("%x\n", mac.Sum(nil))
	return hashCode
}

func (c *Client) processParams() {
	// copy params from defaultCommentUv
	c.params.CopyParams(defaultCommentUv)
	// 覆盖默认参数
	if c.dopts.action != "" {
		c.params.uv.Set("Action", c.dopts.action)
	}
	if c.dopts.service != "" {
		c.params.uv.Set("Service", c.dopts.service)
	}
	if c.dopts.signatureMethod != "" {
		c.params.uv.Set("SignatureMethod", c.dopts.signatureMethod)
	}
	if c.dopts.signatureVersion != "" {
		c.params.uv.Set("SignatureVersion", c.dopts.signatureVersion)
	}
	if c.dopts.version != "" {
		c.params.uv.Set("Version", c.dopts.version)
	}
	c.params.SetAccessKey(c.AccessKey)
	c.params.SetSignName(c.SignName)
	c.params.SetTplId(c.TplId)
}

func (c *Client) beforeSend() {
	sign := c.genSha256(c.params.Encode(), c.SecretKey)
	c.signature = sign
}

func (c *Client) SetTplParams(key string, val interface{}) *Client {
	c.tplParams[key] = val
	return c
}

func (c *Client) Marshal() (err error) {
	if b, err := json.Marshal(c.tplParams); err != nil {
	} else {
		c.params.SetTplParams(string(b))
	}
	return err
}

func (c *Client) SendBatchSms(mobile []string) (result []byte, err error) {
	m := strings.Join(mobile, ",")
	c.params.SetAction(actionBatchSendSmsKey)
	result, err = c.SendSms(m)
	return
}

func (c *Client) SendSms(mobile string) (result []byte, err error) {
	if err = c.Marshal(); err != nil {
		return
	}
	c.params.SetMobile(mobile)
	c.beforeSend()
	p := fmt.Sprintf("%s&Signature=%s", c.params.Encode(), c.signature)
	p = strings.TrimSpace(p)
	req, err := http.NewRequest("POST", postUrl, strings.NewReader(p))
	if err != nil {
		return
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if result, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return
}
