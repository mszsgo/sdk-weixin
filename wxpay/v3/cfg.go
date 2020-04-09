package v3

// 微信支付V3接口调用参数配置项

type Cfg struct {

	// ****** 以下参数通过微信后台获取
	// 接口服务地址
	ServiceUrl string
	// 商户号
	Mchid string
	// 微信V3 secret
	WxV3Secret string
	// 微信支付商户私钥
	MerPrivateKey string
	// 微信商户证书序列号
	MchidSerialNo string

	// ****** 私有属性，需要通过微信证书下载接口获取，小于12小时获取一次
	// 微信支付平台公钥
	wxpayPublicKey string
	// 微信支付平台公钥证书序列号
	wxpayPublicKeySeriaNo string
}

func NewCfg() *Cfg {
	//通过配置中心或者配置文件读取配置信息
	c := &Cfg{
		ServiceUrl: "https://api.mch.weixin.qq.com",
		Mchid:      "1526826661",
		MerPrivateKey: `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQC++j0Bs8K3c7St
tO82KbTrrqwRklrQqKwLoXfHz3i+kGojhYMX5Yg4xzglLh/gg3mtb4Mdma5GT7N8
wMeQnPvFbHNGYcVvkNTNN0a98NI7S22fn1pXU4CieMvdWble/q0N9ATd5OflVyYv
7aUTmUgt/kGnknt+GPbVl56KRQOxJ7Hg6VV77MYykLVOGdehp3MvKEctrj6y9DuU
n2QtTVhG6qNMzGR+sQ0kx987VYjx3lcOJjybteHDbZk5MG0pn05Ckzkbi1PCBSqr
pWaYScDQP/srrR8abPEkDDi33H2dReqsTbQX64xkuJVvs7w/mdusOGciGkViXa3/
IjTJj1yNAgMBAAECggEAe4vtPZWqKP3sa9e6TOLMTQ/R9wgkQgpsSpXppFSeu0E3
uMUdmq794iVXOA5PxvGnHMEgJF0QNiJHbjRUpkQ4SG6xYF3N2S/rytqwpB4QExXn
1DTFv+qgM/tLf2YVGtRM5dLv9xCnyGEJrkXh3fzyifaLSLnltBy4SlNs1+xvyAFS
uGvBVX8oxM07JBsSWPEyuVcB9rA+Yll7HJxpqCuOcc+we3vOlGGo8twSmfbArc3f
wzMOcfOw0K4O+3Z6kQKgd2iF2E9xHpjyX35DQmq/QInjyQF3VZIHUGuwqxh7UZCJ
6fSMzaTDCWdyBpQ5LjF41NoVUNeR84jMZsDzeQ3nwQKBgQD6Ju7RidTia1OfV3YG
JfKu9AqCKiTiHpksAzYRdj8szlWgwOh1f+ayfwdw/dNikf27Tf3e8bw6nrh4nr5d
mPrqupnWDHB2g65C6caciCBdaMbpXieS1aRfZG5e/kRpkC2QC+Ggg8AcEtwvsDZp
oaOptwbiuSyS9+9wWskNhg7WfQKBgQDDcSrjfGsqn9DBZZ0xEsP9oerqEmUuNwTC
trZHIunsz4+duFwUn7x59U91uKkfUxpzWUx1R3fBW93t5EPt1sMdN25hgI3jfiZ/
Xew5yH0XYD6QVtWCXzXXuNkJOPM4zh+96MAinNZKXnEJTCvTvA7ZTCfTMLOqV5Mw
Tm9lh3irUQKBgAEx2GglyV9/dbnIGCc1XTBauAYhH2X5EXA1X7e7odeb8KrA4RtW
jgqCMs3mWHEbE2QmnjTHYMfC5EynLW+TAHfIhl2QV6UpQdbN+QXcXJM1oeWRbozz
+kH+X8ySWE9MwfrzI5O5rVw09to/dDMS844m8qB4k+7rwjf+JwGqhz2dAoGAKq5E
F3nMTXqpNBLkyRq4AmOh0YxC3FzXhU4xcEeHnleVnXPtZ/OaTWfs+mBhTp3vYNFX
iSUaWfed952p1/7WjULVsCVK1yttbNMuC1BlQP2brBnKdrYkJAASJZlyRC1/cRGr
I+PsSEFnnggsagjflUS0TcKM+d42Ho6CdUGocIECgYAGQxhXWFxNUVaC6zVmTtGT
+OMudpk/vozIVt/Tg6ZBOs4jMdsGWuJfDpBiqCEHglIwPraZe617R7aT/vsxmxAJ
BIFPlCVzoh99pGSnCDFxan2gjKYZ/pVOtdy6GkKGgygInLvfQ73ykD2E6hDMvBET
MlCkAH+u7DZXSFvUTFmKAA==
-----END PRIVATE KEY-----`,
		MchidSerialNo: "21C36DEB8D8F9767B7AA84EBDBC2826587988125",
		WxV3Secret:    "4fg7dfgsd2er32ghk23o4kj23h4c2j3n",
	}

	//获取微信平台公钥
	//c.WxpayPublicKey = ""
	//c.WxpayPublicKeySeriaNo = ""

	return c
}

func (c *Cfg) WxpayPublicKey() []byte {
	//微信平台公钥不存在时，需要重新获取
	if c.wxpayPublicKey == "" {

	}
	return nil
}

func (c *Cfg) WxpayPublicKeySeriaNo(seriaNo string) (ok bool, err error) {
	// 序列号不对应时需要重新获取
	if c.wxpayPublicKeySeriaNo != seriaNo {

	}
	return true, nil
}
