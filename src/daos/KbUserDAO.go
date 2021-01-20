package daos

import (
	"fmt"
	"gorm.io/gorm"
	"knowledgeBase/src/models/KbUserModel"
)

type KbUserDAO struct {
	DB *gorm.DB `inject:"-"`
}

func NewKbUserDao() *KbUserDAO {
	return &KbUserDAO{}
}

func (this *KbUserDAO) FindKbsByUserID(r *KbUserModel.GetKbsRequest) []*KbUserModel.KbUserImpl {
	var kbus []*KbUserModel.KbUserImpl
	fmt.Println(kbus)
	this.DB.Table("kb_users").Where("user_id = ?", r.UserID).Limit(r.Size).Offset(r.Size * (r.Page - 1)).Find(&kbus)
	return kbus
}

func (this *KbUserDAO) GetKbDetail(kbID int) {

}
