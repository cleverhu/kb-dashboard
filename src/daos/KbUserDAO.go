package daos

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"knowledgeBase/src/common"
	"knowledgeBase/src/grpcServices"
	"knowledgeBase/src/models/DocGrpModel"
	"knowledgeBase/src/models/KbModel"
	"knowledgeBase/src/models/KbUserModel"

	"strings"
	"time"
)

type KbUserDAO struct {
	DB                  *gorm.DB                         `inject:"-"`
	KbInfoServiceClient grpcServices.KbInfoServiceClient `inject:"-"`
}

func NewKbUserDao() *KbUserDAO {
	return &KbUserDAO{}
}

func (this *KbUserDAO) FindKbsByUserID(r *KbUserModel.GetKbsRequest) []*KbUserModel.KbUserResp {
	var kb []*KbUserModel.KbUserResp
	this.DB.Raw(`select  kb_users.kb_id,kb_users.join_time,kb_users.can_edit,kbs.kb_name as kb_name,kbs.creator_id as creator_id ,kb_kinds.kind as kind
from kb_users 
right join kbs on kb_users.kb_id = kbs.kb_id
right join kb_kinds on kbs.kb_kind_id = kb_kinds.id
where kb_users.user_id = ? limit ? offset ? `, r.UserID, r.Size, r.Size*(r.Page-1)).Find(&kb)
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

func (this *KbUserDAO) PutKb(req *KbModel.KbInputRequest) string {
	err := this.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Table("kbs").Create(&req).Error
		if err != nil {
			return err
		}
		err = tx.Table("kb_users").Exec(`insert into kb_users (kb_id,user_id,join_time,can_edit) values(?,?,?,?)`,
			req.ID, req.CreatorID, time.Now(), "Y").Error
		return err
	})
	if err != nil {
		return "添加失败"
	}
	return "添加成功"
}

func (this *KbUserDAO) GroupDetailByID(kbID, userID int) []*DocGrpModel.DocGrpResponseImpl {

	c := &struct {
		C int `gorm:"column:c"`
	}{}

	this.DB.Table("kb_users").Raw("select count(*) as c from kb_users where kb_id = ? and user_id = ?", kbID, userID).Find(&c)

	if c.C == 0 {
		return nil
	}

	kbName := &struct {
		Name string `gorm:"column:kb_name"`
	}{}

	this.DB.Table("kbs").Raw("select kb_name from kbs where kb_id = ?", kbID).Find(&kbName.Name)

	var dgm []*DocGrpModel.DocGrpResponseImpl

	this.groupDetailByID(kbName.Name, kbID, 0, &dgm)
	return dgm
}

func (this *KbUserDAO) groupDetailByID(kbName string, kbID, groupID int, result *[]*DocGrpModel.DocGrpResponseImpl) []*DocGrpModel.DocGrpResponseImpl {
	//先找到分组
	this.DB.Table("doc_grps").Raw(`select group_id,group_name,shorturl from doc_grps 
where kb_id = ? and pid = ? 
order by group_order`, kbID, groupID).Find(&result)

	for _, v := range *result {
		//寻找子分组 这个数据是临时的，不会返回真实数据
		var grp []*DocGrpModel.DocGrpResponseImpl
		this.groupDetailByID(kbName, kbID, v.GroupID, &grp)
		for _, item := range grp {
			v.Children = append(v.Children, item)
		}
	}
	return *result
}

func (this *KbUserDAO) DeleteGroupByID(groupID, userID int) string {

	kb := &struct {
		ID int64 `gorm:"column:kb_id"`
	}{}

	this.DB.Table("doc_grps").Raw("select kb_id from doc_grps where group_id = ?", groupID).Find(&kb)
	if kb.ID == 0 {
		return "分组不存在,删除失败"
	}
	c := &struct {
		C int `gorm:"column:c"`
	}{}

	this.DB.Table("kb_users").Raw("select count(*) as c from kb_users where user_id = ? and can_edit ='Y' and kb_id = ?", userID, kb.ID).Find(&c)

	if c.C == 0 {
		return "您无权限操作该知识库"
	}

	c.C = 0
	this.DB.Table("docs").Raw("select count(*) as c from docs where group_id=?", groupID).Find(&c)

	if c.C > 0 {
		return "分组中存在文章,请先删除再删除分组"
	}

	if this.DB.Table("doc_grps").Exec("delete  from doc_grps where group_id =?", groupID).Error != nil {
		return "删除失败"
	}
	go func() {
		res, _ := this.KbInfoServiceClient.UpdateKbDetailList(context.Background(), &grpcServices.KbInfoRequest{Id: []int64{kb.ID}})
		fmt.Println(kb.ID, "更新结果:", res.Result)
	}()

	return "删除成功"
}

func (this *KbUserDAO) UpdateGroupByID(req *DocGrpModel.DocGroupInsertRequest, userID int) string {
	fmt.Println(req)
	if strings.TrimSpace(req.Title) == "" {
		return "标题为空,修改失败"
	}

	kb := &struct {
		ID int64 `gorm:"column:kb_id"`
	}{}

	this.DB.Table("doc_grps").Raw("select kb_id from doc_grps where group_id = ?", req.GroupID).Find(&kb)
	if kb.ID == 0 {
		return "分组不存在,修改失败"
	}

	c := &struct {
		C int `gorm:"column:c"`
	}{}

	this.DB.Table("kb_users").Raw("select count(*) as c from kb_users where user_id = ? and can_edit ='Y' and kb_id = ?", userID, kb.ID).Find(&c)

	if c.C == 0 {
		return "您无权限操作该知识库"
	}

	if this.DB.Table("doc_grps").Exec("update  doc_grps set group_name = ? where group_id =?", req.Title, req.GroupID).Error != nil {
		return "修改失败"
	}
	go func() {
		fmt.Println(this.KbInfoServiceClient)
		res, err := this.KbInfoServiceClient.UpdateKbDetailList(context.Background(), &grpcServices.KbInfoRequest{Id: []int64{kb.ID}})
		if err == nil {
			fmt.Println(kb.ID, "更新结果:", res.Result)
		} else {
			fmt.Println(err)
		}

	}()

	return "修改成功"

}

func (this *KbUserDAO) InsertGroupByID(req *DocGrpModel.DocGroupInsertRequest, userID int) string {

	kb := &struct {
		ID int64 `gorm:"column:kb_id"`
	}{}
	if req.GroupID != 0 {
		if strings.TrimSpace(req.SonTitle) == "" {
			return "标题为空,添加失败"
		}

		this.DB.Table("doc_grps").Raw("select kb_id from doc_grps where group_id = ?", req.GroupID).Find(&kb)
		if kb.ID == 0 {
			return "父分组不存在,添加失败"
		}
	} else {
		kb.ID = req.KbID

		req.SonTitle = req.Title
	}

	c := &struct {
		C int `gorm:"column:c"`
	}{}
	this.DB.Table("kb_users").Raw("select count(*) as c from kb_users where user_id = ? and can_edit ='Y' and kb_id = ?", userID, kb.ID).Find(&c)

	if c.C == 0 {
		return "您无权限操作该知识库"
	}

	if this.DB.Table("doc_grps").Exec(`insert into 
doc_grps (group_name,kb_id,create_time,creator_id,shorturl,pid) 
values(?,?,?,?,?,?)`, req.SonTitle, kb.ID, time.Now(), userID, common.ShotURL(time.Now().UnixNano()), req.GroupID).Error != nil {
		return "添加失败"
	}

	go func() {
		res, _ := this.KbInfoServiceClient.UpdateKbDetailList(context.Background(), &grpcServices.KbInfoRequest{Id: []int64{kb.ID}})
		fmt.Println(kb.ID, "更新结果:", res.Result)
	}()

	return "添加成功"

}
