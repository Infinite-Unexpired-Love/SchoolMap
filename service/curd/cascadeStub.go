package curd

import (
	"TGU-MAP/models"
	"gorm.io/gorm"
	"reflect"
)

type CascadeStub struct {
	db       *gorm.DB
	elemType reflect.Type
}

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

func (stub *CascadeStub) NewSlice() interface{} {
	// 创建一个新的切片实例
	sliceType := reflect.SliceOf(stub.elemType)
	newSlice := reflect.MakeSlice(sliceType, 0, 0).Interface()
	return newSlice
}

func (stub *CascadeStub) FetchData() interface{} {
	items := stub.NewSlice()
	fetchData(stub.db, &items)
	return &items
}

func (stub *CascadeStub) InsertNodeByPath(item models.Cascade, path ...string) {
	if len(path) == 0 {
		insertNode(stub.db, nil, item)
	} else {
		insertNode(stub.db, stub.FindElemID(path), item)
	}
}

func (stub *CascadeStub) InsertNodeByID(item models.Cascade, parentID uint) {
	insertNode(stub.db, &parentID, item)
}

func (stub *CascadeStub) UpdateNodeByPath(item models.Updatable, path ...string) {
	target := stub.NewInstance()
	updateNode(stub.db, stub.FindElemID(path), target, item)
}

func (stub *CascadeStub) UpdateNodeByID(item *models.ListItem, elemID uint) {
	target := stub.NewInstance()
	updateNode(stub.db, &elemID, target, item)
}

func (stub *CascadeStub) DeleteNodeByID(elemID uint) {
	target := stub.NewInstance()
	children := stub.NewSlice()
	deleteNode(stub.db, &elemID, target, children)
}

func (stub *CascadeStub) DeleteNodeByPath(path ...string) {
	elemID := stub.FindElemID(path)
	target := stub.NewInstance()
	children := stub.NewSlice()
	deleteNode(stub.db, elemID, target, children)
}

func (stub *CascadeStub) FindElemID(path []string) *uint {
	target := stub.NewInstance()
	return findElemID(stub.db, target, path...)
}
