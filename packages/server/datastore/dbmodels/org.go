package dbmodels

type Organization struct {
	BaseModel    `gorm:"embedded"`
	Name         string `json:"name" gorm:"not null"`
	Handle       string `json:"handle" gorm:"unique"`
	Avatar       string `json:"avatar"`
	Description  string `json:"description"`
	APITokenHash string `json:"api_token_hash"`
	JoinCode     string `json:"join_code" gorm:"not null"`

	Users   []User   `gorm:"many2many:user_organizations;"` // many to many
	Secrets []Secret `gorm:"foreignKey:OrgUUID"`
}
