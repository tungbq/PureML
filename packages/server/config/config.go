package config

import (
	"os"
)

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

func HasAdminAccess(email string) bool {
	_, ok := adminAccess[email]
	return ok
}

func TokenSigningSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}