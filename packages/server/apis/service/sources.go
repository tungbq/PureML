package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

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
func GetR2Secret(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	result, err := datastore.GetSourceSecret(orgId, "R2")
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "R2 secrets")
	return response
}

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
func GetS3Secret(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	result, err := datastore.GetSourceSecret(orgId, "S3")
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, result, "S3 secrets")
	return response
}

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
func ConnectR2Secret(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := request.GetOrgId()
	accountId := request.GetParsedBodyAttribute("account_id")
	if accountId == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Account Id not found in request body")
	} else if accountId.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Account Id cannot be empty")
	}
	accountIdData := accountId.(string)
	accessKeyId := request.GetParsedBodyAttribute("access_key_id")
	if accessKeyId == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Id not found in request body")
	} else if accessKeyId.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Id cannot be empty")
	}
	accessKeyIdData := accessKeyId.(string)
	accessKeySecret := request.GetParsedBodyAttribute("access_key_secret")
	if accessKeySecret == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Secret not found in request body")
	} else if accessKeySecret.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Secret cannot be empty")
	}
	accessKeySecretData := accessKeySecret.(string)
	bucketName := request.GetParsedBodyAttribute("bucket_name")
	if bucketName == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Bucket name not found in request body")
	} else if bucketName.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Bucket name cannot be empty")
	}
	bucketNameData := bucketName.(string)
	publicURL := request.GetParsedBodyAttribute("public_url")
	if publicURL == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Public URL not found in request body")
	} else if publicURL.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Public URL cannot be empty")
	}
	publicURLData := publicURL.(string)
	// Delete existing secrets
	err := datastore.DeleteR2Secrets(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	// Create new secrets
	createdR2Secret, err := datastore.CreateR2Secrets(orgId, accountIdData, accessKeyIdData, accessKeySecretData, bucketNameData, publicURLData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	err = createdR2Secret.CreateBucketIfNotExists()
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	createdSource, err := datastore.CreateR2Source(orgId, publicURLData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, createdSource, "R2 connected successfully")
}

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
func ConnectS3Secret(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := request.GetOrgId()
	accessKeyId := request.GetParsedBodyAttribute("access_key_id")
	if accessKeyId == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Id not found in request body")
	} else if accessKeyId.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Id cannot be empty")
	}
	accessKeyIdData := accessKeyId.(string)
	accessKeySecret := request.GetParsedBodyAttribute("access_key_secret")
	if accessKeySecret == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Secret not found in request body")
	} else if accessKeySecret.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Access Key Secret cannot be empty")
	}
	accessKeySecretData := accessKeySecret.(string)
	bucketName := request.GetParsedBodyAttribute("bucket_name")
	if bucketName == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Bucket name not found in request body")
	} else if bucketName.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Bucket name cannot be empty")
	}
	bucketNameData := bucketName.(string)
	bucketLocation := request.GetParsedBodyAttribute("bucket_location")
	if bucketLocation == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Public URL not found in request body")
	} else if bucketLocation.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Public URL cannot be empty")
	}
	bucketLocationData := bucketLocation.(string)
	// Delete existing secrets
	err := datastore.DeleteS3Secrets(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	// Create new secrets
	createdS3Secret, err := datastore.CreateS3Secrets(orgId, accessKeyIdData, accessKeySecretData, bucketNameData, bucketLocationData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	err = createdS3Secret.CreateBucketIfNotExists()
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	createdSource, err := datastore.CreateS3Source(orgId, createdS3Secret.PublicURL)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, createdSource, "S3 connected successfully")
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
func DeleteR2Secrets(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	err := datastore.DeleteR2Secrets(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, nil, "R2 disconnected")
	return response
}

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
func DeleteS3Secrets(request *models.Request) *models.Response {
	orgId := request.GetOrgId()
	err := datastore.DeleteS3Secrets(orgId)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, nil, "S3 disconnected")
	return response
}
