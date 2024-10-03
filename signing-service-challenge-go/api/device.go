package api

import (
	"encoding/json"
	"net/http"

	"github.com/fiskaly/coding-challenges/signing-service-challenge/crypto"
	"github.com/fiskaly/coding-challenges/signing-service-challenge/domain"
	uuid "github.com/google/uuid"
)

// TODO: REST endpoints ...
func (s *Server) CreateSignatureDevice(response http.ResponseWriter, request *http.Request) {
	// if request.Method != http.MethodPost {
	// 	WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
	// 		http.StatusText(http.StatusMethodNotAllowed),
	// 	})
	// 	return
	// }
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
		// TODO validate signDevice
		if signDevice.SignatureAlgorithm == "" {
			WriteErrorResponse(response, http.StatusBadRequest, []string{
				"signature_algorithm is required",
			})
			return
		}
		if signDevice.SignatureAlgorithm != "RSA" && signDevice.SignatureAlgorithm != "ECC" {
			WriteErrorResponse(response, http.StatusBadRequest, []string{
				"signature_algorithm must be RSA or ECC",
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
		}
		// persist signDevice
		if s.persistenceLayer != nil {
			s.persistenceLayer.Save(string(signDevice.Id), signDevice)
		} else {
			WriteInternalError(response)
			return
		}
	case http.MethodGet:
		// get all devices
		if s.persistenceLayer != nil {
			devices := s.persistenceLayer.GetAll()
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
}
