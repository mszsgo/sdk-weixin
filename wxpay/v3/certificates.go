package v3

import (
	"io/ioutil"
	"log"
)

// 获取平台证书列表
// GET /v3/certificates

/*
获取商户当前可用的平台证书列表。微信支付提供该接口，帮助商户后台系统实现平台证书的平滑更换。
该请求无需身份认证信息之外的其他参数，请点击Response查看应答示例。

注意事项
如果自行实现验证平台签名逻辑的话，需要注意以下事项:
程序实现定期更新平台证书的逻辑，不要硬编码验证应答消息签名的平台证书
定期调用该接口，间隔时间小于12 小时
加密请求消息中的敏感信息时，使用最新的平台证书（即：证书启用时间较晚的证书）
*/

func (p *Protocol) V3Certificates() error {
	resp, err := p.Do("GET", "/v3/certificates", "")
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println("获取证书列表:" + string(bytes))
	return nil
}
