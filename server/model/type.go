package model

type Recipe struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	ProducerID  int          `json:"producer_id"`
	Difficulty  int          `json:"difficulty"`
	Time        string       `json:"time"`
	Ingredients []Ingredient `json:"ingredients"`
	Method      []Method     `json:"method"`
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

type Method struct {
	ImageID int    `json:"image_id"`
	Content string `json:"content"`
}
