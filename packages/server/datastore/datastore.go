package datastore

import (
	"github.com/PriyavKaneria/PureML/service/datastore/impl"
	"github.com/PriyavKaneria/PureML/service/models"
)

var ds *impl.SQLiteDatastore = nil

func Init() {
	ds = impl.NewSQLiteDatastore()
}

func GetAllAdminOrgs() ([]models.Organization, error) {
	return ds.GetAllAdminOrgs()
}

func CreateOrganization(org models.Organization) error {
	return ds.CreateOrganization(org)
}

type Datastore interface {
	GetAllAdminOrgs() ([]models.Organization, error)
	CreateOrganization(org models.Organization) error
}
