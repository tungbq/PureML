package dbmodels

import uuid "github.com/satori/go.uuid"

type Dataset struct {
	BaseModel        `gorm:"embedded"`
	Name             string    `json:"name" gorm:"unique"`
	Wiki             string    `json:"wiki"`
	OrganizationUUID uuid.UUID `json:"org_uuid" gorm:"type:uuid;not null"`
	CreatedBy        uuid.UUID `json:"created_by" gorm:"type:uuid;not null"`
	UpdatedBy        uuid.UUID `json:"updated_by" gorm:"type:uuid;not null"`
	IsPublic         bool      `json:"is_public" gorm:"not null;default:false"`

	Org           Organization `gorm:"foreignKey:OrganizationUUID"`
	CreatedByUser User         `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User         `gorm:"foreignKey:UpdatedBy"`

	Readme   Readme          `gorm:"foreignKey:DatasetUUID"`
	Branches []DatasetBranch `gorm:"foreignKey:DatasetUUID"`
	Users    []User          `gorm:"many2many:dataset_users;"`
}

type DatasetUser struct {
	DatasetUUID uuid.UUID `json:"dataset_uuid" gorm:"type:uuid;primaryKey"`
	UserUUID    uuid.UUID `json:"user_uuid" gorm:"type:uuid;primaryKey"`
	Role        string    `json:"role" gorm:"not null default:member"`
}

type DatasetBranch struct {
	BaseModel   `gorm:"embedded"`
	Name        string    `json:"name" gorm:"not null;index:idx_dataset_branch,unique"`
	DatasetUUID uuid.UUID `json:"dataset_uuid" gorm:"type:uuid;not null;index:idx_dataset_branch,unique"`
	IsDefault   bool      `json:"is_default" default:"false"`

	Dataset Dataset `gorm:"foreignKey:DatasetUUID"`

	Versions []DatasetVersion `gorm:"foreignKey:BranchUUID"`
}

type DatasetVersion struct {
	BaseModel   `gorm:"embedded"`
	Version     string    `json:"version" gorm:"not null;index:idx_dataset_branch_version,unique"`
	BranchUUID  uuid.UUID `json:"branch_uuid" gorm:"type:uuid;not null;index:idx_dataset_branch_version,unique"`
	LineageUUID uuid.UUID `json:"lineage_uuid"`
	Hash        string    `json:"hash" gorm:"not null"`
	PathUUID    uuid.UUID `json:"path_uuid" gorm:"type:uuid;"`
	IsEmpty     bool      `json:"is_empty"`

	Branch  DatasetBranch `gorm:"foreignKey:BranchUUID"`
	Lineage Lineage       `gorm:"foreignKey:LineageUUID"`
	Path    Path          `gorm:"foreignKey:PathUUID"`
}

type Lineage struct {
	BaseModel `gorm:"embedded"`
	Lineage   string `json:"lineage"`
}

type DatasetReview struct {
	BaseModel      `gorm:"embedded"`
	FromBranchUUID uuid.UUID `json:"from_branch_uuid" gorm:"type:uuid;not null"`
	ToBranchUUID   uuid.UUID `json:"to_branch_uuid" gorm:"type:uuid;not null"`
	Title          string    `json:"title" gorm:"not null"`
	Description    string    `json:"description"`
	CreatedBy      uuid.UUID `json:"created_by" gorm:"type:uuid;not null"`
	AssignedTo     uuid.UUID `json:"assigned_to" gorm:"type:uuid;"`
	IsComplete     bool      `json:"is_complete" default:"false"`
	IsAccepted     bool      `json:"is_accepted" default:"false"`

	FromBranch     DatasetBranch `gorm:"foreignKey:FromBranchUUID"`
	ToBranch       DatasetBranch `gorm:"foreignKey:ToBranchUUID"`
	CreatedByUser  User          `gorm:"foreignKey:CreatedBy"`
	AssignedToUser User          `gorm:"foreignKey:AssignedTo"`
}
