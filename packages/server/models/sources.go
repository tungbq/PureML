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

// Response models

type PathResponse struct {
	ID         uint               `json:"id"`
	SourcePath string             `json:"source_path"`
	SourceType SourceTypeResponse `json:"source_type"`
}

type SourceTypeResponse struct {
	Name      string `json:"name"`
	PublicURL string `json:"public_url"`
}
