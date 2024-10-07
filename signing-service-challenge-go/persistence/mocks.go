package persistence

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/stretchr/testify/mock"
)

type MockDeviceStoreRepo struct {
	mock.Mock
}

func (m *MockDeviceStoreRepo) Save(signDevice *domain.SignatureDevice) {
	m.Called(signDevice)
}

func (m *MockDeviceStoreRepo) GetById(id string) *domain.SignatureDevice {
	args := m.Called(id)
	return args.Get(0).(*domain.SignatureDevice)
}

func (m *MockDeviceStoreRepo) GetAll() []interface{} {
	args := m.Called()
	return args.Get(0).([]interface{})
}

func (m *MockDeviceStoreRepo) IncrementCounter(deviceId string) {
	m.Called(deviceId)
}

type MockTransactionStoreRepo struct {
	mock.Mock
}

func (m *MockTransactionStoreRepo) Save(value interface{}) {
	m.Called(value)
}

func (m *MockTransactionStoreRepo) GetByDevice(deviceId string) []*domain.Transaction {
	args := m.Called(deviceId)
	return args.Get(0).([]*domain.Transaction)
}
