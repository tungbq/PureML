package datastore

import (
	"os"

	"github.com/PureML-Inc/PureML/server/datastore/impl"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
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

func GetOrgById(orgId uuid.UUID) (*models.OrganizationResponse, error) {
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

func GetUserOrganizationByOrgIdAndEmail(orgId uuid.UUID, email string) (*models.UserOrganizationsResponse, error) {
	return ds.GetUserOrganizationByOrgIdAndEmail(orgId, email)
}

func CreateUserOrganizationFromEmailAndOrgId(email string, orgId uuid.UUID) (*models.UserOrganizationsResponse, error) {
	return ds.CreateUserOrganizationFromEmailAndOrgId(email, orgId)
}

func DeleteUserOrganizationFromEmailAndOrgId(email string, orgId uuid.UUID) error {
	return ds.DeleteUserOrganizationFromEmailAndOrgId(email, orgId)
}

func CreateUserOrganizationFromEmailAndJoinCode(email string, joinCode string) (*models.UserOrganizationsResponse, error) {
	return ds.CreateUserOrganizationFromEmailAndJoinCode(email, joinCode)
}

func UpdateOrg(orgId uuid.UUID, orgName string, orgDesc string, orgAvatar string) (*models.OrganizationResponse, error) {
	return ds.UpdateOrg(orgId, orgName, orgDesc, orgAvatar)
}

func GetUser(email string) (*models.UserResponse, error) {
	return ds.GetUser(email)
}

type Datastore interface {
	GetAllAdminOrgs() ([]models.OrganizationResponse, error)
	GetOrgByID(orgId uuid.UUID) (*models.OrganizationResponse, error)
	GetOrgByJoinCode(joinCode string) (*models.OrganizationResponse, error)
	CreateOrgFromEmail(email string, orgName string, orgDesc string, orgHandle string) (*models.OrganizationResponse, error)
	GetUserOrganizationsByEmail(email string) ([]models.UserOrganizationsResponse, error)
	GetUserOrganizationByOrgIdAndEmail(orgId uuid.UUID, email string) (*models.UserOrganizationsResponse, error)
	CreateUserOrganizationFromEmailAndOrgId(email string, orgId uuid.UUID) (*models.UserOrganizationsResponse, error)
	DeleteUserOrganizationFromEmailAndOrgId(email string, orgId uuid.UUID) error
	CreateUserOrganizationFromEmailAndJoinCode(email string, joinCode string) (*models.UserOrganizationsResponse, error)
	UpdateOrg(orgId uuid.UUID, orgName string, orgDesc string, orgAvatar string) (*models.OrganizationResponse, error)
	GetUser(email string) (*models.UserResponse, error)
}
