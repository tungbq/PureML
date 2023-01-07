package impl

import (
	"fmt"

	"github.com/PriyavKaneria/PureML/service/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMSQLDatastore() *MySQLDatastore {
	db, err := gorm.Open(mysql.Open(""), &gorm.Config{}) //Todo
	if err != nil {
		fmt.Println(err)
		panic("Error connecting to database")
	}
	return &MySQLDatastore{
		DB: db,
	}
}

type MySQLDatastore struct {
	DB *gorm.DB
}

func (ds *MySQLDatastore) GetAllAdminOrgs() ([]models.Organization, error) {
	//Todo Interact with DB via GORM
	return []models.Organization{}, nil
}

func (ds *MySQLDatastore) GetOrgByID(orgId string) (*models.Organization, error) {
	return nil, nil
}

func (ds *MySQLDatastore) GetOrgsByUserMail(mailId string) ([]models.OrgAccess, error) {
	return []models.OrgAccess{}, nil
}
