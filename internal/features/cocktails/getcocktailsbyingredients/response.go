package getcocktailsbyingredients

type Response struct {
	Cocktails []cocktailDto `json:"cocktails"`
}

type cocktailDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
