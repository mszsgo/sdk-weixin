package combine

import (
	"github.com/google/uuid"
	v3 "sdk-weixin/wxpay/v3"
	"strings"
	"time"
)

// JS调起支付API
// 文档： https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/combine/chapter3_6.shtml

func JspaySign(wx *v3.Wxpay, appId string, packageStr string) (*JspayParams, error) {
	var r = &JspayParams{
		AppId:     appId,
		TimeStamp: string(time.Now().Unix()),
		NonceStr:  strings.ReplaceAll(uuid.New().String(), "-", ""),
		Package:   packageStr,
		SignType:  "RSA",
		PaySign:   "",
	}
	sign, err := v3.RsaSignWithSha256(r.AppId+"\n"+r.TimeStamp+"\n"+r.NonceStr+"\n"+r.Package+"\n", wx.Cfg().MerPrivateKey)
	if err != nil {
		return nil, err
	}
	r.PaySign = sign
	return r, nil
}

type JspayParams struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}
