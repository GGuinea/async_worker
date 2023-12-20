package config

type DbConfig struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DbName string
}

type Config struct {
	Db DbConfig
}

func NewConfig() *Config {
	dbConfig := DbConfig{
		Host:   "localhost",
		Port:   "5432",
		User:   "postgres",
		Pass:   "postgres",
		DbName: "postgres",
	}
	return &Config{Db: dbConfig}
}
