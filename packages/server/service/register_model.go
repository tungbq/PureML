package service

import (
	_ "fmt"
	"net/http"
	"strings"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

var defaultModelBranchNames = []string{"main", "development"}

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
	var modelHash string
	if request.FormValues["hash"] != nil && len(request.FormValues["hash"]) > 0 {
		modelHash = request.FormValues["hash"][0]
	} else {
		return models.NewErrorResponse(http.StatusBadRequest, "Hash is required")
	}
	var modelWiki string
	if request.FormValues["wiki"] != nil && len(request.FormValues["wiki"]) > 0 {
		modelWiki = request.FormValues["wiki"][0]
	}
	var modelSourceType string
	if request.FormValues["storage"] != nil && len(request.FormValues["storage"]) > 0 {
		modelSourceType = strings.ToUpper(request.FormValues["storage"][0])
	}
	fileHeader := request.GetFormFile("file")
	if fileHeader == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "File is required")
	}
	model, err := datastore.GetModelByName(orgId, modelName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if model == nil {
		// Create model and default branches as model does not exist
		model, err := datastore.CreateModel(orgId, modelName, modelWiki, userUUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		modelBranches, err := datastore.CreateModelBranches(model.UUID, defaultModelBranchNames)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		modelVersion, err := datastore.UploadAndRegisterModelFile(orgId, modelBranches[1].UUID, fileHeader, modelHash, modelSourceType)
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
			modelBranches, err := datastore.CreateModelBranches(model.UUID, defaultModelBranchNames)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			modelVersion, err := datastore.UploadAndRegisterModelFile(orgId, modelBranches[1].UUID, fileHeader, modelHash, modelSourceType)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			return models.NewDataResponse(http.StatusOK, modelVersion, "Model successfully created")
		} else {
			// Model branches exists (defaultBranches)
			var developmentBranch models.ModelBranchResponse
			for _, branch := range modelBranches {
				if branch.Name == "development" {
					developmentBranch = branch
					break
				}
			}
			if developmentBranch.UUID == uuid.Nil {
				return models.NewErrorResponse(http.StatusConflict, "Model development branch not found")
			}
			modelVersion, err := datastore.UploadAndRegisterModelFile(orgId, developmentBranch.UUID, fileHeader, modelHash, modelSourceType)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			return models.NewDataResponse(http.StatusOK, modelVersion, "Model successfully created")
		}
	}
}
