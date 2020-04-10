// 微信支付API v3 SDK golang版本
package v3

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	GET  = "GET"
	POST = "POST"
)

type Wxpay struct {
	cfg         *Config
	wxpayConfig *WxpayConfig
}

func NowWxpay(c *Config) *Wxpay {
	wx := &Wxpay{cfg: c}
	wx.timerRefreshCertificates()
	return wx
}

func (p *Wxpay) Cfg() *Config {
	return p.cfg
}

// 微信支付v3接口规则
// 文档：https://wechatpay-api.gitbook.io/wechatpay-api-v3/
// 为了保证安全性，微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密。API v3密钥是加密时使用的对称密钥。

// AES-256-GCM 加密
/*func (p *Protocol) Aes256GcmEncrypt(nonce,associatedDate,ciphertext string) (string,error)  {

}*/

// AES-256-GCM 解密
func (p *Wxpay) Aes256GcmDecrypt(nonce, associatedData, ciphertext string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(p.cfg.V3Secret))
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCMWithNonceSize(block, len(nonce))
	if err != nil {
		return nil, err
	}
	cipherdata, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	plaindata, err := aesgcm.Open(nil, []byte(nonce), cipherdata, []byte(associatedData))
	if err != nil {
		return nil, err
	}
	return plaindata, err
}

// 计算请求报文签名
func (p *Wxpay) Sign(method, path, body, timestamp, nonce_str string) (sign string, err error) {
	targetStr := method + "\n" + path + "\n" + timestamp + "\n" + nonce_str + "\n" + body + "\n"
	log.Println("签名原始字符串：\n" + targetStr)
	sign, err = RsaSignWithSha256(targetStr, p.cfg.MerPrivateKey)
	log.Println("签名结果字符串：" + sign)
	return
}

// 验签响应报文签名
func (p *Wxpay) Very(signature, serial, time, nonce, body string) (ok bool, err error) {
	if serial == "" {
		return false, errors.New("微信平台公钥证书序列号为空")
	}
	if serial != p.wxpayConfig.wxpayPublicKeySeriaNo {
		err = p.refreshCertificates()
		if err != nil {
			return false, err
		}
	}
	if serial != p.wxpayConfig.wxpayPublicKeySeriaNo {
		return false, errors.New("验签使用的证书编号与微信平台公钥证书编号不匹配")
	}

	checkStr := time + "\n" + nonce + "\n" + body + "\n"
	blocks, _ := pem.Decode(p.wxpayConfig.wxpayPublicKey)
	if blocks == nil || blocks.Type != "PUBLIC KEY" {
		log.Println("failed to decode PUBLIC KEY")
		return false, err
	}
	oldSign, err := base64.StdEncoding.DecodeString(signature)
	pub, err := x509.ParsePKIXPublicKey(blocks.Bytes)
	hashed := sha256.Sum256([]byte(checkStr))
	err = rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hashed[:], oldSign)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 拼接权限验证字符串
func (p *Wxpay) Authorization(method, path, body string) (authStr string, err error) {
	authorization := "WECHATPAY2-SHA256-RSA2048" //固定字符串
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce_str := strings.ReplaceAll(uuid.New().String(), "-", "")
	signature, err := p.Sign(method, path, body, timestamp, nonce_str)
	if err != nil {
		return "", err
	}
	mchid := p.cfg.Mchid
	serial_no := p.cfg.MchidSerialNo
	authorization = fmt.Sprintf(`%s mchid="%s",nonce_str="%s",signature="%s",timestamp="%s",serial_no="%s"`, authorization, mchid, nonce_str, signature, timestamp, serial_no)
	return authorization, nil
}

//发送请求
func (p *Wxpay) Do(method, path, body string) (*http.Response, error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, p.cfg.Url+path, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "Golang http client")
	request.Header.Set("Content-Type", "application/json")
	authorization, err := p.Authorization(method, path, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", authorization)
	return client.Do(request)
}

//发送请求并验证签名
func (p *Wxpay) Call(method, path string, i interface{}, o interface{}) error {
	body := ""
	if i != nil {
		bytes, err := json.Marshal(i)
		if err != nil {
			return err
		}
		body = string(bytes)
	}

	resp, err := p.Do(method, path, body)
	if err != nil {
		return err
	}
	requestId := resp.Header.Get("Request-ID")
	signature := resp.Header.Get("Wechatpay-Signature")
	serial := resp.Header.Get("Wechatpay-Serial")
	timestamp := resp.Header.Get("Wechatpay-Timestamp")
	nonce := resp.Header.Get("Wechatpay-Nonce")
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resBody := string(bytes)
	log.Println("请求ID：" + requestId + "  响应报文：" + resBody)
	log.Println("resp.StatusCode=" + strconv.Itoa(resp.StatusCode))
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		return errors.New("weixin http status " + resp.Status)
	}
	ok, err := p.Very(signature, serial, timestamp, nonce, resBody)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("签名校验失败")
	}

	err = json.Unmarshal(bytes, o)
	if err != nil {
		return err
	}
	return nil
}
