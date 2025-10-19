package helpers

import (
	"net/http"
	"regexp"
	"strings"
	"time"
)

func IsValidDomain(domain string) bool {
	if len(domain) == 0 {
		return false
	}
	domainRegex := `^(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(domainRegex, domain)
	return matched
}

func HasSubdomain(domain string) bool {
	parts := strings.Split(domain, ".")
	return len(parts) > 2
}

func IsDomainAccessible(domain string) bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := client.Get("http://" + domain)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	return response.StatusCode >= 200 && response.StatusCode < 400
}
