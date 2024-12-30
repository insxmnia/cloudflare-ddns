package api

import (
	"cloudflare-ddns/pkg/cloudflare/internal/constants"
	"cloudflare-ddns/pkg/cloudflare/internal/models"
	"cloudflare-ddns/pkg/net"
	"encoding/json"
	"fmt"
	"net/http"
)

type Api struct {
	Credentials *models.MCredentials
}

func (api *Api) GetDomainID(domain string) (string, error) {
	var RequestResponse models.MDomainIDResult

	url := fmt.Sprintf(constants.CF_DOMAIN_ID_URL, domain)
	request, err := net.CreateRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	AddHeaders(request, api.Credentials)
	if response := net.ExecuteRequest(request); response.Error != nil || response.Status != http.StatusOK {
		return "", fmt.Errorf("error fetching domain ID: %v", response.Error)
	} else {
		err := json.Unmarshal(response.Body, &RequestResponse)
		if err != nil {
			return "", err
		}
	}

	return RequestResponse.Result[0].ID, nil
}

func (api *Api) GetAllRecords(domain_id string) (models.MDNSRecords, error) {
	var RequestResponse models.MDNSRecords

	url := fmt.Sprintf(constants.CF_DNS_ALL_URL, domain_id)
	request, err := net.CreateRequest("GET", url, nil)
	if err != nil {
		return models.MDNSRecords{}, err
	}
	AddHeaders(request, api.Credentials)
	if response := net.ExecuteRequest(request); response.Error != nil || response.Status != http.StatusOK {
		return models.MDNSRecords{}, fmt.Errorf("error fetching domain records: %v", response.Error)
	} else {
		err := json.Unmarshal(response.Body, &RequestResponse)
		if err != nil {
			return models.MDNSRecords{}, err
		}
	}

	return RequestResponse, nil
}

func (api *Api) UpdateRecord(record_id string, domain_id string, new_ip string) error {
	url := fmt.Sprintf(constants.CF_DNS_UPDATE_URL, domain_id, record_id)
	body, err := json.Marshal(map[string]interface{}{
		"content": new_ip,
	})
	if err != nil {
		return err
	}

	request, err := net.CreateRequest("PATCH", url, body)
	if err != nil {
		return err
	}
	AddHeaders(request, api.Credentials)
	if response := net.ExecuteRequest(request); response.Error != nil {
		return fmt.Errorf("error updating DNS record: %v", response.Error)
	} else if response.Status != http.StatusOK {
		return fmt.Errorf("error updating DNS record: %s", response.Body)
	}

	return nil
}
