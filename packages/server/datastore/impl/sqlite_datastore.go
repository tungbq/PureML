package impl

import (
	"fmt"

	"github.com/PureML-Inc/PureML/server/models"
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

func (ds *SQLiteDatastore) CreateOrganization(org models.Organization) error {
	result := ds.DB.Create(&org)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
