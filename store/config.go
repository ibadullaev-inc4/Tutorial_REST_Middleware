package store

// Config ...
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// DatabaseURL: "host=127.0.0.1 port=5432 user=nariman dbname=restapi_dev sslmode=disable",
func NewConfig() *Config {
	return &Config{}
}
