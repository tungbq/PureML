package service

import (
	_ "fmt"
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

var defaultModelBranchNames = []string{"main", "development"}

// CreateModel godoc
// @Security ApiKeyAuth
// @Summary Register model
// @Description Register model file. Create model and default branches if not exists
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/create [post]
// @Param orgId path string true "Organization UUID"
// @Param modelName path string true "Model name"
// @Param data body models.CreateModelRequest true "Model details"
func CreateModel(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := request.GetOrgId()
	userUUID := request.GetUserUUID()
	modelName := request.GetPathParam("modelName")
	modelWiki := request.GetParsedBodyAttribute("wiki")
	var modelWikiData string
	if modelWiki == nil {
		modelWikiData = ""
	} else {
		modelWikiData = modelWiki.(string)
	}
	modelIsPublic := request.GetParsedBodyAttribute("is_public")
	var modelIsPublicData bool
	if modelIsPublic == nil {
		modelIsPublicData = false
	} else {
		modelIsPublicData = modelIsPublic.(bool)
	}
	modelBranchNames := request.GetParsedBodyAttribute("branches")
	var modelBranchNamesData []string
	if modelBranchNames == nil {
		modelBranchNamesData = defaultModelBranchNames
	} else {
		modelBranchNamesData = modelBranchNames.([]string)
	}
	modelReadme := request.GetParsedBodyAttribute("readme")
	var modelReadmeData *models.ReadmeRequest
	if modelReadme == nil {
		modelReadmeData = &models.ReadmeRequest{
			FileType: "markdown",
			Content:  "",
		}
	} else {
		modelReadmeData = &models.ReadmeRequest{
			FileType: modelReadme.(map[string]interface{})["file_type"].(string),
			Content:  modelReadme.(map[string]interface{})["content"].(string),
		}
	}
	model, err := datastore.GetModelByName(orgId, modelName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if model != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Model already exists")
	}
	model, err = datastore.CreateModel(orgId, modelName, modelWikiData, modelIsPublicData, modelReadmeData, userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	_, err = datastore.CreateModelBranches(model.UUID, modelBranchNamesData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, model, "Model & branches successfully created")
}
