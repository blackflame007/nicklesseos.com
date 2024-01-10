package models

// GameContent is the meta data for a game
type GameContent struct {
	Game        string `json:"game"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	URL         string `json:"-"` // The "-" tag means this field is not affected by encoding/json
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}
