package tests

import (
	"net/http"
	"testing"

	"github.com/PureMLHQ/PureML/packages/purebackend/test"
)

func TestHealthCheck(t *testing.T) {
	scenarios := []test.ApiScenario{
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
