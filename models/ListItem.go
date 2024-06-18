package models

type ListItem struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"type:varchar(100)" json:"title"`
	Desc      string     `json:"desc"`
	Contact   *string    `json:"contact,omitempty"`
	Latitude  *float64   `json:"latitude,omitempty"`
	Longitude *float64   `json:"longitude,omitempty"`
	IconName  *string    `json:"iconName,omitempty"`
	ParentID  *uint      `json:"-"`
	Children  []ListItem `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}

func (target *ListItem) Update(item interface{}) {
	updateStructFields(target, item)
}

func (target *ListItem) SetParentID(parentID *uint) {
	target.ParentID = parentID
}

func (target *ListItem) GetParentID() *uint {
	return target.ParentID
}

func (target *ListItem) GetID() uint {
	return target.ID
}

func (target *ListItem) GetTitle() string {
	return target.Title
}
