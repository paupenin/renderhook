package config

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     Env("DB_HOST", "localhost"),
		Port:     Env("DB_PORT", "5432"),
		User:     Env("DB_USER", "user"),
		Password: Env("DB_PASSWORD", "password"),
		DBName:   Env("DB_NAME", "dbname"),
	}
}
