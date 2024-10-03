package domain

import "time"

type SignatureAlgorithm string

const (
	RSA   SignatureAlgorithm = "RSA"
	ECDSA SignatureAlgorithm = "ECC"
)

// KeyPair interface
type KeyPair interface {
	PublicKey() interface{}
	PrivateKey() interface{}
}

// TODO: signature device domain model ...
type SignatureDevice struct {
	Id                 string             `json:"id"`
	SignatureAlgorithm SignatureAlgorithm `json:"signature_algorithm"`
	KeyPair            KeyPair            `json:"key_pair"`
	Label              string             `json:"label"`
	signatureCounter   int
}

func (d *SignatureDevice) IncrementCounter() {
	d.signatureCounter++
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
