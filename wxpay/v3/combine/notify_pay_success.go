package combine

import (
	"encoding/json"
	"net/http"
	v3 "sdk-weixin/wxpay/v3"
)

func NotifyPaySuccessHandle(wx *v3.Wxpay, callback func(data *NotifyPaySuccessParams) error) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		err := payNotify(wx, r.Body, func(data []byte) error {
			var params *NotifyPaySuccessParams
			err := json.Unmarshal(data, &params)
			if err != nil {
				return err
			}
			err = callback(params)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			resp.WriteHeader(500)
			resp.Write([]byte(`{"code": "ERROR","message": "` + err.Error() + `"}`))
		} else {
			resp.WriteHeader(200)
			resp.Write([]byte(`{"code": "SUCCESS","message": "OK"}`))
		}
	})
}

type NotifyPaySuccessParams struct {
	Combine_appid        string                              `json:"combine_appid"`
	Combine_mchid        string                              `json:"combine_mchid"`
	Combine_out_trade_no string                              `json:"combine_out_trade_no"`
	Scene_info           *NotifySceneInfo                    `json:"scene_info"`
	Sub_orders           *NotifySubOrdersObjectResult        `json:"sub_orders"`
	Combine_payer_info   *NotifyCombinePayerInfoObjectResult `json:"combine_payer_info"`
}

type NotifySceneInfo struct {
	Device_id string `json:"device_id"`
}

//子单信息
type NotifySubOrdersObjectResult struct {
	Mchid         string                    `json:"mchid"` // 子单发起方商户号，必须与发起方appid有绑定关系。
	TradeType     string                    `json:"trade_type"`
	TradeState    string                    `json:"trade_state"`
	BankType      string                    `json:"bank_type"`
	Attach        string                    `json:"attach"` //附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	SuccessTime   string                    `json:"success_time"`
	TransactionId string                    `json:"transaction_id"`
	OutTradeNo    string                    `json:"out_trade_no"`
	SubMchid      string                    `json:"sub_mchid"`
	Amount        *NotifyAmountObjectResult `json:"amount"`
}

//订单金额
type NotifyAmountObjectResult struct {
	TotalAmount   int64  `json:"total_amount"`
	Currency      string `json:"currency"` //符合ISO 4217标准的三位字母代码，人民币：CNY。 示例值：CNY
	PayerAmount   int64  `json:"payer_amount"`
	PayerCurrency string `json:"payer_currency"`
}

// 支付者信息
type NotifyCombinePayerInfoObjectResult struct {
	Openid string `json:"openid"`
}
