package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetModelAllBranches godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get all branches of a model
//	@Description	Get all branches of a model
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
func GetModelAllBranches(request *models.Request) *models.Response {
	var response *models.Response
	modelUUID := request.GetModelUUID()
	allOrgs, err := datastore.GetModelAllBranches(modelUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	} else {
		response = models.NewDataResponse(http.StatusOK, allOrgs, "All model branches")
	}
	return response
}
