package models

import uuid "github.com/satori/go.uuid"

// Request models

type LogRequest struct {
	Data string `json:"data"`
}

// Response models

type ActivityResponse struct {
	UUID     uuid.UUID           `json:"uuid"`
	Name     string              `json:"name"`
	Activity string              `json:"activity"`
	User     UserHandleResponse  `json:"user"`
	Model    ModelNameResponse   `json:"model"`
	Dataset  DatasetNameResponse `json:"dataset"`
}

type TagResponse struct {
	Tag     string                     `json:"tag"`
	Model   ModelNameResponse          `json:"model"`
	Dataset DatasetNameResponse        `json:"dataset"`
	Org     OrganizationHandleResponse `json:"org"`
}

type LogResponse struct {
	Data           string                     `json:"data"`
	ModelVersion   ModelVersionNameResponse   `json:"model_version"`
	DatasetVersion DatasetVersionNameResponse `json:"dataset_version"`
}
