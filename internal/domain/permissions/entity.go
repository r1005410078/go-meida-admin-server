package permissions

type Permission struct {
	ID   *string `json:"id"`
	Name string 	`json:"name"`
	Description string `json:"description"`
	Action string `json:"action"`
}