package dbmodels

import uuid "github.com/satori/go.uuid"

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Handle   string `json:"handle" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`

	Orgs []Organization `gorm:"many2many:user_organizations;"` // many to many
}

type UserOrganizations struct {
	UserUUID         uuid.UUID `json:"user_uuid" gorm:"primaryKey"`
	OrganizationUUID uuid.UUID `json:"organization_uuid" gorm:"primaryKey"`
	Role             string    `json:"role" gorm:"not null;default:member"`
}
