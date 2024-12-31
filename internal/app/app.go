package app

import (
	"cloudflare-ddns/api"
	"cloudflare-ddns/internal/pkg/backgrounder"
	cf "cloudflare-ddns/pkg/cloudflare"
	"cloudflare-ddns/pkg/slogger"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

func Start() {
	// Get domains
	domains := viper.GetStringSlice("domains")
	if len(domains) == 0 {
		slogger.Fatal("Cloudflare DDNS", "no-domains", "No domains provided in the configuration file.")
	}
	// Get subdomains
	subdomains := viper.GetStringMapStringSlice("sub-domains")
	if len(subdomains) == 0 {
		slogger.Fatal("Cloudflare DDNS", "no-subdomains", "No subdomains provided in the configuration file.")
	}

	// Initialize the Cloudflare client with provided credentials.
	cloudflare := cf.Cloudflare{}
	if err := cloudflare.Initialize(viper.GetString("credentials.email"), viper.GetString("credentials.key")); err.Message != "" {
		slogger.Fatal("Cloudflare", "cloudflare-error", err.Message, "error-code", err.Code)
	}

	// Start a new background worker for each domain and initilize a waitGroup
	wg := sync.WaitGroup{}
	workers := make(map[string]*backgrounder.Worker)
	for _, domain := range domains {
		slogger.Info("Cloudflare Worker", "domain", domain, "state", "starting")
		clean_domain := strings.Split(domain, ".")[0]
		worker := backgrounder.NewBackgroundWorker(&wg)
		worker.RunWithInterval(clean_domain, func() {
			slogger.Info("Cloudflare Worker", "domain", domain, "state", "running")
			// This function is for every domains, it will loop through all subdomains and update their DNS records.
			// As well as freshly grab the host IP address and domain's DNS records, this makes sure to always have the most up-to-date data
			public_ip := api.GetPublicIP()

			// get the domain ID
			domain_id, err := cloudflare.GetDomainID(domain)
			if err.Message != "" {
				slogger.Error("cloudflare", "error", err.Message, "error-code", err.Code)
				return
			}
			slogger.Info("Cloudflare DDNS", "domain", domain, "id", domain_id)
			// get the domain's DNS records and compare the subdomains to the records received from Cloudflare
			dns_records, err := cloudflare.GetDNSRecords(domain_id)
			if err.Message != "" {
				slogger.Error("cloudflare", "error", err.Message, "error-code", err.Code)
				return
			}
			// Update the root DNS record (domain name)
			if ok, record := dns_records.Includes(domain); ok {
				slogger.Info("Cloudflare DDNS", "subdomain", record.Name, "ip", record.IP)
				if record.IP == public_ip {
					return
				}
				slogger.Info("Cloudflare DDNS", "updating", record.Name, "to", public_ip)
				err := cloudflare.UpdateRecord(domain_id, record.ID, public_ip)
				if err.Message != "" {
					slogger.Error("cloudflare", "error", err.Message, "error-code", err.Code)
				}
				slogger.Info("Cloudflare DDNS", "updated", record.Name, "to", public_ip)
			} else {
				slogger.Info("Cloudflare DDNS", "subdomain", domain, "status", "not found")
			}
			// Update the subdomains
			for _, subdomain := range subdomains[clean_domain] {
				slogger.Info("Cloudflare DDNS", "subdomain", subdomain, "status", "checking")
				formatted_subdomain := strings.Join([]string{subdomain, domain}, ".")
				if ok, record := dns_records.Includes(formatted_subdomain); ok {
					slogger.Info("Cloudflare DDNS", "subdomain", record.Name, "ip", record.IP)
					if record.IP == public_ip {
						continue
					}
					slogger.Info("Cloudflare DDNS", "updating", record.Name, "to", public_ip)
					err := cloudflare.UpdateRecord(domain_id, record.ID, public_ip)
					if err.Message != "" {
						slogger.Error("cloudflare", "error", err.Message, "error-code", err.Code)
					}
					slogger.Info("Cloudflare DDNS", "updated", record.Name, "to", public_ip)
				} else {
					slogger.Info("Cloudflare DDNS", "subdomain", subdomain, "status", "not found")
				}
			}

		}, 60000) // 1 Minute interval
		workers[domain] = worker
	}
	wg.Wait()
	slogger.Info("Cloudflare DDNS", "status", "completed")
}
