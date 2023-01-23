package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// UpdateModelReadme godoc
// @Security ApiKeyAuth
// @Summary Update readme of a model for a category
// @Description Update readme of a model for a category
// @Tags Model
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/{orgId}/model/{modelName}/readme [post]
// @Param orgId path string true "Organization Id"
// @Param modelName path string true "Model Name"
// @Param data body models.ReadmeRequest true "Data"
func UpdateModelReadme(request *models.Request) *models.Response {
	request.ParseJsonBody()
	modelUUID := request.GetModelUUID()
	modelFileType := request.GetParsedBodyAttribute("file_type")
	var modelFileTypeData string
	if modelFileType == nil {
		modelFileTypeData = ""
	} else {
		modelFileTypeData = modelFileType.(string)
	}
	modelContent := request.GetParsedBodyAttribute("content")
	var modelContentData string
	if modelContent == nil {
		modelContentData = ""
	} else {
		modelContentData = modelContent.(string)
	}
	readme, err := datastore.UpdateModelReadme(modelUUID, modelFileTypeData, modelContentData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, readme, "Model readme updated")
}
