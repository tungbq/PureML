package service

import (
	"fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

var defaultBranchNames = []string{"main", "development"}

// RegisterModel godoc
// @Security ApiKeyAuth
// @Summary Register model
// @Description Register model file. Create model and default branches if not exists
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/register [post]
// @Param file formData file true "Model file"
// @Param orgId path string true "Organization UUID"
// @Param modelName path string true "Model name"
// @Param data formData models.RegisterModelRequest true "Model details"
func RegisterModel(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	userUUID := request.GetUserUUID()
	modelName := request.GetPathParam("modelName")
	modelHash := request.FormValues["hash"][0]
	modelWiki := request.FormValues["wiki"][0]
	fileHeader := request.GetFormFile("file")
	if fileHeader == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "File is required")
	}
	model, err := datastore.GetModelByName(orgId, modelName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	fmt.Println(model)
	if model == nil {
		// Create model and default branches as model does not exist
		model, err := datastore.CreateModel(orgId, modelName, modelWiki, userUUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		modelBranches, err := datastore.CreateModelBranches(model.UUID, defaultBranchNames)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		modelVersion, err := datastore.UploadAndRegisterModelFile(modelBranches[1].UUID, fileHeader, modelHash, "R2")
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		return models.NewDataResponse(http.StatusOK, modelVersion, "Model successfully created")
	} else {
		// Model exists
		modelBranches, err := datastore.GetModelAllBranches(model.UUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		if len(modelBranches) == 0 {
			// Create default branches as model branches does not exist
			modelBranches, err := datastore.CreateModelBranches(model.UUID, defaultBranchNames)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			modelVersion, err := datastore.UploadAndRegisterModelFile(modelBranches[1].UUID, fileHeader, modelHash, "R2")
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			return models.NewDataResponse(http.StatusOK, modelVersion, "Model successfully created")
		} else {
			// Model branches exists (defaultBranches)
			var developementBranch models.ModelBranchResponse
			for _, branch := range modelBranches {
				if branch.Name == "developement" {
					developementBranch = branch
					break
				}
			}
			if developementBranch.UUID == uuid.Nil {
				return models.NewErrorResponse(http.StatusConflict, "Model developement branch not found")
			}
			modelVersion, err := datastore.UploadAndRegisterModelFile(developementBranch.UUID, fileHeader, modelHash, "R2")
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			return models.NewDataResponse(http.StatusOK, modelVersion, "Model successfully created")
		}
	}
}
