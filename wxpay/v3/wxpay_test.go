package v3

import (
	"testing"
)

// 注意：测试方法使用的testConfig在全局定义，是微信支付配置信息，不对外提供。
// 配置信息从微信支付账户中心获取
/*var testConfig = &Config{
	Url:           "https://api.mch.weixin.qq.com",
	Mchid:         "",
	V3Secret:      "",
	MerPrivateKey: "",
	MchidSerialNo: "",
}*/

func Test_wxpay_Certificates(t *testing.T) {
	wx := NowWxpay(testConfig)
	crs, err := wx.Certificates()
	if err != nil {
		t.Log(err)
	}
	t.Log(crs)
}
