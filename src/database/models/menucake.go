package models

type MenuCake struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	Image       string `json:"image"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
