// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	storage "github.com/kyma-project/test-infra/development/tools/pkg/gcscleaner/storage"
	mock "github.com/stretchr/testify/mock"
)

// ObjectIterator is an autogenerated mock type for the ObjectIterator type
type ObjectIterator struct {
	mock.Mock
}

// Next provides a mock function with given fields:
func (_m *ObjectIterator) Next() (storage.ObjectAttrs, error) {
	ret := _m.Called()

	var r0 storage.ObjectAttrs
	if rf, ok := ret.Get(0).(func() storage.ObjectAttrs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.ObjectAttrs)
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
