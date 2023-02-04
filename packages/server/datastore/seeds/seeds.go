package seeds

import (
	"github.com/PureML-Inc/PureML/server/datastore/dbmodels"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var defaultUUID = uuid.Must(uuid.FromString("11111111-1111-1111-1111-111111111111"))
var defaultUUID2 = uuid.Must(uuid.FromString("22222222-2222-2222-2222-222222222222"))

func All() []Seed {
	return []Seed{
		{
			Name: "CreateDemoUserAndOrg",
			Run: func(d *gorm.DB) error {
				return CreateUser(d, "Demo User", "demo@aztlan.in", "demo", "demo", "Demo User Bio", "")
			},
		},
		{
			Name: "CreateDemoModel",
			Run: func(d *gorm.DB) error {
				return CreateModel(d, "Demo Model", "Demo Model Wiki")
			},
		},
		{
			Name: "CreateDemoDataset",
			Run: func(d *gorm.DB) error {
				return CreateDataset(d, "Demo Dataset", "Demo Dataset Wiki")
			},
		},
	}
}

func CreateUser(db *gorm.DB, name string, email string, handle string, password string, bio string, avatar string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	return db.Create(&dbmodels.User{
		BaseModel: dbmodels.BaseModel{
			UUID: defaultUUID,
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
					UUID: defaultUUID,
				},
				Name:        "Demo Org",
				Handle:      "demo",
				Avatar:      "",
				Description: "Demo Org Description",
				JoinCode:    "iwanttojoin",
			},
		},
	}).Error
}

func CreateModel(db *gorm.DB, name string, wiki string) error {
	return db.Create(&dbmodels.Model{
		BaseModel: dbmodels.BaseModel{
			UUID: defaultUUID,
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
		IsPublic: true,
		Branches: []dbmodels.ModelBranch{
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
					},
				},
			},
		},
		Readme: dbmodels.Readme{
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
		},
	}).Error
}

func CreateDataset(db *gorm.DB, name string, wiki string) error {
	return db.Create(&dbmodels.Dataset{
		BaseModel: dbmodels.BaseModel{
			UUID: defaultUUID,
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
		IsPublic: true,
		Branches: []dbmodels.DatasetBranch{
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
					},
				},
			},
		},
		Readme: dbmodels.Readme{
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
		},
	}).Error
}
