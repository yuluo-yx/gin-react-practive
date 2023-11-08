package serializer

import (
	"encoding/json"
)

// Response 响应序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// TokenData token序列化器
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

func BuildListResponse(item interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  item,
			Total: total,
		},
		Msg: "ok",
	}
}

func ErrorResponse(err error) Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return Response{
			Status: 400,
			Msg:    "json类型不匹配",
			Error:  err.Error(),
		}
	}

	return Response{
		Status: 400,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}
