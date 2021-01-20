package DocModel

import "time"

type DocModelAttrFunc func(d *DocImpl)

type DocModelAttrFuncs []DocModelAttrFunc

func WithDocID(id int) DocModelAttrFunc {
	return func(d *DocImpl) {
		d.ID = id
	}
}

func WithKbID(kbId int) DocModelAttrFunc {
	return func(d *DocImpl) {
		d.KbID = kbId
	}
}

func WithDocTitle(title string) DocModelAttrFunc {
	return func(d *DocImpl) {
		d.Title = title
	}
}

func WithTitleUrl(titleUrl string) DocModelAttrFunc {
	return func(d *DocImpl) {
		d.TitleUrl = titleUrl
	}
}

func WithContent(content string) DocModelAttrFunc {
	return func(d *DocImpl) {
		d.Content = content
	}
}

func WithCreatorID(creatorID int) DocModelAttrFunc {
	return func(d *DocImpl) {
		d.CreatorID = creatorID
	}
}

func WithLastEditorID(lastEditorID int) DocModelAttrFunc {
	return func(d *DocImpl) {
		d.LastEditorID = lastEditorID
	}
}

func WithUpdateTime(UpdateTime time.Time) DocModelAttrFunc {
	return func(d *DocImpl) {
		d.UpdatedAt = UpdateTime
	}
}

func WithRemoved(removed string) DocModelAttrFunc {
	return func(d *DocImpl) {
		d.Removed = removed
	}
}

func WithGroupID(groupID int) DocModelAttrFunc {
	return func(d *DocImpl) {
		d.GroupID = groupID
	}
}

func (this DocModelAttrFuncs) Apply(d *DocImpl) {
	for _, f := range this {
		f(d)
	}
}
