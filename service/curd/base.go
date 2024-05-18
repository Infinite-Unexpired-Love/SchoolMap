package curd

import (
	"TGU-MAP/models"
	"TGU-MAP/utils"
	"gorm.io/gorm"
	"reflect"

	"fmt"
)

// fetchData 从数据库中获取所有item，不构造树形结构
func fetchData(db *gorm.DB, items interface{}) *models.CustomError {
	result := db.Find(items)
	if result.Error != nil {
		return models.SQLError(fmt.Sprintf("failed to fetch data: %v", result.Error))
	}
	return nil
}

// findElemID 根据路径查找节点 ID
func findElemID(db *gorm.DB, target interface{}, path ...string) (uint, *models.CustomError) {
	var elemID uint

	// 查找根节点
	result := db.Where("title = ? AND parent_id IS NULL", path[0]).First(target)
	if result.Error != nil {
		return 0, models.SQLError(fmt.Sprintf("failed to find parent node: %v", result.Error))
	}
	elemID = target.(models.BaseInfo).GetID()

	// 遍历路径的每个部分查找对应的子节点
	for _, title := range path[1:] {
		// TODO: 适配其他数据库
		result := db.Raw("SELECT * FROM list_items WHERE title = ? AND parent_id = ? Limit 1", title, elemID).Scan(target)
		if result.Error != nil {
			return 0, models.SQLError(fmt.Sprintf("failed to find parent node: %v", result.Error))
		}
		elemID = target.(models.BaseInfo).GetID()
	}
	return elemID, nil
}

// insertNode 插入新节点
func insertNode(db *gorm.DB, parentID uint, item models.BaseInfo) *models.CustomError {
	//var count int64
	//db.Model(item).Where("parent_id = ? AND title = ?", parentID, item.GetTitle()).Count(&count)
	//if count > 0 {
	//	return models.SQLError(fmt.Sprintf("duplicate title: %v under parentNode", item.GetTitle()))
	//}

	item.SetParentID(parentID)
	result := db.Create(item)
	if result.Error != nil {
		return models.SQLError(fmt.Sprintf("failed to insert node: %v", result.Error))
	} else {
		return nil
	}
}

// updateNode 更新节点
func updateNode(db *gorm.DB, elemID uint, target interface{}, item models.Updatable) *models.CustomError {
	// 查找要更新的节点
	result := db.Where("id = ?", elemID).First(target)
	if result.Error != nil {
		return models.SQLError(fmt.Sprintf("failed to find node: %v", result.Error))
	}
	// 更新节点
	target.(models.Updatable).Update(item)

	result = db.Save(target)
	if result.Error != nil {
		return models.SQLError(fmt.Sprintf("failed to find node: %v", result.Error))
	} else {
		return nil
	}
}

// deleteNode 递归删除节点及其子节点
func deleteNode(db *gorm.DB, elemID uint, target interface{}, children interface{}) *models.CustomError {
	// 复制一个空白的children用来下一次递归调用
	newChildren, err := utils.GetVoidSlice(children)
	if err != nil {
		return models.InvalidArgError(err.Error())
	}
	// 查找子节点
	result := db.Where("parent_id = ?", elemID).Find(&children)
	if result.Error != nil {
		return models.SQLError(fmt.Sprintf("failed to find children nodes: %v", result.Error))
	}
	childrenValue := reflect.ValueOf(children)
	// 遍历子节点并递归删除
	for i := 0; i < childrenValue.Len(); i++ {
		child := childrenValue.Index(i).Addr().Interface()
		childID := child.(models.BaseInfo).GetID()
		if err := deleteNode(db, childID, target, newChildren); err != nil {
			return err
		}
	}

	// 删除当前节点
	result = db.Delete(target, elemID)
	if result.Error != nil {
		return models.SQLError(fmt.Sprintf("failed to delete node: %v", result.Error))
	} else {
		return nil
	}
}
