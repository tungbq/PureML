package dbmodels

import (
	uuid "github.com/satori/go.uuid"
)

type Readme struct {
	BaseModel   `gorm:"embedded"`
	ModelUUID   uuid.NullUUID `json:"model_uuid" gorm:"type:uuid"`
	DatasetUUID uuid.NullUUID `json:"dataset_uuid" gorm:"type:uuid"`

	ReadmeVersions []ReadmeVersion `gorm:"foreignKey:ReadmeUUID"`
}

type ReadmeVersion struct {
	BaseModel  `gorm:"embedded"`
	ReadmeUUID uuid.UUID `json:"readme_uuid" gorm:"type:uuid;not null;index:idx_readme_version,unique"`
	FileType   string    `json:"file_type"`
	Content    string    `json:"content"`
	Version    string    `json:"version" gorm:"not null;index:idx_readme_version,unique"`

	Readme Readme `gorm:"foreignKey:ReadmeUUID"`
}
