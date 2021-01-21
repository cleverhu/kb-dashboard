package daos

import (
	"gorm.io/gorm"
	"knowledgeBase/src/models/KbModel"
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
	this.DB.Raw("select  kb_users.kb_id,kb_users.join_time,kb_users.can_edit,kbs.kb_name as kb_name from kb_users join kbs on kb_users.kb_id = kbs.kb_id where kb_users.user_id = ? limit ? offset ? ", r.UserID, r.Size, r.Size*(r.Page-1)).Find(&kb)
	return kb
}

func (this *KbUserDAO) GetKbDetail(kbID int) *KbModel.KbDetail {
	kd := new(KbModel.KbDetail)
	kbInfo := &KbModel.KbImpl{}
	this.DB.Table("kbs").Where("kb_id = ?", kbID).Find(&kbInfo)
	kd.KbInfo = kbInfo

	var ids []*KbModel.UserID
	this.DB.Table("kb_users").Raw("select user_id from kb_users where kb_id = ? and can_edit = 'Y'", kbID).Find(&ids)
	kd.UserID = ids

	return kd
}
