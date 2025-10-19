package services

import (
	"os/exec"
	"strings"
)

func CheckTLSVersion(domain string) (map[string]bool, error) {
	result := make(map[string]bool)

	var cmdBuilder func(string) *exec.Cmd
	if strings.Contains(domain, "/") || strings.Count(domain, ".") > 1 {
		cmdBuilder = func(tlsVersion string) *exec.Cmd {
			return exec.Command("curl", "-k", tlsVersion, "https://"+domain)
		}
	} else {
		cmdBuilder = func(tlsVersion string) *exec.Cmd {
			return exec.Command("openssl", "s_client", "-connect", domain+":443", "-"+tlsVersion)
		}
	}

	for _, tlsVersion := range []string{"tls1", "tls1_1", "tls1_2", "tls1_3", "--tlsv1.0", "--tlsv1.1", "--tlsv1.2", "--tlsv1.3"} {
		cmd := cmdBuilder(tlsVersion)
		if err := cmd.Run(); err == nil {
			result[tlsVersion] = true
		} else {
			result[tlsVersion] = false
		}
	}

	return result, nil
}
