package service

import (
	"net/http"

	"github.com/PriyavKaneria/PureML/service/datastore"
	"github.com/PriyavKaneria/PureML/service/models"
)

func LeaveOrg(request *models.Request) *models.Response {
	request.ParseJsonBody()
	mailId := request.GetParsedBodyAttribute("email").(string)
	orgId := request.GetOrgId()
	_, err := datastore.DeleteOrgAccessFromMailIdAndOrgId(mailId, orgId)
	if err != nil {
		return models.NewErrorResponse(err)
	}
	response := &models.Response{}
	response.StatusCode = http.StatusOK
	response.Message = "User left organization"
	return response
}
