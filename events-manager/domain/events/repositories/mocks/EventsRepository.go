// Code generated by mockery v2.29.0. DO NOT EDIT.

package mocks

import (
	context "context"
	models "events-manager/domain/events/models"

	mock "github.com/stretchr/testify/mock"
)

// EventsRepository is an autogenerated mock type for the EventsRepository type
type EventsRepository struct {
	mock.Mock
}

// AddAttendeToEventById provides a mock function with given fields: ctx, id, attendeeEmail
func (_m *EventsRepository) AddAttendeToEventById(ctx context.Context, id string, attendeeEmail string) ([]string, error) {
	ret := _m.Called(ctx, id, attendeeEmail)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]string, error)); ok {
		return rf(ctx, id, attendeeEmail)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []string); ok {
		r0 = rf(ctx, id, attendeeEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, id, attendeeEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEvent provides a mock function with given fields: ctx, event
func (_m *EventsRepository) CreateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	ret := _m.Called(ctx, event)

	var r0 models.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Event) (models.Event, error)); ok {
		return rf(ctx, event)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Event) models.Event); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Event) error); ok {
		r1 = rf(ctx, event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEventById provides a mock function with given fields: ctx, id
func (_m *EventsRepository) DeleteEventById(ctx context.Context, id string) (models.Event, error) {
	ret := _m.Called(ctx, id)

	var r0 models.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.Event, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Event); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllEvents provides a mock function with given fields: ctx
func (_m *EventsRepository) GetAllEvents(ctx context.Context) ([]models.Event, error) {
	ret := _m.Called(ctx)

	var r0 []models.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.Event, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.Event); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllEventsByOrganizerEmail provides a mock function with given fields: ctx, id
func (_m *EventsRepository) GetAllEventsByOrganizerEmail(ctx context.Context, id string) ([]models.Event, error) {
	ret := _m.Called(ctx, id)

	var r0 []models.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]models.Event, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.Event); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventById provides a mock function with given fields: ctx, id
func (_m *EventsRepository) GetEventById(ctx context.Context, id string) (models.Event, error) {
	ret := _m.Called(ctx, id)

	var r0 models.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.Event, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Event); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEvent provides a mock function with given fields: ctx, event
func (_m *EventsRepository) UpdateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	ret := _m.Called(ctx, event)

	var r0 models.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Event) (models.Event, error)); ok {
		return rf(ctx, event)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Event) models.Event); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Get(0).(models.Event)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Event) error); ok {
		r1 = rf(ctx, event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEventsRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventsRepository creates a new instance of EventsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventsRepository(t mockConstructorTestingTNewEventsRepository) *EventsRepository {
	mock := &EventsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
