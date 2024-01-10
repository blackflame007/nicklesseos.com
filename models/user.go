package models

// create user struct
type UserInfo struct {
	Email     string `json:"email"`
	FullName  string `json:"name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Picture   string `json:"picture"`
}
