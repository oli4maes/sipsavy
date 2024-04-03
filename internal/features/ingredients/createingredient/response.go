package createingredient

type Response struct {
	Ingredient ingredientDto `json:"ingredient"`
}

type ingredientDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
