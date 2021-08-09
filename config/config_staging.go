// +build staging

package config

const (
	DB_USER                = "street_market_api"
	DB_PASSWORD            = "{db_password}"
	DB_DATABASE            = "street_market_db"
	DB_HOST                = "127.0.0.1"
	API_PORT               = 8080
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091/"
)
