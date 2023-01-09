package impl

import (
	"fmt"

	"github.com/PureML-Inc/PureML/server/datastore/dbmodels"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
	"github.com/teris-io/shortid"
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
		return &SQLiteDatastore{}
	}
	return &SQLiteDatastore{
		DB: db,
	}
}

type SQLiteDatastore struct {
	DB *gorm.DB
}

//////////////////////////////////////////////////////////////////////////////
//////////////////////////// ORGANIZATION METHODS ////////////////////////////
//////////////////////////////////////////////////////////////////////////////

func (ds *SQLiteDatastore) GetAllAdminOrgs() ([]models.OrganizationResponse, error) {
	var organizations []dbmodels.Organization
	ds.DB.Find(&organizations)
	var responseOrganizations []models.OrganizationResponse
	for _, org := range organizations {
		responseOrganizations = append(responseOrganizations, models.OrganizationResponse{
			UUID:        org.UUID,
			Name:        org.Name,
			Handle:      org.Handle,
			Avatar:      org.Avatar,
			Description: org.Description,
			JoinCode:    org.JoinCode,
		})
	}
	return responseOrganizations, nil
}

func (ds *SQLiteDatastore) GetOrgByID(orgId uuid.UUID) (*models.OrganizationResponse, error) {
	org := dbmodels.Organization{
		BaseModel: dbmodels.BaseModel{
			UUID: orgId,
		},
	}
	result := ds.DB.First(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.OrganizationResponse{
		UUID:        org.UUID,
		Name:        org.Name,
		Handle:      org.Handle,
		Avatar:      org.Avatar,
		Description: org.Description,
		JoinCode:    org.JoinCode,
	}, nil
}

func (ds *SQLiteDatastore) GetOrgByJoinCode(joinCode string) (*models.OrganizationResponse, error) {
	var org dbmodels.Organization
	result := ds.DB.Where("join_code = ?", joinCode).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.OrganizationResponse{
		UUID:        org.UUID,
		Name:        org.Name,
		Handle:      org.Handle,
		Avatar:      org.Avatar,
		Description: org.Description,
		JoinCode:    org.JoinCode,
	}, nil
}

func (ds *SQLiteDatastore) CreateOrgFromEmail(email string, orgName string, orgDesc string, orgHandle string) (*models.OrganizationResponse, error) {
	org := dbmodels.Organization{
		Name:         orgName,
		Handle:       orgHandle,
		Avatar:       "",
		Description:  orgDesc,
		JoinCode:     shortid.MustGenerate(),
		APITokenHash: "",
	}
	user := dbmodels.User{
		Email: email,
	}
	err := ds.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&org)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Where("email = ?", email).First(&user)
		if result.Error != nil {
			return result.Error
		}
		userOrg := dbmodels.UserOrganizations{
			UserUUID:         user.UUID,
			OrganizationUUID: org.UUID,
			Role:             "owner",
		}
		result = tx.Create(&userOrg)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &models.OrganizationResponse{
		UUID:        org.UUID,
		Name:        org.Name,
		Handle:      org.Handle,
		Avatar:      org.Avatar,
		Description: org.Description,
		JoinCode:    org.JoinCode,
	}, nil
}

func (ds *SQLiteDatastore) GetOrgByHandle(handle string) (*models.OrganizationResponse, error) {
	var org dbmodels.Organization
	result := ds.DB.Where("handle = ?", handle).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.OrganizationResponse{
		UUID:        org.UUID,
		Name:        org.Name,
		Handle:      org.Handle,
		Avatar:      org.Avatar,
		Description: org.Description,
		JoinCode:    org.JoinCode,
	}, nil
}

func (ds *SQLiteDatastore) GetUserOrganizationsByEmail(email string) ([]models.UserOrganizationsResponse, error) {
	var orgs []models.UserOrganizationsResponse
	var tableOrgs []struct {
		UUID   uuid.UUID
		Handle string
		Name   string
		Avatar string
		Role   string
	}
	result := ds.DB.Table("organizations").Select("organizations.uuid, organizations.handle, organizations.name, organizations.avatar, user_organizations.role").Joins("JOIN user_organizations ON user_organizations.organization_uuid = organizations.uuid").Joins("JOIN users ON users.uuid = user_organizations.user_uuid").Where("users.email = ?", email).Scan(&tableOrgs)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, org := range tableOrgs {
		orgs = append(orgs, models.UserOrganizationsResponse{
			Org: models.OrganizationHandleResponse{
				UUID:   org.UUID,
				Handle: org.Handle,
				Name:   org.Name,
				Avatar: org.Avatar,
			},
			Role: org.Role,
		})
	}
	return orgs, nil
}

func (ds *SQLiteDatastore) GetUserOrganizationByOrgIdAndEmail(orgId uuid.UUID, email string) (*models.UserOrganizationsResponse, error) {
	var org models.UserOrganizationsResponse
	result := ds.DB.Table("organizations").Select("organizations.uuid, organizations.handle, organizations.name, organizations.avatar, user_organization.role").Joins("JOIN user_organizations ON user_organizations.organization_uuid = organizations.uuid").Joins("JOIN users ON users.uuid = user_organizations.user_uuid").Where("users.email = ?", email).Where("organizations.uuid = ?", orgId).Scan(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &org, nil
}

func (ds *SQLiteDatastore) CreateUserOrganizationFromEmailAndOrgId(email string, orgId uuid.UUID) (*models.UserOrganizationsResponse, error) {
	var org dbmodels.Organization
	result := ds.DB.First(&org, orgId)
	if result.Error != nil {
		return nil, result.Error
	}
	var user dbmodels.User
	result = ds.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	userOrganization := dbmodels.UserOrganizations{
		OrganizationUUID: org.UUID,
		UserUUID:         user.UUID,
		Role:             "member",
	}
	result = ds.DB.Create(&userOrganization)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserOrganizationsResponse{
		Org: models.OrganizationHandleResponse{
			UUID:   org.UUID,
			Name:   org.Name,
			Handle: org.Handle,
			Avatar: org.Avatar,
		},
		Role: userOrganization.Role,
	}, nil
}

func (ds *SQLiteDatastore) DeleteUserOrganizationFromEmailAndOrgId(email string, orgId uuid.UUID) error {
	var org dbmodels.Organization
	result := ds.DB.First(&org, orgId)
	if result.Error != nil {
		return result.Error
	}
	var user dbmodels.User
	result = ds.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return result.Error
	}
	result = ds.DB.Where("organization_uuid = ?", org.UUID).Where("user_uuid = ?", user.UUID).Delete(&dbmodels.UserOrganizations{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ds *SQLiteDatastore) CreateUserOrganizationFromEmailAndJoinCode(email string, joinCode string) (*models.UserOrganizationsResponse, error) {
	var org dbmodels.Organization
	result := ds.DB.Where("join_code = ?", joinCode).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	var user dbmodels.User
	result = ds.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	userOrganization := dbmodels.UserOrganizations{
		OrganizationUUID: org.UUID,
		UserUUID:         user.UUID,
		Role:             "member",
	}
	result = ds.DB.Create(&userOrganization)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserOrganizationsResponse{
		Org: models.OrganizationHandleResponse{
			UUID:   org.UUID,
			Name:   org.Name,
			Handle: org.Handle,
			Avatar: org.Avatar,
		},
		Role: userOrganization.Role,
	}, nil
}

func (ds *SQLiteDatastore) UpdateOrg(orgId uuid.UUID, orgName string, orgDesc string, orgAvatar string) (*models.OrganizationResponse, error) {
	var org dbmodels.Organization
	result := ds.DB.First(&org, orgId)
	if result.Error != nil {
		return nil, result.Error
	}
	org.Name = orgName
	org.Description = orgDesc
	org.Avatar = orgAvatar
	result = ds.DB.Save(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.OrganizationResponse{
		UUID:        org.UUID,
		Name:        org.Name,
		Handle:      org.Handle,
		Avatar:      org.Avatar,
		Description: org.Description,
		JoinCode:    org.JoinCode,
	}, nil
}

//////////////////////////////////////////////////////////////////////////////
/////////////////////////////// USER METHODS /////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

func (ds *SQLiteDatastore) GetUserByEmail(email string) (*models.UserResponse, error) {
	var user dbmodels.User
	result := ds.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserResponse{
		Name:   user.Name,
		Email:  user.Email,
		Handle: user.Handle,
		Bio:    user.Bio,
		Avatar: user.Avatar,
	}, nil
}

func (ds *SQLiteDatastore) GetUserByHandle(handle string) (*models.UserResponse, error) {
	var user dbmodels.User
	result := ds.DB.Where("handle = ?", handle).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserResponse{
		Name:   user.Name,
		Email:  user.Email,
		Handle: user.Handle,
		Bio:    user.Bio,
		Avatar: user.Avatar,
	}, nil
}