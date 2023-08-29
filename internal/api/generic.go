package api

func SliceToMap(s []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range s {
		m[v] = true
	}

	return m
}

func RemoveStringSlice(population []string, remove []string) []string {
	m := make(map[string]bool)
	for _, v := range remove {
		m[v] = true
	}

	result := []string{}
	for _, v := range population {
		if !m[v] {
			result = append(result, v)
		}
	}

	return result
}
