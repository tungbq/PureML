package models

// Request models

type LogRequest struct {
	Data               string `json:"data"`
	ModelVersionUUID   string `json:"model_version_uuid"`
	DatasetVersionUUID string `json:"dataset_version_uuid"`
}

// Response models

type ActivityResponse struct {
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
