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

var ListItemClient *curd.CascadeStub

func init() {
	if db, err := initDB("gorm_test"); err != nil {
		panic(err)
	} else {
		ListItemClient = curd.NewCascadeStub(db, models.ListItem{})
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

//func insertData() {
//	// 示例数据
//	data := []models.ListItem{
//		{
//			Title:     "软件园",
//			Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//			Latitude:  ptrFloat64(39.134213),
//			Longitude: ptrFloat64(117.212123),
//			Children: []models.ListItem{
//				{
//					Title:     "软件学院",
//					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//					Latitude:  ptrFloat64(39.134213),
//					Longitude: ptrFloat64(117.212123),
//				},
//				{
//					Title:     "大软公寓",
//					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//					Latitude:  ptrFloat64(39.134123),
//					Longitude: ptrFloat64(117.221321),
//				},
//				{
//					Title:     "大软食堂",
//					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//					Latitude:  ptrFloat64(39.134123),
//					Longitude: ptrFloat64(117.212321),
//				},
//				{
//					Title:     "北苑公寓",
//					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//					Latitude:  ptrFloat64(39.134231),
//					Longitude: ptrFloat64(117.221321),
//				},
//			},
//		},
//		{
//			Title:     "西苑",
//			Desc:      "这里是西苑，有众多学生宿舍，学习氛围浓厚，是学生学习、生活的好地方",
//			Latitude:  ptrFloat64(39.134213),
//			Longitude: ptrFloat64(117.212123),
//			Children: []models.ListItem{
//				{
//					Title:     "软件学院",
//					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//					Latitude:  ptrFloat64(39.134213),
//					Longitude: ptrFloat64(117.212123),
//				},
//				{
//					Title:     "大软公寓",
//					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//					Latitude:  ptrFloat64(39.134123),
//					Longitude: ptrFloat64(117.221321),
//				},
//				{
//					Title:     "大软食堂",
//					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//					Latitude:  ptrFloat64(39.134123),
//					Longitude: ptrFloat64(117.212321),
//				},
//				{
//					Title:     "北苑公寓",
//					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//					Latitude:  ptrFloat64(39.134231),
//					Longitude: ptrFloat64(117.221321),
//				},
//			},
//		},
//		{
//			Title:     "其他",
//			Desc:      "这里是其他地方",
//			Latitude:  ptrFloat64(39.134213),
//			Longitude: ptrFloat64(117.212123),
//			Children: []models.ListItem{
//				{
//					Title:     "软件学院",
//					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//					Latitude:  ptrFloat64(39.134213),
//					Longitude: ptrFloat64(117.212123),
//				},
//				{
//					Title:     "大软公寓",
//					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//					Latitude:  ptrFloat64(39.134123),
//					Longitude: ptrFloat64(117.221321),
//				},
//				{
//					Title:     "大软食堂",
//					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//					Latitude:  ptrFloat64(39.134123),
//					Longitude: ptrFloat64(117.212321),
//				},
//				{
//					Title:     "北苑公寓",
//					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
//					Latitude:  ptrFloat64(39.134231),
//					Longitude: ptrFloat64(117.221321),
//				},
//			},
//		},
//	}
//
//	result := Db.Create(&data)
//	if result.Error != nil {
//		log.Fatalf("failed to insert data: %v", result.Error)
//	} else {
//		fmt.Println("Data inserted successfully")
//	}
//}

func ptrFloat64(f float64) *float64 {
	return &f
}
