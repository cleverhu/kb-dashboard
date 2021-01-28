package services

import (
	"knowledgeBase/src/daos"
	"knowledgeBase/src/models/DocGrpModel"
	"knowledgeBase/src/models/KbModel"
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

func (this *KbUserService) KbByID(kbID int) *KbModel.KbDetail {
	return this.KbUserDao.GetKbDetail(kbID)
}

func (this *KbUserService) GroupDetailByID(kbID, userID int) []*DocGrpModel.DocGrpResponseImpl{
	return this.KbUserDao.GroupDetailByID(kbID, userID)
}

func (this *KbUserService) DeleteGroupByID(groupID, userID int) string {
	return this.KbUserDao.DeleteGroupByID(groupID, userID)
}

func (this *KbUserService) UpdateGroupByID(request *DocGrpModel.DocGroupInsertRequest, userID int) string{
	return this.KbUserDao.UpdateGroupByID(request, userID)
}

func (this *KbUserService) InsertGroup(request *DocGrpModel.DocGroupInsertRequest, userID int) string{
	return this.KbUserDao.InsertGroupByID(request, userID)
}
