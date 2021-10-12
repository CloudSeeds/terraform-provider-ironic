package nwd

import (
	"encoding/json"
	"net"
	"strings"
)

func parseListOfStrings(jsonString string) []string {
	var result []string
	json.Unmarshal([]byte(jsonString), &result)
	return result
}

func parseNumber(jsonString string) int64 {
	var result int64
	json.Unmarshal([]byte(jsonString), &result)
	return result
}

func parseIPAddress(nw map[string]string) (string, string) {
	if nw["type"] == "ipv6" {
		return nw["ip_address"], ""
	}
	if strings.Contains(nw["ip_address"], "/") {
		ipAddress, ipNetmask, err := net.ParseCIDR(nw["ip_address"])
		if err != nil {
			panic(err)
		}

		return ipAddress.String(), ipNetmask.Mask.String()
	}

	return nw["ip_address"], nw["netmask"]
}

func parseRoutes(s string) []map[string]string {
	var parsed []map[string]string

	json.Unmarshal([]byte(s), &parsed)

	for idx, route := range parsed {
		// If the network is currently stored in CIDR notation we split it into address and mask, otherwise we leave it as is
		if strings.Contains(route["network"], "/") {
			_, ipNetmask, err := net.ParseCIDR(route["network"])
			if err != nil {
				panic(err)
			}

			parsed[idx] = map[string]string{
				"network": ipNetmask.IP.String(),
				"netmask": ipNetmask.Mask.String(),
				"gateway": route["gateway"],
			}
		}
	}

	return parsed
}
