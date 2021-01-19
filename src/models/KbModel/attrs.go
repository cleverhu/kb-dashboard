package KbModel

import "time"

type KbModelAttrFunc func(kb *KbImpl)

type KbModelAttrFuncs []KbModelAttrFunc


func WithKbID(kbId int) KbModelAttrFunc {
	return func(kb *KbImpl) {
		kb.ID = kbId
	}
}

func WithKbName(name string) KbModelAttrFunc {
	return func(kb *KbImpl) {
		kb.Name = name
	}
}

func WithDesc(desc string) KbModelAttrFunc {
	return func(kb *KbImpl) {
		kb.Desc = desc
	}
}

func WithKind(kind int) KbModelAttrFunc {
	return func(kb *KbImpl) {
		kb.Kind = kind
	}
}

func WithCreatorID(creatorID int) KbModelAttrFunc {
	return func(kb *KbImpl) {
		kb.CreatorId = creatorID
	}
}

func WithLastEditorID(isPrivate string) KbModelAttrFunc {
	return func(kb *KbImpl) {
		kb.IsPrivate = isPrivate
	}
}

func WithCreateTime(createTime time.Time) KbModelAttrFunc {
	return func(kb *KbImpl) {
		kb.CreateTime = createTime
	}
}

func WithState(state string) KbModelAttrFunc {
	return func(kb *KbImpl) {
		kb.State = state
	}
}


func (this KbModelAttrFuncs) Apply(kb *KbImpl) {
	for _, f := range this {
		f(kb)
	}
}