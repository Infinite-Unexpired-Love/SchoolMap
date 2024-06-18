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

// paginate 分页获取数据
func paginate(db *gorm.DB, offset, limit int, items interface{}) *models.CustomError {
	result := db.Offset(offset).Limit(limit).Find(items)
	if result.Error != nil {
		return models.SQLError(fmt.Sprintf("failed to paginate data: %v", result.Error))
	}
	return nil
}

// count 统计数据总数
func count(db *gorm.DB, item interface{}) (int64, *models.CustomError) {
	var total int64
	result := db.Model(item).Count(&total)
	if result.Error != nil {
		return 0, models.SQLError(fmt.Sprintf("failed to count data: %v", result.Error))
	}
	return total, nil
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

func findNode(db *gorm.DB, dest interface{}, conds ...interface{}) *models.CustomError {
	result := db.First(dest, conds...)
	if result.Error != nil {
		return models.SQLError(fmt.Sprintf("failed to find node: %v", result.Error))
	} else {
		return nil
	}
}

func findNodes(db *gorm.DB, dest interface{}, conds ...interface{}) *models.CustomError {
	result := db.Find(dest, conds...)
	if result.Error != nil {
		return models.SQLError(fmt.Sprintf("failed to find nodes: %v", result.Error))
	} else {
		return nil
	}
}
