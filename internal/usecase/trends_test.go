package usecase

import (
	"errors"
	"testing"

	"github.com/IamVladlen/trend-bot/internal/entity"
	"github.com/IamVladlen/trend-bot/internal/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var errDummyTrends = errors.New("dummy trends error")

type testDeps struct {
	repo *mocks.MockTrendsRepo
	api  *mocks.MockTrendsWebAPI
}

func newTestTrendsUC(t *testing.T) (*TrendsUC, *testDeps) {
	t.Helper()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	deps := &testDeps{
		mocks.NewMockTrendsRepo(ctrl),
		mocks.NewMockTrendsWebAPI(ctrl),
	}
	uc := newTrendsUC(deps.repo, deps.api)

	return uc, deps
}

func TestTrendsUC_GetTrends(t *testing.T) {
	t.Parallel()

	uc, deps := newTestTrendsUC(t)

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		want    entity.Trends
		err     error
	}{
		{
			name: "unsigned user",
			prepare: func() {
				deps.repo.EXPECT().GetCountry(1).Return("", errDummyCountry).Times(1)
				deps.api.EXPECT().GetTrends(1).Return(entity.Trends{}, errDummyTrends).Times(1)
			},
			args: args{
				id: 1,
			},
			want: entity.Trends{},
			err:  errDummyCountry,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.prepare()

			got, err := uc.GetTrends(tt.args.id)

			assert.EqualValues(t, got, tt.want)
			assert.ErrorIs(t, err, tt.err)
		})
	}
}
