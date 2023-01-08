package impl

import "github.com/PureML-Inc/PureML/server/models"

func NewTestDatastore() *TestDatastore {
	return &TestDatastore{}
}

type TestDatastore struct {
}

func (ds *TestDatastore) GetAllAdminOrgs() ([]models.Organization, error) {
	return []models.Organization{}, nil
}

func (ds *TestDatastore) GetOrgByID(orgId string) (*models.Organization, error) {
	return nil, nil
}

func (ds *TestDatastore) GetOrgsByUserMail(mailId string) ([]models.OrgAccess, error) {
	return []models.OrgAccess{}, nil
}
