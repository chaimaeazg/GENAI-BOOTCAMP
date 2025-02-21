package config

type Config struct {
	DBPath string
}

func Load() *Config {
	return &Config{
		DBPath: "db/words.db",
	}
}
