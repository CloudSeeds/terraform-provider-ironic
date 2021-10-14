package nwd

func cleanPort(port map[string]interface{}) interface{} {
	result := make(map[string]interface{})

	for k, vIface := range port {
		v := vIface.(string)
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

func cleanNetwork(nw map[string]interface{}) interface{} {
	result := make(map[string]interface{})

	for k, vIface := range nw {
		v := vIface.(string)
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
