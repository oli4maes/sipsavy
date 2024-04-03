package getallingredients

type Response struct {
	Ingredients []ingredientDto `json:"ingredients"`
}

type ingredientDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
