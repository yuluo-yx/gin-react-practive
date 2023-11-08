package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "参数错误",

	// 用户错误类型
	ErrorExitUser:                  "用户名已经存在",
	ErrorFailEncryption:            "用户密码加密失败",
	ErrorExitUserNotFound:          "用户不存在",
	ErrorNotCompare:                "用户密码检验错误",
	ErrorAuthToken:                 "token认证失败",
	ErrorAuthCheckTokenTimeout:     "token 已过期",
	ErrorAuthCheckTokenFail:        "token 验证失败",
	ErrorAuthInsufficientAuthority: "token 认证权限不足",

	// 留言错误类型
	ErrorMsgSave:   "留言保存失败",
	ErrorMsgGet:    "留言获取失败",
	ErrorMsgDelete: "留言删除失败",
}

// GetMsg 获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}

	return msg
}
