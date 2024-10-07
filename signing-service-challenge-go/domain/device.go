package domain

import (
	"fmt"
	"sync"
	"time"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
)

type SignatureAlgorithm string

const (
	RSA   SignatureAlgorithm = "RSA"
	ECDSA SignatureAlgorithm = "ECC"
)

type DeviceInterface interface {
	GenerateKeyPair() error
	IncrementCounter()
	Counter() int
}

// signature device domain model ...
type SignatureDevice struct {
	Id                 string             `json:"id"`
	SignatureAlgorithm SignatureAlgorithm `json:"signature_algorithm"`
	KeyPair            crypto.KeyPair     `json:"key_pair"`
	Label              string             `json:"label"`
	signatureCounter   int
	mu                 *sync.Mutex
}

func NewSignatureDevice(algorithm SignatureAlgorithm, label string) (*SignatureDevice, error) {
	dev := &SignatureDevice{
		SignatureAlgorithm: algorithm,
		Label:              label,
	}
	err := dev.GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	dev.mu = &sync.Mutex{}
	return dev, nil
}

func (d *SignatureDevice) GenerateKeyPair() error {
	switch d.SignatureAlgorithm {
	case RSA:
		gen := crypto.NewRSAGenerator()
		key, err := gen.Generate()
		if err != nil {
			return err
		}
		d.KeyPair = key
	case ECDSA:
		gen := crypto.NewECCGenerator()
		key, err := gen.Generate()
		if err != nil {
			return err
		}
		d.KeyPair = key
	default:
		return fmt.Errorf("signature_algorithm must be RSA or ECC")
	}
	return nil
}

func (d *SignatureDevice) IncrementCounter() {
	d.mu.Lock()
	d.signatureCounter++
	d.mu.Unlock()
}

func (d *SignatureDevice) Counter() int {
	return d.signatureCounter
}

type Transaction struct {
	DeviceId string    `json:"device_id"`
	Data     string    `json:"data_to_be_signed"`
	SignedAt time.Time `json:"signed_at"`
}

type SignatureResponse struct {
	Signature  *Transaction `json:"transaction"`
	SignedData string       `json:"signed_data"`
}
