package service

import (
	"net/http"

	"github.com/PriyavKaneria/PureML/service/datastore"
	"github.com/PriyavKaneria/PureML/service/models"
)

func AddOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	mailId := request.GetParsedBodyAttribute("email").(string)
	user, err := datastore.GetUser(mailId)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response := &models.Response{}
	if user == nil {
		response.StatusCode = http.StatusNotFound
		response.Message = "User not found"
		return response
	}
	orgId := request.GetOrgId()
	_, err = datastore.CreateOrgAccessFromMailIdAndOrgId(mailId, orgId)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response.StatusCode = http.StatusOK
	response.Message = "User added to organization"
	return response
}
