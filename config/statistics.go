package config

func GetStatAllowSite() map[string]int {
	source := Get("statistics:site").([]interface{})
	allowSite := make(map[string]int, len(data))

	for i := 0; i < len(data); i++ {
		allowSite[source[i].(string)] = 0
	}

	return allowSite
}
