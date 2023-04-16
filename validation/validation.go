package validation

import (
	"github.com/asaskevich/govalidator"
)

func Validate(request interface{}) (err error) {
	_, err = govalidator.ValidateStruct(request)
	return err
}
