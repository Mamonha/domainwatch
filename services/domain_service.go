package services

import (
	"gin-quickstart/helpers"

	"github.com/likexian/whois"
)

func ValidateDomain(domain string) string {
	if !helpers.IsValidDomain(domain) {
		return "Invalid domain format."
	}

	if !helpers.IsDomainAccessible(domain) {
		return "Domain is not accessible."
	}

	return "Yes, the domain " + domain + " is valid and accessible."
}

func Whois(domain string) string {

	if !helpers.IsValidDomain(domain) {
		return "Invalid domain format."
	}

	if !helpers.IsDomainAccessible(domain) {
		return "Domain is not accessible."
	}

	result, err := whois.Whois(domain)
	if err != nil {
		return "Error retrieving WHOIS data: " + err.Error()
	}
	return result
}
