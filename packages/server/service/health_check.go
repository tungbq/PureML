package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/models"
)

func HealthCheck(request *models.Request) *models.Response {
	return &models.Response{
		StatusCode: http.StatusOK,
		Body: models.ResponseBody{
			Status:  http.StatusOK,
			Message: "Server is up and running🚀🎉",
			Data:    nil,
		},
	}
}
