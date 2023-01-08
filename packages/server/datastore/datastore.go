package datastore

import (
	"os"

	"github.com/PureML-Inc/PureML/server/datastore/impl"
	"github.com/PureML-Inc/PureML/server/models"
)

var ds *impl.SQLiteDatastore = nil

func init() {
	stage := os.Getenv("STAGE")
	if stage == "Testing" {
		//For testing
		ds = impl.NewTestSQLiteDatastore()
	} else {
		//Real db
		ds = impl.NewSQLiteDatastore()
	}
}

func GetAllAdminOrgs() ([]models.Organization, error) {
	return ds.GetAllAdminOrgs()
}

func GetOrgById(orgId string) (*models.Organization, error) {
	// return ds.GetOrgByID(orgId)
	return nil, nil
}

func GetOrgByJoinCode(joinCode string) (*models.Organization, error) {
	return nil, nil
}

func CreateOrgFromEmail(email string, orgName string) (*models.Organization, error) {
	return nil, nil
}

func GetUserOrganizationsByEmail(email string) ([]models.UserOrganizations, error) {
	// return ds.GetOrgsByUserMail(email)
	return nil, nil
}

func GetUserOrganizationByOrgIdAndEmail(orgId string, email string) (*models.UserOrganizations, error) {
	return nil, nil
}

func CreateUserOrganizationFromEmailAndOrgId(email string, orgId string) (*models.UserOrganizations, error) {
	return nil, nil
}

func DeleteUserOrganizationFromEmailAndOrgId(email string, orgId string) (*models.UserOrganizations, error) {
	return nil, nil
}

func CreateOrgAcessFromEmailAndJoinCode(email string, joinCode string) (*models.UserOrganizations, error) {
	return nil, nil
}

func UpdateOrg(orgId string, orgName string) (*models.Organization, error) {
	return nil, nil
}

func GetUser(email string) (*models.User, error) {
	return nil, nil
}

func GetUserWithUserOrganization(email string, orgId string) (*models.User, error) {
	return nil, nil
}

type Datastore interface {
	GetAllAdminOrgs() ([]models.Organization, error)
	GetOrgByID(orgId string) (*models.Organization, error)
	GetOrgsByUserMail(email string) ([]models.UserOrganizations, error)
}
