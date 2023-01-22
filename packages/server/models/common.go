package models

import uuid "github.com/satori/go.uuid"

// Request models

type ReadmeRequest struct {
	FileType string `json:"file_type"`
	Content  string `json:"content"`
}

type LogRequest struct {
	Data string `json:"data"`
}

// Response models

type ActivityResponse struct {
	UUID     uuid.UUID           `json:"uuid"`
	Category string              `json:"category"`
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

type HashRequest struct {
	Hash string `json:"hash"`
}

type ActivityRequest struct {
	Activity string `json:"activity"`
}

type ReadmeResponse struct {
	UUID          uuid.UUID             `json:"uuid"`
	LatestVersion ReadmeVersionResponse `json:"latest_version"`
}

type ReadmeVersionResponse struct {
	UUID     uuid.UUID `json:"uuid"`
	FileType string    `json:"file_type"`
	Content  string    `json:"content"`
	Version  string    `json:"version"`
}
