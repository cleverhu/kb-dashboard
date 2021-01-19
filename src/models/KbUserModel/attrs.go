package KbUserModel

import "time"

type KbUserModelAttrFunc func(kb *KbUserImpl)

type KbUserModelAttrFuncs []KbUserModelAttrFunc

func WithKbID(kbId int) KbUserModelAttrFunc {
	return func(ks *KbUserImpl) {
		ks.KbID = kbId
	}
}

func WithUserID(userID int) KbUserModelAttrFunc {
	return func(ks *KbUserImpl) {
		ks.UserID = userID
	}
}

func WithJoinTime(joinTime time.Time) KbUserModelAttrFunc {
	return func(ks *KbUserImpl) {
		ks.JoinTime = joinTime
	}
}

func WithCanEdit(canEdit string) KbUserModelAttrFunc {
	return func(ks *KbUserImpl) {
		ks.CanEdit = canEdit
	}
}

func (this KbUserModelAttrFuncs) Apply(ks *KbUserImpl) {
	for _, f := range this {
		f(ks)
	}
}
