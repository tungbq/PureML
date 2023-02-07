package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/core"
	"github.com/PureML-Inc/PureML/server/middlewares"
	"github.com/PureML-Inc/PureML/server/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// BindUserApi registers the admin api endpoints and the corresponding handlers.
func BindUserApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	userGroup := rg.Group("/user")
	userGroup.GET("/profile", api.DefaultHandler(GetProfile), middlewares.AuthenticateJWT(api.app))
	userGroup.GET("/profile/:userHandle", api.DefaultHandler(GetProfileByHandle))
	userGroup.POST("/profile", api.DefaultHandler(UpdateProfile), middlewares.AuthenticateJWT(api.app))
	userGroup.POST("/signup", api.DefaultHandler(UserSignUp))
	userGroup.POST("/login", api.DefaultHandler(UserLogin))
	userGroup.POST("/forgot-password", api.DefaultHandler(UserForgotPassword))
	userGroup.POST("/reset-password", api.DefaultHandler(UserResetPassword)) //TODO To complete the logic here and update middlewares
}

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
func (api *Api) UserSignUp(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email")
	var emailData string
	if email == nil {
		emailData = ""
	} else {
		emailData = email.(string)
	}
	if emailData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Email is required")
	}
	if addr, ok := ValidateMailAddress(emailData); ok {
		emailData = addr
	} else {
		return models.NewErrorResponse(http.StatusBadRequest, "Email is invalid")
	}
	handle := request.GetParsedBodyAttribute("handle")
	var handleData string
	if handle == nil {
		handleData = ""
	} else {
		handleData = handle.(string)
	}
	if handleData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Handle is required")
	}
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
	password := request.GetParsedBodyAttribute("password")
	var passwordData string
	if password == nil {
		passwordData = ""
	} else {
		passwordData = password.(string)
	}
	if passwordData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Password is required")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordData), 10)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	user, err := api.app.Dao().CreateUser(name.(string), emailData, handleData, bio.(string), avatar.(string), string(hashedPassword))
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, user, "User created")
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
func (api *Api) UserLogin(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email")
	var emailData string
	if email == nil {
		emailData = ""
	} else {
		emailData = email.(string)
	}
	handle := request.GetParsedBodyAttribute("handle")
	var handleData string
	if handle == nil {
		handleData = ""
	} else {
		handleData = handle.(string)
	}
	if emailData == "" && handleData == "" {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Email or handle is required")
	}
	if emailData != "" {
		if addr, ok := ValidateMailAddress(emailData); ok {
			emailData = addr
		} else {
			return models.NewErrorResponse(http.StatusBadRequest, "Email is invalid")
		}
	}
	password := request.GetParsedBodyAttribute("password")
	var passwordData string
	if password == nil {
		passwordData = ""
	} else {
		passwordData = password.(string)
	}
	if passwordData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Password is required")
	}
	var user *models.UserResponse
	var err error
	if email != nil {
		email := email.(string)
		user, err = api.app.Dao().GetSecureUserByEmail(email)
	} else {
		handle := handle.(string)
		user, err = api.app.Dao().GetSecureUserByHandle(handle)
	}
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if user == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "User not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordData))
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
//
//	@Security		ApiKeyAuth
//	@Summary		User reset password.
//	@Description	User can reset password by providing old password and new password.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/reset-password [post]
func (api *Api) UserResetPassword(request *models.Request) *models.Response {
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
func (api *Api) UserForgotPassword(request *models.Request) *models.Response {
	request.ParseJsonBody()
	email := request.GetParsedBodyAttribute("email").(string)
	user, err := api.app.Dao().GetUserByEmail(email)
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
func (api *Api) GetProfile(request *models.Request) *models.Response {
	userUUID := request.GetUserUUID()
	user, err := api.app.Dao().GetUserProfileByUUID(userUUID)
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
func (api *Api) GetProfileByHandle(request *models.Request) *models.Response {
	userHandle := request.GetPathParam("userHandle")
	user, err := api.app.Dao().GetUserByHandle(userHandle)
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
func (api *Api) UpdateProfile(request *models.Request) *models.Response {
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
	user, err := api.app.Dao().UpdateUser(email, updatedAttributes["name"], updatedAttributes["avatar"], updatedAttributes["bio"])
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
//
//	@Security		ApiKeyAuth
//	@Summary		Delete user profile.
//	@Description	Delete user profile.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/delete [delete]
func (api *Api) DeleteProfile(request *models.Request) *models.Response {
	return nil
}

var UserSignUp ServiceFunc = (*Api).UserSignUp
var UserLogin ServiceFunc = (*Api).UserLogin
var UserResetPassword ServiceFunc = (*Api).UserResetPassword
var UserForgotPassword ServiceFunc = (*Api).UserForgotPassword
var GetProfile ServiceFunc = (*Api).GetProfile
var GetProfileByHandle ServiceFunc = (*Api).GetProfileByHandle
var UpdateProfile ServiceFunc = (*Api).UpdateProfile
var DeleteProfile ServiceFunc = (*Api).DeleteProfile
