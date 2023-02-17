package dbmodels

type Path struct {
	BaseModel      `gorm:"embedded"`
	SourceTypeUUID string `json:"source_type_uuid" gorm:"not null"`
	SourcePath     string `json:"source_path" gorm:"unique;not null"`

	SourceType SourceType `gorm:"foreignKey:SourceTypeUUID"`
}

type SourceType struct {
	BaseModel `gorm:"embedded"`
	OrgUUID   string `json:"org_uuid" gorm:"not null;index:idx_org_source_type,unique"`
	Name      string `json:"name" gorm:"not null;index:idx_org_source_type,unique"`
	PublicURL string `json:"public_url"`

	Org Organization `gorm:"foreignKey:OrgUUID"`
}

type Secret struct {
	BaseModel `gorm:"embedded"`
	OrgUUID   string `json:"org_uuid" gorm:"not null;index:idx_org_secret,unique"`
	Name      string `json:"name" gorm:"not null;index:idx_org_secret,unique"`
	Value     string `json:"value" gorm:"not null"`

	Org Organization `gorm:"foreignKey:OrgUUID"`
}
