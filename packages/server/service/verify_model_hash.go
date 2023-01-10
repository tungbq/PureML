package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

func VerifyModelHashStatus(request *models.Request) *models.Response {
	orgId := request.GetPathParam("orgId")
	modelName := request.GetPathParam("modelName")
	request.ParseJsonBody()
	hashValue := request.GetParsedBodyAttribute("hash").(string)
	versions, err := datastore.GetAllModelVersions(uuid.Must(uuid.FromString(orgId)), modelName)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := true
	for _, version := range versions {
		if version.Hash == hashValue {
			response = false
			break
		}
	}
	return models.NewDataResponse(http.StatusOK, response, "Hash value retrieved")
}
