package dbmodels

import (
	commondbmodels "github.com/PureMLHQ/PureML/packages/purebackend/core/common/dbmodels"
	datasetdbmodels "github.com/PureMLHQ/PureML/packages/purebackend/dataset/dbmodels"
	modeldbmodels "github.com/PureMLHQ/PureML/packages/purebackend/model/dbmodels"
	userorgdbmodels "github.com/PureMLHQ/PureML/packages/purebackend/user_org/dbmodels"
	uuid "github.com/satori/go.uuid"
)

type Activity struct {
	commondbmodels.BaseModel `gorm:"embedded"`
	UserUUID                 uuid.UUID     `json:"user_uuid" gorm:"type:uuid;not null"`
	Category                 string        `json:"category"`
	Activity                 string        `json:"activity"`
	ModelUUID                uuid.NullUUID `json:"model_uuid" gorm:"type:uuid;"`
	DatasetUUID              uuid.NullUUID `json:"dataset_uuid" gorm:"type:uuid;"`

	User    userorgdbmodels.User    `gorm:"foreignKey:UserUUID"`
	Model   modeldbmodels.Model     `gorm:"foreignKey:ModelUUID"`
	Dataset datasetdbmodels.Dataset `gorm:"foreignKey:DatasetUUID"`
}

type Tag struct {
	ModelUUID        uuid.NullUUID `json:"model_uuid" gorm:"type:uuid;primaryKey"`
	DatasetUUID      uuid.NullUUID `json:"dataset_uuid" gorm:"type:uuid;primaryKey"`
	OrganizationUUID uuid.UUID     `json:"organization_uuid" gorm:"type:uuid;not null;index:idx_org_tag,unique"`
	Tag              string        `json:"tag" gorm:"not null;index:idx_org_tag,unique"`

	Model   modeldbmodels.Model          `gorm:"foreignKey:ModelUUID"`
	Dataset datasetdbmodels.Dataset      `gorm:"foreignKey:DatasetUUID"`
	Org     userorgdbmodels.Organization `gorm:"foreignKey:OrganizationUUID"`
}

type Log struct {
	commondbmodels.BaseModel `gorm:"embedded"`
	Key                      string        `json:"key" gorm:""`
	Data                     string        `json:"data"`
	Type                     string        `json:"type"`
	ModelVersionUUID         uuid.NullUUID `json:"model_version_uuid" gorm:"type:uuid;"`
	DatasetVersionUUID       uuid.NullUUID `json:"dataset_version_uuid" gorm:"type:uuid;"`

	ModelVersion   modeldbmodels.ModelVersion     `gorm:"foreignKey:ModelVersionUUID"`
	DatasetVersion datasetdbmodels.DatasetVersion `gorm:"foreignKey:DatasetVersionUUID"`
}
