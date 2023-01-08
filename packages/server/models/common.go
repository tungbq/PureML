package models

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

// Response models

type ActivityResponse struct {
	Activity string              `json:"activity"`
	User     UserHandleResponse  `json:"user"`
	Model    ModelNameResponse   `json:"model"`
	Dataset  DatasetNameResponse `json:"dataset"`
}

type TagResponse struct {
	Tag     string                     `json:"tag"`
	Model   ModelNameResponse          `json:"model"`
	Dataset DatasetNameResponse        `json:"dataset"`
	Org     OrganizationHandleResponse `json:"org"`
}

type LogResponse struct {
	Data           string                     `json:"data"`
	ModelVersion   ModelVersionNameResponse   `json:"model_version"`
	DatasetVersion DatasetVersionNameResponse `json:"dataset_version"`
}
