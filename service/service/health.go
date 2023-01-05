package service

import (
	"net/http"

	"github.com/PriyavKaneria/PureML/service/models"
)

func Health(request *models.Request) *models.Response {
	response := &models.Response{}
	response.StatusCode = http.StatusOK
	response.Body = "Service is up and running🚀🎉"
	return response
}
