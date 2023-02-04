package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// GetModelReviews godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get model reviews
//	@Description	Get model reviews
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/review [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			modelName	path	string	true	"Model Name"
func GetModelReviews(request *models.Request) *models.Response {
	modelUUID := request.GetModelUUID()
	reviews, err := datastore.GetModelReviews(modelUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	response := models.NewDataResponse(http.StatusOK, reviews, "Model review version")
	return response
}

// CreateModelReview godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Create a new review request for model
//	@Description	Create a new review request for model
//	@Description	From and To branch names are required (Not UUID)
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/review/create [post]
//	@Param			orgId		path	string						true	"Organization Id"
//	@Param			modelName	path	string						true	"Model Name"
//	@Param			data		body	models.ModelReviewRequest	true	"Review"
func CreateModelReview(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := request.GetOrgId()
	modelName := request.GetModelName()
	modelUUID := request.GetModelUUID()
	userUUID := request.GetUserUUID()
	fromBranch := request.GetParsedBodyAttribute("from_branch")
	if fromBranch == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "From Branch not found in request body")
	} else if fromBranch.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "From Branch cannot be empty")
	}
	fromBranchData := fromBranch.(string)
	fromBranchDb, err := datastore.GetModelBranchByName(orgId, modelName, fromBranchData)
	if err != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "From Branch not found")
	}
	fromBranchUUID := fromBranchDb.UUID
	fromBranchVersion := request.GetParsedBodyAttribute("from_branch_version")
	if fromBranchVersion == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "From Branch Version not found in request body")
	} else if fromBranchVersion.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "From Branch Version cannot be empty")
	}
	fromBranchVersionData := fromBranchVersion.(string)
	fromBranchVersionDb, err := datastore.GetModelBranchVersion(fromBranchUUID, fromBranchVersionData)
	if err != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "From Branch not found")
	}
	fromBranchVersionUUID := fromBranchVersionDb.UUID
	toBranch := request.GetParsedBodyAttribute("to_branch")
	if toBranch == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "To Branch not found in request body")
	} else if toBranch.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "To Branch cannot be empty")
	}
	toBranchData := toBranch.(string)
	toBranchDb, err := datastore.GetModelBranchByName(orgId, modelName, toBranchData)
	if err != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "To Branch not found")
	}
	toBranchUUID := toBranchDb.UUID
	title := request.GetParsedBodyAttribute("title")
	if title == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Title not found in request body")
	} else if title.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "Title cannot be empty")
	}
	titleData := title.(string)
	description := request.GetParsedBodyAttribute("description")
	if description == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "Description not found in request body")
	}
	descriptionData := description.(string)
	isComplete := request.GetParsedBodyAttribute("is_complete")
	var isCompleteData bool
	if isComplete == nil {
		isCompleteData = false
	}
	isCompleteData = isComplete.(bool)
	IsAccepted := request.GetParsedBodyAttribute("is_accepted")
	var isAcceptedData bool
	if IsAccepted == nil {
		isAcceptedData = false
	}
	isAcceptedData = IsAccepted.(bool)
	createdReview, err := datastore.CreateModelReview(modelUUID, userUUID, fromBranchUUID, fromBranchVersionUUID, toBranchUUID, titleData, descriptionData, isCompleteData, isAcceptedData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, []models.ModelReviewResponse{*createdReview}, "Model review created")
}

// UpdateModelReview godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Update review of a model
//	@Description	Update review of a model
//	@Tags			Model
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/model/{modelName}/review/{reviewId}/update [post]
//	@Param			orgId		path	string							true	"Organization Id"
//	@Param			modelName	path	string							true	"Model Name"
//	@Param			reviewId	path	string							true	"Review UUID"
//	@Param			review		body	models.ModelReviewUpdateRequest	true	"Review"
func UpdateModelReview(request *models.Request) *models.Response {
	request.ParseJsonBody()
	reviewUUID := uuid.Must(uuid.FromString(request.GetPathParam("reviewId")))
	review, err := datastore.GetModelReview(reviewUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	if review == nil {
		return models.NewErrorResponse(http.StatusNotFound, "Review with given ID not found")
	}
	title := request.GetParsedBodyAttribute("title")
	description := request.GetParsedBodyAttribute("description")
	isComplete := request.GetParsedBodyAttribute("is_complete")
	isAccepted := request.GetParsedBodyAttribute("is_accepted")
	updatedAttributes := map[string]any{}
	if title != nil {
		updatedAttributes["title"] = title.(string)
	}
	if description != nil {
		updatedAttributes["description"] = description.(string)
	}
	if isComplete != nil {
		updatedAttributes["is_complete"] = isComplete.(bool)
	}
	if isAccepted != nil {
		updatedAttributes["is_accepted"] = isAccepted.(bool)
	}
	updatedDbReview, err := datastore.UpdateModelReview(reviewUUID, updatedAttributes)
	if err != nil {
		if err.Error() == "review already complete" {
			return models.NewErrorResponse(http.StatusBadRequest, err.Error())
		}
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, []models.ModelReviewResponse{*updatedDbReview}, "Model review updated")
}
