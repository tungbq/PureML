package tests

import (
	"net/http"
	"strings"
	"testing"

	coretests "github.com/PureMLHQ/PureML/packages/purebackend/core/apis/service/tests"
	"github.com/PureMLHQ/PureML/packages/purebackend/tests"
	"github.com/labstack/echo/v4"
)

func TestGetAllPublicModels(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "get all public models + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/public/model",
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + coretests.ValidAdminUserUuid.String() + `"`,
				`"name":"Demo Model"`,
				`"wiki":"Demo Model Wiki"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"message":"Models successfully retrieved"`,
			},
		},
		{
			Name:   "get all public models + invalid token",
			Method: http.MethodGet,
			Url:    "/api/public/model",
			RequestHeaders: map[string]string{
				"Authorization": coretests.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get all public models + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/public/model",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidTokenNoUser,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + coretests.ValidAdminUserOrgUuid.String() + `"`,
				`"name":"Demo Model"`,
				`"wiki":"Demo Model Wiki"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"message":"Models successfully retrieved"`,
			},
		},
		{
			Name:   "get all public models + valid token + user found",
			Method: http.MethodGet,
			Url:    "/api/public/model",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + coretests.ValidAdminUserOrgUuid.String() + `"`,
				`"name":"Demo Model"`,
				`"wiki":"Demo Model Wiki"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"message":"Models successfully retrieved"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGetAllModels(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "get all org models + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/all",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get all org models + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/all",
			RequestHeaders: map[string]string{
				"Authorization": coretests.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get all org models + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/all",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "get all org models + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.InvalidOrgUuidString + "/model/all",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "get all org models + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidNoOrgUuid.String() + "/model/all",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "get all org models + valid token + user not owner",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/all",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidUserToken,
			},
			BeforeTestFunc: func(t *testing.T, app *tests.TestApp, e *echo.Echo) {
				// Make notadmin a "member" of the admin user org
				_, err := app.Dao().CreateUserOrganizationFromEmailAndOrgId("notadmin@aztlan.in", coretests.ValidAdminUserOrgUuid)
				if err != nil {
					t.Fatal(err)
				}
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + coretests.ValidAdminUserOrgUuid.String() + `"`,
				`"name":"Demo Model"`,
				`"wiki":"Demo Model Wiki"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"message":"Models successfully retrieved"`,
			},
		},
		{
			Name:   "get all org models + valid token + user is owner",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/all",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + coretests.ValidAdminUserUuid.String() + `"`,
				`"name":"Demo Model"`,
				`"wiki":"Demo Model Wiki"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"uuid":"` + coretests.ValidUserUuid.String() + `"`,
				`"name":"Demo Private Model"`,
				`"wiki":"Demo Private Model Wiki"`,
				`"is_public":false`,
				`"created_by":{`,
				`"updated_by":{`,
				`"message":"Models successfully retrieved"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGetModel(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "create model + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "create model + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model",
			RequestHeaders: map[string]string{
				"Authorization": coretests.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "create model + valid token + invalid org uuid",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.InvalidOrgUuidString + "/model/Demo%20Model",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "create model + valid token + org not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidNoOrgUuid.String() + "/model/Demo%20Model",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "create model + valid token + org found + model not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/No%20Model",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model not found`,
			},
		},
		{
			Name:   "create model + valid token + org found + model found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + coretests.ValidAdminUserOrgUuid.String() + `"`,
				`"name":"Demo Model"`,
				`"wiki":"Demo Model Wiki"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"message":"Model successfully retrieved"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestCreateModel(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "create model + unauthorized",
			Method:         http.MethodPost,
			Url:            "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/test/create",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "create model + invalid token",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/test/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "create model + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.InvalidOrgUuidString + "/model/test/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "create model + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidNoOrgUuid.String() + "/model/test/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "create model + valid token + org found + model name already exists",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"wiki":"test",
				"is_public":true,
				"branch_names":["main","test"],
				"readme":{"file_type":"markdown","content":"test"}
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Model already exists"`,
			},
		},
		{
			Name:   "create model + valid token + org found + custom branch names without main",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/test/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"wiki":"test",
				"is_public":true,
				"branch_names":["dev","test"],
				"readme":{"file_type":"markdown","content":"test"}
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Branch names must contain 'main'"`,
			},
		},
		{
			Name:   "create model + valid token + org found + model and branches created successfully",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/model/test/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"wiki":"test",
				"is_public":true,
				"branch_names":["main","test"],
				"readme":{"file_type":"markdown","content":"test"}
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"`,
				`"name":"test"`,
				`"wiki":"test"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"readme":{`,
				`"latest_version":{`,
				`"file_type":"markdown"`,
				`"content":"test"`,
				`"org":{`,
				`"message":"Model and branches successfully created"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
