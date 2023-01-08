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