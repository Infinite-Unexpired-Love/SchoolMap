package utils

import (
	"TGU-MAP/models"
	"encoding/json"
)

func Marshal(src interface{}) []byte {
	if b, err := json.Marshal(src); err != nil {
		panic(models.SerializeError())
	} else {
		return b
	}
}
