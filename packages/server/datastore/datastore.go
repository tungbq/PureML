package datastore

import (
	"mime/multipart"
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

func GetUserByEmail(email string) (*models.UserResponse, error) {
	return ds.GetUserByEmail(email)
}

func GetUserByHandle(email string) (*models.UserResponse, error) {
	return ds.GetUserByHandle(email)
}

func CreateUser(name string, email string, handle string, bio string, avatar string, hashedPassword string) (*models.UserResponse, error) {
	return ds.CreateUser(name, email, handle, bio, avatar, hashedPassword)
}

func UpdateUser(email string, name string, avatar string, bio string) (*models.UserResponse, error) {
	return ds.UpdateUser(email, name, avatar, bio)
}

func CreateLogForModelVersion(data string, modelVersionUUID uuid.UUID) (*models.LogResponse, error) {
	return ds.CreateLogForModelVersion(data, modelVersionUUID)
}

func CreateLogForDatasetVersion(data string, datasetVersionUUID uuid.UUID) (*models.LogResponse, error) {
	return ds.CreateLogForDatasetVersion(data, datasetVersionUUID)
}

func GetAllModels(orgId uuid.UUID) ([]models.ModelResponse, error) {
	return ds.GetAllModels(orgId)
}

func GetModelByName(orgId uuid.UUID, modelName string) (*models.ModelResponse, error) {
	return ds.GetModelByName(orgId, modelName)
}

func CreateModel(orgId uuid.UUID, name string, wiki string, userUUID uuid.UUID) (*models.ModelResponse, error) {
	return ds.CreateModel(orgId, name, wiki, userUUID)
}

func CreateModelBranch(modelUUID uuid.UUID, branchName string) (*models.ModelBranchResponse, error) {
	return ds.CreateModelBranch(modelUUID, branchName)
}

func CreateModelBranches(modelUUID uuid.UUID, branchNames []string) ([]models.ModelBranchResponse, error) {
	var branches []models.ModelBranchResponse

	for _, branchName := range branchNames {
		branch, err := CreateModelBranch(modelUUID, branchName)
		if err != nil {
			return nil, err
		}
		branches = append(branches, *branch)
	}

	return branches, nil
}

func UploadAndRegisterModelFile(modelBranchUUID uuid.UUID, file *multipart.FileHeader, hash string, source string) (*models.ModelVersionResponse, error) {
	return ds.UploadAndRegisterModelFile(modelBranchUUID, file, hash, source)
}

func GetModelAllBranches(modelUUID uuid.UUID) ([]models.ModelBranchResponse, error) {
	return ds.GetModelAllBranches(modelUUID)
}

func GetModelAllVersions(modelUUID uuid.UUID) ([]models.ModelVersionResponse, error) {
	return ds.GetModelAllVersions(modelUUID)
}

func GetBranchByName(orgId uuid.UUID, modelName string, branchName string) (*models.ModelBranchResponse, error) {
	return ds.GetBranchByName(orgId, modelName, branchName)
}

func GetBranchByUUID(modelbranchUUID uuid.UUID) (*models.ModelBranchResponse, error) {
	return ds.GetBranchByUUID(modelbranchUUID)
}

func GetModelBranchAllVersions(modelbranchUUID uuid.UUID) ([]models.ModelVersionResponse, error) {
	return ds.GetModelBranchAllVersions(modelbranchUUID)
}

func GetModelBranchVersion(modelbranchUUID uuid.UUID, version string) (*models.ModelVersionResponse, error) {
	return ds.GetModelBranchVersion(modelbranchUUID, version)
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
	GetUserByEmail(email string) (*models.UserResponse, error)
	GetUserByHandle(email string) (*models.UserResponse, error)
	CreateUser(name string, email string, handle string, bio string, avatar string, hashedPassword string) (*models.UserResponse, error)
	UpdateUser(email string, name string, avatar string, bio string) (*models.UserResponse, error)
	CreateLogForModelVersion(data string, modelVersionUUID uuid.UUID) (*models.LogResponse, error)
	CreateLogForDatasetVersion(data string, datasetVersionUUID uuid.UUID) (*models.LogResponse, error)
}
