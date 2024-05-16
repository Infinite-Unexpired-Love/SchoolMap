package curd

import (
	"TGU-MAP/models"
	"gorm.io/gorm"
	"reflect"

	"fmt"
	"log"
)

func fetchData(db *gorm.DB, items interface{}) {
	result := db.Preload("Children").Where("parent_id IS NULL").Find(items)
	if result.Error != nil {
		log.Fatalf("failed to fetch data: %v", result.Error)
	}

}

func findElemID(db *gorm.DB, cur interface{}, path ...string) *uint {
	var elemID *uint

	result := db.Where("title = ? AND parent_id IS NULL", path[0]).First(cur)
	if result.Error != nil {
		log.Fatalf("failed to find parent node: %v", result.Error)
	}
	elemID = cur.(models.Cascade).GetID()
	for _, title := range path[1:] {
		//result := db.Where("title = ?", title).Where("parent_id = ?", elemID).First(&parent)
		result := db.Raw("SELECT * FROM list_items WHERE title = ? AND parent_id = ? Limit 1", title, *elemID).Scan(cur)
		if result.Error != nil {

			log.Fatalf("failed to find parent node: %v", result.Error)
		}
		elemID = cur.(models.Cascade).GetID()
	}
	return elemID
}

func insertNode(db *gorm.DB, parentID *uint, item models.Cascade) {
	item.SetParentID(parentID)
	result := db.Create(item)
	if result.Error != nil {
		log.Fatalf("failed to insert node: %v", result.Error)
	} else {
		fmt.Println("Node inserted successfully")
	}
}

func updateNode(db *gorm.DB, elemID uint, target interface{}, item models.Updatable) {
	result := db.Where("id = ?", elemID).First(target)
	if result.Error != nil {
		log.Fatalf("failed to find node: %v", result.Error)
	}
	target.(models.Updatable).Update(item)

	result = db.Save(target)
	if result.Error != nil {
		log.Fatalf("failed to update node: %v", result.Error)
	}

	fmt.Println("Node updated successfully")
}

func deleteNode(db *gorm.DB, elemID uint, target interface{}, children interface{}) {
	//result := db.Where("parent_id = ?", elemID).Find(&children)
	//if result.Error != nil {
	//	log.Fatalf("failed to find children nodes: %v", result.Error)
	//}
	//
	//for _, child := range children {
	//	deleteNode(db, *child.(models.Cascade).GetID(), target, children)
	//}
	//
	//result = db.Delete(target, elemID)
	//fmt.Println("Node deleted successfully")

	// 通过反射确保 children 是一个切片
	childrenValue := reflect.ValueOf(children)
	//if childrenValue.Kind() != reflect.Ptr || childrenValue.Elem().Kind() != reflect.Slice {
	//	log.Fatalf("children should be a pointer to a slice")
	//}
	childrenValue = childrenValue.Elem()

	// 查找子节点
	result := db.Where("parent_id = ?", elemID).Find(children)
	if result.Error != nil {
		log.Fatalf("failed to find children nodes: %v", result.Error)
	}

	// 遍历子节点并递归删除
	for i := 0; i < childrenValue.Len(); i++ {
		child := childrenValue.Index(i).Interface()
		childID := child.(models.Cascade).GetID()
		deleteNode(db, *childID, target, children)
	}

	// 删除当前节点
	result = db.Delete(target, elemID)
	if result.Error != nil {
		log.Fatalf("failed to delete node: %v", result.Error)
	} else {
		fmt.Println("Node deleted successfully")
	}
}
