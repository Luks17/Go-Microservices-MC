package validators

import (
	"github.com/Luks17/Go-Microservices-MC/db/repository/sqlc"
	"github.com/go-playground/validator/v10"
)

var ValidCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(sqlc.Currencies); ok {
		return currency.Valid()
	}
	return false
}
