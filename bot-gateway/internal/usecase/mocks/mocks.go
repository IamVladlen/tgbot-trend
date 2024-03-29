// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/IamVladlen/trend-bot/bot-gateway/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockCountryRepo is a mock of CountryRepo interface.
type MockCountryRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCountryRepoMockRecorder
}

// MockCountryRepoMockRecorder is the mock recorder for MockCountryRepo.
type MockCountryRepoMockRecorder struct {
	mock *MockCountryRepo
}

// NewMockCountryRepo creates a new mock instance.
func NewMockCountryRepo(ctrl *gomock.Controller) *MockCountryRepo {
	mock := &MockCountryRepo{ctrl: ctrl}
	mock.recorder = &MockCountryRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCountryRepo) EXPECT() *MockCountryRepoMockRecorder {
	return m.recorder
}

// ChangeCountry mocks base method.
func (m *MockCountryRepo) ChangeCountry(id int, country string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeCountry", id, country)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeCountry indicates an expected call of ChangeCountry.
func (mr *MockCountryRepoMockRecorder) ChangeCountry(id, country interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeCountry", reflect.TypeOf((*MockCountryRepo)(nil).ChangeCountry), id, country)
}

// GetCountry mocks base method.
func (m *MockCountryRepo) GetCountry(id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountry", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountry indicates an expected call of GetCountry.
func (mr *MockCountryRepoMockRecorder) GetCountry(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountry", reflect.TypeOf((*MockCountryRepo)(nil).GetCountry), id)
}

// MockTrendsRepo is a mock of TrendsRepo interface.
type MockTrendsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTrendsRepoMockRecorder
}

// MockTrendsRepoMockRecorder is the mock recorder for MockTrendsRepo.
type MockTrendsRepoMockRecorder struct {
	mock *MockTrendsRepo
}

// NewMockTrendsRepo creates a new mock instance.
func NewMockTrendsRepo(ctrl *gomock.Controller) *MockTrendsRepo {
	mock := &MockTrendsRepo{ctrl: ctrl}
	mock.recorder = &MockTrendsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrendsRepo) EXPECT() *MockTrendsRepoMockRecorder {
	return m.recorder
}

// GetCountry mocks base method.
func (m *MockTrendsRepo) GetCountry(id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountry", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountry indicates an expected call of GetCountry.
func (mr *MockTrendsRepoMockRecorder) GetCountry(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountry", reflect.TypeOf((*MockTrendsRepo)(nil).GetCountry), id)
}

// MockTrendsWebAPI is a mock of TrendsWebAPI interface.
type MockTrendsWebAPI struct {
	ctrl     *gomock.Controller
	recorder *MockTrendsWebAPIMockRecorder
}

// MockTrendsWebAPIMockRecorder is the mock recorder for MockTrendsWebAPI.
type MockTrendsWebAPIMockRecorder struct {
	mock *MockTrendsWebAPI
}

// NewMockTrendsWebAPI creates a new mock instance.
func NewMockTrendsWebAPI(ctrl *gomock.Controller) *MockTrendsWebAPI {
	mock := &MockTrendsWebAPI{ctrl: ctrl}
	mock.recorder = &MockTrendsWebAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrendsWebAPI) EXPECT() *MockTrendsWebAPIMockRecorder {
	return m.recorder
}

// GetTrends mocks base method.
func (m *MockTrendsWebAPI) GetTrends(country string) (entity.Trends, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrends", country)
	ret0, _ := ret[0].(entity.Trends)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrends indicates an expected call of GetTrends.
func (mr *MockTrendsWebAPIMockRecorder) GetTrends(country interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrends", reflect.TypeOf((*MockTrendsWebAPI)(nil).GetTrends), country)
}

// MockTrendsMicroservice is a mock of TrendsMicroservice interface.
type MockTrendsMicroservice struct {
	ctrl     *gomock.Controller
	recorder *MockTrendsMicroserviceMockRecorder
}

// MockTrendsMicroserviceMockRecorder is the mock recorder for MockTrendsMicroservice.
type MockTrendsMicroserviceMockRecorder struct {
	mock *MockTrendsMicroservice
}

// NewMockTrendsMicroservice creates a new mock instance.
func NewMockTrendsMicroservice(ctrl *gomock.Controller) *MockTrendsMicroservice {
	mock := &MockTrendsMicroservice{ctrl: ctrl}
	mock.recorder = &MockTrendsMicroserviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrendsMicroservice) EXPECT() *MockTrendsMicroserviceMockRecorder {
	return m.recorder
}

// GetScheduledMessages mocks base method.
func (m *MockTrendsMicroservice) GetScheduledMessages(ctx context.Context, interval string) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScheduledMessage", ctx, interval)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScheduledMessage indicates an expected call of GetScheduledMessage.
func (mr *MockTrendsMicroserviceMockRecorder) GetScheduledMessage(ctx, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduledMessage", reflect.TypeOf((*MockTrendsMicroservice)(nil).GetScheduledMessages), ctx, interval)
}

// SetChatSchedule mocks base method.
func (m *MockTrendsMicroservice) SetChatSchedule(ctx context.Context, chatId int64, interval string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetChatSchedule", ctx, chatId, interval)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetChatSchedule indicates an expected call of SetChatSchedule.
func (mr *MockTrendsMicroserviceMockRecorder) SetChatSchedule(ctx, chatId, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetChatSchedule", reflect.TypeOf((*MockTrendsMicroservice)(nil).SetChatSchedule), ctx, chatId, interval)
}
