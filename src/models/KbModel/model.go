package KbModel

import (
	"fmt"
	"time"
)

type KbImpl struct {
	ID         int       `gorm:"column:kb_id;primary_key"`
	Name       string    `gorm:"column:kb_name"`
	Desc       string    `gorm:"column:kb_desc"`
	Kind       int       `gorm:"-"`
	CreatorID  int       `gorm:"column:creator_id"`
	IsPrivate  string    `gorm:"-"`
	CreateTime time.Time `gorm:"-"`
	State      string    `gorm:"-"`
}

type KbDetail struct {

}

func (this *KbImpl) String() string {
	str := fmt.Sprintf(`{kbId:%d,kbName:%s,desc:%s,kind:%d,creatorID:%d,isPrivate:%s,createTime:%s,state:%s}`,
		this.ID, this.Name, this.Desc, this.Kind, this.CreatorID, this.IsPrivate, this.CreateTime.Format("2006-01-02 15:04:05"), this.State)
	return str
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

