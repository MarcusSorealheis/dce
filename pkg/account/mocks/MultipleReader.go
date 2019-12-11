// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/Optum/dce/pkg/model"

// MultipleReader is an autogenerated mock type for the MultipleReader type
type MultipleReader struct {
	mock.Mock
}

// GetAccounts provides a mock function with given fields:
func (_m *MultipleReader) GetAccounts() (*model.Accounts, error) {
	ret := _m.Called()

	var r0 *model.Accounts
	if rf, ok := ret.Get(0).(func() *model.Accounts); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Accounts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountsByPrincipalID provides a mock function with given fields: principalID
func (_m *MultipleReader) GetAccountsByPrincipalID(principalID string) (*model.Accounts, error) {
	ret := _m.Called(principalID)

	var r0 *model.Accounts
	if rf, ok := ret.Get(0).(func(string) *model.Accounts); ok {
		r0 = rf(principalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Accounts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(principalID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountsByStatus provides a mock function with given fields: status
func (_m *MultipleReader) GetAccountsByStatus(status string) (*model.Accounts, error) {
	ret := _m.Called(status)

	var r0 *model.Accounts
	if rf, ok := ret.Get(0).(func(string) *model.Accounts); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Accounts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}