package config

import (

	"os"


)






type Config struct {
	CORS_ALLOWED_ORIGIN string
	REDIS_ADDR          string
	Channelname string
	Jwtsecret string
	Db_name string 
	
	port string
}
func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
func LoadConfig() (*Config, error) {
	var config Config


	
	config.port =  "8080"
	config.Db_name= "gotest.db"
	config.CORS_ALLOWED_ORIGIN = "CORS_ALLOWED_ORIGIN"
	config.REDIS_ADDR = "localhost:6379"
	return &config, nil
}