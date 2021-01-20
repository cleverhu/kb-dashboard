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

func (this *KbUserService) UserKbs(r *KbUserModel.GetKbsRequest) []*KbUserModel.KbUserImpl {
	//fmt.Println(r)
	return this.KbUserDao.FindKbsByUserID(r)
}



