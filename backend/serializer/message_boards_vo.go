package serializer

import "backend/model"

type Message struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Text     string `json:"text"`
	CreateAt int64  `json:"create_at"`
}

func BuildMsg(msg *model.Message) Message {
	return Message{
		Id:       msg.ID,
		UserId:   msg.UserId,
		Text:     msg.Text,
		CreateAt: msg.CreatedAt.Unix(),
	}
}

func BuildMsgs(items []*model.Message) (msgs []Message) {
	for _, item := range items {
		msgs = append(msgs, BuildMsg(item))
	}
	return
}
