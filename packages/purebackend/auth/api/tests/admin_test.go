package tests

import (
	"net/http"
	"testing"

	coretests "github.com/PureMLHQ/PureML/packages/purebackend/core/apis/service/tests"
	"github.com/PureMLHQ/PureML/packages/purebackend/tests"
)

func TestGetAllAdminOrgs(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "get all admin orgs + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/all",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get all admin orgs + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/all",
			RequestHeaders: map[string]string{
				"Authorization": coretests.InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get all admin orgs + valid token + not admin user",
			Method: http.MethodGet,
			Url:    "/api/org/all",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidUserToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`"status":403`,
				`"data":null`,
				`"message":"Forbidden"`,
			},
		},
		{
			Name:   "get all admin orgs + valid token + admin user",
			Method: http.MethodGet,
			Url:    "/api/org/all",
			RequestHeaders: map[string]string{
				"Authorization": coretests.ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + coretests.ValidAdminUserOrgUuid.String() + `"`,
				`"name":"Demo Org"`,
				`"handle":"demo"`,
				`"avatar":""`,
				`"description":"Demo Org Description"`,
				`"join_code":"iwanttojoindemo"`,
				`"uuid":"` + coretests.ValidUserOrgUuid.String() + `"`,
				`"name":"Demo Org"`,
				`"handle":"notadmin"`,
				`"avatar":""`,
				`"description":"Demo Org Description"`,
				`"join_code":"iwanttojoinnotadmin"`,
				`"message":"All organizations"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
