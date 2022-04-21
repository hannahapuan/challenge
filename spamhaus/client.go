package spamhaus

import (
	dc "challenge/dnsbl-check"
	dnsblcheck "challenge/dnsbl-check"
	"challenge/dnsbl-check/dnsblprovider"
)

const spamhausRBL = "zen.spamhaus.org"

// Query queries zen.spamhaus.org to see if the given ip address is on the RBL
func Query(addr string) (string, error) {
	spamhausProvider := dnsblprovider.GeneralProvider{
		URL: spamhausRBL,
	}
	lr := dc.Lookup(addr, spamhausProvider)

	return dnsblcheck.ProcessLookupResult(lr), nil
}
