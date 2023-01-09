package dbmodels

import uuid "github.com/satori/go.uuid"

type Model struct {
	BaseModel
	Name             string    `json:"name" gorm:"not null;unique:idx_org_model_name"`
	Wiki             string    `json:"wiki"`
	OrganizationUUID uuid.UUID `json:"org_uuid" gorm:"not null;unique:idx_org_model_name"`
	CreatedBy        uuid.UUID `json:"created_by" gorm:"not null"`
	UpdatedBy        uuid.UUID `json:"updated_by"`
	IsPublic         bool      `json:"is_public" default:"false"`

	Org           Organization `gorm:"foreignKey:OrganizationUUID"`
	CreatedByUser User         `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User         `gorm:"foreignKey:UpdatedBy"`

	Branches []ModelBranch `gorm:"foreignKey:ModelUUID"`
	Users    []User        `gorm:"many2many:model_users;"`
}

type ModelUser struct {
	ModelUUID uuid.UUID `json:"model_uuid" gorm:"primaryKey"`
	UserUUID  uuid.UUID `json:"user_uuid" gorm:"primaryKey"`
	Role      string
}

type ModelBranch struct {
	BaseModel
	Name      string    `json:"name" gorm:"not null;unique:idx_model_branch"`
	ModelUUID uuid.UUID `json:"model_uuid" gorm:"not null;unique:idx_model_branch"`
	IsDefault bool      `json:"is_default" default:"false"`

	Model Model `gorm:"foreignKey:ModelUUID"`

	Versions []ModelVersion `gorm:"foreignKey:BranchUUID"`
}

type ModelVersion struct {
	BaseModel
	Version    string    `json:"version" gorm:"not null;unique:idx_branch_version"`
	BranchUUID uuid.UUID `json:"branch_uuid" gorm:"not null;unique:idx_branch_version"`
	Hash       string    `json:"hash" gorm:"not null;unique"`
	PathUUID   uuid.UUID `json:"path_uuid"`

	Branch ModelBranch `gorm:"foreignKey:BranchUUID"`
	Path   Path        `gorm:"foreignKey:PathUUID"`
}

type ModelReview struct {
	BaseModel
	FromBranchUUID uuid.UUID `json:"from_branch_uuid" gorm:"not null"`
	ToBranchUUID   uuid.UUID `json:"to_branch_uuid" gorm:"not null"`
	Title          string    `json:"title" gorm:"not null"`
	Description    string    `json:"description"`
	CreatedBy      uuid.UUID `json:"created_by" gorm:"not null"`
	AssignedTo     uuid.UUID `json:"assigned_to"`
	IsComplete     bool      `json:"is_complete" default:"false"`
	IsAccepted     bool      `json:"is_accepted" default:"false"`

	FromBranch     ModelBranch `gorm:"foreignKey:FromBranchUUID"`
	ToBranch       ModelBranch `gorm:"foreignKey:ToBranchUUID"`
	CreatedByUser  User        `gorm:"foreignKey:CreatedBy"`
	AssignedToUser User        `gorm:"foreignKey:AssignedTo"`
}
