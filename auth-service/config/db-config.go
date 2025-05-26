package config

type DataBaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// var Load interface {
// 	Load( ) DataBaseConfig
// }

// func Load() DataBaseConfig {}
func Load() DataBaseConfig {
	return DataBaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "auth_user",
		Password: "secret123",
		DBName:   "auth_db",
	}
}
