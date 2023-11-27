package serializer

import (
	"backend/dao"
)

type MsgResult struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Text     string `json:"text"`
	CreateAt int64  `json:"create_at"`
	UserName string `json:"user_name"`
}

func BuildMsg(msg *dao.MsgResult) MsgResult {
	return MsgResult{
		Id:       msg.Id,
		UserId:   msg.UserId,
		Text:     msg.Text,
		CreateAt: msg.CreatedAt.Unix(),
		UserName: msg.UserName,
	}
}

func BuildMsgs(items []*dao.MsgResult) (msgs []MsgResult) {
	for _, item := range items {
		msgs = append(msgs, BuildMsg(item))
	}
	return
}
