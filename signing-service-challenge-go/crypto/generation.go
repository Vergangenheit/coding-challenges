package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
)

type Generator interface {
	Generate() (KeyPair, error)
}

// RSAGenerator generates a RSA key pair.
type RSAGenerator struct{}

func NewRSAGenerator() Generator {
	return &RSAGenerator{}
}

// Generate generates a new RSAKeyPair.
func (g *RSAGenerator) Generate() (KeyPair, error) {
	// Security has been ignored for the sake of simplicity.
	key, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		return nil, err
	}

	return &RSAKeyPair{
		Public:  &key.PublicKey,
		Private: key,
	}, nil
}

// ECCGenerator generates an ECC key pair.
type ECCGenerator struct{}

func NewECCGenerator() Generator {
	return &ECCGenerator{}
}

// Generate generates a new ECCKeyPair.
func (g *ECCGenerator) Generate() (KeyPair, error) {
	// Security has been ignored for the sake of simplicity.
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &ECCKeyPair{
		Public:  &key.PublicKey,
		Private: key,
	}, nil
}
