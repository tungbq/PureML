package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
)

// UpdateProfile godoc
// @Summary User update profile.
// @Description User can update profile by providing name, avatar, bio.
// @Tags User
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/signup [post]
// @Param org body models.UserUpdateRequest true "User details"
func UpdateProfile(request *models.Request) *models.Response {
	request.ParseJsonBody()
	name := request.GetParsedBodyAttribute("name")
	avatar := request.GetParsedBodyAttribute("avatar")
	bio := request.GetParsedBodyAttribute("bio")
	updatedAttributes := map[string]string{}
	if name != nil {
		updatedAttributes["name"] = name.(string)
	}
	if avatar != nil {
		updatedAttributes["avatar"] = avatar.(string)
	}
	if bio != nil {
		updatedAttributes["bio"] = bio.(string)
	}
	email := request.GetUserMail()
	user, err := datastore.UpdateUser(email, updatedAttributes)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := &models.Response{}
	response.StatusCode = http.StatusOK
	response.Body.Message = "User profile updated"
	response.Body.Data = []map[string]interface{}{
		{
			"email":  user.Email,
			"avatar": user.Avatar,
			"name":   user.Avatar,
		},
	}
	return response
}
