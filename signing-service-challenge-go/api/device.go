package api

import (
	"encoding/json"
	"net/http"
)

type SignatureDevice struct{}

// TODO: REST endpoints ...
func (s *Server) CreateSignatureDevice(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}
	// decode body
	signDevice := &SignatureDevice{}
	if err := json.NewDecoder(request.Body).Decode(signDevice); err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{
			http.StatusText(http.StatusBadRequest),
		})
		return
	}
	// TODO validate signDevice
}

func (s *Server) SignTransaction(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		WriteErrorResponse(response, http.StatusMethodNotAllowed, []string{
			http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}
}
