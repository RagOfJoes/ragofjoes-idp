// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	registration "github.com/RagOfJoes/mylo/flow/registration"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/gofrs/uuid"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, newFlow
func (_m *Repository) Create(ctx context.Context, newFlow registration.Flow) (*registration.Flow, error) {
	ret := _m.Called(ctx, newFlow)

	var r0 *registration.Flow
	if rf, ok := ret.Get(0).(func(context.Context, registration.Flow) *registration.Flow); ok {
		r0 = rf(ctx, newFlow)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registration.Flow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, registration.Flow) error); ok {
		r1 = rf(ctx, newFlow)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *Repository) Get(ctx context.Context, id string) (*registration.Flow, error) {
	ret := _m.Called(ctx, id)

	var r0 *registration.Flow
	if rf, ok := ret.Get(0).(func(context.Context, string) *registration.Flow); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registration.Flow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByFlowID provides a mock function with given fields: ctx, flowID
func (_m *Repository) GetByFlowID(ctx context.Context, flowID string) (*registration.Flow, error) {
	ret := _m.Called(ctx, flowID)

	var r0 *registration.Flow
	if rf, ok := ret.Get(0).(func(context.Context, string) *registration.Flow); ok {
		r0 = rf(ctx, flowID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registration.Flow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, flowID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, updateFlow
func (_m *Repository) Update(ctx context.Context, updateFlow registration.Flow) (*registration.Flow, error) {
	ret := _m.Called(ctx, updateFlow)

	var r0 *registration.Flow
	if rf, ok := ret.Get(0).(func(context.Context, registration.Flow) *registration.Flow); ok {
		r0 = rf(ctx, updateFlow)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*registration.Flow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, registration.Flow) error); ok {
		r1 = rf(ctx, updateFlow)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
