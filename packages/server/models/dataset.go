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

// Response models

type DatasetNameResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type DatasetResponse struct {
	ID        uint                 `json:"id"`
	Name      string               `json:"name"`
	Wiki      string               `json:"wiki"`
	Org       OrganizationResponse `json:"org"`
	CreatedBy UserHandleResponse   `json:"created_by"`
	UpdatedBy UserHandleResponse   `json:"updated_by"`
	IsPublic  bool                 `json:"is_public"`
}

type DatasetUserResponse struct {
	User UserHandleResponse `json:"user"`
	Role string             `json:"role"`
}

type DatasetBranchNameResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type DatasetBranchResponse struct {
	ID        uint                `json:"id"`
	Name      string              `json:"name"`
	Dataset   DatasetNameResponse `json:"dataset"`
	IsDefault bool                `json:"is_default"`
}

type DatasetVersionNameResponse struct {
	ID      uint   `json:"id"`
	Version string `json:"version"`
}

type DatasetVersionResponse struct {
	ID      uint                      `json:"id"`
	Version string                    `json:"version"`
	Branch  DatasetBranchNameResponse `json:"branch"`
	Lineage LineageResponse           `json:"lineage"`
	Hash    string                    `json:"hash"`
	Path    PathResponse              `json:"path"`
}

type LineageResponse struct {
	ID      uint   `json:"id"`
	Lineage string `json:"lineage"`
}

type DatasetReviewResponse struct {
	ID          uint                  `json:"id"`
	FromBranch  DatasetBranchResponse `json:"from_branch"`
	ToBranch    DatasetBranchResponse `json:"to_branch"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	CreatedBy   UserHandleResponse    `json:"created_by"`
	AssignedTo  UserHandleResponse    `json:"assigned_to"`
	IsComplete  bool                  `json:"is_complete"`
	IsAccepted  bool                  `json:"is_accepted"`
}
