package pkg

import (
	"encoding/json"
	"log"
)

func StructBuilder[T any](data []byte) (T, error) {
	var v T
	err := json.Unmarshal(data, &v)
	if err != nil {
		log.Println("Failed to unmarshal JSON", err)
		return v, err
	}

	return v, nil
}
