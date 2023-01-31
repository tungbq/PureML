package datastore

import (
	"mime/multipart"

	"github.com/PureML-Inc/PureML/server/config"
	"github.com/PureML-Inc/PureML/server/datastore/impl"
	"github.com/PureML-Inc/PureML/server/models"
	uuid "github.com/satori/go.uuid"
)

var ds *impl.Datastore

func InitDB() {
	databaseType := config.GetDatabaseType()
	if databaseType == "local" {
		//SQLite db for local
		ds = impl.NewSQLiteDatastore()
	} else if databaseType == "postgres" {
		//Postgres db
		ds = impl.NewPostgresDatastore()
	}
}

func GetAllAdminOrgs() ([]models.OrganizationResponse, error) {
	return ds.GetAllAdminOrgs()
}

func GetOrgById(orgId uuid.UUID) (*models.OrganizationResponse, error) {
	return ds.GetOrgByID(orgId)
}

func GetOrgByHandle(orgHandle string) (*models.OrganizationResponse, error) {
	return ds.GetOrgByHandle(orgHandle)
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

func GetUserByHandle(handle string) (*models.UserResponse, error) {
	return ds.GetUserByHandle(handle)
}

func GetSecureUserByEmail(email string) (*models.UserResponse, error) {
	return ds.GetSecureUserByEmail(email)
}

func GetSecureUserByHandle(handle string) (*models.UserResponse, error) {
	return ds.GetSecureUserByHandle(handle)
}

func GetUserByUUID(userUUID uuid.UUID) (*models.UserResponse, error) {
	return ds.GetUserByUUID(userUUID)
}

func GetUserProfileByUUID(userUUID uuid.UUID) (*models.UserProfileResponse, error) {
	return ds.GetUserProfileByUUID(userUUID)
}

func CreateUser(name string, email string, handle string, bio string, avatar string, hashedPassword string) (*models.UserResponse, error) {
	return ds.CreateUser(name, email, handle, bio, avatar, hashedPassword)
}

func UpdateUser(email string, name string, avatar string, bio string) (*models.UserResponse, error) {
	return ds.UpdateUser(email, name, avatar, bio)
}

func GetLogForModelVersion(modelVersionUUID uuid.UUID) ([]models.LogResponse, error) {
	return ds.GetLogForModelVersion(modelVersionUUID)
}

func GetKeyLogForModelVersion(modelVersionUUID uuid.UUID, key string) ([]models.LogResponse, error) {
	return ds.GetKeyLogForModelVersion(modelVersionUUID, key)
}

func CreateLogForModelVersion(key string, data string, modelVersionUUID uuid.UUID) (*models.LogResponse, error) {
	return ds.CreateLogForModelVersion(key, data, modelVersionUUID)
}

func GetLogForDatasetVersion(datasetVersionUUID uuid.UUID) ([]models.LogResponse, error) {
	return ds.GetLogForDatasetVersion(datasetVersionUUID)
}

func GetKeyLogForDatasetVersion(datasetVersionUUID uuid.UUID, key string) ([]models.LogResponse, error) {
	return ds.GetKeyLogForDatasetVersion(datasetVersionUUID, key)
}

func CreateLogForDatasetVersion(key string, data string, datasetVersionUUID uuid.UUID) (*models.LogResponse, error) {
	return ds.CreateLogForDatasetVersion(key, data, datasetVersionUUID)
}

func GetAllModels(orgId uuid.UUID) ([]models.ModelResponse, error) {
	return ds.GetAllModels(orgId)
}

func GetModelByName(orgId uuid.UUID, modelName string) (*models.ModelResponse, error) {
	return ds.GetModelByName(orgId, modelName)
}

func CreateModel(orgId uuid.UUID, name string, wiki string, isPublic bool, readmeData *models.ReadmeRequest, userUUID uuid.UUID) (*models.ModelResponse, error) {
	return ds.CreateModel(orgId, name, wiki, isPublic, readmeData, userUUID)
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

func UploadAndRegisterModelFile(orgId uuid.UUID, modelBranchUUID uuid.UUID, file *multipart.FileHeader, isEmpty bool, hash string, source string) (*models.ModelBranchVersionResponse, error) {
	return ds.UploadAndRegisterModelFile(orgId, modelBranchUUID, file, isEmpty, hash, source)
}

func GetModelAllBranches(modelUUID uuid.UUID) ([]models.ModelBranchResponse, error) {
	return ds.GetModelAllBranches(modelUUID)
}

func GetModelAllVersions(modelUUID uuid.UUID) ([]models.ModelBranchVersionResponse, error) {
	return ds.GetModelAllVersions(modelUUID)
}

func GetModelBranchByName(orgId uuid.UUID, modelName string, branchName string) (*models.ModelBranchResponse, error) {
	return ds.GetModelBranchByName(orgId, modelName, branchName)
}

func GetModelBranchByUUID(modelBranchUUID uuid.UUID) (*models.ModelBranchResponse, error) {
	return ds.GetModelBranchByUUID(modelBranchUUID)
}

func GetModelBranchAllVersions(modelBranchUUID uuid.UUID) ([]models.ModelBranchVersionResponse, error) {
	return ds.GetModelBranchAllVersions(modelBranchUUID)
}

func GetModelBranchVersion(modelBranchUUID uuid.UUID, version string) (*models.ModelBranchVersionResponse, error) {
	return ds.GetModelBranchVersion(modelBranchUUID, version)
}

func GetAllDatasets(orgId uuid.UUID) ([]models.DatasetResponse, error) {
	return ds.GetAllDatasets(orgId)
}

func GetDatasetByName(orgId uuid.UUID, datasetName string) (*models.DatasetResponse, error) {
	return ds.GetDatasetByName(orgId, datasetName)
}

func CreateDataset(orgId uuid.UUID, name string, wiki string, isPublic bool, readmeData *models.ReadmeRequest, userUUID uuid.UUID) (*models.DatasetResponse, error) {
	return ds.CreateDataset(orgId, name, wiki, isPublic, readmeData, userUUID)
}

func CreateDatasetBranch(datasetUUID uuid.UUID, branchName string) (*models.DatasetBranchResponse, error) {
	return ds.CreateDatasetBranch(datasetUUID, branchName)
}

func CreateDatasetBranches(datasetUUID uuid.UUID, branchNames []string) ([]models.DatasetBranchResponse, error) {
	var branches []models.DatasetBranchResponse

	for _, branchName := range branchNames {
		branch, err := CreateDatasetBranch(datasetUUID, branchName)
		if err != nil {
			return nil, err
		}
		branches = append(branches, *branch)
	}

	return branches, nil
}

func UploadAndRegisterDatasetFile(orgId uuid.UUID, datasetBranchUUID uuid.UUID, file *multipart.FileHeader, isEmpty bool, hash string, source string, lineage string) (*models.DatasetBranchVersionResponse, error) {
	return ds.UploadAndRegisterDatasetFile(orgId, datasetBranchUUID, file, isEmpty, hash, source, lineage)
}

func GetDatasetAllBranches(datasetUUID uuid.UUID) ([]models.DatasetBranchResponse, error) {
	return ds.GetDatasetAllBranches(datasetUUID)
}

func GetDatasetAllVersions(datasetUUID uuid.UUID) ([]models.DatasetBranchVersionResponse, error) {
	return ds.GetDatasetAllVersions(datasetUUID)
}

func GetDatasetBranchByName(orgId uuid.UUID, datasetName string, branchName string) (*models.DatasetBranchResponse, error) {
	return ds.GetDatasetBranchByName(orgId, datasetName, branchName)
}

func GetDatasetBranchByUUID(datasetBranchUUID uuid.UUID) (*models.DatasetBranchResponse, error) {
	return ds.GetDatasetBranchByUUID(datasetBranchUUID)
}

func GetDatasetBranchAllVersions(datasetBranchUUID uuid.UUID) ([]models.DatasetBranchVersionResponse, error) {
	return ds.GetDatasetBranchAllVersions(datasetBranchUUID)
}

func GetDatasetBranchVersion(datasetBranchUUID uuid.UUID, version string) (*models.DatasetBranchVersionResponse, error) {
	return ds.GetDatasetBranchVersion(datasetBranchUUID, version)
}

func GetModelActivity(modelUUID uuid.UUID, category string) (*models.ActivityResponse, error) {
	return ds.GetModelActivity(modelUUID, category)
}

func CreateModelActivity(modelUUID uuid.UUID, userUUID uuid.UUID, category string, activity string) (*models.ActivityResponse, error) {
	return ds.CreateModelActivity(modelUUID, userUUID, category, activity)
}

func UpdateModelActivity(activityUUID uuid.UUID, updatedAttributes map[string]string) (*models.ActivityResponse, error) {
	return ds.UpdateModelActivity(activityUUID, updatedAttributes)
}

func DeleteModelActivity(activityUUID uuid.UUID) error {
	return ds.DeleteModelActivity(activityUUID)
}

func GetDatasetActivity(datasetUUID uuid.UUID, category string) (*models.ActivityResponse, error) {
	return ds.GetDatasetActivity(datasetUUID, category)
}

func CreateDatasetActivity(datasetUUID uuid.UUID, userUUID uuid.UUID, category string, activity string) (*models.ActivityResponse, error) {
	return ds.CreateDatasetActivity(datasetUUID, userUUID, category, activity)
}

func UpdateDatasetActivity(activityUUID uuid.UUID, updatedAttributes map[string]string) (*models.ActivityResponse, error) {
	return ds.UpdateDatasetActivity(activityUUID, updatedAttributes)
}

func DeleteDatasetActivity(activityUUID uuid.UUID) error {
	return ds.DeleteDatasetActivity(activityUUID)
}

func GetSourceSecret(orgId uuid.UUID, source string) (*models.SourceSecrets, error) {
	return ds.GetSourceSecret(orgId, source)
}

func CreateR2Secrets(orgId uuid.UUID, accountId string, accessKeyId string, accessKeySecret string, bucketName string, publicURL string) (*impl.R2Secrets, error) {
	return ds.CreateR2Secrets(orgId, accountId, accessKeyId, accessKeySecret, bucketName, publicURL)
}

func CreateR2Source(orgId uuid.UUID, publicURL string) (*models.SourceTypeResponse, error) {
	return ds.CreateR2Source(orgId, publicURL)
}

func DeleteR2Secrets(orgId uuid.UUID) error {
	return ds.DeleteR2Secrets(orgId)
}

func CreateS3Secrets(orgId uuid.UUID, accessKeyId string, accessKeySecret string, bucketName string, bucketLocation string) (*impl.S3Secrets, error) {
	return ds.CreateS3Secrets(orgId, accessKeyId, accessKeySecret, bucketName, bucketLocation)
}

func CreateS3Source(orgId uuid.UUID, publicURL string) (*models.SourceTypeResponse, error) {
	return ds.CreateS3Source(orgId, publicURL)
}

func DeleteS3Secrets(orgId uuid.UUID) error {
	return ds.DeleteS3Secrets(orgId)
}

func GetModelReadmeVersion(modelUUID uuid.UUID, version string) (*models.ReadmeVersionResponse, error) {
	return ds.GetModelReadmeVersion(modelUUID, version)
}

func GetModelReadmeAllVersions(modelUUID uuid.UUID) ([]models.ReadmeVersionResponse, error) {
	return ds.GetModelReadmeAllVersions(modelUUID)
}

func UpdateModelReadme(modelUUID uuid.UUID, fileType string, content string) (*models.ReadmeVersionResponse, error) {
	return ds.UpdateModelReadme(modelUUID, fileType, content)
}

func GetDatasetReadmeVersion(modelUUID uuid.UUID, version string) (*models.ReadmeVersionResponse, error) {
	return ds.GetDatasetReadmeVersion(modelUUID, version)
}

func GetDatasetReadmeAllVersions(modelUUID uuid.UUID) ([]models.ReadmeVersionResponse, error) {
	return ds.GetDatasetReadmeAllVersions(modelUUID)
}

func UpdateDatasetReadme(modelUUID uuid.UUID, fileType string, content string) (*models.ReadmeVersionResponse, error) {
	return ds.UpdateDatasetReadme(modelUUID, fileType, content)
}

func GetModelReview(reviewUUID uuid.UUID) (*models.ModelReviewResponse, error) {
	return ds.GetModelReview(reviewUUID)
}

func GetModelReviews(modelUUID uuid.UUID) ([]models.ModelReviewResponse, error) {
	return ds.GetModelReviews(modelUUID)
}

func CreateModelReview(modelUUID uuid.UUID, userUUID uuid.UUID, fromBranch uuid.UUID, fromBranchVersion uuid.UUID, toBranch uuid.UUID, title string, desc string, isComplete bool, isAccepted bool) (*models.ModelReviewResponse, error) {
	return ds.CreateModelReview(modelUUID, userUUID, fromBranch, fromBranchVersion, toBranch, title, desc, isComplete, isAccepted)
}

func UpdateModelReview(reviewUUID uuid.UUID, updatedAttributes map[string]any) (*models.ModelReviewResponse, error) {
	return ds.UpdateModelReview(reviewUUID, updatedAttributes)
}

func GetDatasetReview(reviewUUID uuid.UUID) (*models.DatasetReviewResponse, error) {
	return ds.GetDatasetReview(reviewUUID)
}

func GetDatasetReviews(datasetUUID uuid.UUID) ([]models.DatasetReviewResponse, error) {
	return ds.GetDatasetReviews(datasetUUID)
}

func CreateDatasetReview(datasetUUID uuid.UUID, userUUID uuid.UUID, fromBranch uuid.UUID, fromBranchVerison uuid.UUID, toBranch uuid.UUID, title string, desc string, isComplete bool, isAccepted bool) (*models.DatasetReviewResponse, error) {
	return ds.CreateDatasetReview(datasetUUID, userUUID, fromBranch, fromBranchVerison, toBranch, title, desc, isComplete, isAccepted)
}

func UpdateDatasetReview(reviewUUID uuid.UUID, updatedAttributes map[string]any) (*models.DatasetReviewResponse, error) {
	return ds.UpdateDatasetReview(reviewUUID, updatedAttributes)
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
