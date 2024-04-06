package dto

import "github.com/go-playground/validator/v10"

type ProductCreateDTO struct {
	Nombre string  `json:"name" validate:"required,min=2,max=100"`
	Precio float64 `json:"price" validate:"required,gt=0"`
}

func (p *ProductCreateDTO) Validate() []*validator.FieldError {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		var fieldErrors []*validator.FieldError

		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, fe := range ve {
				fieldError := fe
				fieldErrors = append(fieldErrors, &fieldError)
			}
			return fieldErrors
		}
	}
	return nil
}
