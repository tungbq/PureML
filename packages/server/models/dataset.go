package models

type Dataset struct {
	BaseModel
	Name      string `json:"name" gorm:"unique"`
	Wiki      string `json:"wiki"`
	OrgID     uint   `json:"org_id" gorm:"not null"`
	CreatedBy uint   `json:"created_by" gorm:"not null"`
	UpdatedBy uint   `json:"updated_by" gorm:"not null"`
	IsPublic  bool   `json:"is_public" gorm:"not null;default:false"`

	Org           Organization `gorm:"foreignKey:OrgID"`
	CreatedByUser User         `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User         `gorm:"foreignKey:UpdatedBy"`

	Branches []DatasetBranch `gorm:"foreignKey:DatasetID"`
	Users    []User          `gorm:"many2many:dataset_users;"`
}

type DatasetUser struct {
	DatasetID uint   `json:"dataset_id" gorm:"primaryKey"`
	UserID    uint   `json:"user_id" gorm:"primaryKey"`
	Role      string `json:"role" gorm:"not null default:member"`
}

type DatasetBranch struct {
	BaseModel
	Name      string `json:"name" gorm:"not null;unique:idx_dataset_branch"`
	DatasetID uint   `json:"dataset_id" gorm:"not null;unique:idx_dataset_branch"`
	IsDefault bool   `json:"is_default" default:"false"`

	Dataset Dataset `gorm:"foreignKey:DatasetID"`

	Versions []DatasetVersion `gorm:"foreignKey:BranchID"`
}

type DatasetVersion struct {
	BaseModel
	Version   string `json:"version" gorm:"not null;unique:idx_branch_version"`
	BranchID  uint   `json:"branch_id" gorm:"not null;unique:idx_branch_version"`
	LineageID uint   `json:"lineage_id"`
	Hash      string `json:"hash" gorm:"not null"`
	PathID    uint   `json:"path_id"`

	Branch  DatasetBranch `gorm:"foreignKey:BranchID"`
	Lineage Lineage       `gorm:"foreignKey:LineageID"`
	Path    Path          `gorm:"foreignKey:PathID"`
}

type Lineage struct {
	BaseModel
	Lineage string `json:"lineage"`
}

type DatasetReview struct {
	BaseModel
	FromBranchID uint   `json:"from_branch_id" gorm:"not null"`
	ToBranchID   uint   `json:"to_branch_id" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Description  string `json:"description"`
	CreatedBy    uint   `json:"created_by" gorm:"not null"`
	AssignedTo   uint   `json:"assigned_to"`
	IsComplete   bool   `json:"is_complete" default:"false"`
	IsAccepted   bool   `json:"is_accepted" default:"false"`

	FromBranch     DatasetBranch `gorm:"foreignKey:FromBranchID"`
	ToBranch       DatasetBranch `gorm:"foreignKey:ToBranchID"`
	CreatedByUser  User          `gorm:"foreignKey:CreatedBy"`
	AssignedToUser User          `gorm:"foreignKey:AssignedTo"`
}