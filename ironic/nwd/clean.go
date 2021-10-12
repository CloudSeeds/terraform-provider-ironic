package nwd

func cleanPort(port map[string]string) interface{} {
	result := make(map[string]interface{})

	for k, v := range port {
		switch k {
		case "bond_links":
			result[k] = parseListOfStrings(v)
		case "bond_miimon":
			result[k] = parseNumber(v)
		case "vlan_id":
			result[k] = parseNumber(v)
		default:
			result[k] = v
		}
	}

	return result
}

func cleanNetwork(nw map[string]string) interface{} {
	result := make(map[string]interface{})

	for k, v := range nw {
		switch k {
		case "ip_address":
			address, netmask := parseIPAddress(nw)

			result["ip_address"] = address
			if netmask != "" {
				result["netmask"] = netmask
			}
		case "netmask": // We silently drop this field, it is always covered by `ip_address`
			break
		case "routes":
			result[k] = parseRoutes(v)
		default:
			result[k] = v
		}
	}

	return result
}
