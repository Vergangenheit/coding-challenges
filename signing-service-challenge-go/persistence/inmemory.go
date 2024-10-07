package persistence

import (
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

type DeviceStore interface {
	Save(value *domain.SignatureDevice)
	GetById(id string) *domain.SignatureDevice
	GetAll() []interface{}
}

// TODO: in-memory persistence ...
type InMemoryDeviceStore map[string]interface{}

func NewInMemoryDeviceStore() DeviceStore {
	return &InMemoryDeviceStore{}
}

func (p *InMemoryDeviceStore) Save(value *domain.SignatureDevice) {
	if value.Id == "" {
		value.Id = uuid.New().String()
	}
	(*p)[value.Id] = value
}

func (p *InMemoryDeviceStore) GetById(id string) *domain.SignatureDevice {
	inter := (*p)[id]
	device, ok := inter.(*domain.SignatureDevice)
	if !ok {
		return nil
	}
	return device
}

func (p *InMemoryDeviceStore) GetAll() []interface{} {
	values := make([]interface{}, 0, len(*p))
	for _, value := range *p {
		values = append(values, value)
	}
	return values
}

type TransactionStore interface {
	Save(value interface{})
	GetByDevice(deviceId string) []*domain.Transaction
}

type InMemoryTransactionStore map[string]interface{}

func NewInMemoryTransactionStore() TransactionStore {
	return &InMemoryTransactionStore{}
}

func (p *InMemoryTransactionStore) Save(value interface{}) {
	// create key as uuid string
	id := uuid.New().String()
	(*p)[id] = value
}

// extract all transactions handled by a specific device
func (p *InMemoryTransactionStore) GetByDevice(deviceId string) []*domain.Transaction {
	deviceTransactions := make([]*domain.Transaction, 0)
	for _, value := range *p {
		transaction := value.(*domain.Transaction)
		if transaction.DeviceId == deviceId {
			deviceTransactions = append(deviceTransactions, transaction)
		}
	}
	return deviceTransactions
}
