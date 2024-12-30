package constants

const (
	// Cloudflare token verification endpoint
	CF_TOKEN_VERIFICATION_URL = "https://api.cloudflare.com/client/v4/user/tokens/verify"
	// Cloudflare zone ID URL
	CF_DOMAIN_ID_URL = "https://api.cloudflare.com/client/v4/zones?name=%s"
	// Cloudflare DNS record URL
	CF_DNS_RECORD_URL = "https://api.cloudflare.com/client/v4/zones/%s/dns_records?name=%s"
	// Cloudflare DNS all records URL for type A
	CF_DNS_ALL_URL = "https://api.cloudflare.com/client/v4/zones/%s/dns_records?type=A"
	// Cloudflare DNS record update URL
	CF_DNS_UPDATE_URL = "https://api.cloudflare.com/client/v4/zones/%s/dns_records/%s"
)
