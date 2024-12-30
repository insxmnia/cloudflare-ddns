package models

/*
Here you define the structs and interfaces for the Cloudflare API package
Prefix the definitions with the ones below for clearer identification.

'M' prefix stands for 'Model' for struct definitions.
'I' prefix stands for 'Interface' for interface definitions.

If you have regions enabled in the IDE use //region to easier separate different models
*/

// region Cloudflare Package Specific

type MCFError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// endregion
