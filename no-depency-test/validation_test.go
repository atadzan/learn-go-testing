package validation

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateError(t *testing.T) {
	cases := []struct {
		name   string
		in     *Input
		expErr error
	}{
		{
			name:   "bad_user_id",
			in:     &Input{UserId: 0},
			expErr: errors.New("invalid id param"),
		},
		{
			name:   "bad_payment_type",
			in:     &Input{UserId: 2, PaymentType: "bad"},
			expErr: errors.New("payment type must be card"),
		},
	}
	for _, tCase := range cases {
		err := Validate(*tCase.in)
		require.Error(t, err)
		require.EqualError(t, tCase.expErr, err.Error())
	}
}
