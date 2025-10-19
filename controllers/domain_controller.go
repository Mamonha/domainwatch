package controllers

import (
	"encoding/json"
	"gin-quickstart/services"
	"net/http"
	"strings"
)

func Whois(domain string, w http.ResponseWriter) {
	result := services.Whois(domain)

	parsedWhois := parseWhoisResponse(result)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parsedWhois)
}

func parseWhoisResponse(rawWhois string) map[string]string {
	parsed := make(map[string]string)
	lines := strings.Split(rawWhois, "\n")

	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			key := strings.TrimSpace(parts[0])
			if len(parts) > 1 {
				value := strings.TrimSpace(parts[1])
				parsed[key] = value
			}
		}
	}

	return parsed
}

func Validate(domain string, w http.ResponseWriter) {
	responseMessage := services.ValidateDomain(domain)
	response := map[string]string{
		"message": responseMessage,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CheckTLS(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Domain string `json:"domain"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	domain := body.Domain
	tlsResults, err := services.CheckTLSVersion(domain)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error checking TLS versions: " + err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tlsResults)
}
