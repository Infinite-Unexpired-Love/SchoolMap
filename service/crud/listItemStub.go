package crud

import (
	"TGU-MAP/models"
	"gorm.io/gorm"
	"reflect"
)

type ListItemStub struct {
	s *CascadeStub
}

func NewListItemStub(db *gorm.DB) *ListItemStub {
	return &ListItemStub{s: NewCascadeStub(db, models.ListItem{})}
}

func (stub *ListItemStub) setChildren(item *models.ListItem, itemMap map[uint][]models.ListItem) {
	children := itemMap[item.ID]
	for i := range children {
		stub.setChildren(&children[i], itemMap)
	}
	item.Children = children
}

func (stub *ListItemStub) setChildrenIterative(root *models.ListItem, itemMap map[uint][]models.ListItem) {
	stack := []*models.ListItem{root}

	for len(stack) > 0 {
		// Pop
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Get children
		children := itemMap[current.ID]
		current.Children = children

		// Push children onto stack
		for i := len(children) - 1; i >= 0; i-- {
			stack = append(stack, &children[i])
		}
	}
}

// Init 插入ID为0的数据，满足外键约束
func (stub *ListItemStub) Init() *models.CustomError {
	return stub.s.Init(&models.ListItem{ID: 0})
}

// FetchData 返回树形结构数据
func (stub *ListItemStub) FetchData() (*[]models.ListItem, *models.CustomError) {
	pdata, err := stub.s.FetchData()
	if err != nil {
		return nil, err
	}
	data := reflect.ValueOf(pdata).Elem().Interface()
	items := data.([]models.ListItem)

	itemMap := make(map[uint][]models.ListItem)
	var parentID uint
	for _, item := range items {
		if item.ParentID == nil {
			parentID = 0
		} else {
			parentID = *item.ParentID
		}
		itemMap[parentID] = append(itemMap[parentID], item)
	}

	var rootItems []models.ListItem

	for _, item := range itemMap[0] {
		stub.setChildren(&item, itemMap)
		rootItems = append(rootItems, item)
	}
	return &rootItems, nil
}

// InsertData 插入构造好的数据，无法保证插入节点数量，不具备去重功能
func (stub *ListItemStub) InsertData(data *[]models.ListItem) *models.CustomError {
	result := stub.s.db.Create(data)
	if result.Error != nil {
		return models.SQLError("failed to insert data")
	}
	return nil
}

func (stub *ListItemStub) InsertNodeByPath(item *models.ListItem, path ...string) *models.CustomError {
	return stub.s.InsertNodeByPath(item, path...)
}

func (stub *ListItemStub) InsertNodeByID(item *models.ListItem, parentID *uint) *models.CustomError {
	return stub.s.InsertNodeByID(item, parentID)
}

func (stub *ListItemStub) UpdateNodeByPath(item *models.ListItem, path ...string) *models.CustomError {
	return stub.s.UpdateNodeByPath(item, path...)
}

func (stub *ListItemStub) UpdateNodeByID(item *models.ListItem, elemID uint) *models.CustomError {
	return stub.s.UpdateNodeByID(item, elemID)
}

func (stub *ListItemStub) DeleteNodeByID(elemID uint) *models.CustomError {
	return stub.s.DeleteNodeByID(elemID)
}

func (stub *ListItemStub) DeleteNodeByPath(path ...string) *models.CustomError {
	return stub.s.DeleteNodeByPath(path...)
}

func (stub *ListItemStub) FindElemID(path ...string) (uint, *models.CustomError) {
	return stub.s.FindElemID(path)
}
