package models

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