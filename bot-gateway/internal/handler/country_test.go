package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validateCountry(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			name: "valid_country",
			args: args{
				text: "🇺🇸",
			},
			want: "US",
			err:  nil,
		},
		{
			name: "invalid_country",
			args: args{
				text: "invalid",
			},
			want: "",
			err:  errInvalidCountry,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateCountry(tt.args.text)

			assert.EqualValues(t, got, tt.want)
			assert.ErrorIs(t, err, tt.err)
		})
	}
}
