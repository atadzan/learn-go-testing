package validation

import "errors"

type Input struct {
	UserId      int
	PaymentType string
	Items       []string
}

func Validate(input Input) error {
	if input.UserId == 0 {
		return errors.New("invalid id param")
	}

	if input.PaymentType != "card" {
		return errors.New("payment type must be card")
	}
	if len(input.Items) == 0 {
		return errors.New("items can't be blank")
	}
	return nil

}
