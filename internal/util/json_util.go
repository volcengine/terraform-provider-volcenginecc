package util

import "encoding/json"

func JsonString(input any) string {
	marshal, err := json.Marshal(input)
	if err != nil {
		return ""
	}
	return string(marshal)
}
