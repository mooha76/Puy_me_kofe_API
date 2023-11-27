package config

type Config struct {
	DBConfig DatabaseConfig `bson:"dbconfig"`
}

type DatabaseConfig struct {
	DB_HOST     string `bson:"DB_HOST"`
	DB_PORT     string `bson:"DB_PORT"`
	DB_NAME     string `bson:"DB_NAME"`
	DB_USER     string `bson:"DB_USER"`
	DB_PASSWORD string `bson:"DB_PASSWORD"`
	Sslmode     string `bson:"Sslmode"`
}
