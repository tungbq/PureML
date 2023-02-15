package service_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/PureML-Inc/PureML/purebackend/tests"
	"github.com/labstack/echo/v4"
)

func TestGetOrgsForUser(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "get user orgs + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get user orgs + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org",
			RequestHeaders: map[string]string{
				"Authorization": InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get user orgs + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/org",
			RequestHeaders: map[string]string{
				"Authorization": ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "get user orgs + valid token + user found",
			Method: http.MethodGet,
			Url:    "/api/org",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + ValidAdminUserOrgUuid.String() + `"`,
				`"name":"Demo Org"`,
				`"handle":"demo"`,
				`"avatar":""`,
				`"description":"Demo Org Description"`,
				`"message":"User Organizations"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestAddUsersToOrg(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "add user to org + unauthorized",
			Method:         http.MethodPost,
			Url:            "/api/org/" + ValidAdminUserOrgUuid.String() + "/add",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "add user to org + invalid token",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/add",
			RequestHeaders: map[string]string{
				"Authorization": InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "add user to org + valid token + login user not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/add",
			RequestHeaders: map[string]string{
				"Authorization": ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "add user to org + valid token + invalid org uuid",
			Method: http.MethodPost,
			Url:    "/api/org/" + InvalidOrgUuidString + "/add",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "add user to org + valid token + org not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidNoOrgUuid.String() + "/add",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "add user to org + valid token + no email in body",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/add",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"email":""
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Email is required"`,
			},
		},
		{
			Name:   "add user to org + valid token + invalid email",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/add",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"email":"invalidemail"
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Email is invalid"`,
			},
		},
		{
			Name:   "add user to org + valid token + user to add not found",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/add",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"email":"noone@nomail.com"
			}`),
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`"status":404`,
				`"data":null`,
				`"message":"User to add not found"`,
			},
		},
		{
			Name:   "add user to org + valid token + not authorized to add user",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/add",
			RequestHeaders: map[string]string{
				"Authorization": ValidUserToken,
			},
			Body: strings.NewReader(`{
				"email":"test@test.com"
			}`),
			BeforeTestFunc: func(t *testing.T, app *tests.TestApp, e *echo.Echo) {
				// Create 3rd user
				_, err := app.Dao().CreateUser("test", "test@test.com", "test", "", "", "$2a$10$N..OOp8lPw0fRGCXT.HxH.LO8BUKwlncI/ufXK/bLTEvyeFmdCun.")
				if err != nil {
					t.Fatal(err)
				}
				// Make notadmin a "member" of the admin user org
				_, err = app.Dao().CreateUserOrganizationFromEmailAndOrgId("notadmin@aztlan.in", ValidAdminUserOrgUuid)
				if err != nil {
					t.Fatal(err)
				}
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`"status":403`,
				`"data":null`,
				`"message":"You are not authorized to add users this organization"`,
			},
		},
		{
			Name:   "add user to org + valid token + user already added",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/add",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"email":"notadmin@aztlan.in"
			}`),
			BeforeTestFunc: func(t *testing.T, app *tests.TestApp, e *echo.Echo) {
				// Make notadmin a "member" of the admin user org
				_, err := app.Dao().CreateUserOrganizationFromEmailAndOrgId("notadmin@aztlan.in", ValidAdminUserOrgUuid)
				if err != nil {
					t.Fatal(err)
				}
			},
			ExpectedStatus: 409,
			ExpectedContent: []string{
				`"status":409`,
				`"data":null`,
				`"message":"User already added to organization"`,
			},
		},
		{
			Name:   "add user to org + valid token + user added successfully",
			Method: http.MethodPost,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/add",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"email":"notadmin@aztlan.in"
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":null`,
				`"message":"User added to organization"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestJoinOrg(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "join org + unauthorized",
			Method:         http.MethodPost,
			Url:            "/api/org/join",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "join org + invalid token",
			Method: http.MethodPost,
			Url:    "/api/org/join",
			RequestHeaders: map[string]string{
				"Authorization": InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "join org + valid token + login user not found",
			Method: http.MethodPost,
			Url:    "/api/org/join",
			RequestHeaders: map[string]string{
				"Authorization": ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "join org + valid token + no join code in body",
			Method: http.MethodPost,
			Url:    "/api/org/join",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"join_code":""
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Join code is required"`,
			},
		},
		{
			Name:   "join org + valid token + invalid join code",
			Method: http.MethodPost,
			Url:    "/api/org/join",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"join_code":"joincodedoesnotexist"
			}`),
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`"status":404`,
				`"data":null`,
				`"message":"Invalid join code"`,
			},
		},
		{
			Name:   "join org + valid token + user already added",
			Method: http.MethodPost,
			Url:    "/api/org/join",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"join_code":"iwanttojoindemo"
			}`),
			ExpectedStatus: 409,
			ExpectedContent: []string{
				`"status":409`,
				`"data":null`,
				`"message":"User already member of organization"`,
			},
		},
		{
			Name:   "join org + valid token + user added successfully",
			Method: http.MethodPost,
			Url:    "/api/org/join",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"join_code":"iwanttojoinnotadmin"
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":null`,
				`"message":"User joined organization"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestLeaveOrg(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "leave org + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/org/" + ValidAdminUserOrgUuid.String() + "/leave",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "leave org + invalid token",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/leave",
			RequestHeaders: map[string]string{
				"Authorization": InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "leave org + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/leave",
			RequestHeaders: map[string]string{
				"Authorization": ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "leave org + valid token + invalid org uuid",
			Method: http.MethodGet,
			Url:    "/api/org/" + InvalidOrgUuidString + "/leave",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`Invalid UUID format`,
			},
		},
		{
			Name:   "leave org + valid token + org not found",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidNoOrgUuid.String() + "/leave",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`Organization not found`,
			},
		},
		{
			Name:   "leave org + valid token + owner cannot leave",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/leave",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`"status":403`,
				`"data":null`,
				`"message":"Owner can't leave organization"`,
			},
		},
		{
			Name:   "leave org + valid token + leave org successfully",
			Method: http.MethodGet,
			Url:    "/api/org/" + ValidAdminUserOrgUuid.String() + "/leave",
			RequestHeaders: map[string]string{
				"Authorization": ValidUserToken,
			},
			BeforeTestFunc: func(t *testing.T, app *tests.TestApp, e *echo.Echo) {
				// Make notadmin a "member" of the admin user org
				_, err := app.Dao().CreateUserOrganizationFromEmailAndOrgId("notadmin@aztlan.in", ValidAdminUserOrgUuid)
				if err != nil {
					t.Fatal(err)
				}
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":null`,
				`"message":"User left organization"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

// TODO
func TestRemoveOrg(t *testing.T) {
}
