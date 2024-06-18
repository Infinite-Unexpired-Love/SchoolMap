package models

import "gorm.io/gorm"

type AliasItem struct {
	ID       uint       `gorm:"primaryKey" json:"id"`
	Title    string     `gorm:"type:varchar(50);not null" json:"title"`
	Markers  []ListItem `gorm:"many2many:alias_item_list_items" json:"markers"`
	DeleteAt gorm.DeletedAt
}

func (a *AliasItem) GetID() uint {
	return a.ID
}

func (a *AliasItem) GetTitle() string {
	return a.Title
}
