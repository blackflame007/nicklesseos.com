package models

// create UserInfo struct
type UserInfo struct {
	Email     string `json:"email"`
	FullName  string `json:"name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Picture   string `json:"picture"`
}

// create User struct

type User struct {
	UserInfo  UserInfo
	HighScore int `json:"high_score"`
}
