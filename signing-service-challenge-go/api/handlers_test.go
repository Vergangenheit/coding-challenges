package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/persistence"
	"github.com/stretchr/testify/assert"
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
