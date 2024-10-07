package crypto

import "github.com/stretchr/testify/mock"

type MockedKeyPair struct {
	mock.Mock
}

func (m *MockedKeyPair) PublicKey() interface{} {
	args := m.Called()
	return args.Get(0)
}

func (m *MockedKeyPair) PrivateKey() interface{} {
	args := m.Called()
	return args.Get(0)
}

type MockedGenerator struct {
	mock.Mock
}

func (m *MockedGenerator) Generate() (KeyPair, error) {
	args := m.Called()
	return args.Get(0).(KeyPair), args.Error(1)
}
