package command

type RegisterCommand struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verificationCode" binding:"required"`
}
