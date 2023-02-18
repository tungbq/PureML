package service_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/PureML-Inc/PureML/packages/purebackend/tests"
	"github.com/labstack/echo/v4"
)

func TestUserSignUp(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:   "signup + no email",
			Method: http.MethodPost,
			Url:    "/api/user/signup",
			Body: strings.NewReader(`{
				"handle":"test",
				"name":"test",
				"password":"test"
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Email is required"`,
			},
		},
		{
			Name:   "signup + invalid email",
			Method: http.MethodPost,
			Url:    "/api/user/signup",
			Body: strings.NewReader(`{
				"email":"invalidEmail",
				"handle":"test",
				"name":"test",
				"password":"test"
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Email is invalid"`,
			},
		},
		{
			Name:   "signup + no handle",
			Method: http.MethodPost,
			Url:    "/api/user/signup",
			Body: strings.NewReader(`{
				"email":"test@test.com",
				"name":"test",
				"password":"test"
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Handle is required"`,
			},
		},
		{
			Name:   "signup + no password",
			Method: http.MethodPost,
			Url:    "/api/user/signup",
			Body: strings.NewReader(`{
				"email":"test@test.com",
				"handle":"test",
				"name":"test"
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Password is required"`,
			},
		},
		{
			Name:   "signup + user exists with same email or handle",
			Method: http.MethodPost,
			Url:    "/api/user/signup",
			Body: strings.NewReader(`{
				"email":"test@test.com",
				"handle":"test",
				"name":"test",
				"password":"test"
			}`),
			BeforeTestFunc: func(t *testing.T, app *tests.TestApp, e *echo.Echo) {
				// Create user
				_, err := app.Dao().CreateUser("test", "test@test.com", "test", "", "", "$2a$10$N..OOp8lPw0fRGCXT.HxH.LO8BUKwlncI/ufXK/bLTEvyeFmdCun.", true)
				if err != nil {
					t.Fatal(err)
				}
			},
			ExpectedStatus: 409,
			ExpectedContent: []string{
				`"status":409`,
				`"data":null`,
				`"message":"User with email already exists"`,
			},
		},
		{
			Name:   "signup successful",
			Method: http.MethodPost,
			Url:    "/api/user/signup",
			Body: strings.NewReader(`{
				"email":"test@test.com",
				"handle":"test",
				"name":"test",
				"password":"test"
			}`),
			BeforeTestFunc: func(t *testing.T, app *tests.TestApp, e *echo.Echo) {
				// Delete user if exists
				err := app.Dao().ExecuteSQL(fmt.Sprintf("DELETE FROM users WHERE email = '%s'", "test@test.com"))
				if err != nil {
					t.Fatal(err)
				}
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"name":"test",`,
				`"email":"test@test.com",`,
				`"handle":"test",`,
				`"message":"User created"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestUserLogin(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:   "login + no email & no handle",
			Method: http.MethodPost,
			Url:    "/api/user/login",
			Body: strings.NewReader(`{
				"password":"test"
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Email or handle is required"`,
			},
		},
		{
			Name:   "login + no password",
			Method: http.MethodPost,
			Url:    "/api/user/login",
			Body: strings.NewReader(`{
				"email":"test@test.com",
				"handle":"test"
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Password is required"`,
			},
		},
		{
			Name:   "login + invalid email",
			Method: http.MethodPost,
			Url:    "/api/user/login",
			Body: strings.NewReader(`{
				"email":"invalid",
				"password":"test"
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Email is invalid"`,
			},
		},
		{
			Name:   "login + invalid credentials with email",
			Method: http.MethodPost,
			Url:    "/api/user/login",
			Body: strings.NewReader(`{
				"email":"demo@aztlan.in",
				"password":"wrongpassword"
			}`),
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`"status":401`,
				`"data":null`,
				`"message":"Invalid credentials"`,
			},
		},
		{
			Name:   "login + invalid credentials with handle",
			Method: http.MethodPost,
			Url:    "/api/user/login",
			Body: strings.NewReader(`{
				"handle":"demo",
				"password":"wrongpassword"
			}`),
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`"status":401`,
				`"data":null`,
				`"message":"Invalid credentials"`,
			},
		},
		{
			Name:   "login + email not verified",
			Method: http.MethodPost,
			Url:    "/api/user/login",
			Body: strings.NewReader(`{
				"handle":"test",
				"password":"test"
			}`),
			BeforeTestFunc: func(t *testing.T, app *tests.TestApp, e *echo.Echo) {
				// Create user
				_, err := app.Dao().CreateUser("test", "test@test.com", "test", "", "", "$2a$10$N..OOp8lPw0fRGCXT.HxH.LO8BUKwlncI/ufXK/bLTEvyeFmdCun.", true)
				if err != nil {
					t.Fatal(err)
				}
			},
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`"status":401`,
				`"data":null`,
				`"message":"User email is not verified"`,
			},
		},
		{
			Name:   "login successful with email",
			Method: http.MethodPost,
			Url:    "/api/user/login",
			Body: strings.NewReader(`{
				"email":"demo@aztlan.in",
				"password":"demo"
			}`),
			ExpectedStatus: 202,
			ExpectedContent: []string{
				`"status":202`,
				`"data":[{`,
				`"email":"demo@aztlan.in"`,
				`"accessToken":`,
				`"message":"User logged in"`,
			},
		},
		{
			Name:   "login successful with handle",
			Method: http.MethodPost,
			Url:    "/api/user/login",
			Body: strings.NewReader(`{
				"handle":"demo",
				"password":"demo"
			}`),
			ExpectedStatus: 202,
			ExpectedContent: []string{
				`"status":202`,
				`"data":[{`,
				`"email":"demo@aztlan.in"`,
				`"accessToken":`,
				`"message":"User logged in"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

// TODO
func TestUserResetPassword(t *testing.T) {
}

// TODO
func TestUserForgotPassword(t *testing.T) {
}

func TestGetProfile(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "get profile + unauthorized",
			Method:         http.MethodGet,
			Url:            "/api/user/profile",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "get profile + invalid token",
			Method: http.MethodGet,
			Url:    "/api/user/profile",
			RequestHeaders: map[string]string{
				"Authorization": InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get profile + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/user/profile",
			RequestHeaders: map[string]string{
				"Authorization": ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "get profile + valid token + user found",
			Method: http.MethodGet,
			Url:    "/api/user/profile",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"uuid":"` + ValidAdminUserUuid.String() + `"`,
				`"email":"demo@aztlan.in"`,
				`"handle":"demo"`,
				`"name":"Demo User"`,
				`"message":"User profile"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGetProfileByHandle(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "get profile by handle + unauthorized + user not found",
			Method:         http.MethodGet,
			Url:            "/api/user/profile/noone",
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":null`,
				`"message":"Public User profile"`,
			},
		},
		{
			Name:           "get profile by handle + unauthorized + user found",
			Method:         http.MethodGet,
			Url:            "/api/user/profile/demo",
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"email":"demo@aztlan.in"`,
				`"handle":"demo"`,
				`"name":"Demo User"`,
				`"message":"Public User profile"`,
			},
		},
		{
			Name:   "get profile by handle + invalid token",
			Method: http.MethodGet,
			Url:    "/api/user/profile/demo",
			RequestHeaders: map[string]string{
				"Authorization": InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "get profile by handle + valid token + user not found",
			Method: http.MethodGet,
			Url:    "/api/user/profile/noone",
			RequestHeaders: map[string]string{
				"Authorization": ValidTokenNoUser,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":null`,
				`"message":"Public User profile"`,
			},
		},
		{
			Name:   "get profile + valid token + user found",
			Method: http.MethodGet,
			Url:    "/api/user/profile/demo",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"email":"demo@aztlan.in"`,
				`"handle":"demo"`,
				`"name":"Demo User"`,
				`"message":"Public User profile"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestUpdateProfile(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "update profile + unauthorized",
			Method:         http.MethodPost,
			Url:            "/api/user/profile",
			ExpectedStatus: 401,
			ExpectedContent: []string{
				`Authentication token required`,
			},
		},
		{
			Name:   "update profile + invalid token",
			Method: http.MethodPost,
			Url:    "/api/user/profile",
			RequestHeaders: map[string]string{
				"Authorization": InvalidToken,
			},
			ExpectedStatus: 403,
			ExpectedContent: []string{
				`Could not parse authentication token`,
			},
		},
		{
			Name:   "update profile + valid token + user not found",
			Method: http.MethodPost,
			Url:    "/api/user/profile",
			RequestHeaders: map[string]string{
				"Authorization": ValidTokenNoUser,
			},
			ExpectedStatus: 404,
			ExpectedContent: []string{
				`User not found`,
			},
		},
		{
			Name:   "update profile + valid token + user found + empty name",
			Method: http.MethodPost,
			Url:    "/api/user/profile",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"name":"",
				"avatar":"",
				"bio":""
			}`),
			ExpectedStatus: 400,
			ExpectedContent: []string{
				`"status":400`,
				`"data":null`,
				`"message":"Name cannot be empty"`,
			},
		},
		{
			Name:   "update profile + valid token + user found + update name",
			Method: http.MethodPost,
			Url:    "/api/user/profile",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"name":"Demo User New"
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"email":"demo@aztlan.in"`,
				`"handle":"demo"`,
				`"name":"Demo User New"`,
				`"avatar":""`,
				`"bio":"Demo User Bio"`,
				`"message":"User profile updated"`,
			},
		},
		{
			Name:   "update profile + valid token + user found + update avatar",
			Method: http.MethodPost,
			Url:    "/api/user/profile",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"avatar":"test"
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"email":"demo@aztlan.in"`,
				`"handle":"demo"`,
				`"name":"Demo User"`,
				`"avatar":"test"`,
				`"bio":"Demo User Bio"`,
				`"message":"User profile updated"`,
			},
		},
		{
			Name:   "update profile + valid token + user found + update bio",
			Method: http.MethodPost,
			Url:    "/api/user/profile",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"bio":"Demo User Bio New"
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"email":"demo@aztlan.in"`,
				`"handle":"demo"`,
				`"name":"Demo User"`,
				`"avatar":""`,
				`"bio":"Demo User Bio New"`,
				`"message":"User profile updated"`,
			},
		},
		{
			Name:   "update profile + valid token + user found + update all",
			Method: http.MethodPost,
			Url:    "/api/user/profile",
			RequestHeaders: map[string]string{
				"Authorization": ValidAdminToken,
			},
			Body: strings.NewReader(`{
				"name":"Demo User New",
				"avatar":"avatar",
				"bio":"bio"
			}`),
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":[{`,
				`"email":"demo@aztlan.in"`,
				`"handle":"demo"`,
				`"name":"Demo User New"`,
				`"avatar":"avatar"`,
				`"bio":"bio"`,
				`"message":"User profile updated"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

// TODO
func TestDeleteProfile(t *testing.T) {
}
