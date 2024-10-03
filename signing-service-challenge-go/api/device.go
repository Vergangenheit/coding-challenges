package api

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	uuid "github.com/google/uuid"
)

// TODO: REST endpoints ...
func (s *Server) CreateSignatureDevice(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		// decode body
		signDevice := &domain.SignatureDevice{}
		if err := json.NewDecoder(request.Body).Decode(signDevice); err != nil {
			WriteErrorResponse(response, http.StatusBadRequest, []string{
				http.StatusText(http.StatusBadRequest),
			})
			return
		}
		// assign id
		if signDevice.Id == "" {
			signDevice.Id = uuid.New().String()
		}
		// generate key pair
		switch signDevice.SignatureAlgorithm {
		case "RSA":
			gen := crypto.RSAGenerator{}
			key, err := gen.Generate()
			if err != nil {
				WriteErrorResponse(response, http.StatusInternalServerError, []string{
					http.StatusText(http.StatusInternalServerError),
				})
				return
			}
			signDevice.KeyPair = key
		case "ECC":
			gen := crypto.ECCGenerator{}
			key, err := gen.Generate()
			if err != nil {
				WriteErrorResponse(response, http.StatusInternalServerError, []string{
					http.StatusText(http.StatusInternalServerError),
				})
				return
			}
			signDevice.KeyPair = key
		default:
			WriteErrorResponse(response, http.StatusBadRequest, []string{
				"signature_algorithm must be RSA or ECC",
			})
			return
		}
		// persist signDevice
		if s.deviceStore != nil {
			s.deviceStore.Save(string(signDevice.Id), signDevice)
		} else {
			WriteInternalError(response)
			return
		}
		// TODO write response
		WriteAPIResponse(response, http.StatusCreated, signDevice)
	case http.MethodGet:
		// get all devices
		if s.deviceStore != nil {
			devices := s.deviceStore.GetAll()
			WriteAPIResponse(response, http.StatusOK, devices)
		} else {
			WriteInternalError(response)
			return
		}
	default:
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}
}

func (s *Server) SignTransaction(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}
	// decode body
	transactionToBeSigned := &domain.Transaction{}
	if err := json.NewDecoder(request.Body).Decode(transactionToBeSigned); err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			http.StatusText(http.StatusBadRequest),
		})
		return
	}
	if transactionToBeSigned.DeviceId == "" || transactionToBeSigned.Data == "" {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			"device_id and data_to_be_signed must not be empty",
		})
		return
	}
	// get device
	signDevice := s.deviceStore.GetById(transactionToBeSigned.DeviceId)
	if signDevice == nil {
		WriteErrorResponse(response, http.StatusNotFound, []string{
			"device not found",
		})
		return
	}
	// get algo from device
	algo := signDevice.SignatureAlgorithm

	var resp *domain.SignatureResponse
	var err error
	// sign data
	switch algo {
	case "RSA":
		rsaSigner := crypto.NewRSASigner()
		resp, err = s.signData(transactionToBeSigned, signDevice, rsaSigner)
		if err != nil {
			WriteErrorResponse(response, http.StatusInternalServerError, []string{
				http.StatusText(http.StatusInternalServerError),
			})
			return
		}
		// increment counter
		signDevice.IncrementCounter()

	case "ECC":
		eccSigner := crypto.NewECDSASigner()
		resp, err = s.signData(transactionToBeSigned, signDevice, eccSigner)
		if err != nil {
			WriteErrorResponse(response, http.StatusInternalServerError, []string{
				err.Error(),
			})
			return
		}
		// increment counter
		signDevice.IncrementCounter()

	default:
		WriteErrorResponse(response, http.StatusInternalServerError, []string{
			"cannot sign if signature_algorithm not RSA or ECC",
		})
	}
	// response
	WriteAPIResponse(response, http.StatusOK, resp)
}

func (s *Server) signData(transaction *domain.Transaction, signDevice *domain.SignatureDevice, signer crypto.Signer) (*domain.SignatureResponse, error) {
	var lastSignatureEncoded string
	var err error
	// check the last signature on device if any
	deviceTransactions := s.transactionStore.GetByDevice(transaction.DeviceId)
	if len(deviceTransactions) == 0 {
		// use encoded device id
		lastSignatureEncoded = encodeString(transaction.DeviceId)
	} else {
		// sort by signedAt
		// Sort by SignedAt in descending order (latest first)
		sort.Slice(deviceTransactions, func(i, j int) bool {
			return deviceTransactions[i].SignedAt.After(deviceTransactions[j].SignedAt)
		})
		latestTransaction := deviceTransactions[0]
		lastSignatureEncoded, err = encodeStruct(latestTransaction)
		if err != nil {
			return nil, err
		}
	}
	signedData, err := signer.Sign(strconv.Itoa(signDevice.Counter()), transaction.Data, lastSignatureEncoded)
	if err != nil {
		return nil, err
	}
	// persist transaction
	transaction.SignedAt = time.Now()
	s.transactionStore.Save(transaction)
	// response
	resp := &domain.SignatureResponse{
		Signature:  transaction,
		SignedData: signedData,
	}
	return resp, nil

}

func encodeString(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func encodeStruct(data interface{}) (string, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	encodedString := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return encodedString, nil
}
