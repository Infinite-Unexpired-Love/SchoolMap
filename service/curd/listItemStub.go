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

func (stub *ListItemStub) FetchData() *[]models.ListItem {
	ptr := stub.s.FetchData()
	items := reflect.ValueOf(ptr).Elem().Interface()
	listItems := items.([]models.ListItem)
	return &listItems
}

func (stub *ListItemStub) InsertNodeByPath(item *models.ListItem, path ...string) {
	stub.s.InsertNodeByPath(item, path...)
}
