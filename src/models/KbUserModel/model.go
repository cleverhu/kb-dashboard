package KbUserModel

import "time"

type KbUserImpl struct {
	KbID     int       `gorm:"column:kb_id"`
	UserID   int       `gorm:"column:user_id"`
	JoinTime time.Time `gorm:"-"`
	CanEdit  string    `gorm:"-"`
}

func New(attrs ...KbUserModelAttrFunc) *KbUserImpl {
	ks := &KbUserImpl{}
	KbUserModelAttrFuncs(attrs).Apply(ks)
	return ks
}

func (this *KbUserImpl) Mutate(attrs ...KbUserModelAttrFunc) *KbUserImpl {
	KbUserModelAttrFuncs(attrs).Apply(this)
	return this
}
