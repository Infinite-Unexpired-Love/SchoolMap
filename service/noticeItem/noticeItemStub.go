package noticeItem

import (
	"TGU-MAP/models"
	"TGU-MAP/service/crud"
	"gorm.io/gorm"
	"reflect"
)

type NoticeItemStub struct {
	s *crud.NormalStub
}

func NewNoticeItemStub(db *gorm.DB) *NoticeItemStub {
	return &NoticeItemStub{s: crud.NewNormalStub(db, models.NoticeItem{})}
}

func (stub *NoticeItemStub) FetchData() (*[]models.NoticeItem, *models.CustomError) {
	pdata, err := stub.s.FetchData()
	if err != nil {
		return nil, err
	}
	data := reflect.ValueOf(pdata).Elem().Interface()
	items := data.([]models.NoticeItem)
	return &items, nil
}

func (stub *NoticeItemStub) InsertNode(item *models.NoticeItem) *models.CustomError {
	return stub.s.InsertNode(item)
}

func (stub *NoticeItemStub) UpdateNodeByID(item *models.NoticeItem, elemID uint) *models.CustomError {
	return stub.s.UpdateNodeByID(item, elemID)
}

func (stub *NoticeItemStub) DeleteNodeByID(elemID uint) *models.CustomError {
	return stub.s.DeleteNodeByID(elemID)
}
