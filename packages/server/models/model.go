package models

import uuid "github.com/satori/go.uuid"

// Request models

type CreateModelRequest struct {
	Wiki        string        `json:"wiki"`
	IsPublic    bool          `json:"is_public"`
	Readme      ReadmeRequest `json:"readme"`
	BranchNames []string      `json:"branch_names"`
}

type CreateModelBranchRequest struct {
	BranchName string `json:"branch_name"`
}

type RegisterModelRequest struct {
	Hash    string `json:"hash"`
	Storage string `json:"storage"`
	IsEmpty bool   `json:"is_empty"`
}

type ModelReviewRequest struct {
	FromBranch        string `json:"from_branch"`
	FromBranchVersion string `json:"from_branch_version"`
	ToBranch          string `json:"to_branch"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	IsComplete        bool   `json:"is_complete"`
	IsAccepted        bool   `json:"is_accepted"`
}

type ModelReviewUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsComplete  bool   `json:"is_complete"`
	IsAccepted  bool   `json:"is_accepted"`
}

// Response models

type ModelNameResponse struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type ModelResponse struct {
	UUID      uuid.UUID          `json:"uuid"`
	Name      string             `json:"name"`
	Wiki      string             `json:"wiki"`
	CreatedBy UserHandleResponse `json:"created_by"`
	UpdatedBy UserHandleResponse `json:"updated_by"`
	Readme    ReadmeResponse     `json:"readme"`
	IsPublic  bool               `json:"is_public"`
}

type ModelUserResponse struct {
	User UserHandleResponse `json:"user"`
	Role string             `json:"role"`
}

type ModelBranchNameResponse struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type ModelBranchResponse struct {
	UUID      uuid.UUID         `json:"uuid"`
	Name      string            `json:"name"`
	Model     ModelNameResponse `json:"model"`
	IsDefault bool              `json:"is_default"`
}

type ModelBranchVersionNameResponse struct {
	UUID    uuid.UUID `json:"uuid"`
	Version string    `json:"version"`
}

type ModelBranchVersionResponse struct {
	UUID    uuid.UUID               `json:"uuid"`
	Version string                  `json:"version"`
	Branch  ModelBranchNameResponse `json:"branch"`
	Hash    string                  `json:"hash"`
	Path    PathResponse            `json:"path"`
	Logs    []LogDataResponse       `json:"logs"`
	IsEmpty bool                    `json:"is_empty"`
}

type ModelReviewResponse struct {
	UUID              uuid.UUID                      `json:"uuid"`
	Model             ModelNameResponse              `json:"model"`
	FromBranch        ModelBranchNameResponse        `json:"from_branch"`
	FromBranchVersion ModelBranchVersionNameResponse `json:"from_branch_version"`
	ToBranch          ModelBranchNameResponse        `json:"to_branch"`
	Title             string                         `json:"title"`
	Description       string                         `json:"description"`
	CreatedBy         UserHandleResponse             `json:"created_by"`
	AssignedTo        UserHandleResponse             `json:"assigned_to"`
	IsComplete        bool                           `json:"is_complete"`
	IsAccepted        bool                           `json:"is_accepted"`
}
