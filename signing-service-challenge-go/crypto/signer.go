package crypto

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	Sign(dataToBeSigned []byte) ([]byte, error)
}

type RSASigner struct{}

func NewRSASigner() Signer {
	return &RSASigner{}
}

func (s *RSASigner) Sign(dataToBeSigned []byte) ([]byte, error) {

	return dataToBeSigned, nil
}

type ECDSASigner struct{}

func NewECDSASigner() Signer {
	return &ECDSASigner{}
}

func (s *ECDSASigner) Sign(dataToBeSigned []byte) ([]byte, error) {

	return dataToBeSigned, nil
}
