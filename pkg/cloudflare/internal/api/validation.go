package api

import (
	"cloudflare-ddns/pkg/cloudflare/internal/constants"
	"cloudflare-ddns/pkg/cloudflare/internal/models"
	"cloudflare-ddns/pkg/net"
	"net/http"
)

func ValidateCredentials(credentials *models.MCredentials) bool {
	if credentials.APIKey == "" || credentials.Email == "" {
		return false
	}

	request, err := net.CreateRequest("GET", constants.CF_TOKEN_VERIFICATION_URL, nil)
	if err != nil {
		return false
	}

	AddHeaders(request, credentials)

	response := net.ExecuteRequest(request)

	return response.Status == http.StatusOK
}

func AddHeaders(request *http.Request, credentials *models.MCredentials) {
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Auth-Email", credentials.Email)
	request.Header.Add("Authorization", "Bearer "+credentials.APIKey)
}
