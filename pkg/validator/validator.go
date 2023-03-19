package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key string
	Msg string
}

func (v *ValidError) Error() string {
	return v.Msg
}

type ValidErrors []*ValidError

func (vs *ValidErrors) Error() string {
	return strings.Join(vs.Errors(), ",")
}

func (vs *ValidErrors) Errors() []string {
	var errs []string
	for _, v := range *vs {
		errs = append(errs, v.Error())
	}
	return errs
}

func BindAndValid(c *gin.Context, v any) (bool, ValidErrors) {
	var vs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, vs
		}
		for _, v := range verrs {
			vs = append(vs, &ValidError{
				Key: v.Namespace(),
				Msg: v.Error(),
			})
		}
		return false, vs
	}
	return true, nil
}
