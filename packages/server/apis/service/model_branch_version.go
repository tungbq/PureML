package service

import (
	_ "fmt"
	"net/http"
	"strings"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)
// GetModelBranchAllVersions godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get all branch versions of a model
//	@Description	Get all branch versions of a model
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch/{branchName}/version [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
//	@Param			branchName	path	string	true	"Branch Name"
func GetModelBranchAllVersions(request *models.Request) *models.Response {
	var response *models.Response
	branchUUID := request.GetModelBranchUUID()
	allVersions, err := datastore.GetModelBranchAllVersions(branchUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	} else {
		response = models.NewDataResponse(http.StatusOK, allVersions, "All organizations")
	}
	return response
}


// GetModelBranchVersion godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get specific branch version of a model
//	@Description	Get specific branch version of a model
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch/{branchName}/version/{version} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
//	@Param			branchName	path	string	true	"Branch Name"
//	@Param			version		path	string	true	"Version"
func GetModelBranchVersion(request *models.Request) *models.Response {
	branchUUID := request.GetModelBranchUUID()
	versionName := request.PathParams["version"]
	version, err := datastore.GetModelBranchVersion(branchUUID, versionName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if version == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Version not found")
	}
	return models.NewDataResponse(http.StatusOK, version, "Model branch details")
}

// VerifyModelBranchHashStatus godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Verify model hash status
//	@Description	Verify model hash status to determine if model is already uploaded
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch/{branchName}/hash-status [post]
//	@Param			orgId		path	string				true	"Organization Id"
//	@Param			modelName	path	string				true	"Model Name"
//	@Param			branchName	path	string				true	"Branch Name"
//	@Param			hash		body	models.HashRequest	true	"Hash value"
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

// RegisterModel godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Register model
//	@Description	Register model file. Create model and default branches if not exists
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/branch/{branchName}/register [post]
//	@Param			file		formData	file						true	"Model file"
//	@Param			orgId		path		string						true	"Organization UUID"
//	@Param			modelName	path		string						true	"Model name"
//	@Param			branchName	path		string						true	"Branch name"
//	@Param			data		formData	models.RegisterModelRequest	true	"Model details"
func RegisterModel(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	var modelHash string
	if request.FormValues["hash"] != nil && len(request.FormValues["hash"]) > 0 {
		modelHash = request.FormValues["hash"][0]
	} else {
		return models.NewErrorResponse(http.StatusBadRequest, "Hash is required")
	}
	var modelSourceType string
	if request.FormValues["storage"] != nil && len(request.FormValues["storage"]) > 0 {
		modelSourceType = strings.ToUpper(request.FormValues["storage"][0])
	}
	var modelIsEmpty bool
	if request.FormValues["isEmpty"] != nil && len(request.FormValues["isEmpty"]) > 0 {
		modelIsEmpty = request.FormValues["isEmpty"][0] == "true"
	}
	fileHeader := request.GetFormFile("file")
	if fileHeader == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "File is required")
	}
	modelBranchName := request.GetPathParam("branchName")
	if modelBranchName == "main" {
		return models.NewErrorResponse(http.StatusBadRequest, "Cannot register model directly to main branch")
	}
	sourceValid := false
	for source := range models.SupportedSources {
		if models.SupportedSources[source] == modelSourceType {
			sourceValid = true
			break
		}
	}
	if !sourceValid {
		return models.NewErrorResponse(http.StatusBadRequest, "Unsupported model source type")
	}
	modelBranchUUID := request.GetModelBranchUUID()
	versions, err := datastore.GetModelBranchAllVersions(modelBranchUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := false
	for _, version := range versions {
		if version.Hash == modelHash {
			response = true
			break
		}
	}
	if response {
		return models.NewErrorResponse(http.StatusBadRequest, "Model with this hash already exists")
	}
	modelVersion, err := datastore.UploadAndRegisterModelFile(orgId, modelBranchUUID, fileHeader, modelIsEmpty, modelHash, modelSourceType)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, modelVersion, "Model successfully registered")
}
