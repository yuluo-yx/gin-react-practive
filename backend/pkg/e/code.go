package e

// 状态码

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// 用户错误类型
	ErrorExitUser                  = 5001
	ErrorFailEncryption            = 5002
	ErrorExitUserNotFound          = 5003
	ErrorNotCompare                = 5004
	ErrorAuthToken                 = 5005
	ErrorAuthCheckTokenTimeout     = 5006
	ErrorAuthCheckTokenFail        = 5007
	ErrorAuthInsufficientAuthority = 5008
	ErrorUploadFail                = 5009

	// 留言错误类型
	ErrorMsgSave   = 6001
	ErrorMsgGet    = 6002
	ErrorMsgDelete = 6003
)
