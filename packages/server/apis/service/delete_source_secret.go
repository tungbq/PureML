package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// DeleteR2Secrets godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Delete secrets for source type r2
//	@Description	Delete secrets for source type r2
//	@Tags			Secret
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/secret/r2/delete [delete]
//	@Param			orgId	path	string	true	"Organization Id"
func DeleteR2Secrets(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	err := datastore.DeleteR2Secrets(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, nil, "R2 disconnected")
	return response
}

// DeleteS3Secrets godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Delete secrets for source type s3
//	@Description	Delete secrets for source type s3
//	@Tags			Secret
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/secret/s3/delete [delete]
//	@Param			orgId	path	string	true	"Organization Id"
func DeleteS3Secrets(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	err := datastore.DeleteS3Secrets(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, nil, "S3 disconnected")
	return response
}
