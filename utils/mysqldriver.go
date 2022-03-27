package utils

import (
	"Project-REST-API/configs"
	"Project-REST-API/entities"
	"fmt"

	faker "github.com/bxcodec/faker/v3"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)
	fmt.Println(connectionString)
	// "root:@tcp(127.0.0.1:3306)/be5db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	db.Migrator().DropTable(&entities.Task{})
	db.Migrator().DropTable(&entities.Project{})
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Task{})
	db.AutoMigrate(&entities.Project{})
	for i := 0; i < 2000; i++ {
		db.Create(&entities.User{
			Nama:     faker.Name(),
			Email:    faker.Email(),
			Password: "xyz",
		})
	}
}
