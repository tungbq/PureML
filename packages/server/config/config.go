package config

import (
	"github.com/PriyavKaneria/PureML/service/models"
)

var adminAccess = map[string]bool{
	"priyavkkaneria@gmail.com": true,
	"kaneriakesha7@gmail.com":  true,
	"akshilvthumar@gmail.com":  true,
	"test.pureml@gmail.com":    true,
	"demo@aztlan.in":           true,
}

func HasAdminAccess(user models.User) bool {
	_, ok := adminAccess[user.Email]
	return ok
}
