package config

import "git.code.oa.com/trpc-go/trpc-go/errs"

// Status code msg 返回状态
type Status struct {
	Code int32
	Msg  string
}

var (
	ResOk                    = Status{0, "Success"}
	InnerProcessingTimeout   = Status{-100, "内部处理超时"}
	InnerServiceOverload     = Status{-200, "内部服务过载"}
	InnerLogicError          = Status{-300, "内部逻辑出错"}
	InnerFrameworkError      = Status{-321, "内部框架错误"}
	InnerOtherError          = Status{-400, "内部其他问题"}
	InnerReadDbError         = Status{-401, "读取数据库错误"}
	InnerWriteDbError        = Status{-402, "写入数据库错误"}
	InnerDeleteDbError       = Status{-403, "删除数据库错误"}
	InnerMarshalError        = Status{-427, "数据序列化失败"}
	InnerUnmarshalError      = Status{-428, "数据反序列化失败"}
	InnerEncryptError        = Status{-429, "数据加密失败"}
	InnerDecryptError        = Status{-430, "数据解密失败"}
	ClientParamParsingError  = Status{100, "客户端请求参数解析失败"}
	ClientPermissionError    = Status{200, "客户端没有访问权限"}
	ClientInvalidTokenError  = Status{221, "客户端Token无效"}
	ClientForgedIdentity     = Status{224, "客户端伪造身份"}
	ClientContentError       = Status{300, "客户端请求内容问题"}
	ClientRecordNotFound     = Status{301, "记录未找到"}
	ClientCheckUserNameError = Status{302, "用户名重复"}
	ClientUPInvalid          = Status{303, "用户名或密码不规范"}
	ClientLoginError         = Status{304, "登录失败"}
	ClientUserInfoError      = Status{305, "用户信息不存在"}
	ClientInvalidTimeRange   = Status{320, "时间范围无效"}
	ClientInvalidParamError  = Status{321, "客户端请求参数校验失败"}
	ClientExtractTokenError  = Status{326, "解析Token失败"}
	ClientNoTokenError       = Status{401, "客户端缺少Token"}
	ClientLoginExpired       = Status{403, "用户登陆过期"}
)

// New 错误构造方法
func New(err Status) error {
	return errs.New(int(err.Code), err.Msg)
}
