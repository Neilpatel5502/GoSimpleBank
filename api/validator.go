package api

import (
	"github.com/Neilpatel5502/GoSimpleBank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(feildLevel validator.FieldLevel) bool {
	if currency, ok := feildLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
