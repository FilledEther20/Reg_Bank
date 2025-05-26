package api

import (
	"github.com/FilledEther20/Reg_Bank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(field validator.FieldLevel) bool {
	if currency, ok := field.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
