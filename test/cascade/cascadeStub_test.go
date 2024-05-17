package main

import (
	"TGU-MAP/models"
	"TGU-MAP/service/curd"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

var ListItemStub *curd.CascadeStub
var db *gorm.DB

func init() {
	if db, err := initDB("gorm_test"); err != nil {
		panic(err)
	} else {
		ListItemStub = curd.NewCascadeStub(db, models.ListItem{})
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

// 测试函数
func TestCascadeStub(t *testing.T) {
	// 插入测试数据
	item1 := &models.ListItem{Title: "Root"}
	ListItemStub.InsertNodeByPath(item1)

	item2 := &models.ListItem{Title: "Child 1"}
	ListItemStub.InsertNodeByPath(item2, "Root")

	item3 := &models.ListItem{Title: "Child 2"}
	ListItemStub.InsertNodeByPath(item3, "Root", "Child 1")

	// 验证插入
	//var items []models.ListItem
	//fmt.Println("whats wrong")
	//ListItemStub.FetchData()
	//result := db.Where("parent_id IS NULL").First(&items)
	//fmt.Println("no wrong")
	//assert.NoError(t, result.Error, "failed to fetch data")
	//
	//// 确认层次结构
	//assert.Equal(t, 1, len(items))
	//assert.Equal(t, 1, len(items[0].Children))
	//assert.Equal(t, 1, len(items[0].Children[0].Children))

	// 更新测试数据
	itemToUpdate := &models.ListItem{Desc: "Updated Description"}
	ListItemStub.UpdateNodeByPath(itemToUpdate, "Root", "Child 1", "Child 2")

	//var updatedItem models.ListItem
	//result = db.Where("title = ?", "Child 2").First(&updatedItem)
	//assert.NoError(t, result.Error, "failed to fetch updated data")
	//assert.Equal(t, "Updated Description", updatedItem.Desc)

	// 删除测试数据
	ListItemStub.DeleteNodeByPath("Root", "Child 1")

	//result = db.Preload("Children.Children").Where("parent_id IS NULL").Find(&items)
	//assert.NoError(t, result.Error, "failed to fetch data after delete")
	//assert.Equal(t, 1, len(items))
	//assert.Equal(t, 0, len(items[0].Children[0].Children))
}

func main() {
	// 运行测试
	TestCascadeStub(&testing.T{})
}
