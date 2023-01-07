package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"` //TODO: is this needed?
}

// API Models
type Request struct {
	User        User
	Body        []byte
	Headers     map[string]string
	PathParams  map[string]string
	QueryParams map[string]string
}

type ResponseBody struct {
	Status  int
	Data    interface{}
	Message string
}

type Response struct {
	Error      error
	Body       ResponseBody
	StatusCode int
}

// Database Models
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

type Organization struct {
	BaseModel
	Name         string `json:"name" gorm:"not null"`
	Handle       string `json:"handle" gorm:"unique"`
	Avatar       string `json:"avatar"`
	Description  string `json:"description"`
	APITokenHash string `json:"api_token_hash"`
	JoinCode     string `json:"join_code" gorm:"not null"`
}

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

type Model struct {
	BaseModel
	Name      string `json:"name" gorm:"not null;unique:idx_org_model_name"`
	Wiki      string `json:"wiki"`
	OrgID     uint   `json:"org_id" gorm:"not null;unique:idx_org_model_name"`
	CreatedBy uint   `json:"created_by" gorm:"not null"`
	UpdatedBy uint   `json:"updated_by"`
	IsPublic  bool   `json:"is_public" default:"false"`

	Org           Organization `gorm:"foreignKey:OrgID"`
	CreatedByUser User         `gorm:"foreignKey:CreatedBy"`
	UpdatedByUser User         `gorm:"foreignKey:UpdatedBy"`

	Branches []ModelBranch `gorm:"foreignKey:ModelID"`
	Users    []User        `gorm:"many2many:model_users;"`
}

type ModelUser struct {
	ModelID uint `json:"model_id" gorm:"primaryKey"`
	UserID  uint `json:"user_id" gorm:"primaryKey"`
	Role    string
}

type ModelBranch struct {
	BaseModel
	Name      string `json:"name" gorm:"not null;unique:idx_model_branch"`
	ModelID   uint   `json:"model_id" gorm:"not null;unique:idx_model_branch"`
	IsDefault bool   `json:"is_default" default:"false"`

	Model Model `gorm:"foreignKey:ModelID"`

	Versions []ModelVersion `gorm:"foreignKey:BranchID"`
}

type ModelVersion struct {
	BaseModel
	Version  string `json:"version" gorm:"not null;unique:idx_branch_version"`
	BranchID uint   `json:"branch_id" gorm:"not null;unique:idx_branch_version"`
	Hash     string `json:"hash" gorm:"not null;unique"`
	PathID   uint   `json:"path_id"`

	Branch ModelBranch `gorm:"foreignKey:BranchID"`
	Path   Path        `gorm:"foreignKey:PathID"`
}

type ModelReview struct {
	BaseModel
	FromBranchID uint   `json:"from_branch_id" gorm:"not null"`
	ToBranchID   uint   `json:"to_branch_id" gorm:"not null"`
	Title        string `json:"title" gorm:"not null"`
	Description  string `json:"description"`
	CreatedBy    uint   `json:"created_by" gorm:"not null"`
	AssignedTo   uint   `json:"assigned_to"`
	IsComplete   bool   `json:"is_complete" default:"false"`
	IsAccepted   bool   `json:"is_accepted" default:"false"`

	FromBranch     ModelBranch `gorm:"foreignKey:FromBranchID"`
	ToBranch       ModelBranch `gorm:"foreignKey:ToBranchID"`
	CreatedByUser  User        `gorm:"foreignKey:CreatedBy"`
	AssignedToUser User        `gorm:"foreignKey:AssignedTo"`
}

type Path struct {
	BaseModel
	SourceTypeID string `json:"source_type_id" gorm:"not null"`
	SourcePath   string `json:"source_path" gorm:"unique;not null"`

	SourceType SourceType `gorm:"foreignKey:SourceTypeID"`
}

type SourceType struct {
	BaseModel
	Name      string `json:"name" gorm:"not null"`
	PublicURL string `json:"public_url"`
}

type Activity struct {
	BaseModel
	UserID    uint   `json:"user_id" gorm:"not null"`
	Activity  string `json:"activity"`
	ModelID   uint   `json:"model_id"`
	DatasetID uint   `json:"dataset_id"`

	User    User    `gorm:"foreignKey:UserID"`
	Model   Model   `gorm:"foreignKey:ModelID"`
	Dataset Dataset `gorm:"foreignKey:DatasetID"`
}

type Tag struct {
	ModelID   uint   `json:"model_id" gorm:"primaryKey"`
	DatasetID uint   `json:"dataset_id" gorm:"primaryKey"`
	OrgID     uint   `json:"org_id" gorm:"not null;unique:idx_org_tag"`
	Tag       string `json:"tag" gorm:"not null;unique:idx_org_tag"`

	Model   Model        `gorm:"foreignKey:ModelID"`
	Dataset Dataset      `gorm:"foreignKey:DatasetID"`
	Org     Organization `gorm:"foreignKey:OrgID"`
}

type Log struct {
	BaseModel
	Data             string `json:"data"`
	ModelVersionID   uint   `json:"model_version_id"`
	DatasetVersionID uint   `json:"dataset_version_id"`

	ModelVersion   ModelVersion   `gorm:"foreignKey:ModelVersionID"`
	DatasetVersion DatasetVersion `gorm:"foreignKey:DatasetVersionID"`
}
