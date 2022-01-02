package helpers

import (
	"encoding/json"
)

func ToString(data interface{}) string {
	j, _ := json.MarshalIndent(data, "", "   ")
	return string(j)
}
