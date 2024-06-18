package aliasItem

import (
	"TGU-MAP/models"
	"TGU-MAP/service/crud"
	"TGU-MAP/service/listItem"
	"gorm.io/gorm"
	"reflect"
)

type AliasItemStub struct {
	s *crud.NormalStub
}

func NewAliasItemStub(db *gorm.DB) *AliasItemStub {
	return &AliasItemStub{s: crud.NewNormalStub(db, models.AliasItem{})}
}

// TODO:优化
func setFullName(lic *listItem.ListItemStub, li *models.ListItem) *models.CustomError {
	cur := li
	for cur.ParentID != nil {
		parent, err := lic.FindNodeByID(*cur.ParentID)
		if err != nil {
			return err
		}
		li.Title = parent.Title + "-" + li.Title
		cur = parent
	}
	return nil
}

func (stub *AliasItemStub) setMarkers(lic *listItem.ListItemStub, a *models.AliasItem) *models.CustomError {

	var markers []models.ListItem
	if err := stub.s.Db.Model(&a).Association("Markers").Find(&markers); err != nil {
		return models.SQLError("设置markers出错")
	}

	for i := 0; i < len(markers); i++ {
		if err := setFullName(lic, &markers[i]); err != nil {
			return err
		}
	}

	a.Markers = markers

	return nil
}

func (stub *AliasItemStub) FetchData(lic *listItem.ListItemStub) (*[]models.AliasItem, *models.CustomError) {
	pdata, err := stub.s.FetchData()
	if err != nil {
		return nil, err
	}
	data := reflect.ValueOf(pdata).Elem().Interface()
	items := data.([]models.AliasItem)
	for i := 0; i < len(items); i++ {
		if err := stub.setMarkers(lic, &items[i]); err != nil {
			return nil, err
		}
	}
	return &items, nil
}

func (stub *AliasItemStub) InsertNode(item *models.AliasItem) *models.CustomError {
	return stub.s.InsertNode(item)
}

//func (stub *AliasItemStub) UpdateNodeByID(item *models.AliasItem, elemID uint) *models.CustomError {
//	return stub.s.UpdateNodeByID(item, elemID)
//}

func (stub *AliasItemStub) DeleteNodeByID(elemID uint) *models.CustomError {
	//先删除多对多关系记录
	alias, err := stub.s.FindElem("id = ?", elemID)
	if err != nil {
		return err
	}
	aItem := alias.(*models.AliasItem)
	var ls []models.ListItem
	stub.s.Db.Model(aItem).Association("Markers").Find(&ls)
	stub.s.Db.Model(aItem).Association("Markers").Delete(&ls)
	return stub.s.DeleteNodeByID(elemID)
}
