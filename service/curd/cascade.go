package curd

import (
	"TGU-MAP/models"
	"TGU-MAP/utils"
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

func findElemID(db *gorm.DB, target interface{}, path ...string) *uint {
	var elemID *uint

	result := db.Where("title = ? AND parent_id IS NULL", path[0]).First(target)
	if result.Error != nil {
		log.Fatalf("failed to find parent node: %v", result.Error)
	}
	elemID = target.(models.Cascade).GetID()
	for _, title := range path[1:] {
		//result := db.Where("title = ?", title).Where("parent_id = ?", elemID).First(&parent)
		result := db.Raw("SELECT * FROM list_items WHERE title = ? AND parent_id = ? Limit 1", title, *elemID).Scan(target)
		if result.Error != nil {

			log.Fatalf("failed to find parent node: %v", result.Error)
		}
		elemID = target.(models.Cascade).GetID()
	}
	return elemID
}

//func findElem(db *gorm.DB, slice interface{}, path ...string) *uint {
//	var elemID *uint
//	result := db.Where("title = ? AND parent_id IS NULL", path[0]).Find(slice)
//}

func insertNode(db *gorm.DB, parentID *uint, item models.Cascade) {
	item.SetParentID(parentID)
	result := db.Create(item)
	if result.Error != nil {
		log.Fatalf("failed to insert node: %v", result.Error)
	} else {
		fmt.Println("Node inserted successfully")
	}
}

func updateNode(db *gorm.DB, elemID *uint, target interface{}, item models.Updatable) {
	result := db.Where("id = ?", *elemID).First(target)
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

func deleteNode(db *gorm.DB, elemID *uint, target interface{}, children interface{}) {
	//复制一个空白的children用来下一次递归调用
	newChildren, _ := utils.GetVoidSlice(children)

	// 查找子节点
	result := db.Where("parent_id = ?", *elemID).Find(&children)
	if result.Error != nil {
		log.Fatalf("failed to find children nodes: %v", result.Error)
	}
	childrenValue := reflect.ValueOf(children)
	// 遍历子节点并递归删除
	for i := 0; i < childrenValue.Len(); i++ {
		child := childrenValue.Index(i).Addr().Interface()
		childID := child.(models.Cascade).GetID()
		deleteNode(db, childID, target, newChildren)
	}

	// 删除当前节点
	result = db.Delete(target, elemID)
	if result.Error != nil {
		log.Fatalf("failed to delete node: %v", result.Error)
	} else {
		fmt.Println("Node deleted successfully")
	}
}
