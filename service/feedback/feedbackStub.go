package feedback

import (
	"TGU-MAP/models"
	"TGU-MAP/service/crud"
	"gorm.io/gorm"
	"reflect"
)

type FeedbackStub struct {
	s *crud.NormalStub
}

func NewFeedbackStub(db *gorm.DB) *FeedbackStub {
	return &FeedbackStub{s: crud.NewNormalStub(db, models.Feedback{})}
}

func (stub *FeedbackStub) FetchData() (*[]models.Feedback, *models.CustomError) {
	pdata, err := stub.s.FetchData()
	if err != nil {
		return nil, err
	}
	data := reflect.ValueOf(pdata).Elem().Interface()
	items := data.([]models.Feedback)
	return &items, nil
}

func (stub *FeedbackStub) Paginate(offset, pageSize int) (*[]models.Feedback, *models.CustomError) {
	pdata, err := stub.s.Paginate(offset, pageSize)
	if err != nil {
		return nil, err
	}
	data := reflect.ValueOf(pdata).Elem().Interface()
	items := data.([]models.Feedback)
	return &items, nil
}

func (stub *FeedbackStub) Count() (int64, *models.CustomError) {
	return stub.s.Count()
}

func (stub *FeedbackStub) InsertNode(item *models.Feedback) *models.CustomError {
	return stub.s.InsertNode(item)
}

func (stub *FeedbackStub) DeleteNodeByID(elemID uint) *models.CustomError {
	return stub.s.DeleteNodeByID(elemID)
}
