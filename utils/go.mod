module github.com/TeddyCr/priceitt/utils

go 1.22

toolchain go1.22.7

require (
	github.com/jmoiron/sqlx v1.4.0
	github.com/lib/pq v1.10.9
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/stretchr/testify v1.9.0
	github.com/TeddyCr/priceitt/models v0.0.1-alpha
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/TeddyCr/priceitt/models => ../models
