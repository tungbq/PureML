package service_test

import (
	"net/http"
	"testing"

	"github.com/PureML-Inc/PureML/server/tests"
)

func TestHealthCheck(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "health status returns 200",
			Method:         http.MethodGet,
			Url:            "/api/health",
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"code":200`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
