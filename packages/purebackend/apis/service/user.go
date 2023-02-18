package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PureML-Inc/PureML/packages/purebackend/core"
	"github.com/PureML-Inc/PureML/packages/purebackend/middlewares"
	"github.com/PureML-Inc/PureML/packages/purebackend/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// BindUserApi registers the admin api endpoints and the corresponding handlers.
func BindUserApi(app core.App, rg *echo.Group) {
	api := Api{app: app}

	userGroup := rg.Group("/user")
	userGroup.GET("/profile", api.DefaultHandler(GetProfile), middlewares.RequireAuthContext)
	userGroup.GET("/profile/:userHandle", api.DefaultHandler(GetProfileByHandle))
	userGroup.POST("/profile", api.DefaultHandler(UpdateProfile), middlewares.RequireAuthContext)
	userGroup.POST("/signup", api.DefaultHandler(UserSignUp))
	userGroup.POST("/login", api.DefaultHandler(UserLogin))
	userGroup.POST("/verify-email", api.DefaultHandler(UserVerifyEmail))
	userGroup.POST("/forgot-password", api.DefaultHandler(UserForgotPassword))
	userGroup.POST("/reset-password", api.DefaultHandler(UserResetPassword)) //TODO To complete the logic here and update middlewares
}

// UserSignUp godoc
//
//	@Summary		User sign up.
//	@Description	User sign up with email, name, handle and password. The user will receive an email with a verification link if enabled from backend.
//	@Description	Response status code 202 means that the user has been created but the email verification is pending.
//	@Description	Response status code 200 means that the user has been created and no verification of email is needed.
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordData), bcrypt.DefaultCost)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	user, err := api.app.Dao().CreateUser(name.(string), emailData, handleData, bio.(string), avatar.(string), string(hashedPassword), !api.app.Settings().MailService.Enabled)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if api.app.Settings().MailService.Enabled {
		// Send verification email
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uuid":  user.UUID,
			"email": user.Email,
			"exp":   time.Now().Unix() + api.app.Settings().MailVerifificationAuthToken.Duration,
		})
		signedString, err := token.SignedString([]byte(api.app.Settings().MailVerifificationAuthToken.Secret))
		if err != nil {
			panic(err)
		}
		verifyLink := api.app.Settings().Site.BaseURL + "/verify-email?token=" + signedString
		emailTemplate := BaseEmailTemplate("Verify your email address",
			`Hi `+user.Handle+`,<br><br>Thanks for signing up for PureML. Please click the link below to verify your email address.<br><br><a href="`+verifyLink+`">Verify Email</a><br><br>Thanks,<br>Team PureML`,
		)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		subject := "Verify your email address"
		err = api.app.SendMail(user.Email, subject, emailTemplate)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		return models.NewDataResponse(http.StatusAccepted, user, "User created. Please verify your email address")
	}
	return models.NewDataResponse(http.StatusOK, user, "User created")
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
		user, err = api.app.Dao().GetSecureUserByEmail(emailData)
	} else {
		user, err = api.app.Dao().GetSecureUserByHandle(handleData)
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
	if !user.IsVerified {
		return models.NewDataResponse(http.StatusUnauthorized, nil, "User email is not verified")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":   user.UUID,
		"email":  user.Email,
		"handle": user.Handle,
	})
	signedString, err := token.SignedString([]byte(api.app.Settings().AdminAuthToken.Secret))
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

// UserVerifyEmail godoc
//
//	@Summary		User verify email.
//	@Description	User can verify email by providing verification token.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/verify-email [post]
//	@Param			org	body	models.UserVerifyEmailRequest	true	"User details"
func (api *Api) UserVerifyEmail(request *models.Request) *models.Response {
	request.ParseJsonBody()
	token := request.GetParsedBodyAttribute("token")
	var tokenData string
	if token == nil {
		tokenData = ""
	} else {
		tokenData = token.(string)
	}
	if tokenData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Token is required")
	}
	parsedToken, err := jwt.Parse(tokenData, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing algorithm %v", t.Method.Alg())
		}
		return []byte(api.app.Settings().MailVerifificationAuthToken.Secret), nil
	})
	if err != nil {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Could not parse verification token")
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userUUID := uuid.Must(uuid.FromString(claims["uuid"].(string)))
		expireTime := int64(claims["exp"].(float64))
		if expireTime < time.Now().Unix() {
			return models.NewDataResponse(http.StatusBadRequest, nil, "Verification token expired")
		}
		user, err := api.app.Dao().GetUserByUUID(userUUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		if user == nil {
			return models.NewDataResponse(http.StatusNotFound, nil, "User not found")
		}
		if user.IsVerified {
			return models.NewDataResponse(http.StatusBadRequest, nil, "User email is already verified")
		}
		err = api.app.Dao().VerifyUserEmail(userUUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
	} else {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Invalid verification token")
	}
	response := models.NewDataResponse(http.StatusOK, nil, "User verified")
	return response
}

// UserResendVerification godoc
//
//	@Security		ApiKeyAuth
//	@Summary		User resend verification.
//	@Description	User can resend verification email by providing email id.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/resend-verification [post]
func (api *Api) UserResendVerification(request *models.Request) *models.Response {
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
	if emailData != "" {
		if addr, ok := ValidateMailAddress(emailData); ok {
			emailData = addr
		} else {
			return models.NewErrorResponse(http.StatusBadRequest, "Email is invalid")
		}
	}
	user, err := api.app.Dao().GetUserByEmail(emailData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	if user == nil {
		return models.NewDataResponse(http.StatusNotFound, nil, "User not found")
	}
	if user.IsVerified {
		return models.NewDataResponse(http.StatusBadRequest, nil, "User email is already verified")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":  user.UUID,
		"email": user.Email,
		"exp":   time.Now().Unix() + api.app.Settings().MailVerifificationAuthToken.Duration,
	})
	signedString, err := token.SignedString([]byte(api.app.Settings().MailVerifificationAuthToken.Secret))
	if err != nil {
		panic(err)
	}
	verifyLink := api.app.Settings().Site.BaseURL + "/verify-email?token=" + signedString
	emailTemplate := BaseEmailTemplate("Verify your email address",
		`Hi `+user.Handle+`,<br><br>Thanks for signing up for PureML. Please click the link below to verify your email address.<br><br><a href="`+verifyLink+`">Verify Email</a><br><br>Thanks,<br>Team PureML`,
	)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	subject := "Verify your email address"
	err = api.app.SendMail(user.Email, subject, emailTemplate)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	response := models.NewDataResponse(http.StatusOK, nil, "Verification email sent")
	return response
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
//	@Param			org	body	models.UserForgotPasswordRequest	true	"User email"
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":  user.UUID,
		"email": user.Email,
		"exp":   time.Now().Unix() + api.app.Settings().PasswordResetAuthToken.Duration,
	})
	signedString, err := token.SignedString([]byte(api.app.Settings().PasswordResetAuthToken.Secret))
	if err != nil {
		panic(err)
	}
	verifyLink := api.app.Settings().Site.BaseURL + "/reset-password?token=" + signedString
	emailTemplate := BaseEmailTemplate("Reset your password",
		`Hi `+user.Handle+`,<br><br>Someone requested a password reset for your PureML account. If it was you, click the link below to reset your password.<br><br><a href="`+verifyLink+`">Reset Password</a><br><br>Thanks,<br>Team PureML`,
	)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	subject := "Reset your password"
	err = api.app.SendMail(user.Email, subject, emailTemplate)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, nil, "Reset password link sent to your mail")
}

// UserResetPassword godoc
//
//	@Security		ApiKeyAuth
//	@Summary		User reset password.
//	@Description	User can reset password by providing old password and new password.
//	@Tags			User
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/user/reset-password [post]
//	@Param			org	body	models.UserResetPasswordRequest	true	"User email"
func (api *Api) UserResetPassword(request *models.Request) *models.Response {
	request.ParseJsonBody()
	token := request.GetParsedBodyAttribute("token")
	var tokenData string
	if token == nil {
		tokenData = ""
	} else {
		tokenData = token.(string)
	}
	if tokenData == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Token is required")
	}
	parsedToken, err := jwt.Parse(tokenData, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing algorithm %v", t.Method.Alg())
		}
		return []byte(api.app.Settings().PasswordResetAuthToken.Secret), nil
	})
	if err != nil {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Could not parse verification token")
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userUUID := uuid.Must(uuid.FromString(claims["uuid"].(string)))
		expireTime := int64(claims["exp"].(float64))
		if expireTime < time.Now().Unix() {
			return models.NewDataResponse(http.StatusBadRequest, nil, "Password Reset token expired")
		}
		user, err := api.app.Dao().GetSecureUserByUUID(userUUID)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		if user == nil {
			return models.NewDataResponse(http.StatusNotFound, nil, "User not found")
		}
		oldPassword := request.GetParsedBodyAttribute("old_password")
		var oldPasswordData string
		if oldPassword == nil {
			oldPasswordData = ""
		} else {
			oldPasswordData = oldPassword.(string)
		}
		if oldPasswordData == "" {
			return models.NewErrorResponse(http.StatusBadRequest, "Old password is required")
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPasswordData))
		if err != nil {
			return models.NewDataResponse(http.StatusUnauthorized, nil, "Invalid credentials")
		}
		newPassword := request.GetParsedBodyAttribute("new_password")
		var newPasswordData string
		if newPassword == nil {
			newPasswordData = ""
		} else {
			newPasswordData = newPassword.(string)
		}
		if newPasswordData == "" {
			return models.NewErrorResponse(http.StatusBadRequest, "New password is required")
		}
		//TODO: Validate password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPasswordData), bcrypt.DefaultCost)
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
		err = api.app.Dao().UpdateUserPassword(userUUID, string(hashedPassword))
		if err != nil {
			return models.NewServerErrorResponse(err)
		}
	} else {
		return models.NewDataResponse(http.StatusBadRequest, nil, "Invalid verification token")
	}
	return models.NewDataResponse(http.StatusOK, nil, "Password reset successfully")
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
	updatedAttributes := map[string]interface{}{}
	if name != nil {
		updatedAttributes["name"] = name.(string)
		if name.(string) == "" {
			return models.NewErrorResponse(http.StatusBadRequest, "Name cannot be empty")
		}
	}
	if avatar != nil {
		updatedAttributes["avatar"] = avatar.(string)
	}
	if bio != nil {
		updatedAttributes["bio"] = bio.(string)
	}
	email := request.GetUserMail()
	user, err := api.app.Dao().UpdateUser(email, updatedAttributes)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, user, "User profile updated")
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
var UserVerifyEmail ServiceFunc = (*Api).UserVerifyEmail
var UserResetPassword ServiceFunc = (*Api).UserResetPassword
var UserForgotPassword ServiceFunc = (*Api).UserForgotPassword
var GetProfile ServiceFunc = (*Api).GetProfile
var GetProfileByHandle ServiceFunc = (*Api).GetProfileByHandle
var UpdateProfile ServiceFunc = (*Api).UpdateProfile
var DeleteProfile ServiceFunc = (*Api).DeleteProfile
