package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

var defaultBranchNames = []string{"main", "developement"}

func RegisterModel(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := request.GetOrgId().String()
	modelName := request.GetParsedBodyAttribute("name").(string)
	model, err := datastore.GetModelByName(orgId, modelName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if model == nil {
		_, err := datastore.CreateModel(orgId, modelName)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		modelBranches, err := datastore.CreateModelBranches(orgId, modelName, defaultBranchNames)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		fileHeader := request.GetFormFile("file")
		if fileHeader == nil {
			//TODO mandatory file is missing return with appropriate response here
		}
		for _, modelBranch := range modelBranches {
			err = datastore.UploadModelFile(orgId, modelBranch.Name, fileHeader)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
		}
		return models.NewDataResponse(http.StatusOK, model, "Model successfully created")

	}
	return models.NewErrorResponse(http.StatusConflict, "Model with same name already exists")
}
