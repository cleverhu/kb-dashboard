package daos

import (
	"gorm.io/gorm"
	"knowledgeBase/src/models/KbUserModel"
)

type KbUserDAO struct {
	DB *gorm.DB `inject:"-"`
}

func NewKbUserDao() *KbUserDAO {
	return &KbUserDAO{}
}

func (this *KbUserDAO) FindKbsByUserID(r *KbUserModel.GetKbsRequest) []*KbUserModel.KbUserResp {
	var kb []*KbUserModel.KbUserResp
	this.DB.Raw("select  kb_users.kb_id,kb_users.join_time,kb_users.can_edit,kbs.kb_name as kb_name from kb_users join kbs on kb_users.id = kbs.kb_id where kb_users.user_id = ? limit ? offset ? ",r.UserID,r.Size,r.Size * (r.Page - 1)).Find(&kb)
	return kb
}

func (this *KbUserDAO) GetKbDetail(kbID int) {

}
