package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// VerifyModelHashStatus godoc
// @Security ApiKeyAuth
// @Summary Verify model hash status
// @Description Verify model hash status to determine if model is already uploaded
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/hash-status [post]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param hash body string true "Hash value"
func VerifyModelHashStatus(request *models.Request) *models.Response {
	modelName := request.GetModelName()
	modelUUID := request.GetModelUUID()
	if modelName == "" {
		return models.NewDataResponse(http.StatusOK, true, "Hash validity (True - does not exist in db)")
	}
	request.ParseJsonBody()
	hashValue := request.GetParsedBodyAttribute("hash").(string)
	versions, err := datastore.GetModelAllVersions(modelUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := true
	for _, version := range versions {
		if version.Hash == hashValue {
			response = false
			break
		}
	}
	return models.NewDataResponse(http.StatusOK, response, "Hash validity (True - does not exist in db)")
}
