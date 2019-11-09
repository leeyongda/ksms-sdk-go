package ksms

import "net/url"

type params struct {
	uv url.Values
}

func (p *params) CopyParams(uv url.Values) {
	p.uv = uv
}

func (p *params) SetAccessKey(val string) {
	p.uv.Set(accesskey, val)
}

func (p *params) SetSecretKey(val string) {
	p.uv.Set(secretKey, val)
}

func (p *params) SetSignName(val string) {
	p.uv.Set(signNameKey, val)
}

func (p *params) SetTplId(val string) {
	p.uv.Set(tplIdKey, val)
}

func (p *params) SetTplParams(val string) {
	p.uv.Set(tplParamsKey, val)
}

func (p *params) SetMobile(val string) {
	p.uv.Set(mobileKey, val)
}

func (p *params) SetSignature(val string) {
	p.uv.Set(signatureKey, val)
}

func (p *params) SetAction(val string) {
	p.uv.Set(actionKey, val)
}

func (p *params) Encode() string {
	return p.uv.Encode()
}
