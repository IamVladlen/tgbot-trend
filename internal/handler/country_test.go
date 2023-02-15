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
		name    string
		args    args
		want    string
		wantErr bool
		errType error
	}{
		{
			name: "valid_country",
			args: args{
				text: "🇺🇸",
			},
			want:    "US",
			wantErr: false,
			errType: nil,
		},
		{
			name: "invalid_country",
			args: args{
				text: "invalid",
			},
			want:    "",
			wantErr: false,
			errType: errInvalidCountry,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateCountry(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unexpected error: %s", err.Error())
				assert.EqualError(t, tt.errType, err.Error())
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
