package cloudflare

import (
	"cloudflare-ddns/pkg/cloudflare/enums"
	"cloudflare-ddns/pkg/cloudflare/internal/api"
	im "cloudflare-ddns/pkg/cloudflare/internal/models"
	m "cloudflare-ddns/pkg/cloudflare/models"
)

var (
	_cloudflare *Cloudflare
)

func GetInstance() (*Cloudflare, m.MCFError) {
	if _cloudflare == nil {
		return nil, m.MCFError{
			Code:    enums.CF_ERR_NOT_INIT_CODE,
			Message: enums.CF_ERR_NOT_INIT_MSG,
		}
	}
	return _cloudflare, m.MCFError{}
}

type Cloudflare struct {
	initialized bool
	credentials *im.MCredentials
	api         *api.Api
}

func (cf *Cloudflare) Initialize(email string, api_key string) m.MCFError {
	if cf.initialized && _cloudflare != nil {
		return m.MCFError{Code: enums.CF_ERR_ALREADY_INIT_CODE, Message: enums.CF_ERR_ALREADY_INIT_MSG}
	}

	cf.credentials = &im.MCredentials{
		Email:  email,
		APIKey: api_key,
	}

	if !api.ValidateCredentials(cf.credentials) {
		return m.MCFError{Code: enums.CF_ERR_INVALID_CREDS_CODE, Message: enums.CF_ERR_INVALID_CREDS_MSG}
	}

	_cloudflare = cf
	cf.initialized = true
	cf.api = &api.Api{Credentials: cf.credentials}

	return m.MCFError{}
}

func (cf *Cloudflare) GetDomainID(domain string) (string, m.MCFError) {
	if !cf.initialized {
		return "", m.MCFError{Code: enums.CF_ERR_NOT_INIT_CODE, Message: enums.CF_ERR_NOT_INIT_MSG}
	}

	id, err := cf.api.GetDomainID(domain)
	if err != nil {
		return "", m.MCFError{Code: enums.CF_ERR_SERVER_ERROR_CODE, Message: err.Error()}
	}
	return id, m.MCFError{}
}
func (cf *Cloudflare) GetDNSRecords(domain_id string) (im.MDNSRecords, m.MCFError) {
	if !cf.initialized {
		return im.MDNSRecords{}, m.MCFError{Code: enums.CF_ERR_NOT_INIT_CODE, Message: enums.CF_ERR_NOT_INIT_MSG}
	}

	records, err := cf.api.GetAllRecords(domain_id)
	if err != nil {
		return im.MDNSRecords{}, m.MCFError{Code: enums.CF_ERR_SERVER_ERROR_CODE, Message: err.Error()}
	}
	return records, m.MCFError{}
}

func (cf *Cloudflare) UpdateRecord(domain_id string, record_id string, record_ip string) m.MCFError {
	if !cf.initialized {
		return m.MCFError{Code: enums.CF_ERR_NOT_INIT_CODE, Message: enums.CF_ERR_NOT_INIT_MSG}
	}

	err := cf.api.UpdateRecord(record_id, domain_id, record_ip)
	if err != nil {
		return m.MCFError{Code: enums.CF_ERR_SERVER_ERROR_CODE, Message: err.Error()}
	}
	return m.MCFError{}
}
