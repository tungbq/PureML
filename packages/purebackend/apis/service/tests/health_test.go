package service_test

import (
	"net/http"
	"testing"

	"github.com/PureML-Inc/PureML/purebackend/tests"
)

func TestHealthCheck(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:           "health status returns 200",
			Method:         http.MethodGet,
			Url:            "/api/health",
			ExpectedStatus: 200,
			ExpectedContent: []string{
				`"status":200`,
				`"data":`,
				`"message":"Server is up and running🚀🎉"`,
			},
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
