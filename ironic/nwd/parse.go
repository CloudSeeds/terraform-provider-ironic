package nwd

import (
	"encoding/json"
	"net"
	"strings"
)

func parseListOrSingleDNSAddress(jsonString string) []string {
	var result []string

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		result = append(result, jsonString)
	}

	return result
}

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

func parseIPAddress(nw map[string]interface{}) (string, string) {
	if nw["type"] == "ipv6" {
		return nw["ip_address"].(string), ""
	}
	if strings.Contains(nw["ip_address"].(string), "/") {
		ipAddress, ipNetmask, err := net.ParseCIDR(nw["ip_address"].(string))
		if err != nil {
			panic(err)
		}

		return ipAddress.String(), ipNetmask.Mask.String()
	}

	return nw["ip_address"].(string), nw["netmask"].(string)
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
