package dnsblcheck

// This folder is a fork of github.com/Ajnasz/dnsbl-check
// This fork changes the command line tool to have
// exported functions that can be used as a dependency
// to look up DNSBL using a provided provider (e.g. zen.spamhaus.org)

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"challenge/dnsbl-check/dnsblprovider"
	"challenge/dnsbl-check/providerlist"
)

// LookupResult stores the query result with reason
type LookupResult struct {
	IsBlacklisted bool
	Address       string
	Reason        string
	Provider      dnsblprovider.DNSBLProvider
	Error         error
}

func Lookup(address string, provider dnsblprovider.DNSBLProvider) LookupResult {
	isListed, err := provider.IsBlacklisted(address)
	if err != nil {
		return LookupResult{
			Provider: provider,
			Address:  address,
			Error:    err,
		}
	}

	if isListed {
		desc, err := provider.GetReason(address)

		return LookupResult{
			Error:         err,
			Address:       address,
			IsBlacklisted: true,
			Provider:      provider,
			Reason:        desc,
		}
	}

	return LookupResult{
		Address:       address,
		IsBlacklisted: false,
		Provider:      provider,
	}
}

func getBlacklists(addresses []string, providers []dnsblprovider.DNSBLProvider) chan LookupResult {
	var wg sync.WaitGroup
	results := make(chan LookupResult)
	for _, address := range addresses {
		for _, provider := range providers {
			wg.Add(1)
			go func(address string, provider dnsblprovider.DNSBLProvider) {
				defer wg.Done()
				results <- Lookup(address, provider)
			}(address, provider)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func ProcessLookupResult(result LookupResult) string {
	if result.Error != nil {
		return fmt.Sprintf("error: %v, %v", result.Address, result.Error)
	}
	if result.IsBlacklisted {
		if result.Reason == "" {
			result.Reason = "unknown reason"
		}

		return fmt.Sprintf("%s: %v", result.Reason, result.Address)
	} else {
		return fmt.Sprintf("OK: %s", result.Address)
	}
}

func main() {
	var domainsFile = flag.String("p", "", "path to file which stores list of dnsbl checks, empty or - for stdin")
	var addressesParam = flag.String("i", "", "IP Address to check, separate by comma for a list")

	flag.Parse()
	list, err := providerlist.GetProvidersChan(*domainsFile)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading domains")
		os.Exit(1)
	}

	var providers []dnsblprovider.DNSBLProvider

	for item := range list {
		provider := dnsblprovider.GeneralProvider{
			URL: item,
		}

		providers = append(providers, provider)
	}

	addresses := providerlist.GetAddresses(*addressesParam)
	for result := range getBlacklists(addresses, providers) {
		ProcessLookupResult(result)
	}
}
