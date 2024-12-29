package command

type LoggedOutCommand struct {
	Token string `json:"token" binding:"required"`
	UserId string `json:"user_id"`
}

