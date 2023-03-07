package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strings"

	authmiddlewares "github.com/PuremlHQ/PureML/packages/purebackend/auth/middlewares"
	orgmiddlewares "github.com/PuremlHQ/PureML/packages/purebackend/org/middlewares"
	modelmiddlewares "github.com/PuremlHQ/PureML/packages/purebackend/model/middlewares"
	datasetmiddlewares "github.com/PuremlHQ/PureML/packages/purebackend/dataset/middlewares"
	"github.com/PuremlHQ/PureML/packages/purebackend/core/models"
	orgmodels "github.com/PuremlHQ/PureML/packages/purebackend/org/models"
	datasetmodels "github.com/PuremlHQ/PureML/packages/purebackend/dataset/models"
	modelmodels "github.com/PuremlHQ/PureML/packages/purebackend/model/models"
	usermodels "github.com/PuremlHQ/PureML/packages/purebackend/user/models"
	"github.com/labstack/echo/v4"
)

func extractRequest(context echo.Context) *models.Request {
	request := &models.Request{}
	if context.Get(authmiddlewares.ContextAuthKey) != nil {
		request.User = context.Get(authmiddlewares.ContextAuthKey).(*usermodels.UserClaims)
	} else {
		request.User = &usermodels.UserClaims{}
	}
	if context.Get(orgmiddlewares.ContextOrgKey) != nil {
		request.Org = context.Get(orgmiddlewares.ContextOrgKey).(*orgmodels.OrganizationHandleResponse)
	} else {
		request.Org = &orgmodels.OrganizationHandleResponse{}
	}
	if context.Get(modelmiddlewares.ContextModelKey) != nil {
		request.Model = context.Get(modelmiddlewares.ContextModelKey).(*modelmodels.ModelNameResponse)
	} else {
		request.Model = &modelmodels.ModelNameResponse{}
	}
	if context.Get(modelmiddlewares.ContextModelBranchKey) != nil {
		request.ModelBranch = context.Get(modelmiddlewares.ContextModelBranchKey).(*modelmodels.ModelBranchNameResponse)
	} else {
		request.ModelBranch = &modelmodels.ModelBranchNameResponse{}
	}
	if context.Get(modelmiddlewares.ContextModelBranchVersionKey) != nil {
		request.ModelBranchVersion = context.Get(modelmiddlewares.ContextModelBranchVersionKey).(*modelmodels.ModelBranchVersionNameResponse)
	} else {
		request.ModelBranchVersion = &modelmodels.ModelBranchVersionNameResponse{}
	}
	if context.Get(datasetmiddlewares.ContextDatasetKey) != nil {
		request.Dataset = context.Get(datasetmiddlewares.ContextDatasetKey).(*datasetmodels.DatasetNameResponse)
	} else {
		request.Dataset = &datasetmodels.DatasetNameResponse{}
	}
	if context.Get(datasetmiddlewares.ContextDatasetBranchKey) != nil {
		request.DatasetBranch = context.Get(datasetmiddlewares.ContextDatasetBranchKey).(*datasetmodels.DatasetBranchNameResponse)
	} else {
		request.DatasetBranch = &datasetmodels.DatasetBranchNameResponse{}
	}
	if context.Get(datasetmiddlewares.ContextDatasetBranchVersionKey) != nil {
		request.DatasetBranchVersion = context.Get(datasetmiddlewares.ContextDatasetBranchVersionKey).(*datasetmodels.DatasetBranchVersionNameResponse)
	} else {
		request.DatasetBranchVersion = &datasetmodels.DatasetBranchVersionNameResponse{}
	}
	request.Headers = extractHeaders(context)
	request.PathParams = extractPathParams(context)
	request.QueryParams = extractQueryParams(context)
	// if content type is multipart formdata
	contentType := strings.Split(context.Request().Header.Get("Content-Type"), ";")[0]
	if contentType == "multipart/form-data" {
		request.FormValues, request.FormFiles = extractFormData(context)
	} else {
		request.Body = extractBody(context)
	}
	return request
}

func extractBody(context echo.Context) []byte {
	requestBody := context.Request().Body
	buffer := bytes.NewBuffer([]byte{})
	_, err := buffer.ReadFrom(requestBody)
	if err != nil {
		panic(err)
	}
	requestBody.Close()
	return buffer.Bytes()
}

func extractHeaders(context echo.Context) map[string]string {
	headers := map[string]string{}
	for k, v := range context.Request().Header {
		if len(v) <= 0 {
			continue
		}
		headers[k] = v[0]
	}
	return headers
}

func extractQueryParams(context echo.Context) map[string]string {
	queryParams := map[string]string{}
	for k, v := range context.QueryParams() {
		if len(v) <= 0 {
			continue
		}
		queryParams[k] = v[0]
	}
	return queryParams
}

func extractPathParams(context echo.Context) map[string]string {
	pathParams := map[string]string{}
	for _, pathParam := range context.ParamNames() {
		if _, ok := pathParams[pathParam]; ok {
			panic("Conflicting Param found")
		}
		pathParams[pathParam] = context.Param(pathParam)
	}
	return pathParams
}

func extractFormData(context echo.Context) (map[string][]string, map[string][]*multipart.FileHeader) {
	formData, err := context.MultipartForm()
	if err != nil {
		fmt.Println(err)
		panic("Could not process formdata for request")
	}
	if formData == nil {
		return map[string][]string{}, map[string][]*multipart.FileHeader{}
	}
	return formData.Value, formData.File
}

func convertToBytes(object interface{}) []byte {
	switch objectType := object.(type) {
	case string:
		return []byte(objectType)
	case []byte:
		return objectType
	default:
		bytes, err := json.Marshal(objectType)
		if err != nil {
			panic(err)
		}
		return bytes
	}
}
