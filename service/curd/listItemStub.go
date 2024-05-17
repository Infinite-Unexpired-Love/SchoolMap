package curd

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

// FetchData 返回树形结构数据
func (stub *ListItemStub) FetchData() (*[]models.ListItem, error) {
	pdata, err := stub.s.FetchData()
	if err != nil {
		return nil, err
	}
	data := reflect.ValueOf(pdata).Elem().Interface()
	items := data.([]models.ListItem)

	itemMap := make(map[uint][]models.ListItem)
	for _, item := range items {
		parentID := item.ParentID
		itemMap[*parentID] = append(itemMap[*parentID], item)
	}

	var rootItems []models.ListItem
	for _, item := range itemMap[0] {
		stub.setChildren(&item, itemMap)
		rootItems = append(rootItems, item)
	}
	return &rootItems, nil
}

func (stub *ListItemStub) InsertNodeByPath(item *models.ListItem, path ...string) *models.CustomError {
	return stub.s.InsertNodeByPath(item, path...)
}
