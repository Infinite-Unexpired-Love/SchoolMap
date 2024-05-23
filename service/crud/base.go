package crud

import (
	"TGU-MAP/models"
	"fmt"
	"gorm.io/gorm"
)

// fetchData 从数据库中获取所有item
func fetchData(db *gorm.DB, items interface{}) *models.CustomError {
	result := db.Find(items)
	if result.Error != nil {
		return models.SQLError(fmt.Sprintf("failed to fetch data: %v", result.Error))
	}
	return nil
}

// insertNode 插入新节点
func insertNode(db *gorm.DB, item models.BaseInfo) *models.CustomError {
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

func deleteNode(db *gorm.DB, elemID uint, target interface{}) *models.CustomError {
	// 删除节点
	result := db.Delete(target, elemID)
	if result.Error != nil {
		return models.SQLError(fmt.Sprintf("failed to delete node: %v", result.Error))
	} else {
		return nil
	}
}
