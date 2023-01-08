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

func GetAllAdminOrgs() ([]models.OrganizationResponse, error) {
	return ds.GetAllAdminOrgs()
}

func GetOrgById(orgId string) (*models.OrganizationResponse, error) {
	return ds.GetOrgByID(orgId)
}

func GetOrgByJoinCode(joinCode string) (*models.OrganizationResponse, error) {
	return ds.GetOrgByJoinCode(joinCode)
}

func CreateOrgFromEmail(email string, orgName string, orgDesc string, orgHandle string) (*models.OrganizationResponse, error) {
	return ds.CreateOrgFromEmail(email, orgName, orgDesc, orgHandle)
}

func GetUserOrganizationsByEmail(email string) ([]models.UserOrganizationsResponse, error) {
	return ds.GetUserOrganizationsByEmail(email)
}

func GetUserOrganizationByOrgIdAndEmail(orgId string, email string) (*models.UserOrganizationsResponse, error) {
	return ds.GetUserOrganizationByOrgIdAndEmail(orgId, email)
}

func CreateUserOrganizationFromEmailAndOrgId(email string, orgId string) (*models.UserOrganizationsResponse, error) {
	return ds.CreateUserOrganizationFromEmailAndOrgId(email, orgId)
}

func DeleteUserOrganizationFromEmailAndOrgId(email string, orgId string) error {
	return ds.DeleteUserOrganizationFromEmailAndOrgId(email, orgId)
}

func CreateUserOrganizationFromEmailAndJoinCode(email string, joinCode string) (*models.UserOrganizationsResponse, error) {
	return ds.CreateUserOrganizationFromEmailAndJoinCode(email, joinCode)
}

func UpdateOrg(orgId string, orgName string, orgDesc string, orgAvatar string) (*models.OrganizationResponse, error) {
	return ds.UpdateOrg(orgId, orgName, orgDesc, orgAvatar)
}

func GetUserByEmail(email string) (*models.UserResponse, error) {
	return ds.GetUser(email)
}

func UpdateUser(email string, updatedAttributes map[string]string) (*models.UserResponse, error) {
	return nil, nil
}

func CreateUser(name string, email string, hashedPassword string, shortId string) (*models.UserResponse, error) {
	return nil, nil
}

type Datastore interface {
	GetAllAdminOrgs() ([]models.Organization, error)
	GetOrgByID(orgId string) (*models.Organization, error)
	GetOrgsByUserMail(email string) ([]models.UserOrganizations, error)
}
