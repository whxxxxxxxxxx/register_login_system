package serializer

import "register_log/model"

type User struct {
	ID       uint   `json:"id"`
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Type     string `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

func BuildUser(user *model.User) *User {
	return &User{
		ID:       user.ID,
		NickName: user.NickName,
		UserName: user.UserName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}
