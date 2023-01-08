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

// Response models

type OrganizationHandleResponse struct {
	Handle string `json:"handle"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type OrganizationResponse struct {
	Name        string `json:"name"`
	Handle      string `json:"handle"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	JoinCode    string `json:"join_code"`
}
