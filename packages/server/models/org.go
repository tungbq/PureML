package models

type Organization struct {
	Id           string      `json:"id"`
	Name         string      `json:"name"`
	APITokenHash string      `json:"api_token_hash"`
	JoinCode     string      `json:"join_code"`
	Users        interface{} `json:"users"`
}

type OrgAccess struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	APITokenHash string `json:"api_token_hash"`
	Role         string `json:"role"`
}
