package model

type Recipe struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	PruducerID  int          `json:"pruducer_id"`
	Difficuly   int          `json:"difficuly"`
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
