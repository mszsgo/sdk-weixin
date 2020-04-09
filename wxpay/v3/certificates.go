package v3

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"time"
)

/*
文档：https://wechatpay-api.gitbook.io/wechatpay-api-v3/jie-kou-wen-dang/ping-tai-zheng-shu

// 获取平台证书列表
// GET /v3/certificates

获取商户当前可用的平台证书列表。微信支付提供该接口，帮助商户后台系统实现平台证书的平滑更换。
该请求无需身份认证信息之外的其他参数，请点击Response查看应答示例。

注意事项
如果自行实现验证平台签名逻辑的话，需要注意以下事项:
程序实现定期更新平台证书的逻辑，不要硬编码验证应答消息签名的平台证书
定期调用该接口，间隔时间小于12 小时
加密请求消息中的敏感信息时，使用最新的平台证书（即：证书启用时间较晚的证书）
*/

// 定时更新证书，间隔时间小于12 小时
func (p *Wxpay) timerRefreshCertificates() {
	err := p.refreshCertificates()
	if err != nil {
		log.Print(err)
	}
	go func() {
		time.Sleep(time.Hour * 11)
		p.timerRefreshCertificates()
	}()
}

// 最新证书信息
func (p *Wxpay) refreshCertificates() (err error) {
	rs, err := p.Certificates()
	if err != nil {
		return err
	}

	// 获取最新的证书
	var encrypt_certificate *EncryptCertificateObject = nil
	serial_no := ""
	effective_time := time.Now()
	for _, item := range rs.Data {
		t, err := time.Parse(time.RFC3339, item.Effective_time)
		if err != nil {
			return err
		}
		if t.Before(effective_time) {
			effective_time = t
			serial_no = item.Serial_no
			encrypt_certificate = item.Encrypt_certificate
		}
	}
	if encrypt_certificate == nil {
		return errors.New("没有获取到微信平台密文公钥证书")
	}

	//解密最新的证书
	decryptCiphertext, err := p.Aes256GcmDecrypt(encrypt_certificate.Nonce, encrypt_certificate.Associated_data, encrypt_certificate.Ciphertext)
	if err != nil {
		return err
	}

	p.wxpayConfig = &WxpayConfig{
		wxpayPublicKeySeriaNo: serial_no,
		wxpayPublicKey:        decryptCiphertext,
	}
	return nil
}

func (p *Wxpay) Certificates() (cert *CertificatesResult, err error) {
	resp, err := p.Do("GET", "/v3/certificates", "")
	if err != nil {
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Println("获取证书列表:" + string(bytes))
	json.Unmarshal(bytes, &cert)
	if err != nil {
		return
	}
	return
}

type CertificatesResult struct {
	Data []*CertificateObject `json:"data"`
}

type CertificateObject struct {
	Serial_no           string                    `json:"serial_no"`
	Effective_time      string                    `json:"effective_time"`
	Expire_time         string                    `json:"expire_time"`
	Encrypt_certificate *EncryptCertificateObject `json:"encrypt_certificate"`
}

type EncryptCertificateObject struct {
	Algorithm       string `json:"algorithm"`
	Nonce           string `json:"nonce"`
	Associated_data string `json:"associated_data"`
	Ciphertext      string `json:"ciphertext"`
}
