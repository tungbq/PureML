package impl

import (
	"fmt"

	"github.com/PureML-Inc/PureML/server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewTestSQLiteDatastore() *SQLiteDatastore {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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

func (ds *SQLiteDatastore) TestGetAllAdminOrgs() ([]models.Organization, error) {
	return []models.Organization{}, nil
}

func (ds *SQLiteDatastore) TestGetOrgByID(orgId string) (*models.Organization, error) {
	return nil, nil
}

func (ds *SQLiteDatastore) TestGetOrgsByUserMail(email string) ([]models.UserOrganizations, error) {
	return []models.UserOrganizations{}, nil
}
