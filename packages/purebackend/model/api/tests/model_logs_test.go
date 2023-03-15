package tests

import (
	"net/http"
	"strings"
	"testing"

	"github.com/PureMLHQ/PureML/packages/purebackend/test"
	"github.com/labstack/echo/v4"
)

func TestGetAllLogsModel(t *testing.T) {
	scenarios := []test.ApiScenario{
		{
			Name:           "get all logs of model + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get all logs of model + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get all logs of model + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "get all logs of model + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.InvalidOrgUuidString + "/model/Demo%20Model/branch/dev/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "get all logs of model + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidNoOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "get all logs of model + valid token + model not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/NoModel/branch/dev/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model not found`,
			},
		},
		{
			Name:   "get all logs of model + valid token + model branch not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/nobranch/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model Branch not found`,
			},
		},
		{
			Name:   "get all logs of model + valid token + model branch version not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v2/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model Branch Version not found`,
			},
		},
		{
			Name:   "get all logs of model + valid token + model branch version found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			BeforeTestFunc: func(t *testing.T, app *test.TestApp, e *echo.Echo) {
				// Make notadmin a "member" of the admin user org
				_, err := app.Dao().CreateLogForModelVersion("accuracy", "accuracyData", test.ValidAdminUserOrgUuid)
				if err != nil {
					t.Fatal(err)
				}
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"key":"accuracy"`,
				`"data":"accuracyData"`,
				`"model_version":{`,
				`"uuid":"` + test.ValidAdminUserUuid.String() + `"`,
				`"version":"v1"`,
				`"message":"Logs for model version"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGetKeyLogsModel(t *testing.T) {
	scenarios := []test.ApiScenario{
		{
			Name:           "get log by key of model + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log/accuracy",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get log by key of model + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log/accuracy",
			RequestHeaders: map[string]string{
				"Authorization": test.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get log by key of model + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log/accuracy",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "get log by key of model + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.InvalidOrgUuidString + "/model/Demo%20Model/branch/dev/version/v1/log/accuracy",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "get log by key of model + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidNoOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log/accuracy",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "get log by key of model + valid token + model not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/NoModel/branch/dev/version/v1/log/accuracy",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model not found`,
			},
		},
		{
			Name:   "get log by key of model + valid token + model branch not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/nobranch/version/v1/log/accuracy",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model Branch not found`,
			},
		},
		{
			Name:   "get log by key of model + valid token + model branch version not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v2/log/accuracy",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model Branch Version not found`,
			},
		},
		{
			Name:   "get log by key of model + valid token + model branch version found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log/accuracy",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			BeforeTestFunc: func(t *testing.T, app *test.TestApp, e *echo.Echo) {
				// Make notadmin a "member" of the admin user org
				_, err := app.Dao().CreateLogForModelVersion("accuracy", "accuracyData", test.ValidAdminUserOrgUuid)
				if err != nil {
					t.Fatal(err)
				}
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"key":"accuracy"`,
				`"data":"accuracyData"`,
				`"model_version":{`,
				`"uuid":"` + test.ValidAdminUserUuid.String() + `"`,
				`"version":"v1"`,
				`"message":"Specific Key Logs for model version"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestLogModel(t *testing.T) {
	scenarios := []test.ApiScenario{
		{
			Name:           "create model branch version log + unauthorized",
			Method:         http.MethodPost,
			Url:            "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "create model branch version log + invalid token",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "create model branch version log + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.InvalidOrgUuidString + "/model/Demo%20Model/branch/dev/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "create model branch version log + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidNoOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "create model branch version log + valid token + model not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/NoModel/branch/dev/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model not found`,
			},
		},
		{
			Name:   "create model branch version log + valid token + model branch not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/nobranch/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model Branch not found`,
			},
		},
		{
			Name:   "create model branch version log + valid token + model branch version not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v2/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model Branch Version not found`,
			},
		},
		{
			Name:   "create model branch version log + valid token + key empty",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"key":"",
				"data":""
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Key is required"`,
			},
		},
		{
			Name:   "create model branch version log + valid token + log created successfully",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/log",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"key":"test",
				"data":"test"
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"key":"test"`,
				`"data":"test"`,
				`"message":"Log created"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestLogFileModel(t *testing.T) {
	emptyFileMultipartBody, emptyFileMultipartContentType, err := test.MockMultipartData(map[string]string{
		"storage": "local",
	})
	if err != nil {
		t.Fatal(err)
	}
	invalidStorageMultipartBody, invalidStorageMultipartContentType, err := test.MockMultipartData(map[string]string{
		"storage": "nope",
	}, "file")
	if err != nil {
		t.Fatal(err)
	}
	validMultipartBody, validMultipartContentType, err := test.MockMultipartData(map[string]string{
		"storage": "local",
	}, "file")
	if err != nil {
		t.Fatal(err)
	}
	validS3MultipartBody, validS3MultipartContentType, err := test.MockMultipartData(map[string]string{
		"storage": "s3",
	}, "file")
	if err != nil {
		t.Fatal(err)
	}
	validR2MultipartBody, validR2MultipartContentType, err := test.MockMultipartData(map[string]string{
		"storage": "r2",
	}, "file")
	if err != nil {
		t.Fatal(err)
	}
	scenarios := []test.ApiScenario{
		{
			Name:           "create model branch version log file + unauthorized",
			Method:         http.MethodPost,
			Url:            "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/logfile",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "create model branch version log file + invalid token",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/logfile",
			RequestHeaders: map[string]string{
				"Authorization": test.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "create model branch version log file + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.InvalidOrgUuidString + "/model/Demo%20Model/branch/dev/version/v1/logfile",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "create model branch version log file + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidNoOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/logfile",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "create model branch version log file + valid token + model not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/NoModel/branch/dev/version/v1/logfile",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model not found`,
			},
		},
		{
			Name:   "create model branch version log file + valid token + model branch not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/nobranch/version/v1/logfile",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model Branch not found`,
			},
		},
		{
			Name:   "create model branch version log file + valid token + model branch version not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v2/logfile",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model Branch Version not found`,
			},
		},
		{
			Name:   "create model branch version log file + valid token + file empty",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/logfile",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
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
			Name:   "create model branch version log file + valid token + unsupported storage",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/logfile",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
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
			Name:   "create model branch version log file + valid token + s3 not enabled",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/logfile",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
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
			Name:   "create model branch version log file + valid token + r2 not enabled",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/logfile",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
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
			Name:   "create model branch version log file + valid token + log created successfully",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/dev/version/v1/logfile",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
				"Content-Type":  validMultipartContentType.FormDataContentType(),
			},
			Body:           validMultipartBody,
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"key":"tmpfile_`,
				`"data":"file:///model-registry/` + test.ValidAdminUserUuid.String() + `/models/` + test.ValidAdminUserUuid.String() + `/` + test.ValidUserUuid.String() + `/logs/tmpfile_`,
				`"message":"Logs created"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
