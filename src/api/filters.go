package api

import "net/url"

func ConvertQueryToFilter(values url.Values, filterFields []string) map[string]interface{} {
	filters := make(map[string]interface{})
	for _, field := range filterFields {
		if value, ok := values[field]; ok {
			filters[field] = value
		}
	}
	return filters
}
