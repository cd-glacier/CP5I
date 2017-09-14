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
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RecipeID int    `json:recipe_id`
	Quantity string `json:"quantity"`
}

type Method struct {
	ID       int    `json:"id"`
	ImageID  int    `json:"image_id"`
	RecipeID int    `json:recipe_id`
	Order    int    `json:method_order`
	Content  string `json:"content"`
}

type Kitchenware struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RecipeID int    `json:recipe_id`
}
