package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type ListItem struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `json:"title"`
	Desc      string     `json:"desc"`
	Contact   *string    `json:"contact,omitempty"`
	Latitude  *float64   `json:"latitude,omitempty"`
	Longitude *float64   `json:"longitude,omitempty"`
	ParentID  *uint      `json:"parent_id,omitempty"`
	Children  []ListItem `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}

var db *gorm.DB

func initDB() {
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold:             time.Second, // Slow SQL threshold
	//		LogLevel:                  logger.Info, // Log level
	//		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	//		ParameterizedQueries:      true,        // Don't include params in the SQL log
	//		Colorful:                  true,        // Disable color
	//	},
	//)
	var err error
	dsn := "root@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&ListItem{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

func main() {
	initDB()

	// 插入数据
	//insertData()

	// 查询数据
	//fetchData()

	// 插入根节点
	/*insertRootNode(ListItem{
		Title: "学部学院",
		Desc:  "学院现有数学一级学科博士点、数学和统计学两个一级学科硕士授权点，数学与应用数学、信息与计算科学、应用统计学、数据科学与大数据技术四个本科专业",
	})*/

	// 插入叶子节点
	/*insertChildNode("软件学院", ListItem{
		Title:     "教学楼B区",
		Desc:      "这里一般是学生们上机实操的地方",
		Latitude:  ptrFloat64(39.134213),
		Longitude: ptrFloat64(117.212123),
	})*/

	//insertNode(&ListItem{
	//	Title:     "教学办",
	//	Desc:      "这里一般是学生们上机实操的地方",
	//	Latitude:  ptrFloat64(39.134213),
	//	Longitude: ptrFloat64(117.212123),
	//}, "软件园", "软件学院", "教学楼B区")

	// 更新节点
	updateNodeByPath(&ListItem{
		Title: "教学办",
		Desc:  "我被修改了",
	}, "软件园", "软件学院", "教学楼B区")
}

func insertData() {
	// 示例数据
	data := []ListItem{
		{
			Title:     "软件园",
			Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
			Latitude:  ptrFloat64(39.134213),
			Longitude: ptrFloat64(117.212123),
			Children: []ListItem{
				{
					Title:     "软件学院",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134213),
					Longitude: ptrFloat64(117.212123),
				},
				{
					Title:     "大软公寓",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134123),
					Longitude: ptrFloat64(117.221321),
				},
				{
					Title:     "大软食堂",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134123),
					Longitude: ptrFloat64(117.212321),
				},
				{
					Title:     "北苑公寓",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134231),
					Longitude: ptrFloat64(117.221321),
				},
			},
		},
		{
			Title:     "西苑",
			Desc:      "这里是西苑，有众多学生宿舍，学习氛围浓厚，是学生学习、生活的好地方",
			Latitude:  ptrFloat64(39.134213),
			Longitude: ptrFloat64(117.212123),
			Children: []ListItem{
				{
					Title:     "软件学院",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134213),
					Longitude: ptrFloat64(117.212123),
				},
				{
					Title:     "大软公寓",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134123),
					Longitude: ptrFloat64(117.221321),
				},
				{
					Title:     "大软食堂",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134123),
					Longitude: ptrFloat64(117.212321),
				},
				{
					Title:     "北苑公寓",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134231),
					Longitude: ptrFloat64(117.221321),
				},
			},
		},
		{
			Title:     "其他",
			Desc:      "这里是其他地方",
			Latitude:  ptrFloat64(39.134213),
			Longitude: ptrFloat64(117.212123),
			Children: []ListItem{
				{
					Title:     "软件学院",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134213),
					Longitude: ptrFloat64(117.212123),
				},
				{
					Title:     "大软公寓",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134123),
					Longitude: ptrFloat64(117.221321),
				},
				{
					Title:     "大软食堂",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134123),
					Longitude: ptrFloat64(117.212321),
				},
				{
					Title:     "北苑公寓",
					Desc:      "这里是软件工程专业、网络空间安全专业学子学习专业课的地方，有众多企业入驻，学习环境舒适，冬暖夏凉",
					Latitude:  ptrFloat64(39.134231),
					Longitude: ptrFloat64(117.221321),
				},
			},
		},
	}

	result := db.Create(&data)
	if result.Error != nil {
		log.Fatalf("failed to insert data: %v", result.Error)
	} else {
		fmt.Println("Data inserted successfully")
	}
}

func fetchData() {
	var items []ListItem
	result := db.Preload("Children").Where("parent_id IS NULL").Find(&items)
	if result.Error != nil {
		log.Fatalf("failed to fetch data: %v", result.Error)
	}

	jsonData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal data: %v", err)
	}

	fmt.Println(string(jsonData))
}

func insertNodeByPath(item *ListItem, path ...string) {
	var parentID *uint
	if len(path) > 0 {
		parentID = findParentID(path)
		if parentID == nil {
			log.Fatalf("failed to find parent node for titles: %v", path)
		}
	}
	item.ParentID = parentID

	result := db.Create(item)
	if result.Error != nil {
		log.Fatalf("failed to insert node: %v", result.Error)
	} else {
		fmt.Println("Node inserted successfully")
	}
}

func insertNodeByID(item *ListItem, parentID uint) {
	item.ParentID = &parentID

	result := db.Create(item)
	if result.Error != nil {
		log.Fatalf("failed to insert node: %v", result.Error)
	} else {
		fmt.Println("Node inserted successfully")
	}
}

func updateNodeByPath(item *ListItem, path ...string) {
	parentID := findParentID(path)
	if parentID == nil {
		log.Fatalf("failed to find parent node for titles: %v", path)
	}

	var existingItem ListItem
	result := db.Where("title = ? AND parent_id = ?", item.Title, parentID).First(&existingItem)
	if result.Error != nil {
		log.Fatalf("failed to find node: %v", result.Error)
	}

	existingItem.Title = item.Title
	existingItem.Desc = item.Desc
	existingItem.Contact = item.Contact
	existingItem.Latitude = item.Latitude
	existingItem.Longitude = item.Longitude

	result = db.Save(&existingItem)
	if result.Error != nil {
		log.Fatalf("failed to update node: %v", result.Error)
	}

	fmt.Println("Node updated successfully")
}

func updateNodeByID(item *ListItem, parentID uint) {
	var existingItem ListItem
	result := db.Where("title = ? AND parent_id = ?", item.Title, parentID).First(&existingItem)
	if result.Error != nil {
		log.Fatalf("failed to find node: %v", result.Error)
	}

	existingItem.Title = item.Title
	existingItem.Desc = item.Desc
	existingItem.Contact = item.Contact
	existingItem.Latitude = item.Latitude
	existingItem.Longitude = item.Longitude

	result = db.Save(&existingItem)
	if result.Error != nil {
		log.Fatalf("failed to update node: %v", result.Error)
	}

	fmt.Println("Node updated successfully")

}

func findParentID(path []string) *uint {
	var parent ListItem
	var parentID *uint

	result := db.Where("title = ?", path[0]).First(&parent)
	if result.Error != nil {
		log.Fatalf("failed to find parent node: %v", result.Error)
	}
	parentID = &parent.ID
	for _, title := range path[1:] {
		//result := db.Where("title = ?", title).Where("parent_id = ?", parentID).First(&parent)
		result := db.Raw("SELECT * FROM list_items WHERE title = ? AND parent_id = ? Limit 1", title, *parentID).Scan(&parent)
		if result.Error != nil {

			log.Fatalf("failed to find parent node: %v", result.Error)
		}
		parentID = &parent.ID
	}

	return parentID
}

func ptrFloat64(f float64) *float64 {
	return &f
}
