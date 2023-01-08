package models

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
	UserID uint   `json:"user_id" gorm:"primaryKey"`
	OrgID  uint   `json:"org_id" gorm:"primaryKey"`
	Role   string `json:"role" gorm:"not null;default:member"`
}

// Response models

type UserHandleResponse struct {
	Handle string `json:"handle"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Handle   string `json:"handle"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
	Password string `json:"-"`
}

type UserOrganizationsResponse struct {
	Org  OrganizationHandleResponse `json:"org"`
	Role string                     `json:"role"`
}
