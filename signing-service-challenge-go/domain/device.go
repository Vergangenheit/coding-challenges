package domain

type SignatureAlgorithm string

const (
	RSA   SignatureAlgorithm = "RSA"
	ECDSA SignatureAlgorithm = "ECC"
)

// TODO: signature device domain model ...
type SignatureDevice struct {
	Id                 string             `json:"id"`
	SignatureAlgorithm SignatureAlgorithm `json:"signature_algorithm"`
	KeyPair            KeyPair            `json:"key_pair"`
	Label              string             `json:"label"`
	signatureCounter   int
}

// KeyPair interface
type KeyPair interface {
	PublicKey() interface{}
	PrivateKey() interface{}
}
