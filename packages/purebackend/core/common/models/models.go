package models

import (
	uuid "github.com/satori/go.uuid"
)

// Request models

type ReadmeRequest struct {
	FileType string `json:"file_type"`
	Content  string `json:"content"`
}

type LogRequest struct {
	Key  string `json:"key"`
	Data string `json:"data"`
}

// Response models

type LogDataResponse struct {
	Key  string `json:"key"`
	Data string `json:"data"`
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
