package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// UserSignUp godoc
//
//	@Summary		User sign up.
//	@Description	User sign up with email, name, handle and password.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/signup [post]
//	@Param			user	body	models.UserSignupRequest	true	"User details"
func UserSignUp(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email").(string)
	handle := request.GetParsedBodyAttribute("handle").(string)
	name := request.GetParsedBodyAttribute("name")
	if name == nil {
		name = ""
	}
	bio := request.GetParsedBodyAttribute("bio")
	if bio == nil {
		bio = ""
	}
	avatar := request.GetParsedBodyAttribute("avatar")
	if avatar == nil {
		avatar = ""
	}
	password := request.GetParsedBodyAttribute("password").(string)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	user, err := datastore.CreateUser(name.(string), email, handle, bio.(string), avatar.(string), string(hashedPassword))
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, []models.UserResponse{*user}, "User created")
	return response
}

// UserLogin godoc
//
//	@Summary		User login.
//	@Description	User login with email and password.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/login [post]
//	@Param			org	body	models.UserLoginRequest	true	"User details"
func UserLogin(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email")
	handle := request.GetParsedBodyAttribute("handle")
	if email == nil && handle == nil {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Email or handle is required")
	}
	password := request.GetParsedBodyAttribute("password").(string)
	var user *models.UserResponse
	var err error
	if email != nil {
		email := email.(string)
		user, err = datastore.GetSecureUserByEmail(email)
	} else {
		handle := handle.(string)
		user, err = datastore.GetSecureUserByHandle(handle)
	}
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if user == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.NewDataResponse(http.StatusUnauthorized, nil, "Invalid credentials")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":   user.UUID,
		"email":  user.Email,
		"handle": user.Handle,
	})
	signedString, err := token.SignedString(config.TokenSigningSecret())
	if err != nil {
		panic(err)
	}
	data := []map[string]string{
		{
			"email":       user.Email,
			"accessToken": signedString,
		},
	}
	return models.NewDataResponse(http.StatusAccepted, data, "User logged in")
}


// TODO: UserResetPassword godoc
//	@Security		ApiKeyAuth
//	@Summary		User reset password.
//	@Description	User can reset password by providing old password and new password.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/reset-password [post]
func UserResetPassword(request *models.Request) *models.Response {
	return nil
}

// UserForgotPassword godoc
//
//	@Security		ApiKeyAuth
//	@Summary		User forgot password.
//	@Description	User can reset password by providing email id to send reset password link.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/forgot-password [post]
//	@Param			org	body	models.UserResetPasswordRequest	true	"User email"
func UserForgotPassword(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email").(string)
	user, err := datastore.GetUserByEmail(email)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if user == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "User with given email not found")
	}
	token := jwt.NewWithClaims(&jwt.SigningMethodHMAC{}, jwt.MapClaims{
		"email": user.Email,
	})
	_, err = token.SignedString(config.TokenSigningSecret())
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	//TODO : Invoke mailService here
	return models.NewDataResponse(http.StatusOK, nil, "Reset password link sent to your mail")
}


// GetProfile godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get logged in user profile.
//	@Description	Get logged in user profile.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/profile [get]
func GetProfile(request *models.Request) *models.Response {
	userUUID := request.GetUserUUID()
	user, err := datastore.GetUserProfileByUUID(userUUID)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, user, "User profile")
}

// GetProfileByHandle godoc
//
//	@Summary		Get user profile by handle.
//	@Description	Get user profile by handle. Accessible without login.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/profile/{userHandle} [get]
//	@Param			userHandle	path	string	true	"User handle"
func GetProfileByHandle(request *models.Request) *models.Response {
	userHandle := request.GetPathParam("userHandle")
	user, err := datastore.GetUserByHandle(userHandle)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, user, "Public User profile")
}

// UpdateProfile godoc
//
//	@Security		ApiKeyAuth
//	@Summary		User update profile.
//	@Description	User can update profile by providing name, avatar, bio.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/profile [post]
//	@Param			org	body	models.UserUpdateRequest	true	"User details"
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
	user, err := datastore.UpdateUser(email, updatedAttributes["name"], updatedAttributes["avatar"], updatedAttributes["bio"])
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

// TODO: DeleteProfile godoc
//	@Security		ApiKeyAuth
//	@Summary		Delete user profile.
//	@Description	Delete user profile.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/delete [delete]
func DeleteProfile(request *models.Request) *models.Response {
	return nil
}
