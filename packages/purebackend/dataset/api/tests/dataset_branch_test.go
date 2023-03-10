package tests

import (
	"net/http"
	"strings"
	"testing"

	coretests "github.com/PureMLHQ/PureML/packages/purebackend/core/apis/service/tests"
	"github.com/PureMLHQ/PureML/packages/purebackend/tests"
)

func TestGetDatasetAllBranches(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "get all branches of dataset + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get all branches of dataset + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch",
			RequestHeaders: map[string]string{
				"Authorization": coretests.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get all branches of dataset + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "get all branches of dataset + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.InvalidOrgUuidString + "/dataset/Demo%20Dataset/branch",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "get all branches of dataset + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidNoOrgUuid.String() + "/dataset/Demo%20Dataset/branch",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "get all branches of dataset + valid token + dataset not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/NoDataset/branch",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Dataset not found`,
			},
		},
		{
			Name:   "get all branches of dataset + valid token + dataset found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + coretests.ValidAdminUserUuid.String() + `"`,
				`"name":"main"`,
				`"dataset":{`,
				`"name":"Demo Dataset"`,
				`"is_default":true`,
				`"uuid":"` + coretests.ValidUserUuid.String() + `"`,
				`"name":"dev"`,
				`"dataset":{`,
				`"name":"Demo Dataset"`,
				`"is_default":false`,
				`"message":"All dataset branches"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGetDatasetBranch(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "get branch of dataset + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch/main",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get branch of dataset + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch/main",
			RequestHeaders: map[string]string{
				"Authorization": coretests.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get branch of dataset + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch/main",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "get branch of dataset + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.InvalidOrgUuidString + "/dataset/Demo%20Dataset/branch/main",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "get branch of dataset + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidNoOrgUuid.String() + "/dataset/Demo%20Dataset/branch/main",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "get branch of dataset + valid token + dataset not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/NoDataset/branch/main",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Dataset not found`,
			},
		},
		{
			Name:   "get branch of dataset + valid token + dataset branch not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch/nobranch",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Dataset Branch not found`,
			},
		},
		{
			Name:   "get branch of dataset + valid token + dataset branch found",
			Method: http.MethodGet,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch/main",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + coretests.ValidAdminUserUuid.String() + `"`,
				`"name":"main"`,
				`"dataset":{`,
				`"name":"Demo Dataset"`,
				`"is_default":true`,
				`"message":"Dataset branch details"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestCreateDatasetBranch(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "create dataset branch + unauthorized",
			Method:         http.MethodPost,
			Url:            "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch/create",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "create dataset branch + invalid token",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "create dataset branch + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.InvalidOrgUuidString + "/dataset/Demo%20Dataset/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "create dataset branch + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidNoOrgUuid.String() + "/dataset/Demo%20Dataset/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidUserToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "create dataset branch + valid token + dataset not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/NoDataset/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Dataset not found`,
			},
		},
		{
			Name:   "create dataset branch + valid token + branch name empty",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
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
			Name:   "create dataset branch + valid token + branch already exists",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
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
			Name:   "create dataset branch + valid token + branch created successfully",
			Method: http.MethodPost,
			Url:    "/api/org/" + coretests.ValidAdminUserOrgUuid.String() + "/dataset/Demo%20Dataset/branch/create",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"branch_name":"test"
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + coretests.ValidAdminUserUuid.String() + `"`,
				`"name":"test"`,
				`"is_default":false`,
				`"message":"Dataset branch created"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

// TODO
func TestUpdateDatasetBranch(t *testing.T) {
}

// TODO
func TestDeleteDatasetBranch(t *testing.T) {
}
