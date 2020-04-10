package combine

import v3 "sdk-weixin/wxpay/v3"

/*
文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/combine/chapter3_3.shtml
合单查询订单API

*/

func Close(wx v3.Wxpay, combine_out_trade_no string, params *OrderCloseQueryParams) (err error) {
	err = wx.Call(v3.GET, "/v3/combine-transactions/out-trade-no/"+combine_out_trade_no+"/close", &params, nil)
	return
}

type OrderCloseQueryParams struct {
	CombineAppid string                  `json:"combine_appid"`
	SubOrders    []*SubOrdersCloseObject `json:"sub_orders"`
}

type SubOrdersCloseObject struct {
	Mchid        string `json:"mchid"`
	Out_trade_no string `json:"out_trade_no"`
	Sub_mchid    string `json:"sub_mchid"`
}
