package KbUserModel

import (
	"fmt"
	"knowledgeBase/src/models/JsonTime"
	"time"
)

type KbUserImpl struct {
	KbID     int       `gorm:"column:kb_id" json:"kb_id"`
	UserID   int       `gorm:"column:user_id" json:"-"`
	JoinTime time.Time `gorm:"column:join_time" json:"join_time"`
	CanEdit  string    `gorm:"column:can_edit" json:"can_edit"`
}

type GetKbsRequest struct {
	Size   int `json:"size"`
	Page   int `json:"page"`
	UserID int `json:"-"`
}

type KbUserResp struct {
	KbID      int               `gorm:"kb_id" json:"kb_id"`
	KbName    string            `gorm:"kb_name"  json:"kb_name"`
	JoinTime  JsonTime.JsonTime `gorm:"join_time"  json:"join_time"`
	CanEdit   string            `gorm:"can_edit"  json:"can_edit"`
	CreatorID int               `gorm:"can_edit"  json:"creator_id"`
	Kind      string            `gorm:"kind"  json:"kind"`
}

func NewGetKbsRequest() *GetKbsRequest {
	return &GetKbsRequest{}
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
