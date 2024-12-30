package api

import (
	"cloudflare-ddns/pkg/slogger"

	"github.com/pion/stun"
)

// Get public IP address, this method uses the STUN protocol, which is more reliable compared to using an API call to an external service which can go down.
func GetPublicIP() string {
	ipAddress := ""
	server := "stun.l.google.com:19302"

	// New connection
	conn, err := stun.Dial("udp", server)
	if err != nil {
		slogger.Fatal("STUN error", "error", err)
	}

	// Send a simple request to get the public IP
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)
	var addr stun.XORMappedAddress
	err = conn.Do(message, func(res stun.Event) {
		if res.Error != nil {
			slogger.Fatal("STUN error", "error", res.Error)
		}

		if err := addr.GetFrom(res.Message); err != nil {
			slogger.Fatal("STUN error", "error", err)
		}

		ipAddress = addr.IP.String()
	})
	if err != nil {
		slogger.Fatal("STUN error", "error", err)
	}
	return ipAddress
}
