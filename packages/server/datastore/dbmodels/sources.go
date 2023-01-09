package dbmodels

type Path struct {
	BaseModel
	SourceTypeUUID string `json:"source_type_uuid" gorm:"not null"`
	SourcePath     string `json:"source_path" gorm:"unique;not null"`

	SourceType SourceType `gorm:"foreignKey:SourceTypeUUID"`
}

type SourceType struct {
	BaseModel
	Name      string `json:"name" gorm:"not null"`
	PublicURL string `json:"public_url"`
}
