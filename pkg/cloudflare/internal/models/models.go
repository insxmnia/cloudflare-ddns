package models

/*
Here you define the structs and interfaces for the Cloudflare API package
Prefix the definitions with the ones below for clearer identification.

'M' prefix stands for 'Model' for struct definitions.
'I' prefix stands for 'Interface' for interface definitions.

If you have regions enabled in the IDE use //region to easier separate different models
*/

// region Expected Payloads

type MDomainIDResult struct {
	Result []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	} `json:"result"`
}

type MDNSRecords struct {
	Result []MDNSRecord `json:"result"`
}

// Includes checks if the MDNSRecords contains a record with the given name and returns ok (found bool) and the record.
func (m MDNSRecords) Includes(record_name string) (bool, *MDNSRecord) {
	for _, record := range m.Result {
		if record.Name == record_name {
			return true, &record
		}
	}
	return false, nil
}

type MDNSRecord struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	IP       string `json:"content"`
	TTL      int    `json:"ttl"`
	ZoneID   string `json:"zone_id"`
	ZoneName string `json:"zone_name"`
	Status   string `json:"status"`
}

// endregion

// region Others

type MCredentials struct {
	Email  string
	APIKey string
}

// endregion
