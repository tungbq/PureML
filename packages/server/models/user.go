package models

import uuid "github.com/satori/go.uuid"

// Response models

type UserHandleResponse struct {
	UUID   uuid.UUID `json:"uuid"`
	Handle string    `json:"handle"`
	Name   string    `json:"name"`
	Avatar string    `json:"avatar"`
	Email  string    `json:"email"`
}

type UserResponse struct {
	UUID   uuid.UUID `json:"uuid"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Handle string    `json:"handle"`
	Bio    string    `json:"bio"`
	Avatar string    `json:"avatar"`
}

type UserOrganizationsResponse struct {
	Org  OrganizationHandleResponse `json:"org"`
	Role string                     `json:"role"`
}
