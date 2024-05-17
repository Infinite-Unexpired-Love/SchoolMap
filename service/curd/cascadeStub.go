package curd

import (
	"TGU-MAP/models"
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

// FetchData 获取数据并返回包含元素的切片
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
func (stub *CascadeStub) InsertNodeByPath(item models.BaseInfo, path ...string) *models.CustomError {
	if len(path) == 0 {
		// 如果路径为空，直接插入节点
		return insertNode(stub.db, 0, item)
	} else {
		// 否则，根据路径查找父节点并插入子节点
		if id, err := stub.FindElemID(path); err != nil {
			return err
		} else {
			return insertNode(stub.db, id, item)
		}

	}
}

// InsertNodeByID 根据父节点 ID 插入节点
func (stub *CascadeStub) InsertNodeByID(item models.BaseInfo, parentID uint) *models.CustomError {
	return insertNode(stub.db, parentID, item)
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
func (stub *CascadeStub) UpdateNodeByID(item *models.ListItem, elemID uint) *models.CustomError {
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
	return deleteNode(stub.db, elemID, target, children)
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
	return deleteNode(stub.db, elemID, target, children)
}

// FindElemID 根据路径查找节点 ID
func (stub *CascadeStub) FindElemID(path []string) (uint, *models.CustomError) {
	// 创建一个新的零值实例作为目标
	target := stub.NewInstance()
	// 查找路径对应的节点 ID
	return findElemID(stub.db, target, path...)
}
