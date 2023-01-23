package impl

import (
	"fmt"

	"github.com/PureML-Inc/PureML/server/datastore/dbmodels"
	uuid "github.com/satori/go.uuid"

	// "github.com/PureML-Inc/PureML/server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewTestSQLiteDatastore() *Datastore {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Error connecting to database")
	}
	err = db.AutoMigrate(
		&dbmodels.Activity{},
		&dbmodels.Dataset{},
		&dbmodels.DatasetBranch{},
		&dbmodels.DatasetReview{},
		&dbmodels.DatasetUser{},
		&dbmodels.DatasetVersion{},
		&dbmodels.Lineage{},
		&dbmodels.Log{},
		&dbmodels.Model{},
		&dbmodels.ModelBranch{},
		&dbmodels.ModelReview{},
		&dbmodels.ModelUser{},
		&dbmodels.ModelVersion{},
		&dbmodels.Organization{},
		&dbmodels.Path{},
		&dbmodels.User{},
		&dbmodels.UserOrganizations{},
	)
	if err != nil {
		return &Datastore{}
	}
	return &Datastore{
		DB: db,
	}
}

type TestDatastore struct {
	DB *gorm.DB
}

func (ds *Datastore) TestGetAllAdminOrgs() ([]dbmodels.Organization, error) {
	return []dbmodels.Organization{}, nil
}

func (ds *Datastore) TestGetOrgByID(orgId uuid.UUID) (*dbmodels.Organization, error) {
	return nil, nil
}

func (ds *Datastore) TestGetOrgsByUserMail(email string) ([]dbmodels.UserOrganizations, error) {
	return []dbmodels.UserOrganizations{}, nil
}
