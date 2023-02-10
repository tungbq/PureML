package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/middlewares"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/PureML-Inc/PureML/server/tools/filesystem"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// BindModelBranchVersionApi registers the admin api endpoints and the corresponding handlers.
func BindModelBranchVersionApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	modelGroup := rg.Group("/org/:orgId/model", middlewares.RequireAuthContext, middlewares.ValidateOrg(api.app))
	modelGroup.POST("/:modelName/branch/:branchName/update", api.DefaultHandler(UpdateModelBranch), middlewares.ValidateModel(api.app), middlewares.ValidateModelBranch(api.app))
	modelGroup.POST("/:modelName/branch/:branchName/hash-status", api.DefaultHandler(VerifyModelBranchHashStatus), middlewares.ValidateModel(api.app))
	modelGroup.POST("/:modelName/branch/:branchName/register", api.DefaultHandler(RegisterModel), middlewares.ValidateModel(api.app), middlewares.ValidateModelBranch(api.app))
	modelGroup.DELETE("/:modelName/branch/:branchName/delete", api.DefaultHandler(DeleteModelBranch), middlewares.ValidateModel(api.app), middlewares.ValidateModelBranch(api.app))
	modelGroup.GET("/:modelName/branch/:branchName/version", api.DefaultHandler(GetModelBranchAllVersions), middlewares.ValidateModel(api.app), middlewares.ValidateModelBranch(api.app))
	modelGroup.GET("/:modelName/branch/:branchName/version/:version", api.DefaultHandler(GetModelBranchVersion), middlewares.ValidateModel(api.app), middlewares.ValidateModelBranch(api.app), middlewares.ValidateModelBranchVersion(api.app))
}

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
//	@Param			withLogs	query	bool	false	"Include logs"
func (api *Api) GetModelBranchAllVersions(request *models.Request) *models.Response {
	var response *models.Response
	branchUUID := request.GetModelBranchUUID()
	withLogs := strings.ToLower(request.GetQueryParam("withLogs")) == "true"
	allVersions, err := api.app.Dao().GetModelBranchAllVersions(branchUUID, withLogs)
	if err != nil {
		return models.NewServerErrorResponse(err)
	} else {
		response = models.NewDataResponse(http.StatusOK, allVersions, "Model branch all version details")
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
func (api *Api) GetModelBranchVersion(request *models.Request) *models.Response {
	branchUUID := request.GetModelBranchUUID()
	versionName := request.PathParams["version"]
	version, err := api.app.Dao().GetModelBranchVersion(branchUUID, versionName)
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
func (api *Api) VerifyModelBranchHashStatus(request *models.Request) *models.Response {
	modelName := request.GetModelName()
	modelBranchName := request.GetPathParam("branchName")
	orgId := uuid.Must(uuid.FromString(request.GetPathParam("orgId")))
	message := "Hash validity (False - does not exist in db)"
	model, err := api.app.Dao().GetModelBranchByName(orgId, modelName, modelBranchName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if model == nil {
		return models.NewDataResponse(http.StatusOK, false, message)
	}
	modelBranchUUID := model.UUID
	request.ParseJsonBody()
	hashValue := request.GetParsedBodyAttribute("hash").(string)
	versions, err := api.app.Dao().GetModelBranchAllVersions(modelBranchUUID, false)
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
func (api *Api) RegisterModel(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	userUUID := request.GetUserUUID()
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
	versions, err := api.app.Dao().GetModelBranchAllVersions(modelBranchUUID, false)
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
	if modelSourceType == "S3" && !api.app.Settings().S3.Enabled {
		return models.NewErrorResponse(http.StatusBadRequest, "S3 source not enabled")
	}
	sourceTypeUUID, err := api.app.Dao().GetSourceTypeByName(orgId, modelSourceType)
	if err != nil {
		return models.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Source %s not connected properly to organization", modelSourceType))
	}
	if sourceTypeUUID == uuid.Nil {
		if modelSourceType == "S3" && api.app.Settings().S3.Enabled {
			publicUrl := fmt.Sprintf("https://%s.%s", api.app.Settings().S3.Bucket, api.app.Settings().S3.Endpoint)
			sourceType, err := api.app.Dao().CreateS3Source(orgId, publicUrl)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			sourceTypeUUID = sourceType.UUID
		} else if modelSourceType == "LOCAL" {
			sourceType, err := api.app.Dao().CreateLocalSource(orgId)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			sourceTypeUUID = sourceType.UUID
		} else {
			return models.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Source %s not enabled", modelSourceType))
		}
	}
	var filePath string
	if !modelIsEmpty {
		file, err := filesystem.NewFileFromMultipart(fileHeader)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		filePath, err = api.app.UploadFile(file, "model-registry")
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
	}
	modelVersion, err := api.app.Dao().RegisterModelFile(modelBranchUUID, sourceTypeUUID, filePath, modelIsEmpty, modelHash, userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, modelVersion, "Model successfully registered")
}

var GetModelBranchAllVersions ServiceFunc = (*Api).GetModelBranchAllVersions
var GetModelBranchVersion ServiceFunc = (*Api).GetModelBranchVersion
var VerifyModelBranchHashStatus ServiceFunc = (*Api).VerifyModelBranchHashStatus
var RegisterModel ServiceFunc = (*Api).RegisterModel
