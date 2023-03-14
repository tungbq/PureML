package tests

import (
	"net/http"
	"testing"

	"github.com/PureMLHQ/PureML/packages/purebackend/test"
)

func TestGetAllAdminOrgs(t *testing.T) {
	scenarios := []test.ApiScenario{
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
				"Authorization": test.InvalidToken,
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
				"Authorization": test.ValidUserToken,
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
				"Authorization": test.ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + test.ValidAdminUserOrgUuid.String() + `"`,
				`"name":"Demo Org"`,
				`"handle":"demo"`,
				`"avatar":""`,
				`"description":"Demo Org Description"`,
				`"join_code":"iwanttojoindemo"`,
				`"uuid":"` + test.ValidUserOrgUuid.String() + `"`,
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
