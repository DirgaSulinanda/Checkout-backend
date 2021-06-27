// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/repository/promo/repository.go

// Package mock_promo is a generated GoMock package.
package mock_promo

import (
	context "context"
	reflect "reflect"

	promo "github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/promo"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetActivePromo mocks base method.
func (m *MockRepository) GetActivePromo(ctx context.Context) ([]promo.Promo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActivePromo", ctx)
	ret0, _ := ret[0].([]promo.Promo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActivePromo indicates an expected call of GetActivePromo.
func (mr *MockRepositoryMockRecorder) GetActivePromo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActivePromo", reflect.TypeOf((*MockRepository)(nil).GetActivePromo), ctx)
}
