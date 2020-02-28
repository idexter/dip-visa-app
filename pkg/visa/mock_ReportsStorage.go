// Code generated by mockery v1.0.0. DO NOT EDIT.

package visa

import mock "github.com/stretchr/testify/mock"

// MockReportsStorage is an autogenerated mock type for the ReportsStorage type
type MockReportsStorage struct {
	mock.Mock
}

// LoadApplicationReport provides a mock function with given fields: id
func (_m *MockReportsStorage) LoadApplicationReport(id int) (*Report, error) {
	ret := _m.Called(id)

	var r0 *Report
	if rf, ok := ret.Get(0).(func(int) *Report); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Report)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveApplicationReport provides a mock function with given fields: _a0
func (_m *MockReportsStorage) SaveApplicationReport(_a0 Report) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(Report) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
