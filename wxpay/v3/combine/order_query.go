package combine

import v3 "sdk-weixin/wxpay/v3"

/*
文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/combine/chapter3_3.shtml
合单查询订单API

*/

func Query(wx *v3.Wxpay, combine_out_trade_no string) (r *OrderQueryResult, err error) {
	err = wx.Call(v3.GET, "/v3/combine-transactions/out-trade-no/"+combine_out_trade_no, nil, &r)
	return
}

// 响应结果
type OrderQueryResult struct {
	CombineAppid      string                        `json:"combine_appid"`        //合单发起方的appid。
	CombineMchid      string                        `json:"combine_mchid"`        //合单发起方商户号。
	CombineOutTradeNo string                        `json:"combine_out_trade_no"` //合单支付总订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	SceneInfo         *SceneInfoObjectResult        `json:"scene_info"`           // 支付场景信息描述
	SubOrders         []*SubOrdersObjectResult      `json:"sub_orders"`           // 子单信息，最多支持子单条数：50
	CombinePayerInfo  *CombinePayerInfoObjectResult `json:"combine_payer_info"`   // 支付者信息
}

// 场景信息
type SceneInfoObjectResult struct {
	DeviceId string `json:"device_id"` //非必填。 终端设备号（门店号或收银设备ID）。特殊规则：长度最小7个字节
}

//子单信息
type SubOrdersObjectResult struct {
	Mchid         string              `json:"mchid"` // 子单发起方商户号，必须与发起方appid有绑定关系。
	TradeType     string              `json:"trade_type"`
	TradeState    string              `json:"trade_state"`
	BankType      string              `json:"bank_type"`
	Attach        string              `json:"attach"` //附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	SuccessTime   string              `json:"success_time"`
	TransactionId string              `json:"transaction_id"`
	OutTradeNo    string              `json:"out_trade_no"`
	SubMchid      string              `json:"sub_mchid"`
	Amount        *AmountObjectResult `json:"amount"`
}

//订单金额
type AmountObjectResult struct {
	TotalAmount   int64  `json:"total_amount"`
	Currency      string `json:"currency"` //符合ISO 4217标准的三位字母代码，人民币：CNY。 示例值：CNY
	PayerAmount   int64  `json:"payer_amount"`
	PayerCurrency string `json:"payer_currency"`
}

// 结算信息
type SettleInfoObjectResult struct {
	ProfitSharing bool  `json:"profit_sharing"`
	SubsidyAmount int64 `json:"subsidy_amount"`
}

// 支付者信息
type CombinePayerInfoObjectResult struct {
	Openid string `json:"openid"`
}
