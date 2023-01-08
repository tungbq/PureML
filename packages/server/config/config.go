package config

import (
	"fmt"

	// "github.com/PureML-Inc/PureML/server/models"
	"github.com/joho/godotenv"
)

var adminAccess = map[string]bool{
	"priyavkkaneria@gmail.com": true,
	"kaneriakesha7@gmail.com":  true,
	"akshilvthumar@gmail.com":  true,
	"test.pureml@gmail.com":    true,
	"demo@aztlan.in":           true,
}

func Environment() map[string]string {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}
	return myEnv
}

func HasAdminAccess(email string) bool {
	_, ok := adminAccess[email]
	return ok
}

func TokenSigningSecret() string {
	env := Environment()
	return env["JWT_SECRET"]
}
