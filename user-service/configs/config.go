package configs

import "github.com/spf13/viper"

type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`
}

type SqlDB struct {
	Host              string `json:"host"`
	Port              string `json:"port"`
	User              string `json:"user"`
	Password          string `json:"password"`
	DBName            string `json:"db_name"`
	MaxOpenConnection int    `json:"db_max_open_connection"`
	MaxIdleConnection int    `json:"db_max_idle_connection"`
}

type Redis struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type RabbitMQ struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Supabase struct {
	URL    string `json:"url"`
	Secret string `json:"secret"`
	Bucket string `json:"bucket"`
}

type Config struct {
	App      App      `json:"app"`
	SqlDB    SqlDB    `json:"sql_db"`
	Redis    Redis    `json:"redis"`
	RabbitMQ RabbitMQ `json:"rabbitmq"`
	Supabase Supabase `json:"supabase"`
}

func NewConfig() *Config {
	return &Config{
		App: App{
			AppPort: viper.GetString("APP_PORT"),
			AppEnv:  viper.GetString("APP_ENV"),
		},
		SqlDB: SqlDB{
			Host:              viper.GetString("DATABASE_HOST"),
			Port:              viper.GetString("DATABASE_PORT"),
			User:              viper.GetString("DATABASE_USER"),
			Password:          viper.GetString("DATABASE_PASSWORD"),
			MaxOpenConnection: viper.GetInt("DATABASE_MAX_OPEN_CONNECTION"),
			MaxIdleConnection: viper.GetInt("DATABASE_MAX_IDLE_CONNECTION"),
		},
		Redis: Redis{
			Host: viper.GetString("REDIS_HOST"),
			Port: viper.GetString("REDIS_PORT"),
		},
		RabbitMQ: RabbitMQ{
			Host:     viper.GetString("RABBITMQ_HOST"),
			Port:     viper.GetString("RABBITMQ_PORT"),
			Username: viper.GetString("RABBITMQ_USERNAME"),
			Password: viper.GetString("RABBITMQ_PASSWORD"),
		},
		Supabase: Supabase{
			URL:    viper.GetString("SUPABASE_URL"),
			Secret: viper.GetString("SUPABASE_SECRET"),
			Bucket: viper.GetString("SUPABASE_BUCKET"),
		},
	}
}
