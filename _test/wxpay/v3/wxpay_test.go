package v3

import (
	v3 "sdk-weixin/wxpay/v3"
	"sdk-weixin/wxpay/v3/combine"
	"testing"
	"time"
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

var wx = v3.NowWxpay(testConfig)

func Test_wxpay_Certificates(t *testing.T) {
	crs, err := wx.Certificates()
	if err != nil {
		t.Log(err)
	}
	t.Log(crs)
}

// jsapi创建订单测试
func Test_wxpay_combine_JsapiCreate(t *testing.T) {
	r, e := combine.JsapiCreate(wx, &combine.JsapiQueryParams{
		CombineAppid:      "wx6342f5c1ad780e13",
		CombineMchid:      wx.Cfg().Mchid,
		CombineOutTradeNo: "1234567891",
		SceneInfo: &combine.SceneInfoObject{
			DeviceId:      "h5.jsapi.pay1",
			PayerClientIp: "121.40.1.2",
		},
		SubOrders: []*combine.SubOrdersObject{{
			Mchid:  wx.Cfg().Mchid,
			Attach: "",
			Amount: &combine.AmountObject{
				TotalAmount: 100,
				Currency:    "CNY",
			},
			OutTradeNo:    "1234567891",
			Detail:        "测试商品信息",
			ProfitSharing: false,
			Description:   "微信支付测试订单",
			SettleInfo: &combine.SettleInfoObject{
				ProfitSharing: false,
				SubsidyAmount: 0,
			},
		}},
		CombinePayerInfo: &combine.CombinePayerInfoObject{
			Openid: "123",
		},
		TimeStart:  time.Now().Format(time.RFC3339),
		TimeExpire: time.Now().Add(time.Minute * 30).Format(time.RFC3339),
		NotifyUrl:  "https://msd.himkt.cn/pay/notify.do",
		LimitPay:   []string{"no_debit"},
	})
	if e != nil {
		t.Error(e)
		return
	}
	t.Log("PrepayId=" + r.PrepayId)
}
