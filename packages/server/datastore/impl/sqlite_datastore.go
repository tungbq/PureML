package impl

import (
	"fmt"

	"github.com/PureML-Inc/PureML/server/models"
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

type SQLiteDatastore struct {
	DB *gorm.DB
}

//////////////////////////////////////////////////////////////////////////////
//////////////////////////// ORGANIZATION METHODS ////////////////////////////
//////////////////////////////////////////////////////////////////////////////

func (ds *SQLiteDatastore) GetAllAdminOrgs() ([]models.OrganizationResponse, error) {
	var organizations []models.OrganizationResponse
	ds.DB.Find(&organizations)
	return organizations, nil
}

func (ds *SQLiteDatastore) GetOrgByID(orgId string) (*models.OrganizationResponse, error) {
	var org models.OrganizationResponse
	result := ds.DB.First(&org, orgId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &org, nil
}

func (ds *SQLiteDatastore) GetOrgByJoinCode(joinCode string) (*models.OrganizationResponse, error) {
	var org models.OrganizationResponse
	result := ds.DB.Where("join_code = ?", joinCode).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &org, nil
}

func (ds *SQLiteDatastore) CreateOrgFromEmail(email string, orgName string, orgDesc string, orgHandle string) (*models.OrganizationResponse, error) {
	org := models.Organization{
		Name:         orgName,
		Handle:       orgHandle,
		Avatar:       "",
		Description:  orgDesc,
		JoinCode:     shortid.MustGenerate(),
		APITokenHash: "",

		Users: []models.User{
			{
				Email: email,
			},
		},
	}
	result := ds.DB.Create(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.OrganizationResponse{
		ID:          org.ID,
		Name:        org.Name,
		Handle:      org.Handle,
		Avatar:      org.Avatar,
		Description: org.Description,
		JoinCode:    org.JoinCode,
	}, nil
}

func (ds *SQLiteDatastore) GetOrgByHandle(handle string) (*models.OrganizationResponse, error) {
	var org models.OrganizationResponse
	result := ds.DB.Where("handle = ?", handle).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &org, nil
}

func (ds *SQLiteDatastore) GetUserOrganizationsByEmail(email string) ([]models.UserOrganizationsResponse, error) {
	var orgs []models.UserOrganizationsResponse
	result := ds.DB.Table("organizations").Select("organizations.id, organizations.handle, organizations.name, organizations.avatar, user_organizations.role").Joins("JOIN user_organizations ON user_organizations.organization_id = organizations.id").Joins("JOIN users ON users.id = user_organizations.user_id").Where("users.email = ?", email).Scan(&orgs)
	if result.Error != nil {
		return nil, result.Error
	}
	return orgs, nil
}

func (ds *SQLiteDatastore) GetUserOrganizationByOrgIdAndEmail(orgId string, email string) (*models.UserOrganizationsResponse, error) {
	var org models.UserOrganizationsResponse
	result := ds.DB.Table("organizations").Select("organizations.id, organizations.handle, organizations.name, organizations.avatar").Joins("JOIN user_organizations ON user_organizations.organization_id = organizations.id").Joins("JOIN users ON users.id = user_organizations.user_id").Where("users.email = ?", email).Where("organizations.id = ?", orgId).Scan(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &org, nil
}

func (ds *SQLiteDatastore) CreateUserOrganizationFromEmailAndOrgId(email string, orgId string) (*models.UserOrganizationsResponse, error) {
	var org models.Organization
	result := ds.DB.First(&org, orgId)
	if result.Error != nil {
		return nil, result.Error
	}
	var user models.User
	result = ds.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	userOrganization := models.UserOrganizations{
		OrgID:  org.ID,
		UserID: user.ID,
		Role:   "owner",
	}
	result = ds.DB.Create(&userOrganization)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserOrganizationsResponse{
		Org: models.OrganizationHandleResponse{
			ID:     org.ID,
			Name:   org.Name,
			Handle: org.Handle,
			Avatar: org.Avatar,
		},
		Role: userOrganization.Role,
	}, nil
}

func (ds *SQLiteDatastore) DeleteUserOrganizationFromEmailAndOrgId(email string, orgId string) error {
	var org models.Organization
	result := ds.DB.First(&org, orgId)
	if result.Error != nil {
		return result.Error
	}
	var user models.User
	result = ds.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return result.Error
	}
	result = ds.DB.Where("organization_id = ?", org.ID).Where("user_id = ?", user.ID).Delete(&models.UserOrganizations{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ds *SQLiteDatastore) CreateUserOrganizationFromEmailAndJoinCode(email string, joinCode string) (*models.UserOrganizationsResponse, error) {
	var org models.Organization
	result := ds.DB.Where("join_code = ?", joinCode).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	var user models.User
	result = ds.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	userOrganization := models.UserOrganizations{
		OrgID:  org.ID,
		UserID: user.ID,
		Role:   "member",
	}
	result = ds.DB.Create(&userOrganization)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserOrganizationsResponse{
		Org: models.OrganizationHandleResponse{
			ID:     org.ID,
			Name:   org.Name,
			Handle: org.Handle,
			Avatar: org.Avatar,
		},
		Role: userOrganization.Role,
	}, nil
}

func (ds *SQLiteDatastore) UpdateOrg(orgId string, orgName string, orgDesc string, orgAvatar string) (*models.OrganizationResponse, error) {
	var org models.Organization
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
		ID:          org.ID,
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

func (ds *SQLiteDatastore) GetUser(email string) (*models.UserResponse, error) {
	var user models.User
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
