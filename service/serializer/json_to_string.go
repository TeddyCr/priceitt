package serializer

import (
	"encoding/json"

	"github.com/TeddyCr/priceitt/service/models/generated"
)

func JsonToString(entity generated.IEntity) ([]byte, error) {
	jsonBytes, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}
