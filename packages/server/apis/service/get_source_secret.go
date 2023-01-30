package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// GetR2Secret godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get secrets for source type r2
//	@Description	Get secrets for source type r2
//	@Tags			Secret
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/secret/r2 [get]
//	@Param			orgId	path	string	true	"Organization Id"
func GetR2Secret(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	result, err := datastore.GetSourceSecret(orgId, "R2")
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "R2 secrets")
	return response
}

// GetS3Secret godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get secrets for source type s3
//	@Description	Get secrets for source type s3
//	@Tags			Secret
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/secret/s3 [get]
//	@Param			orgId	path	string	true	"Organization Id"
func GetS3Secret(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	result, err := datastore.GetSourceSecret(orgId, "S3")
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "S3 secrets")
	return response
}
