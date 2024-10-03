package crypto

import "fmt"

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	// Sign(dataToBeSigned []byte) ([]byte, error)
	Sign(deviceCounter, dataToBeSigned, lastSignatureEncoded string) (string, error)
}

// TODO: implement RSA and ECDSA signing ...
type RSASigner struct{}

func NewRSASigner() *RSASigner {
	return &RSASigner{}
}

func (s *RSASigner) Sign(deviceCounter, dataToBeSigned, lastSignatureEncoded string) (string, error) {
	if deviceCounter == "" || dataToBeSigned == "" || lastSignatureEncoded == "" {
		return "", fmt.Errorf("signature data must not be empty")
	}
	return fmt.Sprintf("%s_%s_%s", deviceCounter, dataToBeSigned, lastSignatureEncoded), nil
}

type ECDSASigner struct{}

func NewECDSASigner() *ECDSASigner {
	return &ECDSASigner{}
}

func (s *ECDSASigner) Sign(deviceCounter, dataToBeSigned, lastSignatureEncoded string) (string, error) {
	if deviceCounter == "" || dataToBeSigned == "" || lastSignatureEncoded == "" {
		return "", fmt.Errorf("signature data must not be empty")
	}
	return fmt.Sprintf("%s_%s_%s", deviceCounter, dataToBeSigned, lastSignatureEncoded), nil
}
