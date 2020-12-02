package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kwstars/Go-000/Week02/module"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func initDB() *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:mysql@tcp(127.0.0.1:3306)/mytest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&module.Product{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	db := initDB()
	s := InitProductService(db)

	r := gin.Default()

	r.GET("/products", s.FindAll)
	r.GET("/products/:id", s.FindByID)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
