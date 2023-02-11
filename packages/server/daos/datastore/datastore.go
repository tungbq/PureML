package impl

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/daos/dbmodels"
	"github.com/PureML-Inc/PureML/server/models"
	puregosqlite "github.com/glebarez/sqlite"
	uuid "github.com/satori/go.uuid"
	"github.com/teris-io/shortid"
	"gorm.io/driver/postgres"
	cgosqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewSQLiteDatastore(optDataDir ...string) *Datastore {
	var dialector gorm.Dialector
	var err error
	var dataDir string
	if len(optDataDir) == 0 || optDataDir[0] == "" {
		// fallback to the default test data directory
		dataDir = config.GetDataDir()
	} else {
		dataDir = optDataDir[0]
	}
	databasePath := fmt.Sprintf("%s/pureml.db", dataDir)
	if config.IsCGOEnabled() {
		dialector = cgosqlite.Open(databasePath)
	} else {
		dialector = puregosqlite.Open(databasePath)
	}
	db, err := gorm.Open(dialector, &gorm.Config{})
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
		&dbmodels.Dataset{},
		&dbmodels.ModelBranch{},
		&dbmodels.ModelReview{},
		&dbmodels.ModelUser{},
		&dbmodels.ModelVersion{},
		&dbmodels.Organization{},
		&dbmodels.Path{},
		&dbmodels.User{},
		&dbmodels.UserOrganizations{},
		&dbmodels.Secret{},
		&dbmodels.Readme{},
		&dbmodels.ReadmeVersion{},
	)
	if err != nil {
		return &Datastore{}
	}
	seedAdminIfNotExists(db)
	return &Datastore{
		DB: db,
	}
}

func NewPostgresDatastore(databaseUrl string) *Datastore {
	dsn := databaseUrl
	// fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
		&dbmodels.Dataset{},
		&dbmodels.ModelBranch{},
		&dbmodels.ModelReview{},
		&dbmodels.ModelUser{},
		&dbmodels.ModelVersion{},
		&dbmodels.Organization{},
		&dbmodels.Path{},
		&dbmodels.User{},
		&dbmodels.UserOrganizations{},
		&dbmodels.Secret{},
		&dbmodels.Readme{},
		&dbmodels.ReadmeVersion{},
	)
	if err != nil {
		return &Datastore{}
	}
	seedAdminIfNotExists(db)
	return &Datastore{
		DB: db,
	}
}

func seedAdminIfNotExists(db *gorm.DB) {
	var user dbmodels.User
	adminDetails := config.GetAdminDetails()
	err := db.Where("uuid = ?", adminDetails["uuid"]).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		// admin user does not exist, create it
		adminUser := dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: adminDetails["uuid"].(uuid.UUID),
			},
			Email:    adminDetails["email"].(string),
			Password: adminDetails["password"].(string),
			Name:     adminDetails["handle"].(string),
			Handle:   adminDetails["handle"].(string),
			Orgs: []dbmodels.Organization{
				{
					BaseModel: dbmodels.BaseModel{
						UUID: adminDetails["uuid"].(uuid.UUID),
					},
					Name:     adminDetails["org_name"].(string),
					Handle:   adminDetails["org_handle"].(string),
					JoinCode: "",
				},
			},
		}
		db.Create(&adminUser)
		var userOrganization dbmodels.UserOrganizations
		db.Where("user_uuid = ? AND organization_uuid = ?", adminUser.UUID, adminUser.UUID).First(&userOrganization)
		userOrganization.Role = "owner"
		db.Save(&userOrganization)
	} else if err != nil {
		fmt.Println(err)
	}
}

type Datastore struct {
	DB *gorm.DB
}

func (ds *Datastore) ExecuteSQL(sql string) error {
	return ds.DB.Exec(sql).Error
}

func (ds *Datastore) Close() error {
	sqlDB, err := ds.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

//////////////////////////// ORGANIZATION METHODS ////////////////////////////

func (ds *Datastore) GetAllAdminOrgs() ([]models.OrganizationResponse, error) {
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

func (ds *Datastore) GetOrgByID(orgId uuid.UUID) (*models.OrganizationResponse, error) {
	org := dbmodels.Organization{
		BaseModel: dbmodels.BaseModel{
			UUID: orgId,
		},
	}
	result := ds.DB.Limit(1).Find(&org)
	if result.RowsAffected == 0 {
		return nil, nil
	}
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

func (ds *Datastore) GetOrgByJoinCode(joinCode string) (*models.OrganizationResponse, error) {
	var org dbmodels.Organization
	result := ds.DB.Where("join_code = ?", joinCode).Limit(1).Find(&org)
	if result.RowsAffected == 0 {
		return nil, nil
	}
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

func (ds *Datastore) CreateOrgFromEmail(email string, orgName string, orgDesc string, orgHandle string) (*models.OrganizationResponse, error) {
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

func (ds *Datastore) GetOrgByHandle(handle string) (*models.OrganizationResponse, error) {
	var org dbmodels.Organization
	result := ds.DB.Where("handle = ?", handle).Limit(1).Find(&org)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.OrganizationResponse{
		UUID:        org.UUID,
		Name:        org.Name,
		Handle:      org.Handle,
		Avatar:      org.Avatar,
		Description: org.Description,
		// JoinCode:    org.JoinCode,
	}, nil
}

func (ds *Datastore) GetUserOrganizationsByEmail(email string) ([]models.UserOrganizationsResponse, error) {
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

func (ds *Datastore) GetUserOrganizationByOrgIdAndUserUUID(orgId uuid.UUID, userUUID uuid.UUID) (*models.UserOrganizationsRoleResponse, error) {
	userOrganizations := dbmodels.UserOrganizations{
		UserUUID:         userUUID,
		OrganizationUUID: orgId,
	}
	result := ds.DB.Limit(1).Find(&userOrganizations)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	userOrgResponse := models.UserOrganizationsRoleResponse{
		UserUUID: userOrganizations.UserUUID,
		OrgUUID:  userOrganizations.OrganizationUUID,
		Role:     userOrganizations.Role,
	}
	return &userOrgResponse, nil
}

func (ds *Datastore) CreateUserOrganizationFromEmailAndOrgId(email string, orgId uuid.UUID) (*models.UserOrganizationsResponse, error) {
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

func (ds *Datastore) DeleteUserOrganizationFromEmailAndOrgId(email string, orgId uuid.UUID) error {
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

func (ds *Datastore) CreateUserOrganizationFromEmailAndJoinCode(email string, joinCode string) (*models.UserOrganizationsResponse, error) {
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

func (ds *Datastore) UpdateOrg(orgId uuid.UUID, updatedAttributes map[string]interface{}) (*models.OrganizationResponse, error) {
	var org dbmodels.Organization
	result := ds.DB.First(&org, orgId)
	if result.Error != nil {
		return nil, result.Error
	}
	if updatedAttributes["name"] != nil {
		org.Name = updatedAttributes["name"].(string)
	}
	if updatedAttributes["desc"] != nil {
		org.Description = updatedAttributes["desc"].(string)
	}
	if updatedAttributes["avatar"] != nil {
		org.Avatar = updatedAttributes["avatar"].(string)
	}
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

func (ds *Datastore) GetOrgAllPublicModels(orgId uuid.UUID) ([]models.ModelResponse, error) {
	var modelsdb []*dbmodels.Model
	result := ds.DB.Preload("CreatedByUser").Preload("UpdatedByUser").Where("is_public = ?", true).Where("organization_uuid = ?", orgId).Find(&modelsdb)
	if result.Error != nil {
		return nil, result.Error
	}
	var returnModels []models.ModelResponse
	for _, model := range modelsdb {
		returnModels = append(returnModels, models.ModelResponse{
			UUID:     model.UUID,
			Name:     model.Name,
			Wiki:     model.Wiki,
			IsPublic: model.IsPublic,
			CreatedBy: models.UserHandleResponse{
				UUID:   model.CreatedByUser.UUID,
				Handle: model.CreatedByUser.Handle,
				Avatar: model.CreatedByUser.Avatar,
				Name:   model.CreatedByUser.Name,
				Email:  model.CreatedByUser.Email,
			},
			UpdatedBy: models.UserHandleResponse{
				UUID:   model.UpdatedByUser.UUID,
				Handle: model.UpdatedByUser.Handle,
				Avatar: model.UpdatedByUser.Avatar,
				Name:   model.UpdatedByUser.Name,
				Email:  model.UpdatedByUser.Email,
			},
		})
	}
	return returnModels, nil
}

func (ds *Datastore) GetOrgAllPublicDatasets(orgId uuid.UUID) ([]models.DatasetResponse, error) {
	var datasetsdb []*dbmodels.Dataset
	result := ds.DB.Preload("CreatedByUser").Preload("UpdatedByUser").Where("is_public = ?", true).Where("organization_uuid = ?", orgId).Find(&datasetsdb)
	if result.Error != nil {
		return nil, result.Error
	}
	var returnDatasets []models.DatasetResponse
	for _, dataset := range datasetsdb {
		returnDatasets = append(returnDatasets, models.DatasetResponse{
			UUID:     dataset.UUID,
			Name:     dataset.Name,
			Wiki:     dataset.Wiki,
			IsPublic: dataset.IsPublic,
			CreatedBy: models.UserHandleResponse{
				UUID:   dataset.CreatedByUser.UUID,
				Handle: dataset.CreatedByUser.Handle,
				Avatar: dataset.CreatedByUser.Avatar,
				Name:   dataset.CreatedByUser.Name,
				Email:  dataset.CreatedByUser.Email,
			},
			UpdatedBy: models.UserHandleResponse{
				UUID:   dataset.UpdatedByUser.UUID,
				Handle: dataset.UpdatedByUser.Handle,
				Avatar: dataset.UpdatedByUser.Avatar,
				Name:   dataset.UpdatedByUser.Name,
				Email:  dataset.UpdatedByUser.Email,
			},
		})
	}
	return returnDatasets, nil
}

/////////////////////////////// USER METHODS /////////////////////////////////

func (ds *Datastore) GetUserByEmail(email string) (*models.UserResponse, error) {
	var user dbmodels.User
	result := ds.DB.Where("email = ?", email).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		return nil, nil
	}
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

func (ds *Datastore) GetUserByHandle(handle string) (*models.UserProfileResponse, error) {
	var user dbmodels.User
	result := ds.DB.Where("handle = ?", handle).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	numberOfDatasets := int64(0)
	ds.DB.Model(&dbmodels.DatasetUser{}).Where("user_uuid = ?", user.BaseModel.UUID).Count(&numberOfDatasets)
	numberOfModel := int64(0)
	ds.DB.Model(&dbmodels.ModelUser{}).Where("user_uuid = ?", user.BaseModel.UUID).Count(&numberOfModel)
	return &models.UserProfileResponse{
		Name:             user.Name,
		Email:            user.Email,
		Handle:           user.Handle,
		Bio:              user.Bio,
		Avatar:           user.Avatar,
		NumberOfModels:   numberOfModel,
		NumberOfDatasets: numberOfDatasets,
	}, nil
}

func (ds *Datastore) GetSecureUserByEmail(email string) (*models.UserResponse, error) {
	var user dbmodels.User
	result := ds.DB.Where("email = ?", email).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserResponse{
		UUID:     user.UUID,
		Name:     user.Name,
		Email:    user.Email,
		Handle:   user.Handle,
		Bio:      user.Bio,
		Avatar:   user.Avatar,
		Password: user.Password,
	}, nil
}

func (ds *Datastore) GetSecureUserByHandle(handle string) (*models.UserResponse, error) {
	var user dbmodels.User
	result := ds.DB.Where("handle = ?", handle).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserResponse{
		UUID:     user.UUID,
		Name:     user.Name,
		Email:    user.Email,
		Handle:   user.Handle,
		Bio:      user.Bio,
		Avatar:   user.Avatar,
		Password: user.Password,
	}, nil
}

func (ds *Datastore) GetUserByUUID(userUUID uuid.UUID) (*models.UserResponse, error) {
	var user dbmodels.User
	result := ds.DB.Limit(1).Find(&user, userUUID)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserResponse{
		UUID:   user.UUID,
		Name:   user.Name,
		Email:  user.Email,
		Handle: user.Handle,
		Bio:    user.Bio,
		Avatar: user.Avatar,
	}, nil
}

func (ds *Datastore) GetUserProfileByUUID(userUUID uuid.UUID) (*models.UserProfileResponse, error) {
	var user dbmodels.User
	result := ds.DB.Limit(1).Find(&user, userUUID)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	numberOfDatasets := int64(0)
	ds.DB.Model(&dbmodels.DatasetUser{}).Where("user_uuid = ?", userUUID).Count(&numberOfDatasets)
	numberOfModel := int64(0)
	ds.DB.Model(&dbmodels.ModelUser{}).Where("user_uuid = ?", userUUID).Count(&numberOfModel)
	return &models.UserProfileResponse{
		UUID:             user.UUID,
		Name:             user.Name,
		Email:            user.Email,
		Handle:           user.Handle,
		Bio:              user.Bio,
		Avatar:           user.Avatar,
		NumberOfModels:   numberOfModel,
		NumberOfDatasets: numberOfDatasets,
	}, nil
}

func (ds *Datastore) CreateUser(name string, email string, handle string, bio string, avatar string, hashedPassword string) (*models.UserResponse, error) {
	user := dbmodels.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Handle:   handle,
		Bio:      bio,
		Avatar:   avatar,

		Orgs: []dbmodels.Organization{
			{
				Name:        "Private",
				Handle:      handle,
				Avatar:      avatar,
				JoinCode:    shortid.MustGenerate(),
				Description: fmt.Sprintf("Private Organization for %s", handle),
			},
		},
	}
	err := ds.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&user)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Table("user_organizations").Where("user_uuid = ?", user.UUID).Where("organization_uuid = ?", user.Orgs[0].UUID).Update("role", "owner")
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &models.UserResponse{
		UUID:   user.UUID,
		Name:   user.Name,
		Email:  user.Email,
		Handle: user.Handle,
		Bio:    user.Bio,
		Avatar: user.Avatar,
	}, nil
}

func (ds *Datastore) UpdateUser(email string, updatedAttributes map[string]interface{}) (*models.UserResponse, error) {
	var user dbmodels.User
	result := ds.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if updatedAttributes["name"] != nil {
		user.Name = updatedAttributes["name"].(string)
	}
	if updatedAttributes["bio"] != nil {
		user.Bio = updatedAttributes["bio"].(string)
	}
	if updatedAttributes["avatar"] != nil {
		user.Avatar = updatedAttributes["avatar"].(string)
	}
	result = ds.DB.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.UserResponse{
		UUID:   user.UUID,
		Name:   user.Name,
		Email:  user.Email,
		Handle: user.Handle,
		Bio:    user.Bio,
		Avatar: user.Avatar,
	}, nil
}

// Helper
func IncrementVersion(latestVersion string) string {
	version := strings.Split(latestVersion, "v")
	versionNumber, _ := strconv.Atoi(version[1])
	newVersionNumber := versionNumber + 1
	newVersion := fmt.Sprintf("v%d", newVersionNumber)
	return newVersion
}

/////////////////////////////// MODEL METHODS/////////////////////////////////

func (ds *Datastore) GetModelByName(orgId uuid.UUID, modelName string) (*models.ModelResponse, error) {
	var model dbmodels.Model
	result := ds.DB.Preload("CreatedByUser").Preload("UpdatedByUser").Preload("Readme.ReadmeVersions", func(db *gorm.DB) *gorm.DB {
		return db.Order("LENGTH(readme_versions.version) DESC").Order("readme_versions.version DESC").Limit(1)
	}).Where("name = ?", modelName).Where("organization_uuid = ?", orgId).Limit(1).Find(&model)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.ModelResponse{
		UUID: model.UUID,
		Name: model.Name,
		Wiki: model.Wiki,
		CreatedBy: models.UserHandleResponse{
			UUID:   model.CreatedByUser.UUID,
			Handle: model.CreatedByUser.Handle,
			Avatar: model.CreatedByUser.Avatar,
			Name:   model.CreatedByUser.Name,
			Email:  model.CreatedByUser.Email,
		},
		UpdatedBy: models.UserHandleResponse{
			UUID:   model.UpdatedByUser.UUID,
			Handle: model.UpdatedByUser.Handle,
			Avatar: model.UpdatedByUser.Avatar,
			Name:   model.UpdatedByUser.Name,
			Email:  model.UpdatedByUser.Email,
		},
		IsPublic: model.IsPublic,
		Readme: models.ReadmeResponse{
			UUID: model.Readme.UUID,
			LatestVersion: models.ReadmeVersionResponse{
				UUID:     model.Readme.ReadmeVersions[0].UUID,
				Version:  model.Readme.ReadmeVersions[0].Version,
				FileType: model.Readme.ReadmeVersions[0].FileType,
				Content:  model.Readme.ReadmeVersions[0].Content,
			},
		},
	}, nil
}

func (ds *Datastore) GetModelByUUID(modelUUID uuid.UUID) (*models.ModelResponse, error) {
	var model dbmodels.Model
	result := ds.DB.Preload("CreatedByUser").Preload("UpdatedByUser").Preload("Readme.ReadmeVersions", func(db *gorm.DB) *gorm.DB {
		return db.Order("LENGTH(readme_versions.version) DESC").Order("readme_versions.version DESC").Limit(1)
	}).Where("uuid = ?", modelUUID).Limit(1).Find(&model)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.ModelResponse{
		UUID: model.UUID,
		Name: model.Name,
		Wiki: model.Wiki,
		CreatedBy: models.UserHandleResponse{
			UUID:   model.CreatedByUser.UUID,
			Handle: model.CreatedByUser.Handle,
			Avatar: model.CreatedByUser.Avatar,
			Name:   model.CreatedByUser.Name,
			Email:  model.CreatedByUser.Email,
		},
		UpdatedBy: models.UserHandleResponse{
			UUID:   model.UpdatedByUser.UUID,
			Handle: model.UpdatedByUser.Handle,
			Avatar: model.UpdatedByUser.Avatar,
			Name:   model.UpdatedByUser.Name,
			Email:  model.UpdatedByUser.Email,
		},
		IsPublic: model.IsPublic,
		Readme: models.ReadmeResponse{
			UUID: model.Readme.UUID,
			LatestVersion: models.ReadmeVersionResponse{
				UUID:     model.Readme.ReadmeVersions[0].UUID,
				Version:  model.Readme.ReadmeVersions[0].Version,
				FileType: model.Readme.ReadmeVersions[0].FileType,
				Content:  model.Readme.ReadmeVersions[0].Content,
			},
		},
	}, nil
}

func (ds *Datastore) GetModelReadmeVersion(modelUUID uuid.UUID, version string) (*models.ReadmeVersionResponse, error) {
	var model dbmodels.Model
	result := ds.DB.Preload("Readme.ReadmeVersions", func(db *gorm.DB) *gorm.DB {
		return db.Where("version = ?", version).Limit(1)
	}).Where("uuid = ?", modelUUID).Limit(1).Find(&model)
	if result.RowsAffected == 0 || len(model.Readme.ReadmeVersions) == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.ReadmeVersionResponse{
		UUID:     model.Readme.ReadmeVersions[0].UUID,
		Version:  model.Readme.ReadmeVersions[0].Version,
		FileType: model.Readme.ReadmeVersions[0].FileType,
		Content:  model.Readme.ReadmeVersions[0].Content,
	}, nil
}

func (ds *Datastore) GetModelReadmeAllVersions(modelUUID uuid.UUID) ([]models.ReadmeVersionResponse, error) {
	var model dbmodels.Model
	result := ds.DB.Preload("Readme.ReadmeVersions").Where("uuid = ?", modelUUID).Limit(1).Find(&model)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	var versions []models.ReadmeVersionResponse
	for _, version := range model.Readme.ReadmeVersions {
		versions = append(versions, models.ReadmeVersionResponse{
			UUID:     version.UUID,
			Version:  version.Version,
			FileType: version.FileType,
			Content:  version.Content,
		})
	}
	return versions, nil
}

func (ds *Datastore) UpdateModelReadme(modelUUID uuid.UUID, fileType string, content string) (*models.ReadmeVersionResponse, error) {
	var model dbmodels.Model
	result := ds.DB.Preload("Readme.ReadmeVersions", func(db *gorm.DB) *gorm.DB {
		return db.Order("LENGTH(version) DESC").Order("version DESC").Limit(1)
	}).Where("uuid = ?", modelUUID).Limit(1).Find(&model)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	var version string
	if len(model.Readme.ReadmeVersions) == 0 {
		version = "v1"
	} else {
		version = IncrementVersion(model.Readme.ReadmeVersions[0].Version)
	}
	readmeVersion := dbmodels.ReadmeVersion{
		Version:  version,
		FileType: fileType,
		Content:  content,
		Readme: dbmodels.Readme{
			BaseModel: dbmodels.BaseModel{
				UUID: model.Readme.UUID,
			},
		},
	}
	result = ds.DB.Create(&readmeVersion)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.ReadmeVersionResponse{
		UUID:     readmeVersion.UUID,
		Version:  readmeVersion.Version,
		FileType: readmeVersion.FileType,
		Content:  readmeVersion.Content,
	}, nil
}

func (ds *Datastore) CreateModel(orgId uuid.UUID, name string, wiki string, isPublic bool, readmeData *models.ReadmeRequest, createdByUser uuid.UUID) (*models.ModelResponse, error) {
	model := dbmodels.Model{
		Name: name,
		Wiki: wiki,
		Org: dbmodels.Organization{
			BaseModel: dbmodels.BaseModel{
				UUID: orgId,
			},
		},
		CreatedByUser: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: createdByUser,
			},
		},
		UpdatedByUser: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: createdByUser,
			},
		},
		IsPublic: isPublic,
		Readme: dbmodels.Readme{
			ReadmeVersions: []dbmodels.ReadmeVersion{
				{
					Version:  "v1",
					FileType: readmeData.FileType,
					Content:  readmeData.Content,
				},
			},
		},
	}
	var user dbmodels.User
	err := ds.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&model)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Where("uuid = ?", createdByUser).First(&user)
		if result.Error != nil {
			return result.Error
		}
		modelUser := dbmodels.ModelUser{
			UserUUID:  user.UUID,
			ModelUUID: model.UUID,
			Role:      "owner",
		}
		result = tx.Create(&modelUser)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &models.ModelResponse{
		UUID: model.UUID,
		Name: model.Name,
		Wiki: model.Wiki,
		CreatedBy: models.UserHandleResponse{
			UUID:   model.CreatedByUser.UUID,
			Handle: model.CreatedByUser.Handle,
			Avatar: model.CreatedByUser.Avatar,
			Name:   model.CreatedByUser.Name,
			Email:  model.CreatedByUser.Email,
		},
		UpdatedBy: models.UserHandleResponse{
			UUID:   model.UpdatedByUser.UUID,
			Handle: model.UpdatedByUser.Handle,
			Avatar: model.UpdatedByUser.Avatar,
			Name:   model.UpdatedByUser.Name,
			Email:  model.UpdatedByUser.Email,
		},
		IsPublic: model.IsPublic,
		Readme: models.ReadmeResponse{
			UUID: model.Readme.UUID,
			LatestVersion: models.ReadmeVersionResponse{
				UUID:     model.Readme.ReadmeVersions[0].UUID,
				Version:  model.Readme.ReadmeVersions[0].Version,
				FileType: model.Readme.ReadmeVersions[0].FileType,
				Content:  model.Readme.ReadmeVersions[0].Content,
			},
		},
	}, nil
}

func (ds *Datastore) GetAllPublicModels() ([]models.ModelResponse, error) {
	var mymodels []dbmodels.Model
	result := ds.DB.Preload("CreatedByUser").Preload("UpdatedByUser").Preload("Org").Where("is_public = ?", true).Find(&mymodels)
	if result.Error != nil {
		return nil, result.Error
	}
	modelResponses := make([]models.ModelResponse, len(mymodels))
	for i, model := range mymodels {
		modelResponses[i] = models.ModelResponse{
			UUID: model.UUID,
			Name: model.Name,
			Wiki: model.Wiki,
			CreatedBy: models.UserHandleResponse{
				UUID:   model.CreatedByUser.UUID,
				Handle: model.CreatedByUser.Handle,
				Avatar: model.CreatedByUser.Avatar,
				Name:   model.CreatedByUser.Name,
				Email:  model.CreatedByUser.Email,
			},
			UpdatedBy: models.UserHandleResponse{
				UUID:   model.UpdatedByUser.UUID,
				Handle: model.UpdatedByUser.Handle,
				Avatar: model.UpdatedByUser.Avatar,
				Name:   model.UpdatedByUser.Name,
				Email:  model.UpdatedByUser.Email,
			},
			IsPublic: model.IsPublic,
		}
	}
	return modelResponses, nil
}

func (ds *Datastore) GetAllModels(orgId uuid.UUID) ([]models.ModelResponse, error) {
	var mymodels []dbmodels.Model
	result := ds.DB.Preload("CreatedByUser").Preload("UpdatedByUser").Where("organization_uuid = ?", orgId).Find(&mymodels)
	if result.Error != nil {
		return nil, result.Error
	}
	modelResponses := make([]models.ModelResponse, len(mymodels))
	for i, model := range mymodels {
		modelResponses[i] = models.ModelResponse{
			UUID: model.UUID,
			Name: model.Name,
			Wiki: model.Wiki,
			CreatedBy: models.UserHandleResponse{
				UUID:   model.CreatedByUser.UUID,
				Handle: model.CreatedByUser.Handle,
				Avatar: model.CreatedByUser.Avatar,
				Name:   model.CreatedByUser.Name,
				Email:  model.CreatedByUser.Email,
			},
			UpdatedBy: models.UserHandleResponse{
				UUID:   model.UpdatedByUser.UUID,
				Handle: model.UpdatedByUser.Handle,
				Avatar: model.UpdatedByUser.Avatar,
				Name:   model.UpdatedByUser.Name,
				Email:  model.UpdatedByUser.Email,
			},
			IsPublic: model.IsPublic,
		}
	}
	return modelResponses, nil
}

func (ds *Datastore) GetModelAllBranches(modelUUID uuid.UUID) ([]models.ModelBranchResponse, error) {
	var modelBranches []dbmodels.ModelBranch
	result := ds.DB.Preload("Model").Where("model_uuid = ?", modelUUID).Find(&modelBranches)
	if result.Error != nil {
		return nil, result.Error
	}
	branches := make([]models.ModelBranchResponse, len(modelBranches))
	for i, branch := range modelBranches {
		branches[i] = models.ModelBranchResponse{
			UUID: branch.UUID,
			Name: branch.Name,
			Model: models.ModelNameResponse{
				UUID: branch.Model.UUID,
				Name: branch.Model.Name,
			},
		}
	}
	return branches, nil
}

func (ds *Datastore) CreateModelBranch(modelUUID uuid.UUID, modelBranchName string) (*models.ModelBranchResponse, error) {
	modelBranch := dbmodels.ModelBranch{
		Name: modelBranchName,
		Model: dbmodels.Model{
			BaseModel: dbmodels.BaseModel{
				UUID: modelUUID,
			},
		},
	}
	err := ds.DB.Create(&modelBranch).Preload("Model").Error
	if err != nil {
		return nil, err
	}
	return &models.ModelBranchResponse{
		UUID: modelBranch.UUID,
		Name: modelBranch.Name,
		Model: models.ModelNameResponse{
			UUID: modelBranch.Model.UUID,
			Name: modelBranch.Model.Name,
		},
	}, nil
}

func (ds *Datastore) RegisterModelFile(modelBranchUUID uuid.UUID, sourceTypeUUID uuid.UUID, filePath string, isEmpty bool, hash string, userUUID uuid.UUID) (*models.ModelBranchVersionResponse, error) {
	sourcePath := dbmodels.Path{
		SourcePath:     filePath,
		SourceTypeUUID: sourceTypeUUID.String(),
	}
	err := ds.DB.Create(&sourcePath).Error
	if err != nil {
		return nil, err
	}
	err = ds.DB.Preload("SourceType").Find(&sourcePath).Error
	if err != nil {
		return nil, err
	}
	latestModelVersion := dbmodels.ModelVersion{
		BranchUUID: modelBranchUUID,
	}
	res := ds.DB.Where(&latestModelVersion).Order("created_at desc").Limit(1).Find(&latestModelVersion)

	var newVersion string
	if res.RowsAffected == 0 {
		newVersion = "v1"
	} else {
		latestVersion := latestModelVersion.Version
		newVersion = IncrementVersion(latestVersion)
	}

	modelVersion := dbmodels.ModelVersion{
		Hash:    hash,
		Version: newVersion,
		Branch: dbmodels.ModelBranch{
			BaseModel: dbmodels.BaseModel{
				UUID: modelBranchUUID,
			},
		},
		CreatedByUser: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: userUUID,
			},
		},
		Path:    sourcePath,
		IsEmpty: isEmpty,
	}
	err = ds.DB.Create(&modelVersion).Preload("Branch").Preload("CreatedByUser").Preload("Path.SourceType").Error
	if err != nil {
		return nil, err
	}

	return &models.ModelBranchVersionResponse{
		UUID:    modelVersion.UUID,
		Hash:    modelVersion.Hash,
		Version: modelVersion.Version,
		Branch: models.ModelBranchNameResponse{
			UUID: modelVersion.Branch.UUID,
			Name: modelVersion.Branch.Name,
		},
		Path: models.PathResponse{
			UUID:       modelVersion.Path.UUID,
			SourcePath: modelVersion.Path.SourcePath,
			SourceType: models.SourceTypeResponse{
				Name:      modelVersion.Path.SourceType.Name,
				PublicURL: modelVersion.Path.SourceType.PublicURL,
			},
		},
		CreatedBy: models.UserHandleResponse{
			UUID:   modelVersion.CreatedByUser.UUID,
			Name:   modelVersion.CreatedByUser.Name,
			Avatar: modelVersion.CreatedByUser.Avatar,
			Email:  modelVersion.CreatedByUser.Email,
			Handle: modelVersion.CreatedByUser.Handle,
		},
		IsEmpty: modelVersion.IsEmpty,
	}, nil
}

func (ds *Datastore) MigrateModelVersionBranch(modelVersion uuid.UUID, toBranch uuid.UUID) (*models.ModelBranchVersionResponse, error) {
	var modelVersionDB dbmodels.ModelVersion
	err := ds.DB.Preload("Branch").Preload("Lineage").Preload("Path").Where("uuid = ?", modelVersion).First(&modelVersionDB).Error
	if err != nil {
		return nil, err
	}
	//Update the branch of the model version
	modelVersionDB.Branch.UUID = toBranch
	modelVersionDB.BaseModel.UUID = uuid.Nil
	err = ds.DB.Save(&modelVersionDB).Error
	if err != nil {
		return nil, err
	}

	return &models.ModelBranchVersionResponse{
		UUID:    modelVersionDB.UUID,
		Hash:    modelVersionDB.Hash,
		Version: modelVersionDB.Version,
		Branch: models.ModelBranchNameResponse{
			UUID: modelVersionDB.BranchUUID,
			Name: modelVersionDB.Branch.Name,
		},
		Path: models.PathResponse{
			UUID:       modelVersionDB.Path.UUID,
			SourcePath: modelVersionDB.Path.SourcePath,
			SourceType: models.SourceTypeResponse{
				Name:      modelVersionDB.Path.SourceType.Name,
				PublicURL: modelVersionDB.Path.SourceType.PublicURL,
			},
		},
		IsEmpty: modelVersionDB.IsEmpty,
	}, nil
}

func (ds *Datastore) GetModelAllVersions(modelUUID uuid.UUID) ([]models.ModelBranchVersionResponse, error) {
	var modelVersions []dbmodels.ModelVersion
	err := ds.DB.Select("model_versions.*").Joins("JOIN model_branches ON model_branches.uuid = model_versions.branch_uuid").Where("model_branches.model_uuid = ?", modelUUID).Find(&modelVersions).Error
	if err != nil {
		return nil, err
	}
	var modelVersionsResponse []models.ModelBranchVersionResponse
	for _, modelVersion := range modelVersions {
		modelVersionsResponse = append(modelVersionsResponse, models.ModelBranchVersionResponse{
			UUID:    modelVersion.UUID,
			Hash:    modelVersion.Hash,
			Version: modelVersion.Version,
			Branch: models.ModelBranchNameResponse{
				UUID: modelVersion.Branch.UUID,
				Name: modelVersion.Branch.Name,
			},
			Path: models.PathResponse{
				UUID:       modelVersion.Path.UUID,
				SourcePath: modelVersion.Path.SourcePath,
				SourceType: models.SourceTypeResponse{
					Name:      modelVersion.Path.SourceType.Name,
					PublicURL: modelVersion.Path.SourceType.PublicURL,
				},
			},
			IsEmpty: modelVersion.IsEmpty,
		})
	}
	return modelVersionsResponse, nil
}

func (ds *Datastore) GetModelBranchByName(orgId uuid.UUID, modelName string, modelBranchName string) (*models.ModelBranchResponse, error) {
	var modelBranch dbmodels.ModelBranch
	model, err := ds.GetModelByName(orgId, modelName)
	if err != nil {
		return nil, err
	}
	res := ds.DB.Where("name = ?", modelBranchName).Where("model_uuid = ?", model.UUID).Preload("Model").Limit(1).Find(&modelBranch)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return &models.ModelBranchResponse{
		UUID: modelBranch.UUID,
		Name: modelBranch.Name,
		Model: models.ModelNameResponse{
			UUID: modelBranch.Model.UUID,
			Name: modelBranch.Model.Name,
		},
		IsDefault: modelBranch.IsDefault,
	}, nil
}

func (ds *Datastore) GetModelBranchByUUID(modelBranchUUID uuid.UUID) (*models.ModelBranchResponse, error) {
	var modelBranch dbmodels.ModelBranch
	res := ds.DB.Where("uuid = ?", modelBranchUUID).Preload("Model").Limit(1).Find(&modelBranch)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return &models.ModelBranchResponse{
		UUID: modelBranch.UUID,
		Name: modelBranch.Name,
		Model: models.ModelNameResponse{
			UUID: modelBranch.Model.UUID,
			Name: modelBranch.Model.Name,
		},
		IsDefault: modelBranch.IsDefault,
	}, nil
}

func (ds *Datastore) GetModelBranchAllVersions(modelBranchUUID uuid.UUID, withLogs bool) ([]models.ModelBranchVersionResponse, error) {
	var modelVersions []dbmodels.ModelVersion
	err := ds.DB.Where("branch_uuid = ?", modelBranchUUID).Preload("Branch").Preload("Path.SourceType").Order("LENGTH(version) DESC").Order("version DESC").Find(&modelVersions).Error
	if err != nil {
		return nil, err
	}
	var modelVersionsResponse []models.ModelBranchVersionResponse
	for _, modelVersion := range modelVersions {
		modelBranchVersion := models.ModelBranchVersionResponse{
			UUID:    modelVersion.UUID,
			Hash:    modelVersion.Hash,
			Version: modelVersion.Version,
			Branch: models.ModelBranchNameResponse{
				UUID: modelVersion.Branch.UUID,
				Name: modelVersion.Branch.Name,
			},
			Path: models.PathResponse{
				UUID:       modelVersion.Path.UUID,
				SourcePath: modelVersion.Path.SourcePath,
				SourceType: models.SourceTypeResponse{
					Name:      modelVersion.Path.SourceType.Name,
					PublicURL: modelVersion.Path.SourceType.PublicURL,
				},
			},
			IsEmpty: modelVersion.IsEmpty,
		}
		if withLogs {
			modelBranchVersion.Logs, err = ds.GetLogForModelVersion(modelVersion.UUID)
			if err != nil {
				return nil, err
			}
		}
		modelVersionsResponse = append(modelVersionsResponse, modelBranchVersion)
	}
	return modelVersionsResponse, nil
}

func (ds *Datastore) GetModelBranchVersion(modelBranchUUID uuid.UUID, version string) (*models.ModelBranchVersionResponse, error) {
	modelVersion := dbmodels.ModelVersion{
		BranchUUID: modelBranchUUID,
	}
	var res *gorm.DB
	if strings.ToLower(version) == "latest" {
		res = ds.DB.Order("created_at desc").Preload("Branch").Preload("Path.SourceType").Limit(1).Find(&modelVersion)
	} else {
		res = ds.DB.Where("version = ?", version).Preload("Branch").Preload("Path.SourceType").Limit(1).Find(&modelVersion)
	}
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &models.ModelBranchVersionResponse{
		UUID:    modelVersion.UUID,
		Hash:    modelVersion.Hash,
		Version: modelVersion.Version,
		Branch: models.ModelBranchNameResponse{
			UUID: modelVersion.Branch.UUID,
			Name: modelVersion.Branch.Name,
		},
		Path: models.PathResponse{
			UUID:       modelVersion.Path.UUID,
			SourcePath: modelVersion.Path.SourcePath,
			SourceType: models.SourceTypeResponse{
				Name:      modelVersion.Path.SourceType.Name,
				PublicURL: modelVersion.Path.SourceType.PublicURL,
			},
		},
		IsEmpty: modelVersion.IsEmpty,
	}, nil
}

/////////////////////////////// DATASET METHODS/////////////////////////////////

func (ds *Datastore) GetDatasetByName(orgId uuid.UUID, datasetName string) (*models.DatasetResponse, error) {
	var dataset dbmodels.Dataset
	result := ds.DB.Preload("CreatedByUser").Preload("UpdatedByUser").Preload("Readme.ReadmeVersions", func(db *gorm.DB) *gorm.DB {
		return db.Order("LENGTH(readme_versions.version) DESC").Order("readme_versions.version DESC").Limit(1)
	}).Where("name = ?", datasetName).Where("organization_uuid = ?", orgId).Limit(1).Find(&dataset)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.DatasetResponse{
		UUID: dataset.UUID,
		Name: dataset.Name,
		Wiki: dataset.Wiki,
		CreatedBy: models.UserHandleResponse{
			UUID:   dataset.CreatedByUser.UUID,
			Handle: dataset.CreatedByUser.Handle,
			Name:   dataset.CreatedByUser.Name,
			Avatar: dataset.CreatedByUser.Avatar,
			Email:  dataset.CreatedByUser.Email,
		},
		UpdatedBy: models.UserHandleResponse{
			UUID:   dataset.UpdatedByUser.UUID,
			Handle: dataset.UpdatedByUser.Handle,
			Name:   dataset.UpdatedByUser.Name,
			Avatar: dataset.UpdatedByUser.Avatar,
			Email:  dataset.UpdatedByUser.Email,
		},
		IsPublic: dataset.IsPublic,
		Readme: models.ReadmeResponse{
			UUID: dataset.Readme.UUID,
			LatestVersion: models.ReadmeVersionResponse{
				UUID:     dataset.Readme.ReadmeVersions[0].UUID,
				Version:  dataset.Readme.ReadmeVersions[0].Version,
				FileType: dataset.Readme.ReadmeVersions[0].FileType,
				Content:  dataset.Readme.ReadmeVersions[0].Content,
			},
		},
	}, nil
}

func (ds *Datastore) GetDatasetByUUID(datasetUUID uuid.UUID) (*models.DatasetResponse, error) {
	var dataset dbmodels.Dataset
	result := ds.DB.Preload("CreatedByUser").Preload("UpdatedByUser").Preload("Readme.ReadmeVersions", func(db *gorm.DB) *gorm.DB {
		return db.Order("LENGTH(readme_versions.version) DESC").Order("readme_versions.version DESC").Limit(1)
	}).Where("uuid = ?", datasetUUID).Limit(1).Find(&dataset)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.DatasetResponse{
		UUID: dataset.UUID,
		Name: dataset.Name,
		Wiki: dataset.Wiki,
		CreatedBy: models.UserHandleResponse{
			UUID:   dataset.CreatedByUser.UUID,
			Handle: dataset.CreatedByUser.Handle,
			Name:   dataset.CreatedByUser.Name,
			Avatar: dataset.CreatedByUser.Avatar,
			Email:  dataset.CreatedByUser.Email,
		},
		UpdatedBy: models.UserHandleResponse{
			UUID:   dataset.UpdatedByUser.UUID,
			Handle: dataset.UpdatedByUser.Handle,
			Name:   dataset.UpdatedByUser.Name,
			Avatar: dataset.UpdatedByUser.Avatar,
			Email:  dataset.UpdatedByUser.Email,
		},
		IsPublic: dataset.IsPublic,
		Readme: models.ReadmeResponse{
			UUID: dataset.Readme.UUID,
			LatestVersion: models.ReadmeVersionResponse{
				UUID:     dataset.Readme.ReadmeVersions[0].UUID,
				Version:  dataset.Readme.ReadmeVersions[0].Version,
				FileType: dataset.Readme.ReadmeVersions[0].FileType,
				Content:  dataset.Readme.ReadmeVersions[0].Content,
			},
		},
	}, nil
}

func (ds *Datastore) GetDatasetReadmeVersion(datasetUUID uuid.UUID, version string) (*models.ReadmeVersionResponse, error) {
	var dataset dbmodels.Dataset
	result := ds.DB.Preload("Readme.ReadmeVersions", func(db *gorm.DB) *gorm.DB {
		return db.Where("version = ?", version).Limit(1)
	}).Where("uuid = ?", datasetUUID).Limit(1).Find(&dataset)
	if result.RowsAffected == 0 || len(dataset.Readme.ReadmeVersions) == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.ReadmeVersionResponse{
		UUID:     dataset.Readme.ReadmeVersions[0].UUID,
		Version:  dataset.Readme.ReadmeVersions[0].Version,
		FileType: dataset.Readme.ReadmeVersions[0].FileType,
		Content:  dataset.Readme.ReadmeVersions[0].Content,
	}, nil
}

func (ds *Datastore) GetDatasetReadmeAllVersions(datasetUUID uuid.UUID) ([]models.ReadmeVersionResponse, error) {
	var dataset dbmodels.Dataset
	result := ds.DB.Preload("Readme.ReadmeVersions").Where("uuid = ?", datasetUUID).Limit(1).Find(&dataset)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	var versions []models.ReadmeVersionResponse
	for _, version := range dataset.Readme.ReadmeVersions {
		versions = append(versions, models.ReadmeVersionResponse{
			UUID:     version.UUID,
			Version:  version.Version,
			FileType: version.FileType,
			Content:  version.Content,
		})
	}
	return versions, nil
}

func (ds *Datastore) UpdateDatasetReadme(datasetUUID uuid.UUID, fileType string, content string) (*models.ReadmeVersionResponse, error) {
	var dataset dbmodels.Dataset
	result := ds.DB.Preload("Readme.ReadmeVersions", func(db *gorm.DB) *gorm.DB {
		return db.Order("LENGTH(version) DESC").Order("version DESC").Limit(1)
	}).Where("uuid = ?", datasetUUID).Limit(1).Find(&dataset)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	var version string
	if len(dataset.Readme.ReadmeVersions) == 0 {
		version = "v1"
	} else {
		version = IncrementVersion(dataset.Readme.ReadmeVersions[0].Version)
	}
	readmeVersion := dbmodels.ReadmeVersion{
		Version:  version,
		FileType: fileType,
		Content:  content,
		Readme: dbmodels.Readme{
			BaseModel: dbmodels.BaseModel{
				UUID: dataset.Readme.UUID,
			},
		},
	}
	result = ds.DB.Create(&readmeVersion)
	if result.Error != nil {
		return nil, result.Error
	}
	return &models.ReadmeVersionResponse{
		UUID:     readmeVersion.UUID,
		Version:  readmeVersion.Version,
		FileType: readmeVersion.FileType,
		Content:  readmeVersion.Content,
	}, nil
}

func (ds *Datastore) CreateDataset(orgId uuid.UUID, name string, wiki string, isPublic bool, readmeData *models.ReadmeRequest, createdByUser uuid.UUID) (*models.DatasetResponse, error) {
	dataset := dbmodels.Dataset{
		Name: name,
		Wiki: wiki,
		Org: dbmodels.Organization{
			BaseModel: dbmodels.BaseModel{
				UUID: orgId,
			},
		},
		CreatedByUser: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: createdByUser,
			},
		},
		UpdatedByUser: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: createdByUser,
			},
		},
		IsPublic: isPublic,
		Readme: dbmodels.Readme{
			ReadmeVersions: []dbmodels.ReadmeVersion{
				{
					Version:  "v1",
					FileType: readmeData.FileType,
					Content:  readmeData.Content,
				},
			},
		},
	}
	var user dbmodels.User
	err := ds.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&dataset)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Where("uuid = ?", createdByUser).First(&user)
		if result.Error != nil {
			return result.Error
		}
		datasetUser := dbmodels.DatasetUser{
			UserUUID:    user.UUID,
			DatasetUUID: dataset.UUID,
			Role:        "owner",
		}
		result = tx.Create(&datasetUser)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &models.DatasetResponse{
		UUID: dataset.UUID,
		Name: dataset.Name,
		Wiki: dataset.Wiki,
		CreatedBy: models.UserHandleResponse{
			UUID:   dataset.CreatedByUser.UUID,
			Handle: dataset.CreatedByUser.Handle,
			Name:   dataset.CreatedByUser.Name,
			Avatar: dataset.CreatedByUser.Avatar,
			Email:  dataset.CreatedByUser.Email,
		},
		UpdatedBy: models.UserHandleResponse{
			UUID:   dataset.UpdatedByUser.UUID,
			Handle: dataset.UpdatedByUser.Handle,
			Name:   dataset.UpdatedByUser.Name,
			Avatar: dataset.UpdatedByUser.Avatar,
			Email:  dataset.UpdatedByUser.Email,
		},
		IsPublic: dataset.IsPublic,
		Readme: models.ReadmeResponse{
			UUID: dataset.Readme.UUID,
			LatestVersion: models.ReadmeVersionResponse{
				UUID:     dataset.Readme.ReadmeVersions[0].UUID,
				Version:  dataset.Readme.ReadmeVersions[0].Version,
				FileType: dataset.Readme.ReadmeVersions[0].FileType,
				Content:  dataset.Readme.ReadmeVersions[0].Content,
			},
		},
	}, nil
}

func (ds *Datastore) GetAllPublicDatasets() ([]models.DatasetResponse, error) {
	var datasets []dbmodels.Dataset
	result := ds.DB.Preload("CreatedByUser").Preload("UpdatedByUser").Preload("Org").Where("is_public = ?", true).Find(&datasets)
	if result.Error != nil {
		return nil, result.Error
	}
	modelResponses := make([]models.DatasetResponse, len(datasets))
	for i, dataset := range datasets {
		modelResponses[i] = models.DatasetResponse{
			UUID: dataset.UUID,
			Name: dataset.Name,
			Wiki: dataset.Wiki,
			CreatedBy: models.UserHandleResponse{
				UUID:   dataset.CreatedByUser.UUID,
				Handle: dataset.CreatedByUser.Handle,
				Avatar: dataset.CreatedByUser.Avatar,
				Name:   dataset.CreatedByUser.Name,
				Email:  dataset.CreatedByUser.Email,
			},
			UpdatedBy: models.UserHandleResponse{
				UUID:   dataset.UpdatedByUser.UUID,
				Handle: dataset.UpdatedByUser.Handle,
				Avatar: dataset.UpdatedByUser.Avatar,
				Name:   dataset.UpdatedByUser.Name,
				Email:  dataset.UpdatedByUser.Email,
			},
			IsPublic: dataset.IsPublic,
		}
	}
	return modelResponses, nil
}

func (ds *Datastore) GetAllDatasets(orgId uuid.UUID) ([]models.DatasetResponse, error) {
	var datasets []dbmodels.Dataset
	result := ds.DB.Preload("CreatedByUser").Preload("UpdatedByUser").Where("organization_uuid = ?", orgId).Find(&datasets)
	if result.Error != nil {
		return nil, result.Error
	}
	datasetResponses := make([]models.DatasetResponse, len(datasets))
	for i, dataset := range datasets {
		datasetResponses[i] = models.DatasetResponse{
			UUID: dataset.UUID,
			Name: dataset.Name,
			Wiki: dataset.Wiki,
			CreatedBy: models.UserHandleResponse{
				UUID:   dataset.CreatedByUser.UUID,
				Handle: dataset.CreatedByUser.Handle,
				Avatar: dataset.CreatedByUser.Avatar,
				Name:   dataset.CreatedByUser.Name,
				Email:  dataset.CreatedByUser.Email,
			},
			UpdatedBy: models.UserHandleResponse{
				UUID:   dataset.UpdatedByUser.UUID,
				Handle: dataset.UpdatedByUser.Handle,
				Avatar: dataset.UpdatedByUser.Avatar,
				Name:   dataset.UpdatedByUser.Name,
				Email:  dataset.UpdatedByUser.Email,
			},
			IsPublic: dataset.IsPublic,
		}
	}
	return datasetResponses, nil
}

func (ds *Datastore) GetDatasetAllBranches(datasetUUID uuid.UUID) ([]models.DatasetBranchResponse, error) {
	var datasetBranches []dbmodels.DatasetBranch
	result := ds.DB.Preload("Dataset").Where("dataset_uuid = ?", datasetUUID).Find(&datasetBranches)
	if result.Error != nil {
		return nil, result.Error
	}
	branches := make([]models.DatasetBranchResponse, len(datasetBranches))
	for i, branch := range datasetBranches {
		branches[i] = models.DatasetBranchResponse{
			UUID: branch.UUID,
			Name: branch.Name,
			Dataset: models.DatasetNameResponse{
				UUID: branch.Dataset.UUID,
				Name: branch.Dataset.Name,
			},
		}
	}
	return branches, nil
}

func (ds *Datastore) CreateDatasetBranch(datasetUUID uuid.UUID, datasetBranchName string) (*models.DatasetBranchResponse, error) {
	datasetBranch := dbmodels.DatasetBranch{
		Name: datasetBranchName,
		Dataset: dbmodels.Dataset{
			BaseModel: dbmodels.BaseModel{
				UUID: datasetUUID,
			},
		},
	}
	err := ds.DB.Create(&datasetBranch).Preload("Dataset").Error
	if err != nil {
		return nil, err
	}
	return &models.DatasetBranchResponse{
		UUID: datasetBranch.UUID,
		Name: datasetBranch.Name,
		Dataset: models.DatasetNameResponse{
			UUID: datasetBranch.Dataset.UUID,
			Name: datasetBranch.Dataset.Name,
		},
	}, nil
}

func (ds *Datastore) RegisterDatasetFile(datasetBranchUUID uuid.UUID, sourceTypeUUID uuid.UUID, filePath string, isEmpty bool, hash string, lineage string, userUUID uuid.UUID) (*models.DatasetBranchVersionResponse, error) {
	sourcePath := dbmodels.Path{
		SourcePath:     filePath,
		SourceTypeUUID: sourceTypeUUID.String(),
	}
	err := ds.DB.Create(&sourcePath).Error
	if err != nil {
		return nil, err
	}
	err = ds.DB.Preload("SourceType").Find(&sourcePath).Error
	if err != nil {
		return nil, err
	}
	latestDatasetVersion := dbmodels.DatasetVersion{
		BranchUUID: datasetBranchUUID,
	}
	res := ds.DB.Where(&latestDatasetVersion).Order("created_at desc").Limit(1).Find(&latestDatasetVersion)

	var newVersion string
	if res.RowsAffected == 0 {
		newVersion = "v1"
	} else {
		latestVersion := latestDatasetVersion.Version
		newVersion = IncrementVersion(latestVersion)
	}

	datasetVersion := dbmodels.DatasetVersion{
		Hash:    hash,
		Version: newVersion,
		Branch: dbmodels.DatasetBranch{
			BaseModel: dbmodels.BaseModel{
				UUID: datasetBranchUUID,
			},
		},
		CreatedByUser: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: userUUID,
			},
		},
		Lineage: dbmodels.Lineage{
			Lineage: lineage,
		},
		Path:    sourcePath,
		IsEmpty: isEmpty,
	}
	err = ds.DB.Create(&datasetVersion).Preload("Lineage").Preload("Branch").Preload("CreatedByUser").Preload("Path.SourceType").Error
	if err != nil {
		return nil, err
	}

	return &models.DatasetBranchVersionResponse{
		UUID:    datasetVersion.UUID,
		Hash:    datasetVersion.Hash,
		Version: datasetVersion.Version,
		Branch: models.DatasetBranchNameResponse{
			UUID: datasetVersion.Branch.UUID,
			Name: datasetVersion.Branch.Name,
		},
		Path: models.PathResponse{
			UUID:       datasetVersion.Path.UUID,
			SourcePath: datasetVersion.Path.SourcePath,
			SourceType: models.SourceTypeResponse{
				Name:      datasetVersion.Path.SourceType.Name,
				PublicURL: datasetVersion.Path.SourceType.PublicURL,
			},
		},
		Lineage: models.LineageResponse{
			UUID:    datasetVersion.Lineage.UUID,
			Lineage: datasetVersion.Lineage.Lineage,
		},
		CreatedBy: models.UserHandleResponse{
			UUID:   datasetVersion.CreatedByUser.UUID,
			Name:   datasetVersion.CreatedByUser.Name,
			Avatar: datasetVersion.CreatedByUser.Avatar,
			Email:  datasetVersion.CreatedByUser.Email,
			Handle: datasetVersion.CreatedByUser.Handle,
		},
		IsEmpty: datasetVersion.IsEmpty,
	}, nil
}

func (ds *Datastore) MigrateDatasetVersionBranch(datasetVersion uuid.UUID, toBranch uuid.UUID) (*models.DatasetBranchVersionResponse, error) {
	var datasetVersionDB dbmodels.DatasetVersion
	err := ds.DB.Preload("Branch").Preload("Lineage").Preload("Path").Where("uuid = ?", datasetVersion).First(&datasetVersionDB).Error
	if err != nil {
		return nil, err
	}
	//Update the branch of the dataset version
	datasetVersionDB.Branch.UUID = toBranch
	datasetVersionDB.BaseModel.UUID = uuid.Nil
	//Create a new lineage record
	datasetVersionDB.Lineage.UUID = uuid.Nil
	err = ds.DB.Save(&datasetVersionDB).Error
	if err != nil {
		return nil, err
	}

	return &models.DatasetBranchVersionResponse{
		UUID:    datasetVersionDB.UUID,
		Hash:    datasetVersionDB.Hash,
		Version: datasetVersionDB.Version,
		Branch: models.DatasetBranchNameResponse{
			UUID: datasetVersionDB.BranchUUID,
			Name: datasetVersionDB.Branch.Name,
		},
		Path: models.PathResponse{
			UUID:       datasetVersionDB.Path.UUID,
			SourcePath: datasetVersionDB.Path.SourcePath,
			SourceType: models.SourceTypeResponse{
				Name:      datasetVersionDB.Path.SourceType.Name,
				PublicURL: datasetVersionDB.Path.SourceType.PublicURL,
			},
		},
		Lineage: models.LineageResponse{
			UUID:    datasetVersionDB.Lineage.UUID,
			Lineage: datasetVersionDB.Lineage.Lineage,
		},
		IsEmpty: datasetVersionDB.IsEmpty,
	}, nil
}

func (ds *Datastore) GetDatasetAllVersions(datasetUUID uuid.UUID) ([]models.DatasetBranchVersionResponse, error) {
	var datasetVersions []dbmodels.DatasetVersion
	err := ds.DB.Select("dataset_versions.*").Joins("JOIN dataset_branches ON dataset_branches.uuid = dataset_versions.branch_uuid").Where("dataset_branches.dataset_uuid = ?", datasetUUID).Find(&datasetVersions).Error
	if err != nil {
		return nil, err
	}
	var datasetVersionsResponse []models.DatasetBranchVersionResponse
	for _, datasetVersion := range datasetVersions {
		datasetVersionsResponse = append(datasetVersionsResponse, models.DatasetBranchVersionResponse{
			UUID:    datasetVersion.UUID,
			Hash:    datasetVersion.Hash,
			Version: datasetVersion.Version,
			Branch: models.DatasetBranchNameResponse{
				UUID: datasetVersion.Branch.UUID,
				Name: datasetVersion.Branch.Name,
			},
			Path: models.PathResponse{
				UUID:       datasetVersion.Path.UUID,
				SourcePath: datasetVersion.Path.SourcePath,
				SourceType: models.SourceTypeResponse{
					Name:      datasetVersion.Path.SourceType.Name,
					PublicURL: datasetVersion.Path.SourceType.PublicURL,
				},
			},
			Lineage: models.LineageResponse{
				UUID:    datasetVersion.Lineage.UUID,
				Lineage: datasetVersion.Lineage.Lineage,
			},
			IsEmpty: datasetVersion.IsEmpty,
		})
	}
	return datasetVersionsResponse, nil
}

func (ds *Datastore) GetDatasetBranchByName(orgId uuid.UUID, datasetName string, datasetBranchName string) (*models.DatasetBranchResponse, error) {
	var datasetBranch dbmodels.DatasetBranch
	dataset, err := ds.GetDatasetByName(orgId, datasetName)
	if err != nil {
		return nil, err
	}
	res := ds.DB.Where("name = ?", datasetBranchName).Where("dataset_uuid = ?", dataset.UUID).Preload("Dataset").Limit(1).Find(&datasetBranch)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &models.DatasetBranchResponse{
		UUID: datasetBranch.UUID,
		Name: datasetBranch.Name,
		Dataset: models.DatasetNameResponse{
			UUID: datasetBranch.Dataset.UUID,
			Name: datasetBranch.Dataset.Name,
		},
		IsDefault: datasetBranch.IsDefault,
	}, nil
}

func (ds *Datastore) GetDatasetBranchByUUID(datasetBranchUUID uuid.UUID) (*models.DatasetBranchResponse, error) {
	var datasetBranch dbmodels.DatasetBranch
	err := ds.DB.Where("uuid = ?", datasetBranchUUID).Preload("Dataset").Find(&datasetBranch).Error
	if err != nil {
		return nil, err
	}
	return &models.DatasetBranchResponse{
		UUID: datasetBranch.UUID,
		Name: datasetBranch.Name,
		Dataset: models.DatasetNameResponse{
			UUID: datasetBranch.Dataset.UUID,
			Name: datasetBranch.Dataset.Name,
		},
		IsDefault: datasetBranch.IsDefault,
	}, nil
}

func (ds *Datastore) GetDatasetBranchAllVersions(datasetBranchUUID uuid.UUID) ([]models.DatasetBranchVersionResponse, error) {
	var datasetVersions []dbmodels.DatasetVersion
	err := ds.DB.Where("branch_uuid = ?", datasetBranchUUID).Preload("Lineage").Preload("Branch").Preload("Path.SourceType").Order("LENGTH(version) DESC").Order("version DESC").Find(&datasetVersions).Error
	if err != nil {
		return nil, err
	}
	var datasetVersionsResponse []models.DatasetBranchVersionResponse
	for _, datasetVersion := range datasetVersions {
		datasetVersionsResponse = append(datasetVersionsResponse, models.DatasetBranchVersionResponse{
			UUID:    datasetVersion.UUID,
			Hash:    datasetVersion.Hash,
			Version: datasetVersion.Version,
			Branch: models.DatasetBranchNameResponse{
				UUID: datasetVersion.Branch.UUID,
				Name: datasetVersion.Branch.Name,
			},
			Path: models.PathResponse{
				UUID:       datasetVersion.Path.UUID,
				SourcePath: datasetVersion.Path.SourcePath,
				SourceType: models.SourceTypeResponse{
					Name:      datasetVersion.Path.SourceType.Name,
					PublicURL: datasetVersion.Path.SourceType.PublicURL,
				},
			},
			Lineage: models.LineageResponse{
				UUID:    datasetVersion.Lineage.UUID,
				Lineage: datasetVersion.Lineage.Lineage,
			},
			IsEmpty: datasetVersion.IsEmpty,
		})
	}
	return datasetVersionsResponse, nil
}

func (ds *Datastore) GetDatasetBranchVersion(datasetBranchUUID uuid.UUID, version string) (*models.DatasetBranchVersionResponse, error) {
	datasetVersion := dbmodels.DatasetVersion{
		BranchUUID: datasetBranchUUID,
	}
	var res *gorm.DB
	if strings.ToLower(version) == "latest" {
		res = ds.DB.Preload("Branch").Preload("Lineage").Preload("Path.SourceType").Order("created_at desc").Limit(1).Find(&datasetVersion)
	} else {
		res = ds.DB.Where("version = ?", version).Preload("Branch").Preload("Lineage").Preload("Path.SourceType").Limit(1).Find(&datasetVersion)
	}
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &models.DatasetBranchVersionResponse{
		UUID:    datasetVersion.UUID,
		Hash:    datasetVersion.Hash,
		Version: datasetVersion.Version,
		Branch: models.DatasetBranchNameResponse{
			UUID: datasetVersion.Branch.UUID,
			Name: datasetVersion.Branch.Name,
		},
		Path: models.PathResponse{
			UUID:       datasetVersion.Path.UUID,
			SourcePath: datasetVersion.Path.SourcePath,
			SourceType: models.SourceTypeResponse{
				Name:      datasetVersion.Path.SourceType.Name,
				PublicURL: datasetVersion.Path.SourceType.PublicURL,
			},
		},
		Lineage: models.LineageResponse{
			UUID:    datasetVersion.Lineage.UUID,
			Lineage: datasetVersion.Lineage.Lineage,
		},
		IsEmpty: datasetVersion.IsEmpty,
	}, nil
}

//////////////////////////////// LOG METHODS /////////////////////////////////

func (ds *Datastore) GetLogForModelVersion(modelVersionUUID uuid.UUID) ([]models.LogDataResponse, error) {
	var logs []dbmodels.Log
	err := ds.DB.Where("model_version_uuid = ?", modelVersionUUID).Preload("ModelVersion").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	var logsResponse []models.LogDataResponse
	for _, log := range logs {
		logsResponse = append(logsResponse, models.LogDataResponse{
			Key:  log.Key,
			Data: log.Data,
		})
	}
	return logsResponse, nil
}

func (ds *Datastore) GetKeyLogForModelVersion(modelVersionUUID uuid.UUID, key string) ([]models.LogResponse, error) {
	var logs []dbmodels.Log
	err := ds.DB.Where("key = ?", key).Where("model_version_uuid = ?", modelVersionUUID).Preload("ModelVersion").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	var logsResponse []models.LogResponse
	for _, log := range logs {
		logsResponse = append(logsResponse, models.LogResponse{
			Key:  log.Key,
			Data: log.Data,
			ModelVersion: models.ModelBranchVersionNameResponse{
				UUID:    log.ModelVersion.UUID,
				Version: log.ModelVersion.Version,
			},
		})
	}
	return logsResponse, nil
}

func (ds *Datastore) CreateLogForModelVersion(key string, data string, modelVersionUUID uuid.UUID) (*models.LogResponse, error) {
	var keyLog dbmodels.Log
	err := ds.DB.Where("key = ?", key).Where("model_version_uuid = ?", modelVersionUUID).First(&keyLog).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	log := dbmodels.Log{
		BaseModel: dbmodels.BaseModel{
			UUID: keyLog.UUID,
		},
		Key:  key,
		Data: data,
		ModelVersion: dbmodels.ModelVersion{
			BaseModel: dbmodels.BaseModel{
				UUID: modelVersionUUID,
			},
		},
	}
	err = ds.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&log).Preload("ModelVersion").Find(&log).Error
	if err != nil {
		return nil, err
	}
	return &models.LogResponse{
		Key:  log.Key,
		Data: log.Data,
		ModelVersion: models.ModelBranchVersionNameResponse{
			UUID:    log.ModelVersion.UUID,
			Version: log.ModelVersion.Version,
		},
	}, nil
}

func (ds *Datastore) GetLogForDatasetVersion(datasetVersion uuid.UUID) ([]models.LogResponse, error) {
	var logs []dbmodels.Log
	err := ds.DB.Where("dataset_version_uuid = ?", datasetVersion).Preload("DatasetVersion").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	var logsResponse []models.LogResponse
	for _, log := range logs {
		logsResponse = append(logsResponse, models.LogResponse{
			Key:  log.Key,
			Data: log.Data,
			DatasetVersion: models.DatasetBranchVersionNameResponse{
				UUID:    log.DatasetVersion.UUID,
				Version: log.DatasetVersion.Version,
			},
		})
	}
	return logsResponse, nil
}

func (ds *Datastore) GetKeyLogForDatasetVersion(datasetVersion uuid.UUID, key string) ([]models.LogResponse, error) {
	var logs []dbmodels.Log
	err := ds.DB.Where("key = ?", key).Where("dataset_version_uuid = ?", datasetVersion).Preload("DatasetVersion").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	var logsResponse []models.LogResponse
	for _, log := range logs {
		logsResponse = append(logsResponse, models.LogResponse{
			Key:  log.Key,
			Data: log.Data,
			DatasetVersion: models.DatasetBranchVersionNameResponse{
				UUID:    log.DatasetVersion.UUID,
				Version: log.DatasetVersion.Version,
			},
		})
	}
	return logsResponse, nil
}

func (ds *Datastore) CreateLogForDatasetVersion(key string, data string, datasetVersionUUID uuid.UUID) (*models.LogResponse, error) {
	var keyLog dbmodels.Log
	err := ds.DB.Where("key = ?", key).Where("dataset_version_uuid = ?", datasetVersionUUID).First(&keyLog).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	log := dbmodels.Log{
		BaseModel: dbmodels.BaseModel{
			UUID: keyLog.UUID,
		},
		Key:  key,
		Data: data,
		DatasetVersion: dbmodels.DatasetVersion{
			BaseModel: dbmodels.BaseModel{
				UUID: datasetVersionUUID,
			},
		},
	}
	err = ds.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&log).Preload("DatasetVersion").Find(&log).Error
	if err != nil {
		return nil, err
	}
	return &models.LogResponse{
		Key:  log.Key,
		Data: log.Data,
		DatasetVersion: models.DatasetBranchVersionNameResponse{
			UUID:    log.DatasetVersion.UUID,
			Version: log.DatasetVersion.Version,
		},
	}, nil
}

//////////////////////////////// ACTIVITY METHODS /////////////////////////////////

func (ds *Datastore) GetModelActivity(modelUUID uuid.UUID, category string) (*models.ActivityResponse, error) {
	var activity dbmodels.Activity
	res := ds.DB.Where("model_uuid = ?", modelUUID).Where("category = ?", category).Preload("Model").Preload("User").Limit(1).Find(&activity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return &models.ActivityResponse{
		UUID:     activity.UUID,
		Category: activity.Category,
		Activity: activity.Activity,
		Model: models.ModelNameResponse{
			UUID: activity.Model.UUID,
			Name: activity.Model.Name,
		},
		User: models.UserHandleResponse{
			UUID: activity.User.UUID,
			Name: activity.User.Name,
		},
	}, nil
}

func (ds *Datastore) CreateModelActivity(modelUUID uuid.UUID, userUUID uuid.UUID, category string, activity string) (*models.ActivityResponse, error) {
	dbactivity := dbmodels.Activity{
		Category: category,
		Activity: activity,
		Model: dbmodels.Model{
			BaseModel: dbmodels.BaseModel{
				UUID: modelUUID,
			},
		},
		User: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: userUUID,
			},
		},
	}
	err := ds.DB.Create(&dbactivity).Preload("Model").Preload("User").Find(&dbactivity).Error
	if err != nil {
		return nil, err
	}
	return &models.ActivityResponse{
		UUID:     dbactivity.UUID,
		Category: dbactivity.Category,
		Activity: dbactivity.Activity,
		Model: models.ModelNameResponse{
			UUID: dbactivity.Model.UUID,
			Name: dbactivity.Model.Name,
		},
		User: models.UserHandleResponse{
			UUID: dbactivity.User.UUID,
			Name: dbactivity.User.Name,
		},
	}, nil
}

func (ds *Datastore) UpdateModelActivity(activityUUID uuid.UUID, updatedAttributes map[string]string) (*models.ActivityResponse, error) {
	var activity dbmodels.Activity
	err := ds.DB.Where("uuid = ?", activityUUID).Preload("Model").Preload("User").Limit(1).Find(&activity).Error
	if err != nil {
		return nil, err
	}
	err = ds.DB.Model(&activity).Updates(updatedAttributes).Error
	if err != nil {
		return nil, err
	}
	return &models.ActivityResponse{
		UUID:     activity.UUID,
		Category: activity.Category,
		Activity: activity.Activity,
		Model: models.ModelNameResponse{
			UUID: activity.Model.UUID,
			Name: activity.Model.Name,
		},
		User: models.UserHandleResponse{
			UUID: activity.User.UUID,
			Name: activity.User.Name,
		},
	}, nil
}

func (ds *Datastore) DeleteModelActivity(activityUUID uuid.UUID) error {
	var activity dbmodels.Activity
	err := ds.DB.Where("uuid = ?", activityUUID).Limit(1).Find(&activity).Error
	if err != nil {
		return err
	}
	err = ds.DB.Delete(&activity).Error
	if err != nil {
		return err
	}
	return nil
}

func (ds *Datastore) GetDatasetActivity(datasetUUID uuid.UUID, category string) (*models.ActivityResponse, error) {
	var activity dbmodels.Activity
	res := ds.DB.Where("dataset_uuid = ?", datasetUUID).Where("category = ?", category).Preload("Dataset").Preload("User").Limit(1).Find(&activity)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return &models.ActivityResponse{
		UUID:     activity.UUID,
		Category: activity.Category,
		Activity: activity.Activity,
		Dataset: models.DatasetNameResponse{
			UUID: activity.Dataset.UUID,
			Name: activity.Dataset.Name,
		},
		User: models.UserHandleResponse{
			UUID: activity.User.UUID,
			Name: activity.User.Name,
		},
	}, nil
}

func (ds *Datastore) CreateDatasetActivity(datasetUUID uuid.UUID, userUUID uuid.UUID, category string, activity string) (*models.ActivityResponse, error) {
	dbactivity := dbmodels.Activity{
		Category: category,
		Activity: activity,
		Dataset: dbmodels.Dataset{
			BaseModel: dbmodels.BaseModel{
				UUID: datasetUUID,
			},
		},
		User: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: userUUID,
			},
		},
	}
	err := ds.DB.Create(&dbactivity).Preload("Dataset").Preload("User").Find(&dbactivity).Error
	if err != nil {
		return nil, err
	}
	return &models.ActivityResponse{
		UUID:     dbactivity.UUID,
		Category: dbactivity.Category,
		Activity: dbactivity.Activity,
		Dataset: models.DatasetNameResponse{
			UUID: dbactivity.Dataset.UUID,
			Name: dbactivity.Dataset.Name,
		},
		User: models.UserHandleResponse{
			UUID: dbactivity.User.UUID,
			Name: dbactivity.User.Name,
		},
	}, nil
}

func (ds *Datastore) UpdateDatasetActivity(activityUUID uuid.UUID, updatedAttributes map[string]string) (*models.ActivityResponse, error) {
	var activity dbmodels.Activity
	err := ds.DB.Where("uuid = ?", activityUUID).Preload("Dataset").Preload("User").Limit(1).Find(&activity).Error
	if err != nil {
		return nil, err
	}
	err = ds.DB.Model(&activity).Updates(updatedAttributes).Error
	if err != nil {
		return nil, err
	}
	return &models.ActivityResponse{
		UUID:     activity.UUID,
		Category: activity.Category,
		Activity: activity.Activity,
		Dataset: models.DatasetNameResponse{
			UUID: activity.Dataset.UUID,
			Name: activity.Dataset.Name,
		},
		User: models.UserHandleResponse{
			UUID: activity.User.UUID,
			Name: activity.User.Name,
		},
	}, nil
}

func (ds *Datastore) DeleteDatasetActivity(activityUUID uuid.UUID) error {
	var activity dbmodels.Activity
	err := ds.DB.Where("uuid = ?", activityUUID).Limit(1).Find(&activity).Error
	if err != nil {
		return err
	}
	err = ds.DB.Delete(&activity).Error
	if err != nil {
		return err
	}
	return nil
}

/////////////////////////////// SECRET API METHODS ///////////////////////////////

func (ds *Datastore) GetSourceTypeByName(orgId uuid.UUID, name string) (uuid.UUID, error) {
	var sourceType dbmodels.SourceType
	res := ds.DB.Where("name = ?", name).Where("org_uuid = ?", orgId).Limit(1).Find(&sourceType)
	if res.RowsAffected == 0 {
		return uuid.Nil, nil
	}
	if res.Error != nil {
		return uuid.Nil, res.Error
	}
	return sourceType.UUID, nil
}

func (ds *Datastore) GetSourceSecret(orgId uuid.UUID, source string) (*models.SourceSecrets, error) {
	var secrets []dbmodels.Secret
	res := ds.DB.Where("org_uuid = ?", orgId).Find(&secrets)
	if res.RowsAffected == 0 {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	var sourceSecret models.SourceSecrets
	switch strings.ToUpper(source) {
	// case "R2":
	// 	var r2Secret R2Secrets
	// 	err := r2Secret.Load(secrets)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	sourceSecret.AccountId = r2Secret.AccountId
	// 	sourceSecret.AccessKeyId = r2Secret.AccessKeyId
	// 	sourceSecret.AccessKeySecret = r2Secret.AccessKeySecret
	// 	sourceSecret.BucketName = r2Secret.BucketName
	// 	sourceSecret.PublicURL = r2Secret.PublicURL
	case "S3":
		for _, secret := range secrets {
			switch strings.ToUpper(secret.Name) {
			case "S3_ACCESS_KEY_ID":
				sourceSecret.AccessKeyId = secret.Value
			case "S3_ACCESS_KEY_SECRET":
				sourceSecret.AccessKeySecret = secret.Value
			case "S3_BUCKET_NAME":
				sourceSecret.BucketName = secret.Value
			case "S3_BUCKET_LOCATION":
				sourceSecret.BucketLocation = secret.Value
			}
		}
		if sourceSecret.AccessKeyId == "" || sourceSecret.AccessKeySecret == "" || sourceSecret.BucketName == "" || sourceSecret.BucketLocation == "" {
			return nil, fmt.Errorf("s3 secrets not found")
		}
	}
	return &sourceSecret, nil
}

// func (ds *Datastore) CreateR2Secrets(orgId uuid.UUID, accountId string, accessKeyId string, accessKeySecret string, bucketName string, publicURL string) (*R2Secrets, error) {
// 	secret := dbmodels.Secret{
// 		Org: dbmodels.Organization{
// 			BaseModel: dbmodels.BaseModel{
// 				UUID: orgId,
// 			},
// 		},
// 	}
// 	err := ds.DB.Transaction(func(tx *gorm.DB) error {
// 		secret.Name = "R2_ACCOUNT_ID"
// 		secret.Value = accountId
// 		err := tx.Create(&secret).Error
// 		if err != nil {
// 			return err
// 		}
// 		secret.UUID = uuid.Nil
// 		secret.Name = "R2_ACCESS_KEY_ID"
// 		secret.Value = accessKeyId
// 		err = tx.Create(&secret).Error
// 		if err != nil {
// 			return err
// 		}
// 		secret.UUID = uuid.Nil
// 		secret.Name = "R2_ACCESS_KEY_SECRET"
// 		secret.Value = accessKeySecret
// 		err = tx.Create(&secret).Error
// 		if err != nil {
// 			return err
// 		}
// 		secret.UUID = uuid.Nil
// 		secret.Name = "R2_BUCKET_NAME"
// 		secret.Value = bucketName
// 		err = tx.Create(&secret).Error
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &R2Secrets{
// 		AccountId:       accountId,
// 		AccessKeyId:     accessKeyId,
// 		AccessKeySecret: accessKeySecret,
// 		BucketName:      bucketName,
// 		PublicURL:       publicURL,
// 	}, nil
// }

func (ds *Datastore) CreateR2Source(orgId uuid.UUID, publicURL string) (*models.SourceTypeResponse, error) {
	sourceType := dbmodels.SourceType{
		Name:      "R2",
		PublicURL: publicURL,
		Org: dbmodels.Organization{
			BaseModel: dbmodels.BaseModel{
				UUID: orgId,
			},
		},
	}
	err := ds.DB.Create(&sourceType).Find(&sourceType).Error
	if err != nil {
		return nil, err
	}
	return &models.SourceTypeResponse{
		UUID:      sourceType.BaseModel.UUID,
		Name:      sourceType.Name,
		PublicURL: sourceType.PublicURL,
	}, nil
}

// func (ds *Datastore) DeleteR2Secrets(orgId uuid.UUID) error {
// 	var secrets []dbmodels.Secret
// 	err := ds.DB.Where("org_uuid = ?", orgId).Where("name LIKE ?", "R2_%").Delete(&secrets).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (ds *Datastore) CreateS3Secrets(orgId uuid.UUID, accessKeyId string, accessKeySecret string, bucketName string, bucketLocation string) (*S3Secrets, error) {
// 	secret := dbmodels.Secret{
// 		Org: dbmodels.Organization{
// 			BaseModel: dbmodels.BaseModel{
// 				UUID: orgId,
// 			},
// 		},
// 	}
// 	err := ds.DB.Transaction(func(tx *gorm.DB) error {
// 		secret.Name = "S3_ACCESS_KEY_ID"
// 		secret.Value = accessKeyId
// 		err := tx.Create(&secret).Error
// 		if err != nil {
// 			return err
// 		}
// 		secret.UUID = uuid.Nil
// 		secret.Name = "S3_ACCESS_KEY_SECRET"
// 		secret.Value = accessKeySecret
// 		err = tx.Create(&secret).Error
// 		if err != nil {
// 			return err
// 		}
// 		secret.UUID = uuid.Nil
// 		secret.Name = "S3_BUCKET_NAME"
// 		secret.Value = bucketName
// 		err = tx.Create(&secret).Error
// 		if err != nil {
// 			return err
// 		}
// 		secret.UUID = uuid.Nil
// 		secret.Name = "S3_BUCKET_LOCATION"
// 		secret.Value = bucketLocation
// 		err = tx.Create(&secret).Error
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	publicURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com", bucketName, bucketLocation)
// 	return &S3Secrets{
// 		AccessKeyId:     accessKeyId,
// 		AccessKeySecret: accessKeySecret,
// 		BucketName:      bucketName,
// 		PublicURL:       publicURL,
// 	}, nil
// }

func (ds *Datastore) CreateS3Source(orgId uuid.UUID, publicURL string) (*models.SourceTypeResponse, error) {
	sourceType := dbmodels.SourceType{
		Name:      "S3",
		PublicURL: publicURL,
		Org: dbmodels.Organization{
			BaseModel: dbmodels.BaseModel{
				UUID: orgId,
			},
		},
	}
	err := ds.DB.Create(&sourceType).Find(&sourceType).Error
	if err != nil {
		return nil, err
	}
	return &models.SourceTypeResponse{
		UUID:      sourceType.BaseModel.UUID,
		Name:      sourceType.Name,
		PublicURL: sourceType.PublicURL,
	}, nil
}

// func (ds *Datastore) DeleteS3Secrets(orgId uuid.UUID) error {
// 	var secrets []dbmodels.Secret
// 	err := ds.DB.Where("org_uuid = ?", orgId).Where("name LIKE ?", "S3_%").Delete(&secrets).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (ds *Datastore) CreateLocalSource(orgId uuid.UUID) (*models.SourceTypeResponse, error) {
	sourceType := dbmodels.SourceType{
		Name: "Local",
		Org: dbmodels.Organization{
			BaseModel: dbmodels.BaseModel{
				UUID: orgId,
			},
		},
		PublicURL: "file://",
	}
	err := ds.DB.Create(&sourceType).Find(&sourceType).Error
	if err != nil {
		return nil, err
	}
	return &models.SourceTypeResponse{
		UUID:      sourceType.UUID,
		Name:      sourceType.Name,
		PublicURL: sourceType.PublicURL,
	}, nil
}

/////////////////////////////// REVIEW API METHODS ///////////////////////////////

func (ds *Datastore) GetModelReview(reviewUUID uuid.UUID) (*models.ModelReviewResponse, error) {
	var review dbmodels.ModelReview
	err := ds.DB.Preload("Model").Preload("FromBranch").Preload("FromBranchVersion").Preload("ToBranch").Preload("CreatedByUser").Where("uuid = ?", reviewUUID).Find(&review).Error
	if err != nil {
		return nil, err
	}
	return &models.ModelReviewResponse{
		UUID: review.UUID,
		Model: models.ModelNameResponse{
			UUID: review.Model.UUID,
			Name: review.Model.Name,
		},
		FromBranch: models.ModelBranchNameResponse{
			UUID: review.FromBranch.UUID,
			Name: review.FromBranch.Name,
		},
		FromBranchVersion: models.ModelBranchVersionNameResponse{
			UUID:    review.FromBranchVersion.UUID,
			Version: review.FromBranchVersion.Version,
		},
		ToBranch: models.ModelBranchNameResponse{
			UUID: review.ToBranch.UUID,
			Name: review.ToBranch.Name,
		},
		Title:       review.Title,
		Description: review.Description,
		IsComplete:  review.IsComplete,
		IsAccepted:  review.IsAccepted,
		CreatedBy: models.UserHandleResponse{
			UUID:   review.CreatedByUser.UUID,
			Handle: review.CreatedByUser.Handle,
			Name:   review.CreatedByUser.Name,
			Avatar: review.CreatedByUser.Avatar,
			Email:  review.CreatedByUser.Email,
		},
	}, nil
}

func (ds *Datastore) GetModelReviews(modelUUID uuid.UUID) ([]models.ModelReviewResponse, error) {
	var reviews []dbmodels.ModelReview
	err := ds.DB.Preload("Model").Preload("FromBranch").Preload("FromBranchVersion").Preload("ToBranch").Preload("CreatedByUser").Where("model_uuid = ?", modelUUID).Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	var reviewResponses []models.ModelReviewResponse
	for _, review := range reviews {
		reviewResponses = append(reviewResponses, models.ModelReviewResponse{
			UUID: review.UUID,
			Model: models.ModelNameResponse{
				UUID: review.Model.UUID,
				Name: review.Model.Name,
			},
			FromBranch: models.ModelBranchNameResponse{
				UUID: review.FromBranch.UUID,
				Name: review.FromBranch.Name,
			},
			FromBranchVersion: models.ModelBranchVersionNameResponse{
				UUID:    review.FromBranchVersion.UUID,
				Version: review.FromBranchVersion.Version,
			},
			ToBranch: models.ModelBranchNameResponse{
				UUID: review.ToBranch.UUID,
				Name: review.ToBranch.Name,
			},
			Title:       review.Title,
			Description: review.Description,
			IsComplete:  review.IsComplete,
			IsAccepted:  review.IsAccepted,
			CreatedBy: models.UserHandleResponse{
				UUID:   review.CreatedByUser.UUID,
				Handle: review.CreatedByUser.Handle,
				Name:   review.CreatedByUser.Name,
				Avatar: review.CreatedByUser.Avatar,
				Email:  review.CreatedByUser.Email,
			},
		})
	}
	return reviewResponses, nil
}

func (ds *Datastore) CreateModelReview(modelUUID uuid.UUID, userUUID uuid.UUID, fromBranch uuid.UUID, fromBranchVersion uuid.UUID, toBranch uuid.UUID, title string, desc string, isComplete bool, isAccepted bool) (*models.ModelReviewResponse, error) {
	review := dbmodels.ModelReview{
		Model: dbmodels.Model{
			BaseModel: dbmodels.BaseModel{
				UUID: modelUUID,
			},
		},
		FromBranch: dbmodels.ModelBranch{
			BaseModel: dbmodels.BaseModel{
				UUID: fromBranch,
			},
		},
		FromBranchVersion: dbmodels.ModelVersion{
			BaseModel: dbmodels.BaseModel{
				UUID: fromBranchVersion,
			},
		},
		ToBranch: dbmodels.ModelBranch{
			BaseModel: dbmodels.BaseModel{
				UUID: toBranch,
			},
		},
		CreatedByUser: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: userUUID,
			},
		},
		Title:       title,
		Description: desc,
		IsComplete:  isComplete,
		IsAccepted:  isAccepted,
	}
	err := ds.DB.Create(&review).Find(&review).Error
	if err != nil {
		return nil, err
	}
	return &models.ModelReviewResponse{
		UUID: review.UUID,
		Model: models.ModelNameResponse{
			UUID: review.Model.UUID,
			Name: review.Model.Name,
		},
		FromBranch: models.ModelBranchNameResponse{
			UUID: review.FromBranch.UUID,
			Name: review.FromBranch.Name,
		},
		FromBranchVersion: models.ModelBranchVersionNameResponse{
			UUID:    review.FromBranchVersion.UUID,
			Version: review.FromBranchVersion.Version,
		},
		ToBranch: models.ModelBranchNameResponse{
			UUID: review.ToBranch.UUID,
			Name: review.ToBranch.Name,
		},
		Title:       review.Title,
		Description: review.Description,
		IsComplete:  review.IsComplete,
		IsAccepted:  review.IsAccepted,
		CreatedBy: models.UserHandleResponse{
			UUID:   review.CreatedByUser.UUID,
			Handle: review.CreatedByUser.Handle,
			Name:   review.CreatedByUser.Name,
			Avatar: review.CreatedByUser.Avatar,
			Email:  review.CreatedByUser.Email,
		},
	}, nil
}

func (ds *Datastore) UpdateModelReview(reviewUUID uuid.UUID, updatedAttributes map[string]any) (*models.ModelReviewResponse, error) {
	var review dbmodels.ModelReview
	err := ds.DB.Preload("Model").Preload("FromBranch").Preload("FromBranchVersion").Preload("ToBranch").Preload("CreatedByUser").Where("uuid = ?", reviewUUID).Find(&review).Error
	if err != nil {
		return nil, err
	}
	if review.IsComplete {
		return nil, errors.New("review already complete")
	}
	alreadyAccepted := review.IsAccepted
	for key, value := range updatedAttributes {
		switch key {
		case "title":
			review.Title = value.(string)
		case "description":
			review.Description = value.(string)
		case "is_complete":
			review.IsComplete = value.(bool)
		case "is_accepted":
			// CANNOT UNACCEPT ALREADY ACCEPTED REVIEW
			if !alreadyAccepted {
				review.IsAccepted = value.(bool)
			}
		}
	}
	err = ds.DB.Transaction(func(tx *gorm.DB) error {
		// CHECK IF REVIEW ACCEPTED
		if review.IsAccepted && !alreadyAccepted {
			// make a new version in to_branch from from_branch specified version
			_, err := ds.MigrateModelVersionBranch(review.FromBranchVersion.UUID, review.ToBranch.UUID)
			if err != nil {
				return err
			}
		}
		result := tx.Save(&review).Find(&review)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &models.ModelReviewResponse{
		UUID: review.UUID,
		Model: models.ModelNameResponse{
			UUID: review.Model.UUID,
			Name: review.Model.Name,
		},
		FromBranch: models.ModelBranchNameResponse{
			UUID: review.FromBranch.UUID,
			Name: review.FromBranch.Name,
		},
		ToBranch: models.ModelBranchNameResponse{
			UUID: review.ToBranch.UUID,
			Name: review.ToBranch.Name,
		},
		Title:       review.Title,
		Description: review.Description,
		IsComplete:  review.IsComplete,
		IsAccepted:  review.IsAccepted,
		CreatedBy: models.UserHandleResponse{
			UUID:   review.CreatedByUser.UUID,
			Handle: review.CreatedByUser.Handle,
			Name:   review.CreatedByUser.Name,
			Avatar: review.CreatedByUser.Avatar,
			Email:  review.CreatedByUser.Email,
		},
	}, nil
}

func (ds *Datastore) GetDatasetReview(reviewUUID uuid.UUID) (*models.DatasetReviewResponse, error) {
	var review dbmodels.DatasetReview
	err := ds.DB.Preload("Dataset").Preload("FromBranch").Preload("FromBranchVersion").Preload("ToBranch").Preload("CreatedByUser").Where("uuid = ?", reviewUUID).Find(&review).Error
	if err != nil {
		return nil, err
	}
	return &models.DatasetReviewResponse{
		UUID: review.UUID,
		Dataset: models.DatasetNameResponse{
			UUID: review.Dataset.UUID,
			Name: review.Dataset.Name,
		},
		FromBranch: models.DatasetBranchNameResponse{
			UUID: review.FromBranch.UUID,
			Name: review.FromBranch.Name,
		},
		FromBranchVersion: models.DatasetBranchVersionNameResponse{
			UUID:    review.FromBranchVersion.UUID,
			Version: review.FromBranchVersion.Version,
		},
		ToBranch: models.DatasetBranchNameResponse{
			UUID: review.ToBranch.UUID,
			Name: review.ToBranch.Name,
		},
		Title:       review.Title,
		Description: review.Description,
		IsComplete:  review.IsComplete,
		IsAccepted:  review.IsAccepted,
		CreatedBy: models.UserHandleResponse{
			UUID:   review.CreatedByUser.UUID,
			Handle: review.CreatedByUser.Handle,
			Name:   review.CreatedByUser.Name,
			Avatar: review.CreatedByUser.Avatar,
			Email:  review.CreatedByUser.Email,
		},
	}, nil
}

func (ds *Datastore) GetDatasetReviews(datasetUUID uuid.UUID) ([]models.DatasetReviewResponse, error) {
	var reviews []dbmodels.DatasetReview
	err := ds.DB.Preload("Dataset").Preload("FromBranch").Preload("FromBranchVersion").Preload("ToBranch").Preload("CreatedByUser").Where("dataset_uuid = ?", datasetUUID).Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	var reviewResponses []models.DatasetReviewResponse
	for _, review := range reviews {
		reviewResponses = append(reviewResponses, models.DatasetReviewResponse{
			UUID: review.UUID,
			Dataset: models.DatasetNameResponse{
				UUID: review.Dataset.UUID,
				Name: review.Dataset.Name,
			},
			FromBranch: models.DatasetBranchNameResponse{
				UUID: review.FromBranch.UUID,
				Name: review.FromBranch.Name,
			},
			FromBranchVersion: models.DatasetBranchVersionNameResponse{
				UUID:    review.FromBranchVersion.UUID,
				Version: review.FromBranchVersion.Version,
			},
			ToBranch: models.DatasetBranchNameResponse{
				UUID: review.ToBranch.UUID,
				Name: review.ToBranch.Name,
			},
			Title:       review.Title,
			Description: review.Description,
			IsComplete:  review.IsComplete,
			IsAccepted:  review.IsAccepted,
			CreatedBy: models.UserHandleResponse{
				UUID:   review.CreatedByUser.UUID,
				Handle: review.CreatedByUser.Handle,
				Name:   review.CreatedByUser.Name,
				Avatar: review.CreatedByUser.Avatar,
				Email:  review.CreatedByUser.Email,
			},
		})
	}
	return reviewResponses, nil
}

func (ds *Datastore) CreateDatasetReview(datasetUUID uuid.UUID, userUUID uuid.UUID, fromBranch uuid.UUID, fromBranchVerison uuid.UUID, toBranch uuid.UUID, title string, desc string, isComplete bool, isAccepted bool) (*models.DatasetReviewResponse, error) {
	review := dbmodels.DatasetReview{
		Dataset: dbmodels.Dataset{
			BaseModel: dbmodels.BaseModel{
				UUID: datasetUUID,
			},
		},
		FromBranch: dbmodels.DatasetBranch{
			BaseModel: dbmodels.BaseModel{
				UUID: fromBranch,
			},
		},
		FromBranchVersion: dbmodels.DatasetVersion{
			BaseModel: dbmodels.BaseModel{
				UUID: fromBranchVerison,
			},
		},
		ToBranch: dbmodels.DatasetBranch{
			BaseModel: dbmodels.BaseModel{
				UUID: toBranch,
			},
		},
		CreatedByUser: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: userUUID,
			},
		},
		Title:       title,
		Description: desc,
		IsComplete:  isComplete,
		IsAccepted:  isAccepted,
	}
	err := ds.DB.Create(&review).Find(&review).Error
	if err != nil {
		return nil, err
	}
	return &models.DatasetReviewResponse{
		UUID: review.UUID,
		Dataset: models.DatasetNameResponse{
			UUID: review.Dataset.UUID,
			Name: review.Dataset.Name,
		},
		FromBranch: models.DatasetBranchNameResponse{
			UUID: review.FromBranch.UUID,
			Name: review.FromBranch.Name,
		},
		FromBranchVersion: models.DatasetBranchVersionNameResponse{
			UUID:    review.FromBranchVersion.UUID,
			Version: review.FromBranchVersion.Version,
		},
		ToBranch: models.DatasetBranchNameResponse{
			UUID: review.ToBranch.UUID,
			Name: review.ToBranch.Name,
		},
		Title:       review.Title,
		Description: review.Description,
		IsComplete:  review.IsComplete,
		IsAccepted:  review.IsAccepted,
		CreatedBy: models.UserHandleResponse{
			UUID:   review.CreatedByUser.UUID,
			Handle: review.CreatedByUser.Handle,
			Name:   review.CreatedByUser.Name,
			Avatar: review.CreatedByUser.Avatar,
			Email:  review.CreatedByUser.Email,
		},
	}, nil
}

func (ds *Datastore) UpdateDatasetReview(reviewUUID uuid.UUID, updatedAttributes map[string]any) (*models.DatasetReviewResponse, error) {
	var review dbmodels.DatasetReview
	err := ds.DB.Preload("Dataset").Preload("FromBranch").Preload("FromBranchVersion").Preload("ToBranch").Preload("CreatedByUser").Where("uuid = ?", reviewUUID).Find(&review).Error
	if err != nil {
		return nil, err
	}
	if review.IsComplete {
		return nil, errors.New("review already complete")
	}
	alreadyAccepted := review.IsAccepted
	for key, value := range updatedAttributes {
		switch key {
		case "title":
			review.Title = value.(string)
		case "description":
			review.Description = value.(string)
		case "is_complete":
			review.IsComplete = value.(bool)
		case "is_accepted":
			// CANNOT UNACCEPT ALREADY ACCEPTED REVIEW
			if !alreadyAccepted {
				review.IsAccepted = value.(bool)
			}
		}
	}
	err = ds.DB.Transaction(func(tx *gorm.DB) error {
		// CHECK IF REVIEW ACCEPTED
		if review.IsAccepted && !alreadyAccepted {
			// make a new version in to_branch from from_branch specified version
			_, err := ds.MigrateDatasetVersionBranch(review.FromBranchVersion.UUID, review.ToBranch.UUID)
			if err != nil {
				return err
			}
		}
		result := tx.Save(&review).Find(&review)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &models.DatasetReviewResponse{
		UUID: review.UUID,
		Dataset: models.DatasetNameResponse{
			UUID: review.Dataset.UUID,
			Name: review.Dataset.Name,
		},
		FromBranch: models.DatasetBranchNameResponse{
			UUID: review.FromBranch.UUID,
			Name: review.FromBranch.Name,
		},
		ToBranch: models.DatasetBranchNameResponse{
			UUID: review.ToBranch.UUID,
			Name: review.ToBranch.Name,
		},
		Title:       review.Title,
		Description: review.Description,
		IsComplete:  review.IsComplete,
		IsAccepted:  review.IsAccepted,
		CreatedBy: models.UserHandleResponse{
			UUID:   review.CreatedByUser.UUID,
			Handle: review.CreatedByUser.Handle,
			Name:   review.CreatedByUser.Name,
			Avatar: review.CreatedByUser.Avatar,
			Email:  review.CreatedByUser.Email,
		},
	}, nil
}
