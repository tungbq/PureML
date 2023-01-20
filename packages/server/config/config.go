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

func GetHost() string {
	env := Environment()
	return env["HOST_URL"]
}

func GetPort() string {
	env := Environment()
	return env["PORT"]
}

func HasAdminAccess(email string) bool {
	_, ok := adminAccess[email]
	return ok
}

func TokenSigningSecret() []byte {
	env := Environment()
	return []byte(env["JWT_SECRET"])
}

// func R2Secrets() []string {
// 	env := Environment()
// 	return []string{env["R2_ACCOUNT_ID"], env["R2_ACCESS_KEY_ID"], env["R2_ACCESS_KEY_SECRET"]}
// }

// func R2BucketName() string {
// 	env := Environment()
// 	return env["R2_BUCKET_NAME"]
// }
