package KbModel

import (
	"fmt"
	"time"
)

type KbImpl struct {
	ID         int       `gorm:"column:kb_id;primary_key" json:"-"`
	Name       string    `gorm:"column:kb_name" json:"kb_name"`
	Desc       string    `gorm:"column:kb_desc" json:"kb_desc"`
	Kind       int       `gorm:"-" json:"-"`
	CreatorID  int       `gorm:"column:creator_id" json:"creator_id"`
	IsPrivate  string    `gorm:"is_private" json:"is_private"`
	CreateTime time.Time `gorm:"create_time" json:"create_time"`
	State      int       `gorm:"-" json:"-"`
}

type KbInputRequest struct {
	ID        int    `gorm:"column:kb_id;primary_key" json:"-"`
	Name      string `gorm:"column:kb_name" json:"kb_name"`
	Desc      string `gorm:"column:kb_desc" json:"kb_desc"`
	Kind      int    `gorm:"column:kb_kind_id" json:"kb_kind_id"`
	IsPrivate string `gorm:"is_private" json:"is_private"`
	CreatorID int    `gorm:"column:creator_id" json:"-"`
}

type User struct {
	Name string `gorm:"column:name" json:"user_name"`
}

type UserID struct {
	ID int `gorm:"column:user_id"`
}

type KbDetail struct {
	KbInfo *KbImpl   `json:"kb_info"`
	UserID []*UserID `json:"users"`
}

func (this *KbImpl) String() string {
	str := fmt.Sprintf(`{kbId:%d,kbName:%s,desc:%s,kind:%d,creatorID:%d,isPrivate:%s,createTime:%s,state:%s}`,
		this.ID, this.Name, this.Desc, this.Kind, this.CreatorID, this.IsPrivate, this.CreateTime, this.State)
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
