package service

import (
	"net/http"

	"github.com/PureML-Inc/PureML/server/datastore"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

// GetDatasetReviews godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get dataset reviews
//	@Description	Get dataset reviews
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/review [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
func GetDatasetReviews(request *models.Request) *models.Response {
	datasetUUID := request.GetDatasetUUID()
	reviews, err := datastore.GetDatasetReviews(datasetUUID)
	if err != nil {
		return models.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	response := models.NewDataResponse(http.StatusOK, reviews, "Dataset Readme version")
	return response
}

// CreateDatasetReview godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Create a new review request for dataset
//	@Description	Create a new review request for dataset
//	@Description	From and To branch names are required (Not UUID)
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/review/create [post]
//	@Param			orgId		path	string						true	"Organization Id"
//	@Param			datasetName	path	string						true	"Dataset Name"
//	@Param			data		body	models.DatasetReviewRequest	true	"Review"
func CreateDatasetReview(request *models.Request) *models.Response {
	request.ParseJsonBody()
	orgId := request.GetOrgId()
	datasetName := request.GetDatasetName()
	datasetUUID := request.GetDatasetUUID()
	userUUID := request.GetUserUUID()
	fromBranch := request.GetParsedBodyAttribute("from_branch")
	if fromBranch == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "From Branch not found in request body")
	} else if fromBranch.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "From Branch cannot be empty")
	}
	fromBranchData := fromBranch.(string)
	fromBranchDb, err := datastore.GetDatasetBranchByName(orgId, datasetName, fromBranchData)
	if err != nil {
		return models.NewErrorResponse(http.StatusBadRequest, "From Branch not found")
	}
	fromBranchUUID := fromBranchDb.UUID
	toBranch := request.GetParsedBodyAttribute("to_branch")
	if toBranch == nil {
		return models.NewErrorResponse(http.StatusBadRequest, "To Branch not found in request body")
	} else if toBranch.(string) == "" {
		return models.NewErrorResponse(http.StatusBadRequest, "To Branch cannot be empty")
	}
	toBranchData := toBranch.(string)
	toBranchDb, err := datastore.GetDatasetBranchByName(orgId, datasetName, toBranchData)
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
	createdReview, err := datastore.CreateDatasetReview(datasetUUID, userUUID, fromBranchUUID, toBranchUUID, titleData, descriptionData, isCompleteData, isAcceptedData)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, []models.DatasetReviewResponse{*createdReview}, "Activity created")
}

// UpdateDatasetReview godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Update review of a dataset
//	@Description	Update review of a dataset
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/review/{reviewId}/update [post]
//	@Param			orgId		path	string								true	"Organization Id"
//	@Param			datasetName	path	string								true	"Dataset Name"
//	@Param			reviewId	path	string								true	"Review UUID"
//	@Param			review		body	models.DatasetReviewUpdateRequest	true	"Review"
func UpdateDatasetReview(request *models.Request) *models.Response {
	request.ParseJsonBody()
	reviewUUID := uuid.Must(uuid.FromString(request.GetPathParam("reviewId")))
	review, err := datastore.GetDatasetReview(reviewUUID)
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
	updatedDbReview, err := datastore.UpdateDatasetReview(reviewUUID, updatedAttributes)
	if err != nil {
		return models.NewServerErrorResponse(err)
	}
	return models.NewDataResponse(http.StatusOK, []models.DatasetReviewResponse{*updatedDbReview}, "Dataset Activity updated")
}
