package service

import (
	// "fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetAllAdminOrgs godoc
// @Summary Get all organizations and their details.
// @Description Get all organizations and their details. Only accessible by admins.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/all [get]
func GetAllAdminOrgs(request *models.Request) *models.Response {
	response := &models.Response{}
	if config.HasAdminAccess(request.User.Email) {
		allOrgs, err := datastore.GetAllAdminOrgs()
		if err != nil {
			return models.NewErrorResponse(err)
		} else {
			response.StatusCode = http.StatusOK
			response.Body.Status = response.StatusCode
			response.Body.Message = "All organizations"
			response.Body.Data = allOrgs
		}
	} else {
		response.StatusCode = http.StatusForbidden
		response.Body.Status = response.StatusCode
		response.Body.Message = "Forbidden"
		response.Body.Data = nil
	}
	return response
}
