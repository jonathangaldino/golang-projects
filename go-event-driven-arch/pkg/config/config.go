package config

type Config struct {
	ServerPort  string
	Environment string
}

func Load() *Config {
	return &Config{
		ServerPort:  "8080",
		Environment: "development",
	}
}
