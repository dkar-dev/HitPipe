package config

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Logger   LoggerConfig   `mapstructure:"logger"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type PostgresConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DB       string `mapstructure:"db"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}
