package DocGrpModel

import "time"

type DocGrpImpl struct {
	ID         int       `gorm:"column:group_id;primary_key"`
	GroupName  string    `gorm:"column:group_name;"`
	KbID       int       `gorm:"column:kb_id;"`
	CreateTime time.Time `gorm:"-"`
	DocCount   int       `gorm:"-"`
	CreatorId  int       `gorm:"creator_id"`
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
