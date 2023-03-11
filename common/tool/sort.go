package tool

import (
	"sort"
)

// Ksort 排序
func Ksort(data map[string]string) map[string]string {
	var dataParams = make(map[string]string)
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		dataParams[k] = data[k]
	}
	return dataParams
}