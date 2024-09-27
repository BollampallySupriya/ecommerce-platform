package application



type Config struct {
	RedisAddress string
	DBUser string
	DBPassword string
	DBName string
	DBHost string
	DBPort uint32
	DBSSLMode string
	ServerAddr string
}
func LoadConfig() Config{
	// load it from env

	cfg := Config{
		RedisAddress: "localhost:6379",
		DBUser: "root",
		DBPassword: "secret",
		DBName: "ecommerce",
		DBHost: "localhost",
		DBPort: 6543,
		DBSSLMode: "true",
		ServerAddr: ":3000",
	}
	return cfg
}