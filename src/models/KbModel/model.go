package KbModel

import "time"

type KbImpl struct {
	ID         int       `gorm:"column:kb_id;primary_key"`
	Name       string    `gorm:"column:kb_name"`
	Desc       string    `gorm:"column:kb_desc"`
	Kind       int       `gorm:"-"`
	CreatorId  int       `gorm:"column:creator_id"`
	IsPrivate  string    `gorm:"-"`
	CreateTime time.Time `gorm:"-"`
	State      string    `gorm:"-"`
}


func New(attrs ...KbModelAttrFunc) *KbImpl {
	kb := &KbImpl{}
	KbModelAttrFuncs(attrs).Apply(kb)
	return kb
}

func (this *KbImpl) Mutate(attrs ...KbModelAttrFunc) *KbImpl {
	KbModelAttrFuncs(attrs).Apply(this)
	return this
}