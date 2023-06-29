package config

type PgConfig struct {
	Host string
	Port string
	User string
	Pwd  string
	Db   string
}

type RedisConfig struct {
	Host               string `validate:"required"`
	Port               string `validate:"required"`
	MinIdleConns       int    `validate:"required"`
	PoolSize           int    `validate:"required"`
	PoolTimeout        int    `validate:"required"`
	Password           string `validate:"required"`
	UseCertificates    bool
	InsecureSkipVerify bool
	CertificatesPaths  struct {
		Cert string
		Key  string
		Ca   string
	}
	DB int
}

type Config struct {
	Pg    PgConfig
	Redis RedisConfig
}

func GetConfig() *Config {
	return &Config{
		Pg: PgConfig{
			Host: "127.0.0.1",
			Port: "5432",
			User: "postgres",
			Pwd:  "postgres",
			Db:   "postgres",
		},
		Redis: RedisConfig{
			Host:               "127.0.0.1",
			Port:               "6379",
			MinIdleConns:       10,
			PoolSize:           10,
			PoolTimeout:        10,
			Password:           "dev",
			UseCertificates:    false,
			InsecureSkipVerify: false,
			CertificatesPaths: struct {
				Cert string
				Key  string
				Ca   string
			}{},
			DB: 1,
		},
	}
}
