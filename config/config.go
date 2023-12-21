package config

type Config struct {
	Db                        DbConfig
	BackgroundProcessorConfig BackgroundProcessorConfig
}

type DbConfig struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DbName string
}

type BackgroundProcessorConfig struct {
	MaxWorkers int
}

func NewConfig() *Config {
	dbConfig := DbConfig{
		Host:   "localhost",
		Port:   "5432",
		User:   "postgres",
		Pass:   "postgres",
		DbName: "postgres",
	}

	backgroundProcessorConfig := BackgroundProcessorConfig{MaxWorkers: 100}
	return &Config{Db: dbConfig, BackgroundProcessorConfig: backgroundProcessorConfig}
}
