![header](https://capsule-render.vercel.app/api?type=waving&height=260&color=0:61D47F,100:6AC963&text=Cloudflare%20DDNS&reversal=false&section=header&textBg=false&fontColor=FFFFFF&fontAlign=40&fontAlignY=30&desc=Cloudflare%20Dynamic%20DNS%20Service&descAlign=40&descAlignY=50)

## Project Purpose
This project is a Go-based Dynamic DNS (DDNS) client for Cloudflare. It automatically updates DNS records in Cloudflare when the IP address changes, keeping your domain pointing to the correct IP.
 
## Table of Contents
See below for quicker & direct access to specific information
<br>
- <a href="#languages">Languages used</a>
- <a href="#inner-works">The inner works</a>
  - <a href="#current-implementation">Current implementation</a>
  - <a href="#basic-implementation-flow">Implementation flow</a>
  - <a href="#key-methods">Key methods</a>
- <a href="#get-started">Get started</a>
  - <a href="#run">How to run</a>
  - <a href="#testing">Testing</a>
- <a href="#roadmap">Roadmap</a>
- <a href="#external-packages">External packages</a>
- <a href="#contact-information">Contact info</a>

## Languages
Languages used in this project
<div align="left">
    <img src="https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&style=for-the-badge" height="30" alt="go logo" />
</div>

## Inner Works
See the complete in's and out's along with current implementation and functionality overview.

### Current Implementation
- Fetch the public IP address of the machine using reliable external service.
- Use Cloudflare's API to update DNS records.
- Allow for multiple domains and records to be updated simultaneously.
- Supports `A` record types.
- Security via API token and custom headers.

### Extra Features
- Supports automatic IP detection and DNS update.
- Provides logging of update operations.
- Requires the following configuration settings:
  - **Cloudflare Credentials**: Your Cloudflare API token with permission to modify DNS records and your Cloudflare account email.
  - **Domains & Subdomains/Records**: List of domains and records to be updated.

### Basic Implementation Flow
![Func Flow](docs/images/implementation-flow.png)

### Key Methods
- `app.Start()`
  - Initializes the application and starts the periodic IP check and DNS update process.
- `cloudflare.UpdateDNS()`
  - Fetches the current public IP and updates the Cloudflare DNS record.
- `api.GetIP()`
  - Fetches the current public IP address using an external service.

## Get Started
Clone the repository:
```bash
git clone https://github.com/insxmnia/cloudflare-ddns.git
```
Inside the repository, install the Go modules:
```bash
go mod tidy
```
Create a `config.yaml` file in the `config` directory using the provided .example.yaml file as a template:
```yaml
# Cloudflare domains
domains: ["example.com"]

# Subdomains/records to change, match domain name without extension
sub-domains: 
  example: ["www", "api"]

# Cloudflare credentials
credentials: 
  email: "info@example.com"
  key: "cloudflare-api-key"
```

#### Run
Run the following command to start the service, entry point is located in `cmd/main.go`:
```bash
go run cmd/main.go
```

#### Testing
To test the functionality:
```bash
go test ./...
```
Display the test coverage of each file:
```bash
go test ./... -cover
```
Check for any race conditions:
```bash
go test ./... -race
```

## Roadmap
- [x] Fetch current public IP address.
- [x] Update `A` DNS records in Cloudflare.
- [x] Logging and reporting DNS updates.
- [ ] Control interval through configuration.
- [ ] Allow wildcard option for subdomains

## External Packages
<div align="left">
    <a href="https://github.com/spf13/viper"><img src="https://img.shields.io/badge/Viper-000?style=for-the-badge&logo=Go&logoColor=white&color=61D47F" height="30" alt="Viper package"></a>
    <a href="https://github.com/stretchr/testify"><img src="https://img.shields.io/badge/Testify-000?style=for-the-badge&logo=Go&logoColor=white&color=00ADD8" height="30" alt="Testify package"></a>
</div>


