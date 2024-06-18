package service

import (
	"TGU-MAP/models"
	"TGU-MAP/service/aliasItem"
	"TGU-MAP/service/feedback"
	"TGU-MAP/service/listItem"
	"TGU-MAP/service/noticeItem"
	"TGU-MAP/service/user"
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

func init() {
	if err := loadConfig(); err != nil {
		panic(err)
	}
	if err := initListItemDB(); err != nil {
		panic(err)
	}
	if err := initAliasItemDB(); err != nil {
		panic(err)
	}
	if err := initNoticeItemDB(); err != nil {
		panic(err)
	}
	if err := initFeedbackDB(); err != nil {
		panic(err)
	}
	if err := initUserDB(); err != nil {
		panic(err)
	}
	//admin := models.User{
	//	Username:   "系统管理员",
	//	Mobile:     "15079994355",
	//	Password:   "wqiej21394heds2!",
	//	Role:       2,
	//	Expiration: nil,
	//}
	//err := UserClient.InsertNode(&admin)
	//if err != nil {
	//	panic(err)
	//}
	initRedis()
}

func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // 配置文件在工作目录下，即项目的根目录
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		return err
	}
	return nil
}

func getConnection() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, " ", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:        time.Second, // Slow SQL threshold
			LogLevel:             logger.Info, // Log level
			ParameterizedQueries: false,       // Don't include params in the SQL log
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
	return db, nil
}

func initListItemDB() error {
	if db, err := getConnection(); err != nil {
		return err
	} else {
		err = db.AutoMigrate(&models.ListItem{})
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
			return err
		}
		ListItemClient = listItem.NewListItemStub(db)
		return nil
	}
}

func initUserDB() error {
	if db, err := getConnection(); err != nil {
		return err
	} else {
		err = db.AutoMigrate(&models.User{})
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
			return err
		}
		UserClient = user.NewUserStub(db)
		return nil
	}
}

func initAliasItemDB() error {
	if db, err := getConnection(); err != nil {
		return err
	} else {
		err = db.AutoMigrate(&models.AliasItem{})
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
			return err
		}
		AliasItemClient = aliasItem.NewAliasItemStub(db)
		return nil
	}
}

func initNoticeItemDB() error {
	if db, err := getConnection(); err != nil {
		return err
	} else {
		err = db.AutoMigrate(&models.NoticeItem{})
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
			return err
		}
		NoticeItemClient = noticeItem.NewNoticeItemStub(db)
		return nil
	}
}

func initFeedbackDB() error {
	if db, err := getConnection(); err != nil {
		return err
	} else {
		err = db.AutoMigrate(&models.Feedback{})
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
			return err
		}
		FeedbackClient = feedback.NewFeedbackStub(db)
		return nil
	}
}

func initRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // 使用默认数据库
	})
}
