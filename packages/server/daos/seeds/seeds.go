package seeds

import (
	"fmt"

	"github.com/PureML-Inc/PureML/server/daos/dbmodels"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var defaultUUID = uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))
var defaultUUID2 = uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222"))

func All() []Seed {
	return []Seed{
		{
			Name: "CreateDemoAdminUserAndOrg",
			Run: func(d *gorm.DB) error {
				return CreateUser(d, defaultUUID, "Demo User", "demo@aztlan.in", "demo", "demo", "Demo User Bio", "")
			},
		},
		{
			Name: "CreateNonAdminUserAndOrg",
			Run: func(d *gorm.DB) error {
				return CreateUser(d, defaultUUID2, "Normal User", "notadmin@aztlan.in", "notadmin", "notadmin", "User Bio", "")
			},
		},
		{
			Name: "CreateDemoModel",
			Run: func(d *gorm.DB) error {
				return CreateModel(d, defaultUUID, "Demo Model", "Demo Model Wiki", true)
			},
		},
		{
			Name: "CreateDemoPrivateModel",
			Run: func(d *gorm.DB) error {
				return CreateModel(d, defaultUUID2, "Demo Private Model", "Demo Private Model Wiki", false)
			},
		},
		{
			Name: "CreateDemoDataset",
			Run: func(d *gorm.DB) error {
				return CreateDataset(d, defaultUUID, "Demo Dataset", "Demo Dataset Wiki", true)
			},
		},
		{
			Name: "CreateDemoPrivateDataset",
			Run: func(d *gorm.DB) error {
				return CreateDataset(d, defaultUUID2, "Demo Private Dataset", "Demo Private Dataset Wiki", false)
			},
		},
	}
}

func CreateUser(db *gorm.DB, uuid uuid.UUID, name string, email string, handle string, password string, bio string, avatar string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	err = db.Create(&dbmodels.User{
		BaseModel: dbmodels.BaseModel{
			UUID: uuid,
		},
		Name:     name,
		Email:    email,
		Handle:   handle,
		Password: string(hashedPassword),
		Bio:      bio,
		Avatar:   avatar,
		Orgs: []dbmodels.Organization{
			{
				BaseModel: dbmodels.BaseModel{
					UUID: uuid,
				},
				Name:        "Demo Org",
				Handle:      handle,
				Avatar:      "",
				Description: "Demo Org Description",
				JoinCode:    fmt.Sprintf("iwanttojoin%s", handle),
			},
		},
	}).Error
	if err != nil {
		return err
	}
	return db.Table("user_organizations").Where("user_uuid = ?", uuid).Where("organization_uuid = ?", uuid).Update("role", "owner").Error
}

func CreateModel(db *gorm.DB, uuid uuid.UUID, name string, wiki string, isPublic bool) error {
	var branches []dbmodels.ModelBranch
	var readme dbmodels.Readme
	if isPublic {
		branches = []dbmodels.ModelBranch{
			{
				BaseModel: dbmodels.BaseModel{
					UUID: defaultUUID,
				},
				Name:      "main",
				IsDefault: true,
			},
			{
				BaseModel: dbmodels.BaseModel{
					UUID: defaultUUID2,
				},
				Name:      "dev",
				IsDefault: false,
				Versions: []dbmodels.ModelVersion{
					{
						BaseModel: dbmodels.BaseModel{
							UUID: defaultUUID,
						},
						Version: "v1",
						Hash:    "1234567890",
						IsEmpty: true,
						CreatedByUser: dbmodels.User{
							BaseModel: dbmodels.BaseModel{
								UUID: defaultUUID,
							},
						},
					},
				},
			},
		}
		readme = dbmodels.Readme{
			BaseModel: dbmodels.BaseModel{
				UUID: defaultUUID,
			},
			ReadmeVersions: []dbmodels.ReadmeVersion{
				{
					BaseModel: dbmodels.BaseModel{
						UUID: defaultUUID,
					},
					Content:  "Demo Readme",
					FileType: "md",
					Version:  "v1",
				},
			},
		}
	}
	return db.Create(&dbmodels.Model{
		BaseModel: dbmodels.BaseModel{
			UUID: uuid,
		},
		Name: name,
		Wiki: wiki,
		Org: dbmodels.Organization{
			BaseModel: dbmodels.BaseModel{
				UUID: defaultUUID,
			},
		},
		CreatedByUser: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: defaultUUID,
			},
		},
		UpdatedByUser: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: defaultUUID,
			},
		},
		IsPublic: isPublic,
		Branches: branches,
		Readme:   readme,
	}).Error
}

func CreateDataset(db *gorm.DB, uuid uuid.UUID, name string, wiki string, isPublic bool) error {
	var branches []dbmodels.DatasetBranch
	var readme dbmodels.Readme
	if isPublic {
		branches = []dbmodels.DatasetBranch{
			{
				BaseModel: dbmodels.BaseModel{
					UUID: defaultUUID,
				},
				Name:      "main",
				IsDefault: true,
			},
			{
				BaseModel: dbmodels.BaseModel{
					UUID: defaultUUID2,
				},
				Name:      "dev",
				IsDefault: false,
				Versions: []dbmodels.DatasetVersion{
					{
						BaseModel: dbmodels.BaseModel{
							UUID: defaultUUID,
						},
						Version: "v1",
						Hash:    "1234567890",
						IsEmpty: true,
						Lineage: dbmodels.Lineage{
							BaseModel: dbmodels.BaseModel{
								UUID: defaultUUID,
							},
							Lineage: "{}",
						},
						CreatedByUser: dbmodels.User{
							BaseModel: dbmodels.BaseModel{
								UUID: defaultUUID,
							},
						},
					},
				},
			},
		}
		readme = dbmodels.Readme{
			BaseModel: dbmodels.BaseModel{
				UUID: defaultUUID,
			},
			ReadmeVersions: []dbmodels.ReadmeVersion{
				{
					BaseModel: dbmodels.BaseModel{
						UUID: defaultUUID,
					},
					Content:  "Demo Readme",
					FileType: "md",
					Version:  "v1",
				},
			},
		}
	}
	return db.Create(&dbmodels.Dataset{
		BaseModel: dbmodels.BaseModel{
			UUID: uuid,
		},
		Name: name,
		Wiki: wiki,
		Org: dbmodels.Organization{
			BaseModel: dbmodels.BaseModel{
				UUID: defaultUUID,
			},
		},
		CreatedByUser: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: defaultUUID,
			},
		},
		UpdatedByUser: dbmodels.User{
			BaseModel: dbmodels.BaseModel{
				UUID: defaultUUID,
			},
		},
		IsPublic: isPublic,
		Branches: branches,
		Readme:   readme,
	}).Error
}