package models

type Organization struct {
	BaseModel
	Name         string `json:"name" gorm:"not null"`
	Handle       string `json:"handle" gorm:"unique"`
	Avatar       string `json:"avatar"`
	Description  string `json:"description"`
	APITokenHash string `json:"api_token_hash"`
	JoinCode     string `json:"join_code" gorm:"not null"`
}