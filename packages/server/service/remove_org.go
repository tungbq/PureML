package service

import "github.com/PureML-Inc/PureML/server/models"

// RemoveOrg godoc
// @Summary Remove user from organization.
// @Description Remove user from organization.
// @Tags Organization
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /org/:orgId/remove [post]
// @Param orgId path string true "Organization ID"
// @Param email body string true "User email"
func RemoveOrg(request *models.Request) *models.Response {
	return nil
}
