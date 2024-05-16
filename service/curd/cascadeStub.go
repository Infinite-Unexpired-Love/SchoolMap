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

func (stub *CascadeStub) NewInstance() interface{} {
	// 创建一个新的零值实例
	newInstance := reflect.New(stub.elemType).Elem().Interface()

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
	parentID := stub.FindElemID(path)
	insertNode(stub.db, &parentID, item)
}

func (stub *CascadeStub) InsertNodeByID(item models.Cascade, parentID uint) {
	insertNode(stub.db, &parentID, item)
}

func (stub *CascadeStub) UpdateNodeByPath(item models.Updatable, path ...string) {
	target := stub.NewInstance()
	updateNode(stub.db, stub.FindElemID(path), &target, item)
}

func (stub *CascadeStub) UpdateNodeByID(item *models.ListItem, elemID uint) {
	target := stub.NewInstance()
	updateNode(stub.db, elemID, &target, item)
}

func (stub *CascadeStub) DeleteNodeByID(elemID uint) {
	target := stub.NewInstance()
	children := stub.NewSlice()
	deleteNode(stub.db, elemID, &target, children)
}

func (stub *CascadeStub) DeleteNodeByPath(path ...string) {
	elemID := stub.FindElemID(path)
	target := stub.NewInstance()
	children := stub.NewSlice()
	deleteNode(stub.db, elemID, &target, children)
}

func (stub *CascadeStub) FindElemID(path []string) uint {
	cur := stub.NewInstance()

	//var elemID *uint
	//
	//result := stub.db.Where("title = ? AND parent_id IS NULL", path[0]).First(&cur)
	//if result.Error != nil {
	//	log.Fatalf("failed to find parent node: %v", result.Error)
	//}
	////TODO: 错误，强转失败
	//elemID = cur.(models.Cascade).GetID()
	//for _, title := range path[1:] {
	//	//result := db.Where("title = ?", title).Where("parent_id = ?", elemID).First(&parent)
	//	result := stub.db.Raw("SELECT * FROM list_items WHERE title = ? AND parent_id = ? Limit 1", title, *elemID).Scan(&cur)
	//	if result.Error != nil {
	//
	//		log.Fatalf("failed to find parent node: %v", result.Error)
	//	}
	//	elemID = cur.(models.Cascade).GetID()
	//}

	return *findElemID(stub.db, &cur, path...)
}
