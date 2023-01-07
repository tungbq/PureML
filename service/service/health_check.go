package service

import (
	"net/http"

	"github.com/PriyavKaneria/PureML/service/models"
)

func HealthCheck(request *models.Request) *models.Response {
	return &models.Response{
		StatusCode: http.StatusOK,
		Body:       "Congratulations!",
		Message:    "Server is up and running🚀🎉",
	}
}
