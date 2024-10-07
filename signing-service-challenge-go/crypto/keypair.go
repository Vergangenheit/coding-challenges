package crypto

// KeyPair is a generic interface that will represent different key pair types (RSA, ECC, etc.)
type KeyPair interface {
	PublicKey() interface{}
	PrivateKey() interface{}
}
