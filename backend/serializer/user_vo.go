package serializer

import "backend/model"

// User 用户信息序列化器  VO view objective
type User struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	CreateAt int64  `json:"create_at"`
}

func BuildUser(user *model.User) User {
	return User{
		Id:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}

func BuildUsers(items []*model.User) (users []User) {
	for _, item := range items {
		users = append(users, BuildUser(item))
	}
	return
}
