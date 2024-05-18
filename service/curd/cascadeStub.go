package curd

import (
	"TGU-MAP/models"
	"TGU-MAP/utils"
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

// CascadeStub 包含数据库连接和元素类型信息
type CascadeStub struct {
	db       *gorm.DB
	elemType reflect.Type
}

// NewCascadeStub 创建一个新的 CascadeStub 实例
func NewCascadeStub(db *gorm.DB, elem interface{}) *CascadeStub {
	// 获取输入的反射类型
	inputType := reflect.TypeOf(elem)
	// 如果输入是指针类型，需要获取其元素类型
	if inputType.Kind() == reflect.Ptr {
		inputType = inputType.Elem()
	}

	return &CascadeStub{db: db, elemType: inputType}
}

// NewInstance 返回一个指向零值实例的指针
func (stub *CascadeStub) NewInstance() interface{} {
	// 创建一个新的零值实例
	newInstance := reflect.New(stub.elemType).Interface()
	return newInstance
}

// NewSlice 返回一个空的元素类型切片
func (stub *CascadeStub) NewSlice() interface{} {
	// 创建一个新的切片实例
	sliceType := reflect.SliceOf(stub.elemType)
	newSlice := reflect.MakeSlice(sliceType, 0, 0).Interface()
	return newSlice
}

// Init 满足外键约束
func (stub *CascadeStub) Init(target models.BaseInfo) *models.CustomError {
	return insertNode(stub.db, target)
}

// FetchData 获取数据并返回包含元素的切片,不进行树形结构化处理
func (stub *CascadeStub) FetchData() (interface{}, *models.CustomError) {
	// 创建一个新的切片实例
	items := stub.NewSlice()
	// 从数据库中获取数据
	if err := fetchData(stub.db, &items); err != nil {
		return nil, err
	}
	return &items, nil
}

// InsertNodeByPath 根据路径插入节点
func (stub *CascadeStub) InsertNodeByPath(item models.CascadeInfo, path ...string) *models.CustomError {
	if len(path) == 0 {
		// 如果路径为空，直接插入节点
		return stub.insertNode(item, nil)
	} else {
		// 否则，根据路径查找父节点并插入子节点
		if id, err := stub.FindElemID(path); err != nil {
			return err
		} else {
			return stub.insertNode(item, &id)
		}

	}
}

// InsertNodeByID 根据父节点 ID 插入节点
func (stub *CascadeStub) InsertNodeByID(item models.CascadeInfo, parentID *uint) *models.CustomError {
	return stub.insertNode(item, parentID)
}

// insertNode 设置父节点ID，保证同一父节点下的Title唯一，辅助函数
func (stub *CascadeStub) insertNode(item models.CascadeInfo, parentID *uint) *models.CustomError {
	item.SetParentID(parentID)
	var count int64
	if parentID != nil {
		stub.db.Model(item).Where("parent_id = ? AND title = ?", parentID, item.GetTitle()).Count(&count)
	} else {
		stub.db.Model(item).Where("parent_id IS NULL AND title = ?", item.GetTitle()).Count(&count)
	}

	if count > 0 {
		return models.SQLError(fmt.Sprintf("duplicate title: %v under one parentNode", item.GetTitle()))
	}
	return insertNode(stub.db, item)
}

// UpdateNodeByPath 根据路径更新节点
func (stub *CascadeStub) UpdateNodeByPath(item models.Updatable, path ...string) *models.CustomError {
	// 创建一个新的零值实例作为目标
	target := stub.NewInstance()
	// 更新节点
	if id, err := stub.FindElemID(path); err != nil {
		return err
	} else {
		return updateNode(stub.db, id, target, item)
	}
}

// UpdateNodeByID 根据节点 ID 更新节点
func (stub *CascadeStub) UpdateNodeByID(item models.Updatable, elemID uint) *models.CustomError {
	// 创建一个新的零值实例作为目标
	target := stub.NewInstance()
	// 更新节点
	return updateNode(stub.db, elemID, target, item)
}

// DeleteNodeByID 根据节点 ID 删除节点及其子节点
func (stub *CascadeStub) DeleteNodeByID(elemID uint) *models.CustomError {
	// 创建新的零值实例和空切片
	target := stub.NewInstance()
	children := stub.NewSlice()
	// 删除节点及其子节点
	return stub.deleteNode(elemID, target, children)
}

// DeleteNodeByPath 根据路径删除节点及其子节点
func (stub *CascadeStub) DeleteNodeByPath(path ...string) *models.CustomError {
	// 查找路径对应的节点 ID
	elemID, err := stub.FindElemID(path)
	if err != nil {
		return err
	}
	// 创建新的零值实例和空切片
	target := stub.NewInstance()
	children := stub.NewSlice()
	// 删除节点及其子节点
	return stub.deleteNode(elemID, target, children)
}

// deleteNode 处理递归调用逻辑，辅助函数
func (stub *CascadeStub) deleteNode(elemID uint, target interface{}, children interface{}) *models.CustomError {
	// 复制一个空白的children用来下一次递归调用
	newChildren, err := utils.GetVoidSlice(children)
	if err != nil {
		return models.InvalidArgError(err.Error())
	}
	// 查找子节点
	result := stub.db.Where("parent_id = ?", elemID).Find(&children)
	if result.Error != nil {
		return models.SQLError(fmt.Sprintf("failed to find children nodes: %v", result.Error))
	}
	childrenValue := reflect.ValueOf(children)
	// 遍历子节点并递归删除
	for i := 0; i < childrenValue.Len(); i++ {
		child := childrenValue.Index(i).Addr().Interface()
		childID := child.(models.BaseInfo).GetID()
		if err := stub.deleteNode(childID, target, newChildren); err != nil {
			return err
		}
	}
	return deleteNode(stub.db, elemID, target)
}

// FindElemID 根据路径查找节点 ID
func (stub *CascadeStub) FindElemID(path []string) (uint, *models.CustomError) {
	// 创建一个新的零值实例作为目标
	target := stub.NewInstance()
	// 查找路径对应的节点 ID
	return stub.findElemID(target, path...)
}

// findElemID 根据路径查找节点 ID，辅助函数
func (stub *CascadeStub) findElemID(target interface{}, path ...string) (uint, *models.CustomError) {
	var elemID uint

	// 查找根节点
	result := stub.db.Where("title = ? AND parent_id IS NULL", path[0]).First(target)
	if result.Error != nil {
		return 0, models.SQLError(fmt.Sprintf("failed to find parent node: %v", result.Error))
	}
	elemID = target.(models.BaseInfo).GetID()

	// 遍历路径的每个部分查找对应的子节点
	for _, title := range path[1:] {
		// TODO: 适配其他数据库
		result := stub.db.Raw("SELECT * FROM list_items WHERE title = ? AND parent_id = ? Limit 1", title, elemID).Scan(target)
		if result.Error != nil {
			return 0, models.SQLError(fmt.Sprintf("failed to find parent node: %v", result.Error))
		}
		elemID = target.(models.BaseInfo).GetID()
	}
	return elemID, nil
}
