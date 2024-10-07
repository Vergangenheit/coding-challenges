package api

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateSignatureDevice_EmptyBody(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/api/v0/device", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	mockDeviceStoreRepo := &persistence.MockDeviceStoreRepo{}
	mockTransactionStoreRepo := &persistence.MockTransactionStoreRepo{}

	s := &Server{
		listenAddress:    ":8081",
		deviceStore:      mockDeviceStoreRepo,
		transactionStore: mockTransactionStoreRepo,
	}

	// Record the response
	rec := httptest.NewRecorder()

	s.SignatureDevice(rec, req)

	// Validate the status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "request body must not be empty")
}

func Test_CreateSignatureDevice_Ok(t *testing.T) {
	createRaw := map[string]interface{}{
		"signature_algorithm": "RSA",
		"label":               "label",
	}
	// Marshal the JSON data into a byte slice
	jsonData, err := json.Marshal(createRaw)
	if err != nil {
		t.Fatalf("Could not marshal JSON: %v", err)
	}
	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/api/v0/device", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	mockDeviceStoreRepo := &persistence.MockDeviceStoreRepo{}
	mockDeviceStoreRepo.On("Save", mock.Anything).Return()
	mockTransactionStoreRepo := &persistence.MockTransactionStoreRepo{}

	s := &Server{
		listenAddress:    ":8081",
		deviceStore:      mockDeviceStoreRepo,
		transactionStore: mockTransactionStoreRepo,
	}

	// Record the response
	rec := httptest.NewRecorder()

	s.SignatureDevice(rec, req)

	// Validate the status code
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func Test_CreateSignatureDevice_CannotDecode(t *testing.T) {
	createRaw := map[string]interface{}{
		"signature": "RSA",
		"labell":    "label",
	}
	// Marshal the JSON data into a byte slice
	jsonData, err := json.Marshal(createRaw)
	if err != nil {
		t.Fatalf("Could not marshal JSON: %v", err)
	}
	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/api/v0/device", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	mockDeviceStoreRepo := &persistence.MockDeviceStoreRepo{}
	mockTransactionStoreRepo := &persistence.MockTransactionStoreRepo{}

	s := &Server{
		listenAddress:    ":8081",
		deviceStore:      mockDeviceStoreRepo,
		transactionStore: mockTransactionStoreRepo,
	}

	// Record the response
	rec := httptest.NewRecorder()

	s.SignatureDevice(rec, req)

	// Validate the status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func Test_CreateSignatureDevice_EmptyDeviceStore(t *testing.T) {
	createRaw := map[string]interface{}{
		"signature_algorithm": "RSA",
		"label":               "label",
	}
	// Marshal the JSON data into a byte slice
	jsonData, err := json.Marshal(createRaw)
	if err != nil {
		t.Fatalf("Could not marshal JSON: %v", err)
	}
	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/api/v0/device", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	mockTransactionStoreRepo := &persistence.MockTransactionStoreRepo{}

	s := &Server{
		listenAddress:    ":8081",
		deviceStore:      nil,
		transactionStore: mockTransactionStoreRepo,
	}

	// Record the response
	rec := httptest.NewRecorder()

	s.SignatureDevice(rec, req)

	// Validate the status code
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func Test_CreateSignatureDevice_get(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/api/v0/device", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	mockDeviceStoreRepo := &persistence.MockDeviceStoreRepo{}
	mockDeviceStoreRepo.On("GetAll").Return([]interface{}{
		map[string]interface{}{
			"id":                  "id1",
			"signature_algorithm": "RSA",
			"label":               "device1",
		},
		map[string]interface{}{
			"id":                  "id2",
			"signature_algorithm": "ECC",
			"label":               "device2",
		},
	})
	mockTransactionStoreRepo := &persistence.MockTransactionStoreRepo{}

	s := &Server{
		listenAddress:    ":8081",
		deviceStore:      mockDeviceStoreRepo,
		transactionStore: mockTransactionStoreRepo,
	}

	// Record the response
	rec := httptest.NewRecorder()

	s.SignatureDevice(rec, req)

	// Validate the status code
	assert.Equal(t, http.StatusOK, rec.Code)
}

func Test_SignTransaction_methodnotallowed(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/api/v0/transaction", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	mockDeviceStoreRepo := &persistence.MockDeviceStoreRepo{}
	mockTransactionStoreRepo := &persistence.MockTransactionStoreRepo{}

	s := &Server{
		listenAddress:    ":8081",
		deviceStore:      mockDeviceStoreRepo,
		transactionStore: mockTransactionStoreRepo,
	}

	// Record the response
	rec := httptest.NewRecorder()

	s.SignTransaction(rec, req)

	// Validate the status code
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
}

func Test_SignTransaction_BadData1(t *testing.T) {
	createRaw := map[string]interface{}{
		"signature_algorithm": "RSA",
		"label":               "label",
	}
	// Marshal the JSON data into a byte slice
	jsonData, err := json.Marshal(createRaw)
	if err != nil {
		t.Fatalf("Could not marshal JSON: %v", err)
	}
	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/api/v0/transaction", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	mockDeviceStoreRepo := &persistence.MockDeviceStoreRepo{}
	mockTransactionStoreRepo := &persistence.MockTransactionStoreRepo{}

	s := &Server{
		listenAddress:    ":8081",
		deviceStore:      mockDeviceStoreRepo,
		transactionStore: mockTransactionStoreRepo,
	}

	// Record the response
	rec := httptest.NewRecorder()

	s.SignTransaction(rec, req)

	// Validate the status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func Test_SignTransaction_BadData2(t *testing.T) {
	createRaw := map[string]interface{}{
		"device_id":         "",
		"data_to_be_signed": "data",
	}
	// Marshal the JSON data into a byte slice
	jsonData, err := json.Marshal(createRaw)
	if err != nil {
		t.Fatalf("Could not marshal JSON: %v", err)
	}
	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/api/v0/transaction", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	mockDeviceStoreRepo := &persistence.MockDeviceStoreRepo{}
	mockTransactionStoreRepo := &persistence.MockTransactionStoreRepo{}

	s := &Server{
		listenAddress:    ":8081",
		deviceStore:      mockDeviceStoreRepo,
		transactionStore: mockTransactionStoreRepo,
	}

	// Record the response
	rec := httptest.NewRecorder()

	s.SignTransaction(rec, req)

	// Validate the status code
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "device_id and data_to_be_signed must not be empty")
}

func Test_SignTransaction_Ok(t *testing.T) {
	createRaw := map[string]interface{}{
		"device_id":         "device_id",
		"data_to_be_signed": "data",
	}
	// Marshal the JSON data into a byte slice
	jsonData, err := json.Marshal(createRaw)
	if err != nil {
		t.Fatalf("Could not marshal JSON: %v", err)
	}
	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/api/v0/transaction", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	mockDeviceStoreRepo := &persistence.MockDeviceStoreRepo{}
	mockDeviceStoreRepo.On("GetById", "device_id").Return(&domain.SignatureDevice{
		Id:                 "device_id",
		SignatureAlgorithm: domain.ECDSA,
		KeyPair: &crypto.ECCKeyPair{
			Public:  &ecdsa.PublicKey{},
			Private: &ecdsa.PrivateKey{},
		},
		Label: "device1",
	})
	mockDeviceStoreRepo.On("IncrementCounter", "device_id").Return()
	mockTransactionStoreRepo := &persistence.MockTransactionStoreRepo{}
	mockTransactionStoreRepo.On("GetByDevice", "device_id").Return([]*domain.Transaction{})
	mockTransactionStoreRepo.On("Save", mock.Anything).Return()

	s := &Server{
		listenAddress:    ":8081",
		deviceStore:      mockDeviceStoreRepo,
		transactionStore: mockTransactionStoreRepo,
	}

	// Record the response
	rec := httptest.NewRecorder()

	s.SignTransaction(rec, req)

	// Validate the status code
	assert.Equal(t, http.StatusOK, rec.Code)
}
