package DocModel

import "time"

type DocImpl struct {
	ID           int       `gorm:"column:doc_id;primary_key"`
	KbID         int       `gorm:"column:kb_id"`
	Title        string    `gorm:"column:doc_title"`
	TitleUrl     string    `gorm:"column:doc_title_url"`
	Content      string    `gorm:"column:doc_content"`
	CreatorId    int       `gorm:"column:creator_id"`
	LastEditorId int       `gorm:"-"`
	UpdatedAt    time.Time `gorm:"-"`
	Removed      string    `gorm:"-"`
	GroupID      int       `gorm:"column:group_id"`
}

func New(attrs ...DocModelAttrFunc) *DocImpl {
	d := &DocImpl{}
	DocModelAttrFuncs(attrs).Apply(d)
	return d
}

func (this *DocImpl) Mutate(attrs ...DocModelAttrFunc) *DocImpl {
	DocModelAttrFuncs(attrs).Apply(this)
	return this
}
