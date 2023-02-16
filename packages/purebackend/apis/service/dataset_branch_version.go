package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PureML-Inc/PureML/purebackend/core"
	"github.com/PureML-Inc/PureML/purebackend/middlewares"
	"github.com/PureML-Inc/PureML/purebackend/models"
	"github.com/PureML-Inc/PureML/purebackend/tools/filesystem"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

// BindDatasetBranchVersionApi registers the admin api endpoints and the corresponding handlers.
func BindDatasetBranchVersionApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	datasetGroup := rg.Group("/org/:orgId/dataset", middlewares.RequireAuthContext, middlewares.ValidateOrg(api.app))
	datasetGroup.POST("/:datasetName/branch/:branchName/hash-status", api.DefaultHandler(VerifyDatasetBranchHashStatus), middlewares.ValidateDataset(api.app))
	datasetGroup.POST("/:datasetName/branch/:branchName/register", api.DefaultHandler(RegisterDataset), middlewares.ValidateDataset(api.app), middlewares.ValidateDatasetBranch(api.app))
	datasetGroup.GET("/:datasetName/branch/:branchName/version", api.DefaultHandler(GetDatasetBranchAllVersions), middlewares.ValidateDataset(api.app), middlewares.ValidateDatasetBranch(api.app))
	datasetGroup.GET("/:datasetName/branch/:branchName/version/:version", api.DefaultHandler(GetDatasetBranchVersion), middlewares.ValidateDataset(api.app), middlewares.ValidateDatasetBranch(api.app), middlewares.ValidateDatasetBranchVersion(api.app))
}

// GetDatasetBranchAllVersions godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get all branch versions of a dataset
//	@Description	Get all branch versions of a dataset
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/{branchName}/version [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
//	@Param			branchName	path	string	true	"Branch Name"
func (api *Api) GetDatasetBranchAllVersions(request *models.Request) *models.Response {
	var response *models.Response
	branchUUID := request.GetDatasetBranchUUID()
	allVersions, err := api.app.Dao().GetDatasetBranchAllVersions(branchUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	} else {
		response = models.NewDataResponse(http.StatusOK, allVersions, "All dataset branch versions")
	}
	return response
}

// GetDatasetBranchVersion godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get specific branch version of a dataset
//	@Description	Get specific branch version of a dataset
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/{branchName}/version/{version} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
//	@Param			branchName	path	string	true	"Branch Name"
//	@Param			version		path	string	true	"Version"
func (api *Api) GetDatasetBranchVersion(request *models.Request) *models.Response {
	branchUUID := request.GetDatasetBranchUUID()
	versionName := request.PathParams["version"]
	version, err := api.app.Dao().GetDatasetBranchVersion(branchUUID, versionName)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if version == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Version not found")
	}
	return models.NewDataResponse(http.StatusOK, version, "Dataset branch version details")
}

// VerifyDatasetBranchHashStatus godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Verify dataset hash status
//	@Description	Verify dataset hash status to determine if dataset is already uploaded
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/{branchName}/hash-status [post]
//	@Param			orgId		path	string				true	"Organization Id"
//	@Param			datasetName	path	string				true	"Dataset Name"
//	@Param			branchName	path	string				true	"Branch Name"
//	@Param			hash		body	models.HashRequest	true	"Hash value"
func (api *Api) VerifyDatasetBranchHashStatus(request *models.Request) *models.Response {
	datasetName := request.GetDatasetName()
	datasetBranchName := request.GetPathParam("branchName")
	orgId := uuid.Must(uuid.FromString(request.GetPathParam("orgId")))
	message := "Hash validity (False - does not exist in db)"
	dataset, err := api.app.Dao().GetDatasetBranchByName(orgId, datasetName, datasetBranchName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if dataset == nil {
		return models.NewDataResponse(http.StatusOK, false, message)
	}
	datasetBranchUUID := dataset.UUID
	request.ParseJsonBody()
	hashValue := request.GetParsedBodyAttribute("hash").(string)
	versions, err := api.app.Dao().GetDatasetBranchAllVersions(datasetBranchUUID)
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

// RegisterDataset godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Register dataset
//	@Description	Register dataset file. Create dataset and default branches if not exists
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/{branchName}/register [post]
//	@Param			file		formData	file							true	"Dataset file"
//	@Param			orgId		path		string							true	"Organization UUID"
//	@Param			datasetName	path		string							true	"Dataset name"
//	@Param			branchName	path		string							true	"Branch name"
//	@Param			data		formData	models.RegisterDatasetRequest	true	"Dataset details"
func (api *Api) RegisterDataset(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	userUUID := request.GetUserUUID()
	var datasetHash string
	if request.FormValues["hash"] != nil && len(request.FormValues["hash"]) > 0 {
		datasetHash = request.FormValues["hash"][0]
	} else {
		return models.NewErrorResponse(http.StatusBadRequest, "Hash is required")
	}
	var datasetSourceType string
	if request.FormValues["storage"] != nil && len(request.FormValues["storage"]) > 0 {
		datasetSourceType = strings.ToUpper(request.FormValues["storage"][0])
		if datasetSourceType == "PUREML-STORAGE" {
			datasetSourceType = "R2"
		}
	}
	var datasetIsEmpty bool
	if request.FormValues["is_empty"] != nil && len(request.FormValues["is_empty"]) > 0 {
		datasetIsEmpty = request.FormValues["is_empty"][0] == "true"
	}
	var datasetLineage string
	if request.FormValues["lineage"] != nil && len(request.FormValues["lineage"]) > 0 {
		datasetLineage = request.FormValues["lineage"][0]
	}
	fileHeader := request.GetFormFile("file")
	if fileHeader == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "File is required")
	}
	datasetBranchName := request.GetPathParam("branchName")
	if datasetBranchName == "main" {
		return models.NewErrorResponse(http.StatusBadRequest, "Cannot register model directly to main branch")
	}
	sourceValid := false
	for source := range models.SupportedSources {
		if models.SupportedSources[source] == datasetSourceType {
			sourceValid = true
			break
		}
	}
	if !sourceValid {
		return models.NewErrorResponse(http.StatusBadRequest, "Unsupported model source type")
	}
	datasetBranchUUID := request.GetDatasetBranchUUID()
	versions, err := api.app.Dao().GetDatasetBranchAllVersions(datasetBranchUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := false
	for _, version := range versions {
		if version.Hash == datasetHash {
			response = true
			break
		}
	}
	if response {
		return models.NewErrorResponse(http.StatusBadRequest, "Dataset with this hash already exists")
	}
	if datasetSourceType == "S3" && !api.app.Settings().S3.Enabled {
		return models.NewErrorResponse(http.StatusBadRequest, "S3 source not enabled")
	}
	if datasetSourceType == "R2" && !api.app.Settings().R2.Enabled {
		return models.NewErrorResponse(http.StatusBadRequest, "R2 source not enabled")
	}
	sourceTypeUUID, err := api.app.Dao().GetSourceTypeByName(orgId, datasetSourceType)
	if err != nil {
		return models.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Source %s not connected properly to organization", datasetSourceType))
	}
	if sourceTypeUUID == uuid.Nil {
		if api.app.Settings().S3.Enabled {
			publicUrl := fmt.Sprintf("https://%s.%s", api.app.Settings().S3.Bucket, api.app.Settings().S3.Endpoint)
			sourceType, err := api.app.Dao().CreateS3Source(orgId, publicUrl)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			sourceTypeUUID = sourceType.UUID
		} else if datasetSourceType == "R2" && api.app.Settings().R2.Enabled {
			publicUrl := fmt.Sprintf("https://%s/%s", api.app.Settings().R2.Endpoint, api.app.Settings().R2.Bucket)
			sourceType, err := api.app.Dao().CreateR2Source(orgId, publicUrl)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			sourceTypeUUID = sourceType.UUID
		} else if datasetSourceType == "LOCAL" {
			sourceType, err := api.app.Dao().CreateLocalSource(orgId)
			if err != nil {
				return models.NewServerErrorResponse(err)
			}
			sourceTypeUUID = sourceType.UUID
		} else {
			return models.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Source %s not enabled", datasetSourceType))
		}
	}
	var filePath string
	if !datasetIsEmpty {
		file, err := filesystem.NewFileFromMultipart(fileHeader)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		filePath, err = api.app.UploadFile(file, "dataset-registry")
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
	}
	datasetVersion, err := api.app.Dao().RegisterDatasetFile(datasetBranchUUID, sourceTypeUUID, filePath, datasetIsEmpty, datasetHash, datasetLineage, userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, datasetVersion, "Dataset successfully registered")
}

var GetDatasetBranchAllVersions ServiceFunc = (*Api).GetDatasetBranchAllVersions
var GetDatasetBranchVersion ServiceFunc = (*Api).GetDatasetBranchVersion
var VerifyDatasetBranchHashStatus ServiceFunc = (*Api).VerifyDatasetBranchHashStatus
var RegisterDataset ServiceFunc = (*Api).RegisterDataset
