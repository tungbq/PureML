package dbmodels

import uuid "github.com/satori/go.uuid"

type Model struct {
	BaseModel        `gorm:"embedded"`
	Name             string        `json:"name" gorm:"not null;index:idx_org_model_name,unique"`
	Wiki             string        `json:"wiki"`
	OrganizationUUID uuid.UUID     `json:"org_uuid" gorm:"type:uuid;not null;index:idx_org_model_name,unique"`
	CreatedBy        uuid.UUID     `json:"created_by" gorm:"type:uuid;not null"`
	UpdatedBy        uuid.NullUUID `json:"updated_by" gorm:"type:uuid;"`
	IsPublic         bool          `json:"is_public" default:"false"`

	Org           Organization `gorm:"foreignKey:OrganizationUUID"`
	CreatedByUser User         `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User         `gorm:"foreignKey:UpdatedBy"`

	Readme   Readme        `gorm:"foreignKey:ModelUUID"`
	Branches []ModelBranch `gorm:"foreignKey:ModelUUID"`
	Users    []User        `gorm:"many2many:model_users;"`
}

type ModelUser struct {
	ModelUUID uuid.UUID `json:"model_uuid" gorm:"type:uuid;primaryKey"`
	UserUUID  uuid.UUID `json:"user_uuid" gorm:"type:uuid;primaryKey"`
	Role      string    `json:"role" gorm:"not null;default:member"`
}

type ModelBranch struct {
	BaseModel `gorm:"embedded"`
	Name      string    `json:"name" gorm:"not null;index:idx_model_branch,unique"`
	ModelUUID uuid.UUID `json:"model_uuid" gorm:"type:uuid;not null;index:idx_model_branch,unique"`
	IsDefault bool      `json:"is_default" default:"false"`
	// IsProtected bool      `json:"is_protected" default:"false"`

	Model Model `gorm:"foreignKey:ModelUUID"`

	Versions []ModelVersion `gorm:"foreignKey:BranchUUID"`
}

type ModelVersion struct {
	BaseModel  `gorm:"embedded"`
	Version    string        `json:"version" gorm:"not null;index:idx_model_branch_version,unique"`
	Hash       string        `json:"hash" gorm:"not null;index:idx_model_branch_hash,unique"`
	IsEmpty    bool          `json:"is_empty"`
	BranchUUID uuid.UUID     `json:"branch_uuid" gorm:"type:uuid;not null;index:idx_model_branch_version,unique;index:idx_model_branch_hash,unique"`
	PathUUID   uuid.NullUUID `json:"path_uuid" gorm:"type:uuid;"`
	CreatedBy  uuid.UUID     `json:"created_by" gorm:"type:uuid;"`

	Branch        ModelBranch `gorm:"foreignKey:BranchUUID"`
	Path          Path        `gorm:"foreignKey:PathUUID"`
	CreatedByUser User        `gorm:"foreignKey:CreatedBy"`
}

type ModelReview struct {
	BaseModel             `gorm:"embedded"`
	ModelUUID             uuid.UUID     `json:"model_uuid" gorm:"type:uuid;not null"`
	FromBranchUUID        uuid.UUID     `json:"from_branch_uuid" gorm:"type:uuid;not null"`
	FromBranchVersionUUID uuid.UUID     `json:"from_branch_version_uuid" gorm:"type:uuid;not null"`
	ToBranchUUID          uuid.UUID     `json:"to_branch_uuid" gorm:"type:uuid;not null"`
	Title                 string        `json:"title" gorm:"not null"`
	Description           string        `json:"description"`
	CreatedBy             uuid.UUID     `json:"created_by" gorm:"type:uuid;not null"`
	AssignedTo            uuid.NullUUID `json:"assigned_to" gorm:"type:uuid;"`
	IsComplete            bool          `json:"is_complete" default:"false"`
	IsAccepted            bool          `json:"is_accepted" default:"false"`

	Model             Model        `gorm:"foreignKey:ModelUUID"`
	FromBranch        ModelBranch  `gorm:"foreignKey:FromBranchUUID"`
	FromBranchVersion ModelVersion `gorm:"foreignKey:FromBranchVersionUUID"`
	ToBranch          ModelBranch  `gorm:"foreignKey:ToBranchUUID"`
	CreatedByUser     User         `gorm:"foreignKey:CreatedBy"`
	AssignedToUser    User         `gorm:"foreignKey:AssignedTo"`
}
