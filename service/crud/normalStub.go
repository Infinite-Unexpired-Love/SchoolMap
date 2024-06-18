package crud

import (
	"TGU-MAP/models"
	"gorm.io/gorm"
	"reflect"
)

type NormalStub struct {
	Db       *gorm.DB
	elemType reflect.Type
}

func NewNormalStub(db *gorm.DB, elem interface{}) *NormalStub {
	// 获取输入的反射类型
	inputType := reflect.TypeOf(elem)
	// 如果输入是指针类型，需要获取其元素类型
	if inputType.Kind() == reflect.Ptr {
		inputType = inputType.Elem()
	}

	return &NormalStub{Db: db, elemType: inputType}
}

// NewInstance 返回一个指向零值实例的指针
func (stub *NormalStub) NewInstance() interface{} {
	// 创建一个新的零值实例
	newInstance := reflect.New(stub.elemType).Interface()
	return newInstance
}

// NewSlice 返回一个空的元素类型切片
func (stub *NormalStub) NewSlice() interface{} {
	// 创建一个新的切片实例
	sliceType := reflect.SliceOf(stub.elemType)
	newSlice := reflect.MakeSlice(sliceType, 0, 0).Interface()
	return newSlice
}

// Init 满足外键约束
//func (stub *CascadeStub) Init(target models.BaseInfo) *models.CustomError {
//	return insertNode(stub.db, target)
//}

// FetchData 获取所有数据并返回包含元素的切片
func (stub *NormalStub) FetchData() (interface{}, *models.CustomError) {
	// 创建一个新的切片实例
	items := stub.NewSlice()
	// 从数据库中获取数据
	if err := fetchData(stub.Db, &items); err != nil {
		return nil, err
	}
	return &items, nil
}

// Paginate 分页查询
func (stub *NormalStub) Paginate(offset, limit int) (interface{}, *models.CustomError) {
	// 创建一个新的切片实例
	items := stub.NewSlice()
	// 从数据库中获取数据
	if err := paginate(stub.Db, offset, limit, &items); err != nil {
		return nil, err
	}
	return &items, nil
}

func (stub *NormalStub) Count() (int64, *models.CustomError) {
	// 创建一个新的切片实例
	item := stub.NewInstance()

	return count(stub.Db, item)
}

func (stub *NormalStub) InsertNode(item models.BaseInfo) *models.CustomError {
	return insertNode(stub.Db, item)
}

func (stub *NormalStub) UpdateNodeByID(item models.Updatable, elemID uint) *models.CustomError {
	// 创建一个新的零值实例指明要操作的表
	target := stub.NewInstance()
	// 更新节点
	return updateNode(stub.Db, elemID, target, item)

}

func (stub *NormalStub) DeleteNodeByID(elemID uint) *models.CustomError {

	target := stub.NewInstance()
	return deleteNode(stub.Db, elemID, target)
}

func (stub *NormalStub) FindElem(conds ...interface{}) (interface{}, *models.CustomError) {
	// 创建一个新的零值实例作为目标
	target := stub.NewInstance()
	// 查找路径对应的节点 ID
	err := findNode(stub.Db, target, conds...)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func (stub *NormalStub) FindElems(conds ...interface{}) (interface{}, *models.CustomError) {
	// 创建一个新的零值实例作为目标
	target := stub.NewSlice()
	// 查找路径对应的节点 ID
	err := findNodes(stub.Db, &target, conds...)
	if err != nil {
		return nil, err
	}
	return target, nil
}
