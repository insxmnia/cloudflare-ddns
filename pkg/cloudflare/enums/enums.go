package enums

/*
These enums are exposed for global usage and not limited to this package.
Prefix the enums with CF_ to indicate that they are Cloudflare-specific.
Use the below extended prefixes for specific use cases:

'CF_ERR_' this is used for the error codes returned by the Cloudflare package.


*/

const (
	CF_ERR_UNKNOWN_CODE           = 1000
	CF_ERR_UNKNOWN_MSG            = "Unknown error occurred"
	CF_ERR_ALREADY_INIT_CODE      = 1001
	CF_ERR_ALREADY_INIT_MSG       = "Cloudflare client has already been initialized"
	CF_ERR_INVALID_TOKEN_CODE     = 1002
	CF_ERR_INVALID_TOKEN_MSG      = "Invalid or expired Cloudflare API token"
	CF_ERR_INVALID_CREDS_CODE     = 1003
	CF_ERR_INVALID_CREDS_MSG      = "Invalid Cloudflare credentials, either the API token or the email address are incorrect."
	CF_ERR_NOT_INIT_CODE          = 1004
	CF_ERR_NOT_INIT_MSG           = "Cloudflare client has not been initialized"
	CF_ERR_EMPTY_CREDENTIALS_CODE = 1005
	CF_ERR_EMPTY_CREDENTIALS_MSG  = "No Cloudflare API token or email address provided"
	CF_ERR_RATE_LIMIT_CODE        = 429
	CF_ERR_RATE_LIMIT_MSG         = "Rate limit exceeded"
	CF_ERR_SERVER_ERROR_CODE      = 500
	CF_ERR_SERVER_ERROR_MSG       = "Server error occurred"
)
