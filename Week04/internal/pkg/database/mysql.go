package database

import (
	"github.com/google/wire"
	"github.com/kwstars/Go-000/Week04/internal/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(New)

func New() (db *gorm.DB, cf func(), err error) {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/mytest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Product{})
	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	cf = func() { sqlDB.Close() }

	return
}
