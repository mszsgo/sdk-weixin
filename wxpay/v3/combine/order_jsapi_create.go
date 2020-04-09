package combine

import v3 "sdk-weixin/wxpay/v3"

/*
文档： https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/combine/chapter3_2.shtml

# 合单下单-JS支付API
使用合单支付接口，用户只输入一次密码，即可完成多个订单的支付。目前最多一次可支持50笔订单进行合单支付。

注意：
• 订单如果需要进行抽佣等，需要在合单中指定需要进行分账（profit_sharing为true）；指定后，交易资金进入二级商户账户，处于冻结状态，可在后续使用分账接口进行分账，利用分账完结进行资金解冻，实现抽佣和对二级商户的账期。
• 合单中同一个二级商户只允许有一笔子订单。

请求URL：https://api.mch.weixin.qq.com/v3/combine-transactions/jsapi
请求方式：POST

*/

// 合单下单-JS支付API
func Jsapi(wx v3.Wxpay, params *JsapiQueryParams) (r *JsapiQueryResult, e error) {
	e = wx.Call(v3.POST, "/v3/combine-transactions/jsapi", params, &r)
	return
}

// 请求参数
type JsapiQueryParams struct {
	Combine_appid        string                  `json:"combine_appid"`        //合单发起方的appid。
	Combine_mchid        string                  `json:"combine_mchid"`        //合单发起方商户号。
	Combine_out_trade_no string                  `json:"combine_out_trade_no"` //合单支付总订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	Scene_info           *SceneInfoObject        `json:"scene_info"`           // 支付场景信息描述
	Sub_orders           []*SubOrdersObject      `json:"sub_orders"`           // 子单信息，最多支持子单条数：50
	Combine_payer_info   *CombinePayerInfoObject `json:"combine_payer_info"`   // 支付者信息
	Time_start           string                  `json:"time_start"`           //订单生成时间，遵循rfc3339标准格式
	Time_expire          string                  `json:"time_expire"`          //订单失效时间，遵循rfc3339标准格式
	Notify_url           string                  `json:"notify_url"`           //接收微信支付异步通知回调地址，通知url必须为直接可访问的URL，不能携带参数。 示例值：https://yourapp.com/notify
	Limit_pay            string                  `json:"limit_pay"`            // 指定支付方式
}

// 场景信息
type SceneInfoObject struct {
	Device_id       string `json:"device_id"`       //非必填。 终端设备号（门店号或收银设备ID）。特殊规则：长度最小7个字节
	Payer_client_ip string `json:"payer_client_ip"` //必填。用户端实际ip
}

//子单信息
type SubOrdersObject struct {
	Mchid          string            `json:"mchid"`  // 子单发起方商户号，必须与发起方appid有绑定关系。
	Attach         string            `json:"attach"` //附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	Amount         *AmountObject     `json:"amount"`
	Out_trade_no   string            `json:"out_trade_no"`
	Sub_mchid      string            `json:"sub_mchid"`
	Detail         string            `json:"detail"`
	Profit_sharing bool              `json:"profit_sharing"`
	Description    string            `json:"description"`
	Settle_info    *SettleInfoObject `json:"settle_info"`
}

//订单金额
type AmountObject struct {
	Total_amount int64  `json:"total_amount"`
	Currency     string `json:"currency"` //符合ISO 4217标准的三位字母代码，人民币：CNY。 示例值：CNY
}

// 结算信息
type SettleInfoObject struct {
	Profit_sharing bool  `json:"profit_sharing"`
	Subsidy_amount int64 `json:"subsidy_amount"`
}

// 支付者信息
type CombinePayerInfoObject struct {
	Openid string `json:"openid"`
}

// 响应参数
type JsapiQueryResult struct {
	Prepay_id string `json:"prepay_id"`
}
