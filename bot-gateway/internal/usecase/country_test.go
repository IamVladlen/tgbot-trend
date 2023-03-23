package usecase

import (
	"errors"
	"testing"

	"github.com/IamVladlen/trend-bot/bot-gateway/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var errDummyCountry = errors.New("dummy country error")

func newTestCountryUC(t *testing.T) (*CountryUC, *mocks.MockCountryRepo) {
	t.Helper()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCountryRepo(ctrl)
	uc := newCountryUC(repo)

	return uc, repo
}

func TestCountryUC_ChangeCountry(t *testing.T) {
	t.Parallel()

	uc, repo := newTestCountryUC(t)

	type args struct {
		id      int
		country string
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		err     error
	}{
		{
			name: "result error",
			prepare: func() {
				repo.EXPECT().ChangeCountry(-5, "🇩🇪").Return(errDummyCountry).Times(1)
			},
			args: args{
				-5,
				"🇩🇪",
			},
			err: errDummyCountry,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.prepare()

			err := uc.ChangeCountry(tt.args.id, tt.args.country)

			assert.ErrorIs(t, err, tt.err)
		})
	}
}
