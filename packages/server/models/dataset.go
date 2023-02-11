package models

import uuid "github.com/satori/go.uuid"

// Request models

type CreateDatasetRequest struct {
	Wiki        string        `json:"wiki"`
	IsPublic    bool          `json:"is_public"`
	Readme      ReadmeRequest `json:"readme"`
	BranchNames []string      `json:"branch_names"`
}

type CreateDatasetBranchRequest struct {
	BranchName string `json:"branch_name"`
}

type RegisterDatasetRequest struct {
	Hash    string `json:"hash"`
	Lineage string `json:"lineage"`
	Storage string `json:"storage"`
	IsEmpty bool   `json:"is_empty"`
}

type DatasetReviewRequest struct {
	FromBranch        string `json:"from_branch"`
	FromBranchVersion string `json:"from_branch_version"`
	ToBranch          string `json:"to_branch"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	IsComplete        bool   `json:"is_complete"`
	IsAccepted        bool   `json:"is_accepted"`
}

type DatasetReviewUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsComplete  bool   `json:"is_complete"`
	IsAccepted  bool   `json:"is_accepted"`
}

// Response models

type DatasetNameResponse struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type DatasetResponse struct {
	UUID      uuid.UUID            `json:"uuid"`
	Name      string               `json:"name"`
	Wiki      string               `json:"wiki"`
	Org       OrganizationResponse `json:"org"`
	CreatedBy UserHandleResponse   `json:"created_by"`
	UpdatedBy UserHandleResponse   `json:"updated_by"`
	Readme    ReadmeResponse       `json:"readme"`
	IsPublic  bool                 `json:"is_public"`
}

type DatasetUserResponse struct {
	User UserHandleResponse `json:"user"`
	Role string             `json:"role"`
}

type DatasetBranchNameResponse struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type DatasetBranchResponse struct {
	UUID      uuid.UUID           `json:"uuid"`
	Name      string              `json:"name"`
	Dataset   DatasetNameResponse `json:"dataset"`
	IsDefault bool                `json:"is_default"`
}

type DatasetBranchVersionNameResponse struct {
	UUID    uuid.UUID `json:"uuid"`
	Version string    `json:"version"`
}

type DatasetBranchVersionResponse struct {
	UUID      uuid.UUID                 `json:"uuid"`
	Version   string                    `json:"version"`
	Branch    DatasetBranchNameResponse `json:"branch"`
	Lineage   LineageResponse           `json:"lineage"`
	Hash      string                    `json:"hash"`
	Path      PathResponse              `json:"path"`
	IsEmpty   bool                      `json:"is_empty"`
	CreatedBy UserHandleResponse        `json:"created_by"`
}

type LineageResponse struct {
	UUID    uuid.UUID `json:"uuid"`
	Lineage string    `json:"lineage"`
}

type DatasetReviewResponse struct {
	UUID              uuid.UUID                        `json:"uuid"`
	Dataset           DatasetNameResponse              `json:"dataset"`
	FromBranch        DatasetBranchNameResponse        `json:"from_branch"`
	FromBranchVersion DatasetBranchVersionNameResponse `json:"from_branch_version"`
	ToBranch          DatasetBranchNameResponse        `json:"to_branch"`
	Title             string                           `json:"title"`
	Description       string                           `json:"description"`
	CreatedBy         UserHandleResponse               `json:"created_by"`
	AssignedTo        UserHandleResponse               `json:"assigned_to"`
	IsComplete        bool                             `json:"is_complete"`
	IsAccepted        bool                             `json:"is_accepted"`
}
