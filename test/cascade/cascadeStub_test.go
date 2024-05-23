package main

import (
	"TGU-MAP/models"
	"TGU-MAP/service/crud"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

var CascadeStub *crud.CascadeStub
var db *gorm.DB
var db2 *gorm.DB

func init() {
	if db, err := initDB("gorm_test"); err != nil {
		panic(err)
	} else {
		CascadeStub = crud.NewCascadeStub(db, models.ListItem{})
	}
	db2, _ = initDB("gorm_test")
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
		return nil, models.SQLError(fmt.Sprintf("failed to connect database: %v", err))
	}

	err = db.AutoMigrate(&models.ListItem{})
	if err != nil {
		return nil, models.SQLError(fmt.Sprintf("failed to migrate database: %v", err))
	}
	return db, nil
}

// 测试函数
func TestCascadeStub(t *testing.T) {
	// 插入测试数据
	item1 := &models.ListItem{Title: "Root"}
	CascadeStub.InsertNodeByPath(item1)

	item2 := &models.ListItem{Title: "Child 1"}
	CascadeStub.InsertNodeByPath(item2, "Root")

	item3 := &models.ListItem{Title: "Child 2"}
	CascadeStub.InsertNodeByPath(item3, "Root", "Child 1")

	// 验证插入
	var items []models.ListItem
	result := db2.Preload("Children.Children").Where("parent_id IS NULL").First(&items)
	assert.NoError(t, result.Error, "failed to fetch data")
	// 确认层次结构
	assert.Equal(t, 1, len(items))
	assert.Equal(t, 1, len(items[0].Children))
	assert.Equal(t, 1, len(items[0].Children[0].Children))

	// 更新测试数据
	itemToUpdate := &models.ListItem{Desc: "Updated Description"}
	CascadeStub.UpdateNodeByPath(itemToUpdate, "Root", "Child 1", "Child 2")

	var updatedItem models.ListItem
	result = db2.Where("title = ?", "Child 2").First(&updatedItem)
	assert.NoError(t, result.Error, "failed to fetch updated data")
	assert.Equal(t, "Updated Description", updatedItem.Desc)

	// 删除测试数据
	CascadeStub.DeleteNodeByPath("Root", "Child 1")

	result = db2.Preload("Children.Children").Where("parent_id IS NULL").Find(&items)
	assert.NoError(t, result.Error, "failed to fetch data after delete")
	assert.Equal(t, 1, len(items))
	assert.Equal(t, 0, len(items[0].Children))
}

func main() {
	// 运行测试
	TestCascadeStub(&testing.T{})
}
