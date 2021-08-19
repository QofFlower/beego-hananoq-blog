package errcode

const (
	SUCCESS                   uint = 0
	INVALID_USERNAME_PASSWORD uint = 21401
	INVALID_TOKEN             uint = 21402
	USER_INFO_EXIST           uint = 21403
	SERVER_INTERNAL_ERROR     uint = 21404
	USERNAME_PASSWORD_EMPTY   uint = 21405
	USERMELEN_PASSWORD_LENOUT uint = 21406
	ERROR                     uint = 21407
	INVALID_PARAM             uint = 21408
	INVALID_JSON_BODY         uint = 21409
	FREQUENCY_REQUEST         uint = 21410
	PERMISSION_DENY           uint = 21411
)

var errMsg map[uint]string

func msgInit() {
	errMsg = make(map[uint]string)
	e := errMsg
	e[SUCCESS] = "成功"
	e[ERROR] = "失败"
	e[INVALID_USERNAME_PASSWORD] = "用户名密码错误"
	e[INVALID_TOKEN] = "无效的token"
	e[USER_INFO_EXIST] = "用户名已被注册"
	e[SERVER_INTERNAL_ERROR] = "服务器内部错误"
	e[USERNAME_PASSWORD_EMPTY] = "用户名和密码为空"
	e[USERMELEN_PASSWORD_LENOUT] = "用户名密码长度超过20"
	e[INVALID_PARAM] = "非法参数"
	e[INVALID_JSON_BODY] = "非法请求body"
	e[FREQUENCY_REQUEST] = "请求过于频繁，请稍后再试"
	e[PERMISSION_DENY] = "权限不足"
}

func init() {
	msgInit()
}

func GetMsg(errCode uint) string {
	return errMsg[errCode]
}
