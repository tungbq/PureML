package models

import uuid "github.com/satori/go.uuid"

// Response models

type PathResponse struct {
	UUID       uuid.UUID          `json:"uuid"`
	SourcePath string             `json:"source_path"`
	SourceType SourceTypeResponse `json:"source_type"`
}

type SourceTypeResponse struct {
	Name      string `json:"name"`
	PublicURL string `json:"public_url"`
}
