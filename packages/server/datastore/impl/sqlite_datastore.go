package impl

import (
	"fmt"

	"github.com/PureML-Inc/PureML/server/models"
	"github.com/teris-io/shortid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDatastore() *SQLiteDatastore {
	db, err := gorm.Open(sqlite.Open("db/pureml.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Error connecting to database")
	}
	err = db.AutoMigrate(
		&models.Activity{},
		&models.Dataset{},
		&models.DatasetBranch{},
		&models.DatasetReview{},
		&models.DatasetUser{},
		&models.DatasetVersion{},
		&models.Lineage{},
		&models.Log{},
		&models.Model{},
		&models.ModelBranch{},
		&models.ModelReview{},
		&models.ModelUser{},
		&models.ModelVersion{},
		&models.Organization{},
		&models.Path{},
		&models.User{},
		&models.UserOrganizations{},
	)
	if err != nil {
		return &SQLiteDatastore{}
	}
	return &SQLiteDatastore{
		DB: db,
	}
}

type SQLiteDatastore struct {
	DB *gorm.DB
}

func (ds *SQLiteDatastore) GetAllAdminOrgs() ([]models.OrganizationResponse, error) {
	var organizations []models.OrganizationResponse
	ds.DB.Find(&organizations)
	return organizations, nil
}

func (ds *SQLiteDatastore) GetOrgByID(orgId string) (*models.OrganizationResponse, error) {
	var org models.OrganizationResponse
	result := ds.DB.First(&org, orgId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &org, nil
}

func (ds *SQLiteDatastore) GetOrgByJoinCode(joinCode string) (*models.OrganizationResponse, error) {
	var org models.OrganizationResponse
	result := ds.DB.Where("join_code = ?", joinCode).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &org, nil
}

func (ds *SQLiteDatastore) CreateOrgFromEmail(email string, orgName string, orgDesc string, orgHandle string) (*models.OrganizationResponse, error) {
	org := models.Organization{
		Name:         orgName,
		Handle:       orgHandle,
		Avatar:       "",
		Description:  orgDesc,
		JoinCode:     shortid.MustGenerate(),
		APITokenHash: "",

		Users: []models.User{
			{
				Email: email,
			},
		},
	}
	result := ds.DB.Create(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.OrganizationResponse{
		ID:          org.ID,
		Name:        org.Name,
		Handle:      org.Handle,
		Avatar:      org.Avatar,
		Description: org.Description,
		JoinCode:    org.JoinCode,
	}, nil
}
