package services

import (
	"os/exec"
	"strings"
)

func CheckTLSVersion(domain string) (map[string]bool, error) {
	result := make(map[string]bool)

	if strings.Contains(domain, "/") {
		for _, tlsVersion := range []string{"--tlsv1.0", "--tlsv1.1", "--tlsv1.2", "--tlsv1.3"} {
			cmd := exec.Command("curl", "-k", tlsVersion, "https://"+domain)
			if err := cmd.Run(); err == nil {
				result[tlsVersion] = true
			} else {
				result[tlsVersion] = false
			}
		}
	} else if strings.Count(domain, ".") == 1 {
		for _, tlsVersion := range []string{"tls1", "tls1_1", "tls1_2", "tls1_3"} {
			cmd := exec.Command("openssl", "s_client", "-connect", domain+":443", "-"+tlsVersion)
			if err := cmd.Run(); err == nil {
				result[tlsVersion] = true
			} else {
				result[tlsVersion] = false
			}
		}
	} else {
		for _, tlsVersion := range []string{"--tlsv1.0", "--tlsv1.1", "--tlsv1.2", "--tlsv1.3"} {
			cmd := exec.Command("curl", "-k", tlsVersion, "https://"+domain)
			if err := cmd.Run(); err == nil {
				result[tlsVersion] = true
			} else {
				result[tlsVersion] = false
			}
		}
	}
	return result, nil
}
