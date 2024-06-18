package models

type Feedback struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Title       string  `gorm:"type:varchar(50);not null" json:"title"`
	Category    uint    `gorm:"not null" json:"category"`
	Contact     *string `gorm:"type:varchar(200)" json:"contact"`
	PublishTime string  `gorm:"type:varchar(50);not null" json:"publishTime"`
	Detail      string  `gorm:"not null" json:"detail"`
}

func (f *Feedback) GetID() uint {
	return f.ID
}

func (f *Feedback) GetTitle() string {
	return f.Title
}
