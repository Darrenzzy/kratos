package ecode

// All common ecode
var (
	OK   = add(0)  // 正确
	Fail = add(-1) // 失败

	NotModified        = add(304)  // 木有改动
	TemporaryRedirect  = add(307)  // 撞车跳转
	RequestErr         = add(400)  // 请求错误
	Unauthorized       = add(401)  // 未认证
	UnauthorizedUser   = add(4010) // 未实名
	UnauthorizedTread  = add(4011) // 未申请支付
	AccessDenied       = add(403)  // 访问权限不足
	NothingFound       = add(404)  // 啥都木有
	MethodNotAllowed   = add(405)  // 不支持该方法
	Conflict           = add(409)  // 冲突
	Canceled           = add(498)  // 客户端取消请求
	ServerErr          = add(500)  // 服务器错误
	ServiceUnavailable = add(503)  // 过载保护,服务暂不可用
	Deadline           = add(504)  // 服务调用超时
	LimitExceed        = add(509)  // 超出限制

)

const (
	ErrLoadCfgMsg  = "获取配置信息失败"
	ErrParamMsg    = "参数验证失败"
	ErrCalcMsg     = "获取价格信息失败"
	ErrNotFoundMsg = "查询数据失败"
)
