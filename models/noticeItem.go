package models

type NoticeItem struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(100);not null" json:"title"`
	PublishTime string `gorm:"type:varchar(50);not null" json:"publishTime"`
	Detail      string `gorm:"not null" json:"detail"`
}

func (n *NoticeItem) GetID() uint {
	return n.ID
}

func (n *NoticeItem) GetTitle() string {
	return n.Title
}

func (n *NoticeItem) Update(item interface{}) {
	updateStructFields(n, item)
}
