package db

import (
	"fmt"
	"goweb/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDatabase 返回一个新的数据库客户端连接
func NewDatabase(conf *conf.Mysql) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.Name, conf.Charset)
	dialector := mysql.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// InitDatabase 初始化了数据库
// func InitDatabase(db *gorm.DB, config *config.Database) error {
// 	db.LogMode(config.Debug) // auto migrate
// 	models := []interface{}{
// 		&models.User{},
// 		// &models.PersonalInfo{},
// 		// &models.Category{},
// 		// &models.Subcategory{},
// 	}
// 	if err := db.AutoMigrate(models...); err != nil {
// 		return err
// 	} // Personal info
// 	// if err := db.Model(&models.User{}).AddForeignKey("account_id", fmt.Sprintf("%s(id)", models.AccountTableName), "CASCADE", "CASCADE").Error; err != nil {
// 	// 	return err
// 	// } // Subcategories
// 	// if err := db.Model(&models.Subcategory{}).AddForeignKey("category_id", fmt.Sprintf("%s(id)", models.CategoryTableName), "CASCADE", "CASCADE").Error; err != nil {
// 	// 	return err
// 	// }
// 	return nil
// }
