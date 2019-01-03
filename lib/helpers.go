package lib

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(message string, object interface{}) {
	jsonResp, err := json.Marshal(object)

	if err != nil {
		fmt.Println("error parsing object", err, object)
	}

	fmt.Println(message, string(jsonResp))
}