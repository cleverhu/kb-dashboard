package DocGrpModel

import (
	"fmt"
	"time"
)

type DocGrpImpl struct {
	ID         int       `gorm:"column:group_id;primary_key"`
	GroupName  string    `gorm:"column:group_name;"`
	KbID       int       `gorm:"column:kb_id;"`
	CreateTime time.Time `gorm:"-"`
	DocCount   int       `gorm:"-"`
	CreatorID  int       `gorm:"creator_id"`
	ShortUrl   string    `gorm:"column:shorturl"`
}



func (this *DocGrpImpl) String() string {
	return fmt.Sprintf("{groupId:%d,groupName:%s,kbId:%d,createTime:%s,docCount:%d,creatorID:%d}", this.ID, this.GroupName, this.KbID, this.CreateTime.Format("2006-01-02 15:04:05"), this.DocCount, this.CreatorID)
}

func New(attrs ...DocGrpModelAttrFunc) *DocGrpImpl {
	d := &DocGrpImpl{}
	DocGrpModelAttrFuncs(attrs).Apply(d)
	return d
}

func (this *DocGrpImpl) Mutate(attrs ...DocGrpModelAttrFunc) *DocGrpImpl {

	DocGrpModelAttrFuncs(attrs).Apply(this)
	return this
}
