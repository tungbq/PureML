package impl

import (
	"fmt"

	"github.com/PriyavKaneria/PureML/service/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewTestSQLiteDatastore() *TestSQLiteDatastore {
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
		return &TestSQLiteDatastore{}
	}
	return &TestSQLiteDatastore{
		DB: db,
	}
}

type TestSQLiteDatastore struct {
	DB *gorm.DB
}

func (ds *TestSQLiteDatastore) GetAllAdminOrgs() ([]models.Organization, error) {
	return []models.Organization{}, nil
}

func (ds *TestDatastore) GetOrgByID(orgId string) (*models.Organization, error) {
	return nil, nil
}

func (ds *TestDatastore) GetOrgsByUserMail(mailId string) ([]models.OrgAccess, error) {
	return []models.OrgAccess{}, nil
}
