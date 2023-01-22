package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// VerifyDatasetBranchHashStatus godoc
// @Security ApiKeyAuth
// @Summary Verify dataset hash status
// @Description Verify dataset hash status to determine if dataset is already uploaded
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/branch/{branchName}/hash-status [post]
// @Param orgId path string true "Organization Id"
// @Param datasetName path string true "Dataset Name"
// @Param branchName path string true "Branch Name"
// @Param hash body models.HashRequest true "Hash value"
func VerifyDatasetBranchHashStatus(request *models.Request) *models.Response {
	datasetName := request.GetDatasetName()
	datasetBranchName := request.GetPathParam("branchName")
	orgId := uuid.Must(uuid.FromString(request.GetPathParam("orgId")))
	message := "Hash validity (True - does not exist in db)"
	dataset, err := datastore.GetDatasetBranchByName(orgId, datasetName, datasetBranchName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if dataset == nil {
		return models.NewDataResponse(http.StatusOK, true, message)
	}
	datasetBranchUUID := dataset.UUID
	request.ParseJsonBody()
	hashValue := request.GetParsedBodyAttribute("hash").(string)
	versions, err := datastore.GetDatasetBranchAllVersions(datasetBranchUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := true
	for _, version := range versions {
		if version.Hash == hashValue {
			response = false
			message = "Hash validity (False - exists in db)"
			break
		}
	}
	return models.NewDataResponse(http.StatusOK, response, message)
}
