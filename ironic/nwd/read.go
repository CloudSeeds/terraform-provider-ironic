package nwd

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func GetNetworkData(d *schema.ResourceData) map[string]interface{} {
	result := make(map[string]interface{})
	l := make([]interface{}, 0)
	n := make([]interface{}, 0)
	s := make([]interface{}, 0)

	if ports, ok := d.GetOk("ports"); ok {
		for _, portIface := range ports.(*schema.Set).List() {
			port := portIface.(map[string]interface{})
			if _, ok := port["id"]; ok {
				l = append(l, cleanPort(port))
			}
		}
	}

	if networks, ok := d.GetOk("networks"); ok {
		for _, networkIface := range networks.(*schema.Set).List() {
			network := networkIface.(map[string]interface{})
			n = append(n, cleanNetwork(network))
		}
	}

	if len(l) == 0 && len(n) == 0 && len(s) == 0 {
		return nil
	}

	return result
}
