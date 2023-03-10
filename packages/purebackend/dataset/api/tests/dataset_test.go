package tests

import (
	"net/http"
	"strings"
	"testing"

	"github.com/PureMLHQ/PureML/packages/purebackend/test"
	"github.com/labstack/echo/v4"
)

func TestGetAllPublicDatasets(t *testing.T) {
	scenarios := []test.ApiScenario{
		{
			Name:           "get all public datasets + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/public/dataset",
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + test.ValidAdminUserUuid.String() + `"`,
				`"name":"Demo Dataset"`,
				`"wiki":"Demo Dataset Wiki"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"message":"Datasets successfully retrieved"`,
			},
		},
		{
			Name:   "get all public datasets + invalid token",
			Method: http.MethodGet,
			Url:    "/api/public/dataset",
			RequestHeaders: map[string]string{
				"Authorization": test.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get all public datasets + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/public/dataset",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidTokenNoUser,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + test.ValidAdminUserOrgUuid.String() + `"`,
				`"name":"Demo Dataset"`,
				`"wiki":"Demo Dataset Wiki"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"message":"Datasets successfully retrieved"`,
			},
		},
		{
			Name:   "get all public datasets + valid token + user found",
			Method: http.MethodGet,
			Url:    "/api/public/dataset",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + test.ValidAdminUserOrgUuid.String() + `"`,
				`"name":"Demo Dataset"`,
				`"wiki":"Demo Dataset Wiki"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"message":"Datasets successfully retrieved"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGetAllDatasets(t *testing.T) {
	scenarios := []test.ApiScenario{
		{
			Name:           "get all org datasets + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/all",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get all org datasets + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/all",
			RequestHeaders: map[string]string{
				"Authorization": test.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get all org datasets + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/all",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "get all org datasets + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.InvalidOrgUuidString + "/dataset/all",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "get all org datasets + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidNoOrgUuid.String() + "/dataset/all",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "get all org datasets + valid token + user not owner",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/all",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidUserToken,
			},
			BeforeTestFunc: func(t *testing.T, app *test.TestApp, e *echo.Echo) {
				// Make notadmin a "member" of the admin user org
				_, err := app.Dao().CreateUserOrganizationFromEmailAndOrgId("notadmin@aztlan.in", test.ValidAdminUserOrgUuid)
				if err != nil {
					t.Fatal(err)
				}
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + test.ValidAdminUserOrgUuid.String() + `"`,
				`"name":"Demo Dataset"`,
				`"wiki":"Demo Dataset Wiki"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"message":"Datasets successfully retrieved"`,
			},
		},
		{
			Name:   "get all org datasets + valid token + user is owner",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/all",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + test.ValidAdminUserUuid.String() + `"`,
				`"name":"Demo Dataset"`,
				`"wiki":"Demo Dataset Wiki"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"uuid":"` + test.ValidUserUuid.String() + `"`,
				`"name":"Demo Private Dataset"`,
				`"wiki":"Demo Private Dataset Wiki"`,
				`"is_public":false`,
				`"created_by":{`,
				`"updated_by":{`,
				`"message":"Datasets successfully retrieved"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGetDataset(t *testing.T) {
	scenarios := []test.ApiScenario{
		{
			Name:           "create dataset + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "create dataset + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset",
			RequestHeaders: map[string]string{
				"Authorization": test.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "create dataset + valid token + invalid org uuid",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.InvalidOrgUuidString + "/dataset/Demo%20Dataset",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "create dataset + valid token + org not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidNoOrgUuid.String() + "/dataset/Demo%20Dataset",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "create dataset + valid token + org found + dataset not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/No%20Dataset",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Dataset not found`,
			},
		},
		{
			Name:   "create dataset + valid token + org found + dataset found",
			Method: http.MethodGet,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + test.ValidAdminUserOrgUuid.String() + `"`,
				`"name":"Demo Dataset"`,
				`"wiki":"Demo Dataset Wiki"`,
				`"is_public":true`,
				`"created_by":{`,
				`"updated_by":{`,
				`"message":"Dataset successfully retrieved"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestCreateDataset(t *testing.T) {
	scenarios := []test.ApiScenario{
		{
			Name:           "create dataset + unauthorized",
			Method:         http.MethodPost,
			Url:            "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/test/create",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "create dataset + invalid token",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/test/create",
			RequestHeaders: map[string]string{
				"Authorization": test.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "create dataset + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.InvalidOrgUuidString + "/dataset/test/create",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "create dataset + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidNoOrgUuid.String() + "/dataset/test/create",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "create dataset + valid token + org found + dataset name already exists",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/create",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
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
				`"message":"Dataset already exists"`,
			},
		},
		{
			Name:   "create dataset + valid token + org found + custom branch names without main",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/test/create",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
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
			Name:   "create dataset + valid token + org found + dataset and branches created successfully",
			Method: http.MethodPost,
			Url:    "/api/org/" + test.ValidAdminUserOrgUuid.String() + "/dataset/test/create",
			RequestHeaders: map[string]string{
				"Authorization": test.ValidAdminToken,
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
				`"message":"Dataset and branches successfully created"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
