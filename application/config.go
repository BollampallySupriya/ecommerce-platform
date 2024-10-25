package application



type Config struct {
	RedisAddress string
	DBUser string
	DBPassword string
	DBName string
	DBHost string
	DBPort uint32
	DBSSLMode string
	ServerAddr uint32
}
func LoadConfig() Config{
	// load it from env

	cfg := Config{
		RedisAddress: "localhost:6379",
		DBUser: "postgres",
		DBPassword: "secret",
		DBName: "ecommerce",
		DBHost: "localhost",
		DBPort: 5432,
		DBSSLMode: "allow",
		ServerAddr: 3000,
	}
	return cfg
}