package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// VerifyModelBranchHashStatus godoc
// @Security ApiKeyAuth
// @Summary Verify model hash status
// @Description Verify model hash status to determine if model is already uploaded
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/org/{orgId}/model/{modelName}/branch/{branchName}/hash-status [post]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param branchName path string true "Branch Name"
// @Param hash body models.HashRequest true "Hash value"
func VerifyModelBranchHashStatus(request *models.Request) *models.Response {
	modelName := request.GetModelName()
	modelBranchName := request.GetPathParam("branchName")
	orgId := uuid.Must(uuid.FromString(request.GetPathParam("orgId")))
	message := "Hash validity (False - does not exist in db)"
	model, err := datastore.GetModelBranchByName(orgId, modelName, modelBranchName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if model == nil {
		return models.NewDataResponse(http.StatusOK, false, message)
	}
	modelBranchUUID := model.UUID
	request.ParseJsonBody()
	hashValue := request.GetParsedBodyAttribute("hash").(string)
	versions, err := datastore.GetModelBranchAllVersions(modelBranchUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := false
	for _, version := range versions {
		if version.Hash == hashValue {
			response = true
			message = "Hash validity (True - exists in db)"
			break
		}
	}
	return models.NewDataResponse(http.StatusOK, response, message)
}
