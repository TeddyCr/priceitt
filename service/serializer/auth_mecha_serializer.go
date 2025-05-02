package serializer

import (
	"encoding/json"

	"github.com/TeddyCr/priceitt/service/models/generated/auth"
)

func DeSerBasicAuthMechanism(authMechanism map[string]interface{}) (auth.Basic, error) {
	var basicAuthMechanism auth.Basic
	jsonBytes, err := json.Marshal(authMechanism)
	if err != nil {
		return auth.Basic{}, err
	}
	err = json.Unmarshal(jsonBytes, &basicAuthMechanism)
	if err != nil {
		return auth.Basic{}, err
	}
	return basicAuthMechanism, nil
}

func DeSerGoogleAuthMechanism(authMechanism map[string]interface{}) (auth.Google, error) {
	var googleAuthMechanism auth.Google
	jsonBytes, err := json.Marshal(authMechanism)
	if err != nil {
		return auth.Google{}, err
	}
	err = json.Unmarshal(jsonBytes, &googleAuthMechanism)
	if err != nil {
		return auth.Google{}, err
	}
	return googleAuthMechanism, nil
}
