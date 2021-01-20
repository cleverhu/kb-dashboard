package KbUserModel

import (
	"fmt"
	"time"
)

type KbUserImpl struct {
	KbID     int       `gorm:"column:kb_id"`
	UserID   int       `gorm:"column:user_id"`
	JoinTime time.Time `gorm:"-"`
	CanEdit  string    `gorm:"-"`
}

func (this *KbUserImpl) String() string {
	str := fmt.Sprintf(`{kbId:%d,userID:%d,joinTime:%s,canEdit:%s}`,
		this.KbID, this.UserID, this.JoinTime.Format("2006-01-02 15:04:05"), this.CanEdit)
	return str
}

func New(attrs ...KbUserModelAttrFunc) *KbUserImpl {
	kbu := &KbUserImpl{}
	KbUserModelAttrFuncs(attrs).Apply(kbu)
	return kbu
}

func (this *KbUserImpl) Mutate(attrs ...KbUserModelAttrFunc) *KbUserImpl {
	KbUserModelAttrFuncs(attrs).Apply(this)
	return this
}
