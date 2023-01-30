package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetAllModels godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get all models of an organization
//	@Description	Get all models of an organization
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/all [get]
//	@Param			orgId	path	string	true	"Organization Id"
func GetAllModels(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	allModels, err := datastore.GetAllModels(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, allModels, "Models successfully retrieved")
}
