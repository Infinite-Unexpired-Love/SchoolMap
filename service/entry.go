package service

import (
	"TGU-MAP/models"
	"TGU-MAP/service/curd"
	"fmt"
	"gorm.io/gorm/logger"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var ListItemClient *curd.ListItemStub

func init() {
	if db, err := initDB("gorm_test"); err != nil {
		panic(err)
	} else {
		ListItemClient = curd.NewListItemStub(db)
	}

}

func initDB(database string) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, " ", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:        time.Second, // Slow SQL threshold
			LogLevel:             logger.Info, // Log level
			ParameterizedQueries: true,        // Don't include params in the SQL log
			Colorful:             true,        // Disable color
		},
	)
	dsn := fmt.Sprintf("root@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.ListItem{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
		return nil, err
	}
	return db, nil
}
