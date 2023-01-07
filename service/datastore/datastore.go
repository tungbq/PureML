package datastore

import (
	"os"

	"github.com/PriyavKaneria/PureML/service/datastore/impl"
	"github.com/PriyavKaneria/PureML/service/models"
)

var ds Datastore = nil

func init() {
	stage := os.Getenv("STAGE")
	if stage == "Testing" {
		//For testing
		ds = impl.NewTestDatastore()
	} else {
		//Real db
		ds = impl.NewMSQLDatastore()
	}
}

func GetAllAdminOrgs() ([]models.Organization, error) {
	return ds.GetAllAdminOrgs()
}

func GetOrgById(orgId string) (*models.Organization, error) {
	return ds.GetOrgByID(orgId)
}

func GetOrgByJoinCode(joinCode string) (*models.Organization, error) {
	return nil, nil
}

func CreateOrgFromMailId(mailId string, orgName string) (*models.Organization, error) {
	return nil, nil
}

func GetOrgAccessesByMailId(mailId string) ([]models.OrgAccess, error) {
	return ds.GetOrgsByUserMail(mailId)
}

func GetOrgAccessByOrgIdAndMailId(orgId string, mailId string) (*models.OrgAccess, error) {
	return nil, nil
}

func CreateOrgAccessFromMailIdAndOrgId(mailId string, orgId string) (*models.OrgAccess, error) {
	return nil, nil
}

func DeleteOrgAccessFromMailIdAndOrgId(mailId string, orgId string) (*models.OrgAccess, error) {
	return nil, nil
}

func CreateOrgAcessFromMailIdAndJoinCode(mailId string, joinCode string) (*models.OrgAccess, error) {
	return nil, nil
}

func UpdateOrg(orgId string, orgName string) (*models.Organization, error) {
	return nil, nil
}

func GetUser(mailId string) (*models.User, error) {
	return nil, nil
}

func GetUserWithOrgAccess(mailId string, orgId string) (*models.User, error) {
	return nil, nil
}

type Datastore interface {
	GetAllAdminOrgs() ([]models.Organization, error)
	GetOrgByID(orgId string) (*models.Organization, error)
	GetOrgsByUserMail(mailId string) ([]models.OrgAccess, error)
}
