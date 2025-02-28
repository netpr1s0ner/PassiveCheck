package runner

import (
	"github.com/projectdiscovery/utils/auth/pdcp"
	updateutils "github.com/projectdiscovery/utils/update"
)

// Version is the current Version of httpx
const Version = `v1.6.8`

// showBanner is used to show the banner to the user
func showBanner() {
	return
}

// GetUpdateCallback returns a callback function that updates httpx
func GetUpdateCallback() func() {
	return func() {
		showBanner()
		updateutils.GetUpdateToolCallback("httpx", Version)()
	}
}

// AuthWithPDCP is used to authenticate with PDCP
func AuthWithPDCP() {
	showBanner()
	pdcp.CheckNValidateCredentials("httpx")
}
