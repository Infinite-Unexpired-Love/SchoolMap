package user

import (
	"TGU-MAP/models"
	"TGU-MAP/service/crud"
	"gorm.io/gorm"
	"reflect"
)

type UserStub struct {
	s *crud.NormalStub
}

func NewUserStub(db *gorm.DB) *UserStub {
	return &UserStub{s: crud.NewNormalStub(db, models.User{})}
}

func (stub *UserStub) FetchData() (*[]models.User, *models.CustomError) {
	pdata, err := stub.s.FetchData()
	if err != nil {
		return nil, err
	}
	data := reflect.ValueOf(pdata).Elem().Interface()
	items := data.([]models.User)
	return &items, nil
}

func (stub *UserStub) InsertNode(item *models.User) *models.CustomError {
	return stub.s.InsertNode(item)
}

func (stub *UserStub) UpdateNodeByID(item *models.User, elemID uint) *models.CustomError {
	return stub.s.UpdateNodeByID(item, elemID)
}

func (stub *UserStub) DeleteNodeByID(elemID uint) *models.CustomError {
	return stub.s.DeleteNodeByID(elemID)
}

func (stub *UserStub) FindElemByID(elemID uint) (*models.User, *models.CustomError) {
	user, err := stub.s.FindElem("id=?", elemID)
	if err != nil {
		return nil, err
	}
	return user.(*models.User), nil
}

func (stub *UserStub) FindElemByMobile(mobile string) (*models.User, *models.CustomError) {
	user, err := stub.s.FindElem("mobile=?", mobile)
	if err != nil {
		return nil, err
	}
	return user.(*models.User), nil
}
