package controllers

import (
	"encoding/json"
	"gin-quickstart/services"
	"net/http"
)

func Whois(domain string) {
	result := services.Whois(domain)
	response := map[string]string{
		"whois": result,
	}
	jsonResponse, _ := json.Marshal(response)
	println(string(jsonResponse))

}

func Validate(domain string, w http.ResponseWriter) {
	responseMessage := services.ValidateDomain(domain)
	response := map[string]string{
		"message": responseMessage,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
