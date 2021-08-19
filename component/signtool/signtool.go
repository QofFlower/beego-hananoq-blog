package signtool

import (
	signutil "github.com/parkingwang/go-sign"
	"time"
)

func CheckSign() {
	requestUri := "/restful/api/numbers?appid=9d8a121ce581499d&nonce_str=ibuaiVcKdpRxkhJA&plate_number=豫A66666" +
		"&time_stamp=1532585241&sign=072defd1a251dc58e4d1799e17ffe7a4"

	// 第一步：创建GoVerifier校验类
	verifier := signutil.NewGoVerifier()

	// 假定从RequestUri中读取校验参数
	if err := verifier.ParseQuery(requestUri); nil != err {
		panic(err)
	}

	// 或者使用verifier.ParseValues(Values)来解析。

	// 第二步：（可选）校验是否包含签名校验必要的参数
	if err := verifier.MustHasOtherKeys("plate_number"); nil != err {
		panic(err)
	}

	// 第三步：检查时间戳是否超时。

	// 时间戳超时：5分钟
	verifier.SetTimeout(time.Minute * 5)
	if err := verifier.CheckTimeStamp(); nil != err {
		panic(err)
	}

	// 第四步: 创建GoSigner来重现客户端的签名信息
	signer := signutil.NewGoSignerMd5()

	// 第五步：从GoVerifier中读取所有请求参数
	signer.SetBody(verifier.GetBodyWithoutSign())

	// 第六步：从数据库读取AppID对应的SecretKey
	// appid := verifier.GetAppId()
	secretKey := "d93047a4d6fe6111"

	// 使用同样的WrapBody方式
	signer.SetAppSecretWrapBody(secretKey)

	// 服务端根据客户端参数生成签名
	sign := signer.GetSignature()

	// 最后，比较服务端生成的签名信息，与客户端提供的签名是否一致即可。
	if verifier.MustString("sign") != sign {
		panic("校验失败")
	}
}
