package service_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/PuremlHQ/PureML/packages/purebackend/tests"
)

func TestGetModelAllBranches(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "get all branches of model + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get all branches of model + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch",
			RequestHeaders: map[string]string{
				"Authorization": InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get all branches of model + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch",
			RequestHeaders: map[string]string{
				"Authorization": ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "get all branches of model + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + InvalidOrgUuidString + "/model/Demo%20Model/branch",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "get all branches of model + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidNoOrgUuid.String() + "/model/Demo%20Model/branch",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "get all branches of model + valid token + model not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/NoModel/branch",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model not found`,
			},
		},
		{
			Name:   "get all branches of model + valid token + model found",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + ValidAdminUserUuid.String() + `"`,
				`"name":"main"`,
				`"model":{`,
				`"name":"Demo Model"`,
				`"is_default":true`,
				`"uuid":"` + ValidUserUuid.String() + `"`,
				`"name":"dev"`,
				`"model":{`,
				`"name":"Demo Model"`,
				`"is_default":false`,
				`"message":"All model branches"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGetModelBranch(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "get branch of model + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/main",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get branch of model + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/main",
			RequestHeaders: map[string]string{
				"Authorization": InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get branch of model + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/main",
			RequestHeaders: map[string]string{
				"Authorization": ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "get branch of model + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + InvalidOrgUuidString + "/model/Demo%20Model/branch/main",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "get branch of model + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidNoOrgUuid.String() + "/model/Demo%20Model/branch/main",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "get branch of model + valid token + model not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/NoModel/branch/main",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model not found`,
			},
		},
		{
			Name:   "get branch of model + valid token + model branch not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/nobranch",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Branch not found`,
			},
		},
		{
			Name:   "get branch of model + valid token + model branch found",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/main",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + ValidAdminUserUuid.String() + `"`,
				`"name":"main"`,
				`"model":{`,
				`"name":"Demo Model"`,
				`"is_default":true`,
				`"message":"Model branch details"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestCreateModelBranch(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "create model branch + unauthorized",
			Method:         http.MethodPost,
			Url:            "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/create",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "create model branch + invalid token",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "create model branch + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + InvalidOrgUuidString + "/model/Demo%20Model/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "create model branch + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidNoOrgUuid.String() + "/model/Demo%20Model/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "create model branch + valid token + model not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/NoModel/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Model not found`,
			},
		},
		{
			Name:   "create model branch + valid token + branch name empty",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"branch_name":""
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Branch name cannot be empty"`,
			},
		},
		{
			Name:   "create model branch + valid token + branch already exists",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"branch_name":"main"
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Branch already exists"`,
			},
		},
		{
			Name:   "create model branch + valid token + branch created successfully",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/model/Demo%20Model/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"branch_name":"test"
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + ValidAdminUserUuid.String() + `"`,
				`"name":"test"`,
				`"is_default":false`,
				`"message":"Model branch created"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

// TODO
func TestUpdateModelBranch(t *testing.T) {
}

// TODO
func TestDeleteModelBranch(t *testing.T) {
}
