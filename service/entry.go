package service

import (
	"TGU-MAP/models"
	"TGU-MAP/service/crud"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var ListItemClient *crud.ListItemStub
var RDB *redis.Client
var GlobalConfig Config

func init() {
	if err := loadConfig(); err != nil {
		panic(err)
	}
	if db, err := initDB(); err != nil {
		panic(err)
	} else {
		ListItemClient = crud.NewListItemStub(db)
	}
	initRedis()
}

type Config struct {
	Database struct {
		Host     string
		Port     int
		Name     string
		User     string
		Password string
	}
	Web struct {
		Port int
	}
	Redis struct {
		Host     string
		Port     int
		Password string
		Db       int
	}
}

func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // 添加配置文件所在路径
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		return err
	}
	return nil
}

func initDB() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, " ", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:        time.Second, // Slow SQL threshold
			LogLevel:             logger.Info, // Log level
			ParameterizedQueries: true,        // Don't include params in the SQL log
			Colorful:             true,        // Disable color
		},
	)
	var dsn string
	if GlobalConfig.Database.Password == "" {
		dsn = fmt.Sprintf("%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", GlobalConfig.Database.User, GlobalConfig.Database.Host, GlobalConfig.Database.Port, GlobalConfig.Database.Name)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", GlobalConfig.Database.User, GlobalConfig.Database.Password, GlobalConfig.Database.Host, GlobalConfig.Database.Port, GlobalConfig.Database.Name)
	}

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

func initRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 请根据实际情况调整
		Password: "",               // 没有密码则为空
		DB:       0,                // 使用默认数据库
	})
}
