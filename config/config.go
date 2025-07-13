package config

type Config struct {
	Port      string
	CacheSize int
}

func NewConfig() *Config {
	return  &Config{
		Port: "8080",
		CacheSize: 1000,
	}
}