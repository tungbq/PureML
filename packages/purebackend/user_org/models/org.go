package models

import (
	uuid "github.com/satori/go.uuid"
)

// Request models

type CreateOrgRequest struct {
	Handle      string `json:"handle"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}

type UpdateOrgRequest struct {
	Name string `json:"name"`
	// Handle      string `json:"handle"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
}

// Response models

type OrganizationHandleResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Handle      string    `json:"handle"`
	Name        string    `json:"name"`
	Avatar      string    `json:"avatar"`
	Description string    `json:"description"`
}

type OrganizationResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Handle      string    `json:"handle"`
	Avatar      string    `json:"avatar"`
	Description string    `json:"description"`
	JoinCode    string    `json:"join_code"`
}

type OrganizationResponseWithMembers struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Handle      string    `json:"handle"`
	Avatar      string    `json:"avatar"`
	Description string    `json:"description"`
	JoinCode    string    `json:"join_code"`

	Members []UserHandleRoleResponse `json:"members"`
}

type UserHandleRoleResponse struct {
	UUID   uuid.UUID `json:"uuid"`
	Handle string    `json:"handle"`
	Name   string    `json:"name"`
	Avatar string    `json:"avatar"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
}