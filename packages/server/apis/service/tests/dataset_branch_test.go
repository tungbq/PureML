package service_test

import (
	"testing"
)

func TestGetDatasetAllBranches(t *testing.T) {
	return
}

// GetDatasetBranch godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Get specific branch of a dataset
//	@Description	Get specific branch of a dataset
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/{branchName} [get]
//	@Param			orgId		path	string	true	"Organization Id"
//	@Param			datasetName	path	string	true	"Dataset Name"
//	@Param			branchName	path	string	true	"Branch Name"
func TestGetDatasetBranch(t *testing.T) {
	return
}

// CreateDatasetBranch godoc
//
//	@Security		ApiKeyAuth
//	@Summary		Create a new branch of a dataset
//	@Description	Create a new branch of a dataset
//	@Tags			Dataset
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Router			/org/{orgId}/dataset/{datasetName}/branch/create [post]
//	@Param			orgId		path	string								true	"Organization Id"
//	@Param			datasetName	path	string								true	"Dataset Name"
//	@Param			branchName	body	models.CreateDatasetBranchRequest	true	"Data"
func TestCreateDatasetBranch(t *testing.T) {
	return
}

// TODO
func TestUpdateDatasetBranch(t *testing.T) {
	return
}

// TODO
func TestDeleteDatasetBranch(t *testing.T) {
	return
}
