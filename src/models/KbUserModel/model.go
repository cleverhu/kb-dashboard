package KbUserModel

import "time"

type KbUserImpl struct {
	KbID     int       `gorm:"column:kb_id"`
	UserID   int       `gorm:"column:user_id"`
	JoinTime time.Time `gorm:"-"`
	CanEdit  string    `gorm:"-"`
}

func New(attrs ...KbUserModelAttrFunc) *KbUserImpl {
	kbs := &KbUserImpl{}
	KbUserModelAttrFuncs(attrs).Apply(kbs)
	return kbs
}

func (this *KbUserImpl) Mutate(attrs ...KbUserModelAttrFunc) *KbUserImpl {
	KbUserModelAttrFuncs(attrs).Apply(this)
	return this
}
