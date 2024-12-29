package command

type ChangePasswordCommand struct {
	UserId       string `json:"userId" binding:"required"`
	Password     string `json:"password" binding:"required"`
}