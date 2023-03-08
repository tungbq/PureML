package service

import (
	"fmt"
	"net/http"

	authmiddlewares "github.com/PureMLHQ/PureML/packages/purebackend/auth/middlewares"
	"github.com/PureMLHQ/PureML/packages/purebackend/core"
	"github.com/PureMLHQ/PureML/packages/purebackend/core/models"
	"github.com/PureMLHQ/PureML/packages/purebackend/core/tools/security"
	orgmiddlewares "github.com/PureMLHQ/PureML/packages/purebackend/user_org/middlewares"
	"github.com/labstack/echo/v4"
)

// BindSecretsApi registers the admin api endpoints and the corresponding handlers.
func BindSecretsApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	secretGroup := rg.Group("/org/:orgId/secret", authmiddlewares.RequireAuthContext, orgmiddlewares.ValidateOrg(api.app))
	// secretGroup.GET("/all", api.DefaultHandler(GetAllSecrets))
	// secretGroup.GET("/r2", api.DefaultHandler(GetR2Secret))
	// secretGroup.POST("/r2/connect", api.DefaultHandler(ConnectR2Secret))
	secretGroup.GET("/r2/test", api.DefaultHandler(TestR2Secret))
	// secretGroup.DELETE("/r2/delete", api.DefaultHandler(DeleteR2Secrets))
	// secretGroup.GET("/s3", api.DefaultHandler(GetS3Secret))
	// secretGroup.POST("/s3/connect", api.DefaultHandler(ConnectS3Secret))
	secretGroup.GET("/s3/test", api.DefaultHandler(TestS3Secret))
	// secretGroup.DELETE("/s3/delete", api.DefaultHandler(DeleteS3Secrets))
}

// GetR2Secret godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get secrets for source type r2
//	@Description	Get secrets for source type r2
//	@Tags			Secret
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/secret/r2 [get]
//	@Param			orgId	path	string	true	"Organization Id"
// func (api *Api) GetR2Secret(request *models.Request) *models.Response {
// 	orgId := request.GetOrgId()
// 	result, err := api.app.Dao().GetSourceSecret(orgId, "R2")
// 	if err != nil {
// 		return models.NewServerErrorResponse(err)
// 	}
// 	response := models.NewDataResponse(http.StatusOK, result, "R2 secrets")
// 	return response
// }

// GetS3Secret godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get secrets for source type s3
//	@Description	Get secrets for source type s3
//	@Tags			Secret
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/secret/s3 [get]
//	@Param			orgId	path	string	true	"Organization Id"
// func (api *Api) GetS3Secret(request *models.Request) *models.Response {
// 	orgId := request.GetOrgId()
// 	result, err := api.app.Dao().GetSourceSecret(orgId, "S3")
// 	if err != nil {
// 		return models.NewServerErrorResponse(err)
// 	}
// 	response := models.NewDataResponse(http.StatusOK, result, "S3 secrets")
// 	return response
// }

// ConnectR2Secret godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Add secrets for source type r2
//	@Description	Add secrets for source type r2
//	@Tags			Secret
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/secret/r2/connect [post]
//	@Param			orgId	path	string					true	"Organization Id"
//	@Param			secret	body	models.R2SecretRequest	true	"Secret"
// func (api *Api) ConnectR2Secret(request *models.Request) *models.Response {
// 	request.ParseJsonBody()
// 	orgId := request.GetOrgId()
// 	accountId := request.GetParsedBodyAttribute("account_id")
// 	if accountId == nil {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Account Id not found in request body")
// 	} else if accountId.(string) == "" {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Account Id cannot be empty")
// 	}
// 	accountIdData := accountId.(string)
// 	accessKeyId := request.GetParsedBodyAttribute("access_key_id")
// 	if accessKeyId == nil {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Id not found in request body")
// 	} else if accessKeyId.(string) == "" {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Id cannot be empty")
// 	}
// 	accessKeyIdData := accessKeyId.(string)
// 	accessKeySecret := request.GetParsedBodyAttribute("access_key_secret")
// 	if accessKeySecret == nil {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Secret not found in request body")
// 	} else if accessKeySecret.(string) == "" {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Secret cannot be empty")
// 	}
// 	accessKeySecretData := accessKeySecret.(string)
// 	bucketName := request.GetParsedBodyAttribute("bucket_name")
// 	if bucketName == nil {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Bucket name not found in request body")
// 	} else if bucketName.(string) == "" {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Bucket name cannot be empty")
// 	}
// 	bucketNameData := bucketName.(string)
// 	publicURL := request.GetParsedBodyAttribute("public_url")
// 	if publicURL == nil {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Public URL not found in request body")
// 	} else if publicURL.(string) == "" {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Public URL cannot be empty")
// 	}
// 	publicURLData := publicURL.(string)
// 	// Delete existing secrets
// 	err := api.app.Dao().DeleteR2Secrets(orgId)
// 	if err != nil {
// 		return models.NewServerErrorResponse(err)
// 	}
// 	// Create new secrets
// 	createdR2Secret, err := api.app.Dao().CreateR2Secrets(orgId, accountIdData, accessKeyIdData, accessKeySecretData, bucketNameData, publicURLData)
// 	if err != nil {
// 		return models.NewServerErrorResponse(err)
// 	}
// 	err = createdR2Secret.CreateBucketIfNotExists()
// 	if err != nil {
// 		return models.NewServerErrorResponse(err)
// 	}
// 	createdSource, err := api.app.Dao().CreateR2Source(orgId, publicURLData)
// 	if err != nil {
// 		return models.NewServerErrorResponse(err)
// 	}
// 	return models.NewDataResponse(http.StatusOK, createdSource, "R2 connected successfully")
// }

// ConnectS3Secret godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Add secrets for source type s3
//	@Description	Add secrets for source type s3
//	@Tags			Secret
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/secret/s3/connect [post]
//	@Param			orgId	path	string					true	"Organization Id"
//	@Param			secret	body	models.S3SecretRequest	true	"Secret"
// func (api *Api) ConnectS3Secret(request *models.Request) *models.Response {
// 	request.ParseJsonBody()
// 	orgId := request.GetOrgId()
// 	accessKeyId := request.GetParsedBodyAttribute("access_key_id")
// 	if accessKeyId == nil {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Id not found in request body")
// 	} else if accessKeyId.(string) == "" {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Id cannot be empty")
// 	}
// 	accessKeyIdData := accessKeyId.(string)
// 	accessKeySecret := request.GetParsedBodyAttribute("access_key_secret")
// 	if accessKeySecret == nil {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Secret not found in request body")
// 	} else if accessKeySecret.(string) == "" {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Secret cannot be empty")
// 	}
// 	accessKeySecretData := accessKeySecret.(string)
// 	bucketName := request.GetParsedBodyAttribute("bucket_name")
// 	if bucketName == nil {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Bucket name not found in request body")
// 	} else if bucketName.(string) == "" {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Bucket name cannot be empty")
// 	}
// 	bucketNameData := bucketName.(string)
// 	bucketLocation := request.GetParsedBodyAttribute("bucket_location")
// 	if bucketLocation == nil {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Public URL not found in request body")
// 	} else if bucketLocation.(string) == "" {
// 		return models.NewErrorResponse(http.StatusBadRequest, "Public URL cannot be empty")
// 	}
// 	bucketLocationData := bucketLocation.(string)
// 	// Delete existing secrets
// 	err := api.app.Dao().DeleteS3Secrets(orgId)
// 	if err != nil {
// 		return models.NewServerErrorResponse(err)
// 	}
// 	// Create new secrets
// 	createdS3Secret, err := api.app.Dao().CreateS3Secrets(orgId, accessKeyIdData, accessKeySecretData, bucketNameData, bucketLocationData)
// 	if err != nil {
// 		return models.NewServerErrorResponse(err)
// 	}
// 	err = createdS3Secret.CreateBucketIfNotExists()
// 	if err != nil {
// 		return models.NewServerErrorResponse(err)
// 	}
// 	createdSource, err := api.app.Dao().CreateS3Source(orgId, createdS3Secret.PublicURL)
// 	if err != nil {
// 		return models.NewServerErrorResponse(err)
// 	}
// 	return models.NewDataResponse(http.StatusOK, createdSource, "S3 connected successfully")
// }

// TestR2Secret godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Test secrets for source type r2
//	@Description	Test secrets for source type r2
//	@Tags			Secret
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/secret/r2/test [get]
//	@Param			orgId	path	string	true	"Organization Id"
func (api *Api) TestR2Secret(request *models.Request) *models.Response {
	if !api.app.Settings().R2.Enabled {
		return models.NewErrorResponse(http.StatusBadRequest, "R2 storage is not enabled.")
	}

	fs, err := api.app.NewFilesystem()
	if err != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Failed to initialize the R2 storage. Raw error: \n"+err.Error())
	}
	defer fs.Close()

	testPrefix := "pureml_settings_test_" + security.PseudorandomString(5)
	testFileKey := testPrefix + "/test.txt"

	// try to upload a test file
	if err := fs.Upload([]byte("test"), testFileKey); err != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Failed to upload a test file. Raw error: \n"+err.Error())
	}

	// test prefix deletion (ensures that both bucket list and delete works)
	if errs := fs.DeletePrefix(testPrefix); len(errs) > 0 {
		return models.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Failed to delete a test file. Raw error: %v", errs))
	}

	return models.NewDataResponse(http.StatusOK, nil, "R2 connected successfully")
}

// TestS3Secret godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Test secrets for source type s3
//	@Description	Test secrets for source type s3
//	@Tags			Secret
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/secret/s3/test [get]
//	@Param			orgId	path	string	true	"Organization Id"
func (api *Api) TestS3Secret(request *models.Request) *models.Response {
	if !api.app.Settings().S3.Enabled {
		return models.NewErrorResponse(http.StatusBadRequest, "S3 storage is not enabled.")
	}

	fs, err := api.app.NewFilesystem()
	if err != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Failed to initialize the S3 storage. Raw error: \n"+err.Error())
	}
	defer fs.Close()

	testPrefix := "pureml_settings_test_" + security.PseudorandomString(5)
	testFileKey := testPrefix + "/test.txt"

	// try to upload a test file
	if err := fs.Upload([]byte("test"), testFileKey); err != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Failed to upload a test file. Raw error: \n"+err.Error())
	}

	// test prefix deletion (ensures that both bucket list and delete works)
	if errs := fs.DeletePrefix(testPrefix); len(errs) > 0 {
		return models.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Failed to delete a test file. Raw error: %v", errs))
	}

	return models.NewDataResponse(http.StatusOK, nil, "S3 connected successfully")
}

// DeleteR2Secrets godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Delete secrets for source type r2
//	@Description	Delete secrets for source type r2
//	@Tags			Secret
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/secret/r2/delete [delete]
//	@Param			orgId	path	string	true	"Organization Id"
// func (api *Api) DeleteR2Secrets(request *models.Request) *models.Response {
// 	orgId := request.GetOrgId()
// 	err := api.app.Dao().DeleteR2Secrets(orgId)
// 	if err != nil {
// 		return models.NewServerErrorResponse(err)
// 	}
// 	response := models.NewDataResponse(http.StatusOK, nil, "R2 disconnected")
// 	return response
// }

// DeleteS3Secrets godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Delete secrets for source type s3
//	@Description	Delete secrets for source type s3
//	@Tags			Secret
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/secret/s3/delete [delete]
//	@Param			orgId	path	string	true	"Organization Id"
// func (api *Api) DeleteS3Secrets(request *models.Request) *models.Response {
// 	orgId := request.GetOrgId()
// 	err := api.app.Dao().DeleteS3Secrets(orgId)
// 	if err != nil {
// 		return models.NewServerErrorResponse(err)
// 	}
// 	response := models.NewDataResponse(http.StatusOK, nil, "S3 disconnected")
// 	return response
// }

// var GetR2Secret ServiceFunc = (*Api).GetR2Secret
// var GetS3Secret ServiceFunc = (*Api).GetS3Secret
// var ConnectR2Secret ServiceFunc = (*Api).ConnectR2Secret
// var ConnectS3Secret ServiceFunc = (*Api).ConnectS3Secret
var TestR2Secret ServiceFunc = (*Api).TestR2Secret
var TestS3Secret ServiceFunc = (*Api).TestS3Secret

// var DeleteR2Secrets ServiceFunc = (*Api).DeleteR2Secrets
// var DeleteS3Secrets ServiceFunc = (*Api).DeleteS3Secrets
