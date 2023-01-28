package config

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	// fmt.Println("Loading environment variables from .env file")
	godotenv.Load("../.env")
}

var adminAccess = map[string]bool{
	"priyavkkaneria@gmail.com": true,
	"kaneriakesha7@gmail.com":  true,
	"akshilvthumar@gmail.com":  true,
	"test.pureml@gmail.com":    true,
	"demo@aztlan.in":           true,
}

func GetHost() string {
	return os.Getenv("HOST_URL")
}

func GetPort() string {
	return os.Getenv("PORT")
}

func GetPureMLR2Secrets() map[string]string {
	return map[string]string{
		"R2_ACCOUNT_ID":       os.Getenv("R2_ACCOUNT_ID"),
		"R2_ACCESS_KEY_ID":    os.Getenv("R2_ACCESS_KEY_ID"),
		"R2_ACCESS_KEY_SECRET": os.Getenv("R2_ACCESS_KEY_SECRET"),
		"R2_BUCKET_NAME":      os.Getenv("R2_BUCKET_NAME"),
		"R2_PUBLIC_URL":       os.Getenv("R2_PUBLIC_URL"),
	}
}

func IsCGOEnabled() bool {
	return os.Getenv("CGO_ENABLED") == "1"
}

func GetDatabaseType() string {
	return os.Getenv("DATABASE")
}

func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}

func GetDataDir() string {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = "./pureml_data"
	}
	return dataDir
}

func GetHttpAddr() string {
	httpAddr := os.Getenv("HTTP_ADDRESS")
	if httpAddr == "" {
		httpAddr = "127.0.0.1:8080"
	}
	return httpAddr
}

func GetHttpsAddr() string {
	return os.Getenv("HTTPS_ADDRESS")
}

func HasAdminAccess(email string) bool {
	_, ok := adminAccess[email]
	return ok
}

func TokenSigningSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}