module priceitt.xyz/edgeAuthorizationServer

go 1.22

toolchain go1.22.7

require (
	github.com/TeddyCr/priceitt/models v0.0.1-alpha
	github.com/TeddyCr/priceitt/utils v0.0.1-alpha
	github.com/fernet/fernet-go v0.0.0-20240119011108-303da6aec611
	github.com/go-chi/chi/v5 v5.1.0
	github.com/google/uuid v1.6.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/lib/pq v1.10.9
	github.com/mitchellh/mapstructure v1.5.0
	github.com/stretchr/testify v1.9.0
	golang.org/x/crypto v0.28.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/TeddyCr/priceitt/utils => ../utils

replace github.com/TeddyCr/priceitt/models => ../models
