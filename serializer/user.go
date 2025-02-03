package serializer

import "ToDoList/model"

type User struct {
	ID        uint   `json:"id" form:"id" example:"1"`
	UserName  string `json:"user_name" form:"user_name" example:"user"`
	Status    string `json:"status" form:"status"`
	CreatedAt int64  `json:"created_at" form:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
