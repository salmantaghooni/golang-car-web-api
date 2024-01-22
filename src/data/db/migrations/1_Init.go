package migrations

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/constants"
	"github.com/salmantaghooni/golang-car-web-api/src/data/db"
	"github.com/salmantaghooni/golang-car-web-api/src/data/models"
	"github.com/salmantaghooni/golang-car-web-api/src/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()
	migrationTabels(database)
}
func Down_1() {

}

func createDefaultInformation(database *gorm.DB) {
	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExist(database, &adminRole)
	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExist(database, &defaultRole)

	adminUser := models.User{FirstName: "salman", LastName: "taghooni", MobileNumber: "09133851769", UserName: constants.DefaultUserName, Email: constants.DefaultEmail}
	psswd := constants.DefaultPassword
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(psswd), bcrypt.DefaultCost)
	adminUser.Password = string(hashedPassword)
	createAdminUserIfNotExist(database, &adminUser, adminRole.Id)

}

func createRoleIfNotExist(database *gorm.DB, r *models.Role) {
	exists := 0
	database.Model(&models.Role{}).
		Select("name").
		Where("name = ?", r.Name).
		First(&exists)
	// if err != nil {
	// 	logger.Error(logging.Postgres, logging.Select, "error in exist role "+r.Name, nil)
	// }
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExist(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.Model(&models.User{}).
		Select("id").
		Where("user_name = ?", u.UserName).First(&exists)
	// if err != nil {
	// 	logger.Error(logging.Postgres, logging.Select, "error in exist useradmin "+u.UserName, nil)
	// }
	if exists == 0 {
		database.Create(u)
		ur := &models.UserRole{RoleId: roleId, UserId: u.Id}
		database.Create(ur)
	}
}

func migrationTabels(database *gorm.DB) {
	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	tables = addNewTable(database, country, tables)
	tables = addNewTable(database, city, tables)
	tables = addNewTable(database, user, tables)
	tables = addNewTable(database, role, tables)
	tables = addNewTable(database, userRole, tables)
	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		// panic("couldn't create table")
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
	createDefaultInformation(database)

}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}
