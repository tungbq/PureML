package service

import (
	"fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

var defaultDatasetBranchNames = []string{"main", "development"}

// RegisterDataset godoc
// @Security ApiKeyAuth
// @Summary Register dataset
// @Description Register dataset file. Create dataset and default branches if not exists
// @Tags Dataset
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/dataset/{datasetName}/register [post]
// @Param file formData file true "Dataset file"
// @Param orgId path string true "Organization UUID"
// @Param datasetName path string true "Dataset name"
// @Param data formData models.RegisterDatasetRequest true "Dataset details"
func RegisterDataset(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	userUUID := request.GetUserUUID()
	datasetName := request.GetPathParam("datasetName")
	var datasetHash string
	if request.FormValues["hash"] != nil && len(request.FormValues["hash"]) > 0 {
		datasetHash = request.FormValues["hash"][0]
	} else {
		return models.NewErrorResponse(http.StatusBadRequest, "Hash is required")
	}
	var datasetWiki string
	if request.FormValues["wiki"] != nil && len(request.FormValues["wiki"]) > 0 {
		datasetWiki = request.FormValues["wiki"][0]
	}
	var datasetLineage string
	if request.FormValues["lineage"] != nil && len(request.FormValues["lineage"]) > 0 {
		datasetLineage = request.FormValues["lineage"][0]
	}
	fileHeader := request.GetFormFile("file")
	if fileHeader == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "File is required")
	}
	dataset, err := datastore.GetDatasetByName(orgId, datasetName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	fmt.Println(dataset)
	if dataset == nil {
		// Create dataset and default branches as dataset does not exist
		dataset, err := datastore.CreateDataset(orgId, datasetName, datasetWiki, userUUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		datasetBranches, err := datastore.CreateDatasetBranches(dataset.UUID, defaultDatasetBranchNames)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		datasetVersion, err := datastore.UploadAndRegisterDatasetFile(datasetBranches[1].UUID, fileHeader, datasetHash, "R2", datasetLineage)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		return models.NewDataResponse(http.StatusOK, datasetVersion, "Dataset successfully created")
	} else {
		// Dataset exists
		datasetBranches, err := datastore.GetDatasetAllBranches(dataset.UUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		if len(datasetBranches) == 0 {
			// Create default branches as dataset branches does not exist
			datasetBranches, err := datastore.CreateDatasetBranches(dataset.UUID, defaultDatasetBranchNames)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			datasetVersion, err := datastore.UploadAndRegisterDatasetFile(datasetBranches[1].UUID, fileHeader, datasetHash, "R2", datasetLineage)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			return models.NewDataResponse(http.StatusOK, datasetVersion, "Dataset successfully created")
		} else {
			// Dataset branches exists (defaultBranches)
			var developementBranch models.DatasetBranchResponse
			for _, branch := range datasetBranches {
				if branch.Name == "developement" {
					developementBranch = branch
					break
				}
			}
			if developementBranch.UUID == uuid.Nil {
				return models.NewErrorResponse(http.StatusConflict, "Dataset developement branch not found")
			}
			datasetVersion, err := datastore.UploadAndRegisterDatasetFile(developementBranch.UUID, fileHeader, datasetHash, "R2", datasetLineage)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			return models.NewDataResponse(http.StatusOK, datasetVersion, "Dataset successfully created")
		}
	}
}
