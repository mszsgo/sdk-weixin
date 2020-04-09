package combine

import v3 "sdk-weixin/wxpay/v3"

/*
文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/combine/chapter3_3.shtml
合单查询订单API

*/

func Close(wx v3.Wxpay, tradeNo string) (err error) {
	err = wx.Call(v3.GET, "/v3/combine-transactions/out-trade-no/"+tradeNo, nil, nil)
	return
}
