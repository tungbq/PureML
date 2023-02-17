package main

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

var (
	functions = map[string]func(args ...string){
		"generateToken": GenerateToken,
	}
)

func GenerateToken(args ...string) {
	if len(args) != 4 {
		panic("GenerateToken: expected 4 arguments (secret, uuid, email, handle), got " + fmt.Sprint(len(args)))
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":   args[1],
		"email":  args[2],
		"handle": args[3],
	})
	signedString, err := token.SignedString([]byte(args[0]))
	if err != nil {
		panic(err)
	}
	println(signedString)
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 1 {
		panic("no function specified")
	}
	function, ok := functions[arguments[0]]
	if !ok {
		panic("unknown function " + arguments[0])
	}
	function(arguments[1:]...)
}
