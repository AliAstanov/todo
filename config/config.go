package config

type PsqlConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

type Config struct {
	PsqlConfig PsqlConfig
}

func Load() Config {
	return Config{
		PsqlConfig: PsqlConfig{
			User: "boot",
			Password: "root",
			Host: "0.0.0.0",
			Port: 5432,
			Name: "my_db",
		},

	}
	
}
