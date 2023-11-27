package config

type Config struct {
	DBConfig    DatabaseConfig `bson:"dbconfig"`
	Sysytemconf Sysytem        `bson:"system"`
	Syslogger   Logger         `bson:"logger"`
}

type DatabaseConfig struct {
	DB_HOST     string `bson:"DB_HOST"`
	DB_PORT     string `bson:"DB_PORT"`
	DB_NAME     string `bson:"DB_NAME"`
	DB_USER     string `bson:"DB_USER"`
	DB_PASSWORD string `bson:"DB_PASSWORD"`
	Sslmode     string `bson:"Sslmode"`
}

type Sysytem struct {
	Port string `bson:"Port"`
}

type Logger struct {
	Path       string `bson:"path"`
	Filename   string `bson:"filename"`
	MaxSize    int    `bson:"maxSize"`
	MaxBackups int    `bson:"maxBackups"`
	MaxAge     int    `bson:"maxAge"`
}
