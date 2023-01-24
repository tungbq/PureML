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

type ModelVersionNameResponse struct {
	UUID    uuid.UUID `json:"uuid"`
	Version string    `json:"version"`
}

type ModelVersionResponse struct {
	UUID    uuid.UUID               `json:"uuid"`
	Version string                  `json:"version"`
	Branch  ModelBranchNameResponse `json:"branch"`
	Hash    string                  `json:"hash"`
	Path    PathResponse            `json:"path"`
	IsEmpty bool                    `json:"is_empty"`
}

type ModelReviewResponse struct {
	UUID        uuid.UUID           `json:"uuid"`
	FromBranch  ModelBranchResponse `json:"from_branch"`
	ToBranch    ModelBranchResponse `json:"to_branch"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	CreatedBy   UserHandleResponse  `json:"created_by"`
	AssignedTo  UserHandleResponse  `json:"assigned_to"`
	IsComplete  bool                `json:"is_complete"`
	IsAccepted  bool                `json:"is_accepted"`
}
