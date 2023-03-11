package turn

import "encoding/json"

func StructToJson(data interface{}) string {
	m, _ := json.Marshal(data)
	return string(m)
}