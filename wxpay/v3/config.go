package v3

// 微信支付V3接口调用参数配置项， 参数通过微信支付商户平台获取
type Config struct {
	WxUrl         string // 接口服务地址，默认：https://api.mch.weixin.qq.com
	WxMchid       string // 商户号
	WxV3Secret    string // 微信V3 secret
	MerPrivateKey string // 微信支付商户私钥字符串，下载证书中的字符串内容，不要做任何修改。
	MchidSerialNo string // 微信商户证书序列号
}

// 需要通过微信证书下载接口获取，小于12小时获取一次
type WxpayConfig struct {
	WxpayPublicKey        string // 微信支付平台公钥
	WxpayPublicKeySeriaNo string // 微信支付平台公钥证书序列号
}

// 获取微信平台公钥与公钥证书序列号
func (c *Config) WxpayPublicKey() *WxpayConfig {

	return nil
}
