package services

import (
	"knowledgeBase/src/daos"
	"knowledgeBase/src/models/KbUserModel"
)

type KbUserService struct {
	KbUserDao *daos.KbUserDAO `inject:"-"`
}

func NewKbUserService() *KbUserService {
	return &KbUserService{}
}

func (this *KbUserService) UserKbs(r *KbUserModel.GetKbsRequest) []*KbUserModel.KbUserResp {
	return this.KbUserDao.FindKbsByUserID(r)
}

func (this *KbUserService) KbByID(kbID int)  {
	this.KbUserDao.GetKbDetail(kbID)
}
