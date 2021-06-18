/**
  @author:panliang
  @data:2021/6/18
  @note
**/
package helpler

import (
	"encoding/json"
)

// json string to a map type

func JsonToMap(str []byte) map[string]interface{}  {
	var jsonMap  map[string]interface{}

	err := json.Unmarshal(str, &jsonMap)
	if err != nil {
		panic(err)
	}
	return jsonMap
}

