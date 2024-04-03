package createingredient

import "github.com/go-ozzo/ozzo-validation"

type Request struct {
	Name string `json:"name"`
}

func (r Request) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.Length(3, 250)))
}
