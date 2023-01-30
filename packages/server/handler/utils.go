package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/PureML-Inc/PureML/server/models"
	"github.com/labstack/echo/v4"
)

func extractRequest(context echo.Context) *models.Request {
	request := &models.Request{}
	if context.Get("User") != nil {
		request.User = context.Get("User").(*models.UserClaims)
	} else {
		request.User = &models.UserClaims{}
	}
	if context.Get("Org") != nil {
		request.Org = context.Get("Org").(*models.OrganizationHandleResponse)
	} else {
		request.Org = &models.OrganizationHandleResponse{}
	}
	if context.Get("Model") != nil {
		request.Model = context.Get("Model").(*models.ModelNameResponse)
	} else {
		request.Model = &models.ModelNameResponse{}
	}
	if context.Get("ModelBranch") != nil {
		request.ModelBranch = context.Get("ModelBranch").(*models.ModelBranchNameResponse)
	} else {
		request.ModelBranch = &models.ModelBranchNameResponse{}
	}
	if context.Get("ModelBranchVersion") != nil {
		request.ModelBranchVersion = context.Get("ModelBranchVersion").(*models.ModelBranchVersionNameResponse)
	} else {
		request.ModelBranchVersion = &models.ModelBranchVersionNameResponse{}
	}
	if context.Get("Dataset") != nil {
		request.Dataset = context.Get("Dataset").(*models.DatasetNameResponse)
	} else {
		request.Dataset = &models.DatasetNameResponse{}
	}
	if context.Get("DatasetBranch") != nil {
		request.DatasetBranch = context.Get("DatasetBranch").(*models.DatasetBranchNameResponse)
	} else {
		request.DatasetBranch = &models.DatasetBranchNameResponse{}
	}
	if context.Get("DatasetBranchVersion") != nil {
		request.DatasetBranchVersion = context.Get("DatasetBranchVersion").(*models.DatasetBranchVersionNameResponse)
	} else {
		request.DatasetBranchVersion = &models.DatasetBranchVersionNameResponse{}
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
