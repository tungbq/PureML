package tests

import (
	"net/http"
	"strings"
	"testing"

	coretests "github.com/PureMLHQ/PureML/packages/purebackend/core/apis/service/tests"
	"github.com/PureMLHQ/PureML/packages/purebackend/tests"
)

func TestGetModelBranchAllVersions(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "get specific version of model branch + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get all versions of model branch + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version",
			RequestHeaders: map[string]string{
				"Authorization": coretests.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get all versions of model branch + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "get all versions of model branch + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.InvalidOrgUuidString + "/model/Demo%20Model/branch/dev/version",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "get all versions of model branch + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidNoOrgUuid.String() + "/model/Demo%20Model/branch/dev/version",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "get all versions of model branch + valid token + model not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/NoModel/branch/dev/version",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model not found`,
			},
		},
		{
			Name:   "get all versions of model branch + valid token + branch not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/nobranch/version",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model Branch not found`,
			},
		},
		{
			Name:   "get all versions of model branch + valid token + model branch found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + coretests.ValidAdminUserUuid.String() + `"`,
				`"hash":"1234567890"`,
				`"version":"v1"`,
				`"branch":{`,
				`"name":"dev"`,
				`"path":{`,
				`"source_path":`,
				`"source_type":{`,
				`"created_by":{`,
				`"name":"Demo User"`,
				`"avatar":`,
				`"email":"demo@aztlan.in"`,
				`"handle":"demo"`,
				`"created_at":`,
				`"is_empty":true`,
				`"message":"All model branch versions"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGetModelBranchVersion(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "get specific version of model branch + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get specific version of model branch + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1",
			RequestHeaders: map[string]string{
				"Authorization": coretests.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get specific version of model branch + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "get specific version of model branch + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.InvalidOrgUuidString + "/model/Demo%20Model/branch/dev/version/v1",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "get specific version of model branch + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidNoOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "get specific version of model branch + valid token + model not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/NoModel/branch/dev/version/v1",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model not found`,
			},
		},
		{
			Name:   "get specific version of model branch + valid token + branch not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/nobranch/version",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model Branch not found`,
			},
		},
		{
			Name:   "get specific version of model branch + valid token + version not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v2",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model Branch Version not found`,
			},
		},
		{
			Name:   "get specific version of model branch + valid token + model branch version found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + coretests.ValidAdminUserUuid.String() + `"`,
				`"hash":"1234567890"`,
				`"version":"v1"`,
				`"branch":{`,
				`"name":"dev"`,
				`"path":{`,
				`"source_path":`,
				`"source_type":{`,
				`"created_by":{`,
				`"name":"Demo User"`,
				`"avatar":`,
				`"email":"demo@aztlan.in"`,
				`"handle":"demo"`,
				`"created_at":`,
				`"is_empty":true`,
				`"message":"Model branch version details"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestVerifyModelBranchHashStatus(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "verify model hash status + unauthorized",
			Method:         http.MethodPost,
			Url:            "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/hash-status",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "verify model hash status + invalid token",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/hash-status",
			RequestHeaders: map[string]string{
				"Authorization": coretests.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "verify model hash status + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.InvalidOrgUuidString + "/model/Demo%20Model/branch/dev/hash-status",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "verify model hash status + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidNoOrgUuid.String() + "/model/Demo%20Model/branch/dev/hash-status",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "verify model hash status + valid token + model not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/noModel/branch/dev/hash-status",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model not found`,
			},
		},
		{
			Name:   "verify model hash status + valid token + hash empty",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/hash-status",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"hash":""
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Hash value is empty"`,
			},
		},
		{
			Name:   "verify model hash status + valid token + branch not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/nobranch/hash-status",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"hash":"test"
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[false]`,
				`"message":"Hash validity (False - does not exist in db)"`,
			},
		},
		{
			Name:   "verify model hash status + valid token + hash does not exist",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/hash-status",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"hash":"test"
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[false]`,
				`"message":"Hash validity (False - does not exist in db)"`,
			},
		},
		{
			Name:   "verify model hash status + valid token + hash exists",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/hash-status",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"hash":"1234567890"
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[true]`,
				`"message":"Hash validity (True - exists in db)"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestRegisterModel(t *testing.T) {
	emptyHashMultipartBody, emptyHashMultipartContentType, err := tests.MockMultipartData(map[string]string{
		"hash":     "",
		"storage":  "local",
		"is_empty": "true",
	}, "file")
	if err != nil {
		t.Fatal(err)
	}
	emptyFileMultipartBody, emptyFileMultipartContentType, err := tests.MockMultipartData(map[string]string{
		"hash":     "uniquehash",
		"storage":  "local",
		"is_empty": "true",
	})
	if err != nil {
		t.Fatal(err)
	}
	invalidStorageMultipartBody, invalidStorageMultipartContentType, err := tests.MockMultipartData(map[string]string{
		"hash":     "uniquehash",
		"storage":  "nope",
		"is_empty": "true",
	}, "file")
	if err != nil {
		t.Fatal(err)
	}
	invalidHashMultipartBody, invalidHashMultipartContentType, err := tests.MockMultipartData(map[string]string{
		"hash":     "1234567890",
		"storage":  "local",
		"is_empty": "true",
	}, "file")
	if err != nil {
		t.Fatal(err)
	}
	validMultipartBody, validMultipartContentType, err := tests.MockMultipartData(map[string]string{
		"hash":     "uniquehash",
		"storage":  "local",
		"is_empty": "true",
	}, "file")
	if err != nil {
		t.Fatal(err)
	}
	validProtectedMultipartBody, validProtectedMultipartContentType, err := tests.MockMultipartData(map[string]string{
		"hash":     "uniquehash",
		"storage":  "local",
		"is_empty": "true",
	}, "file")
	if err != nil {
		t.Fatal(err)
	}
	validS3MultipartBody, validS3MultipartContentType, err := tests.MockMultipartData(map[string]string{
		"hash":     "uniquehash",
		"storage":  "s3",
		"is_empty": "true",
	}, "file")
	if err != nil {
		t.Fatal(err)
	}
	validR2MultipartBody, validR2MultipartContentType, err := tests.MockMultipartData(map[string]string{
		"hash":     "uniquehash",
		"storage":  "r2",
		"is_empty": "true",
	}, "file")
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []tests.ApiScenario{
		{
			Name:           "register model + unauthorized",
			Method:         http.MethodPost,
			Url:            "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/register",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "register model + invalid token",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "register model + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.InvalidOrgUuidString + "/model/Demo%20Model/branch/dev/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "register model + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidNoOrgUuid.String() + "/model/Demo%20Model/branch/dev/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "register model + valid token + model not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/noModel/branch/dev/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model not found`,
			},
		},
		{
			Name:   "register model + valid token + branch not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/nobranch/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model Branch not found`,
			},
		},
		{
			Name:   "register model + valid token + hash empty",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
				"Content-Type":  emptyHashMultipartContentType.FormDataContentType(),
			},
			Body:           emptyHashMultipartBody,
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Hash is required"`,
			},
		},
		{
			Name:   "register model + valid token + file empty",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
				"Content-Type":  emptyFileMultipartContentType.FormDataContentType(),
			},
			Body:           emptyFileMultipartBody,
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"File is required"`,
			},
		},
		{
			Name:   "register model + valid token + protected main branch",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/main/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
				"Content-Type":  validProtectedMultipartContentType.FormDataContentType(),
			},
			Body:           validProtectedMultipartBody,
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Cannot register model directly to main branch"`,
			},
		},
		{
			Name:   "register model + valid token + unsupported storage",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
				"Content-Type":  invalidStorageMultipartContentType.FormDataContentType(),
			},
			Body:           invalidStorageMultipartBody,
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Unsupported model storage"`,
			},
		},
		{
			Name:   "register model + valid token + hash already exists",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
				"Content-Type":  invalidHashMultipartContentType.FormDataContentType(),
			},
			Body:           invalidHashMultipartBody,
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Model with this hash already exists"`,
			},
		},
		{
			Name:   "register model + valid token + s3 not enabled",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
				"Content-Type":  validS3MultipartContentType.FormDataContentType(),
			},
			Body:           validS3MultipartBody,
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"S3 source not enabled"`,
			},
		},
		{
			Name:   "register model + valid token + r2 not enabled",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
				"Content-Type":  validR2MultipartContentType.FormDataContentType(),
			},
			Body:           validR2MultipartBody,
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"R2 source not enabled"`,
			},
		},
		{
			Name:   "register model + valid token + registered successfully",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/register",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
				"Content-Type":  validMultipartContentType.FormDataContentType(),
			},
			Body:           validMultipartBody,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"`,
				`"hash":"uniquehash"`,
				`"version":"v2"`,
				`"branch":{`,
				`"name":"dev"`,
				`"path":{`,
				`"source_path":`,
				`"source_type":{`,
				`"created_by":{`,
				`"name":"Demo User"`,
				`"avatar":`,
				`"email":"demo@aztlan.in"`,
				`"handle":"demo"`,
				`"created_at":`,
				`"is_empty":true`,
				`"message":"Model successfully registered"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}