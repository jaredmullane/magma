// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	storage "magma/lte/cloud/go/services/smsd/storage"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// SMSStorage is an autogenerated mock type for the SMSStorage type
type SMSStorage struct {
	mock.Mock
}

// CreateSMS provides a mock function with given fields: networkID, sms
func (_m *SMSStorage) CreateSMS(networkID string, sms storage.MutableSMS) (string, error) {
	ret := _m.Called(networkID, sms)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, storage.MutableSMS) string); ok {
		r0 = rf(networkID, sms)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, storage.MutableSMS) error); ok {
		r1 = rf(networkID, sms)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSMSs provides a mock function with given fields: networkID, pks
func (_m *SMSStorage) DeleteSMSs(networkID string, pks []string) error {
	ret := _m.Called(networkID, pks)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(networkID, pks)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSMSs provides a mock function with given fields: networkID, pks, imsis, onlyWaiting, startTime, endTime
func (_m *SMSStorage) GetSMSs(networkID string, pks []string, imsis []string, onlyWaiting bool, startTime *time.Time, endTime *time.Time) ([]*storage.SMS, error) {
	ret := _m.Called(networkID, pks, imsis, onlyWaiting, startTime, endTime)

	var r0 []*storage.SMS
	if rf, ok := ret.Get(0).(func(string, []string, []string, bool, *time.Time, *time.Time) []*storage.SMS); ok {
		r0 = rf(networkID, pks, imsis, onlyWaiting, startTime, endTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*storage.SMS)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string, []string, bool, *time.Time, *time.Time) error); ok {
		r1 = rf(networkID, pks, imsis, onlyWaiting, startTime, endTime)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSMSsToDeliver provides a mock function with given fields: networkID, imsis, timeoutThreshold
func (_m *SMSStorage) GetSMSsToDeliver(networkID string, imsis []string, timeoutThreshold time.Duration) ([]*storage.SMS, error) {
	ret := _m.Called(networkID, imsis, timeoutThreshold)

	var r0 []*storage.SMS
	if rf, ok := ret.Get(0).(func(string, []string, time.Duration) []*storage.SMS); ok {
		r0 = rf(networkID, imsis, timeoutThreshold)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*storage.SMS)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string, time.Duration) error); ok {
		r1 = rf(networkID, imsis, timeoutThreshold)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields:
func (_m *SMSStorage) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReportDelivery provides a mock function with given fields: networkID, deliveredMessages, failedMessages
func (_m *SMSStorage) ReportDelivery(networkID string, deliveredMessages map[string][]byte, failedMessages map[string][]storage.SMSFailureReport) error {
	ret := _m.Called(networkID, deliveredMessages, failedMessages)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string][]byte, map[string][]storage.SMSFailureReport) error); ok {
		r0 = rf(networkID, deliveredMessages, failedMessages)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}