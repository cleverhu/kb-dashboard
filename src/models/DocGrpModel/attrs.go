package DocGrpModel

import "time"

type DocGrpModelAttrFunc func(d *DocGrpImpl)

type DocGrpModelAttrFuncs []DocGrpModelAttrFunc

func WithGroupID(id int) DocGrpModelAttrFunc {
	return func(d *DocGrpImpl) {
		d.ID = id
	}
}

func WithGroupName(name string) DocGrpModelAttrFunc {
	return func(d *DocGrpImpl) {
		d.GroupName = name
	}
}

func WithKbID(kbId int) DocGrpModelAttrFunc {
	return func(d *DocGrpImpl) {
		d.KbID = kbId
	}
}

func WithCreateTime(time time.Time) DocGrpModelAttrFunc {
	return func(d *DocGrpImpl) {
		d.CreateTime = time
	}
}

func WithCreatorID(creatorID int) DocGrpModelAttrFunc {
	return func(d *DocGrpImpl) {
		d.CreatorID = creatorID
	}
}

func (this DocGrpModelAttrFuncs) Apply(d *DocGrpImpl) {
	for _, f := range this {
		f(d)
	}
}
