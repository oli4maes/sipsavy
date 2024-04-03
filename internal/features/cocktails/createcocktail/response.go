package createcocktail

type Response struct {
	Cocktail cocktailDto `json:"cocktail"`
}

type cocktailDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
