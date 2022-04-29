package spamhaus

import (
	"challenge/dnsbl-check-fork"
	"challenge/dnsbl-check-fork/dnsblprovider"
)

const spamhausRBL = "zen.spamhaus.org"

// Query queries zen.spamhaus.org to see if the given ip address is on the RBL
func Query(addr string) (string, error) {
	spamhausProvider := dnsblprovider.GeneralProvider{
		URL: spamhausRBL,
	}
	lr := dnsblcheck.Lookup(addr, spamhausProvider)

	return dnsblcheck.ProcessLookupResult(lr), nil
}
