package combine

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	v3 "sdk-weixin/wxpay/v3"
)

/*
支付通知API
文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/combine/chapter3_7.shtml

*/

func PayNotifyHandler(wx v3.Wxpay, callback func(data []byte) error) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
		err := payNotify(wx, r.Body, callback)
		if err != nil {
			resp.WriteHeader(500)
			resp.Write([]byte(`{"code": "ERROR","message": "` + err.Error() + `"}`))
		} else {
			resp.WriteHeader(200)
			resp.Write([]byte(`{"code": "SUCCESS","message": "OK"}`))
		}
	})
}

func payNotify(wx v3.Wxpay, body io.Reader, callback func(data []byte) error) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	var notify *NotifyParams
	err = json.Unmarshal(bytes, &notify)
	if err != nil {
		return err
	}

	resource := notify.Resource
	data, err := wx.Aes256GcmDecrypt(resource.Nonce, resource.AssociatedData, resource.Ciphertext)
	if err != nil {
		return err
	}
	//解密后的json格式业务数据
	err = callback(data)
	if err != nil {
		return err
	}
	return nil
}

type NotifyParams struct {
	Id           string          `json:"id"`
	CreateTime   string          `json:"create_time"`
	EventType    string          `json:"event_type"`
	ResourceType string          `json:"resource_type"`
	Resource     *NotifyResource `json:"resource"`
}

type NotifyResource struct {
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}
