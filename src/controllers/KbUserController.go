package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"gorm.io/gorm"
	"knowledgeBase/src/middlewares"
	"knowledgeBase/src/models/DocGrpModel"
	"knowledgeBase/src/models/KbModel"
	"knowledgeBase/src/models/KbUserModel"
	"knowledgeBase/src/services"
	"strconv"
)

type KbUserController struct {
	Db            *gorm.DB                `inject:"-"`
	KbUserService *services.KbUserService `inject:"-"`
}

func NewKbUserController() *KbUserController {
	return &KbUserController{}
}

func (this *KbUserController) Name() string {
	return "KbUserController"
}

//分页获取用户知识库
func (this *KbUserController) KbsByUserID(ctx *gin.Context) goft.Json {
	k, _ := ctx.Get("_req")
	kbs := this.KbUserService.UserKbs(k.(*KbUserModel.GetKbsRequest))
	return gin.H{"result": kbs, "code": 10000}
}

//获取用户单个知识库信息  未完善
func (this *KbUserController) KbDetailByID(ctx *gin.Context) goft.Json {
	kbID := ctx.Param("id")
	id, err := strconv.Atoi(kbID)
	goft.Error(err, "知识库ID错误")
	kbDetail := this.KbUserService.KbByID(id)
	return gin.H{"result": kbDetail, "code": 10001}
}

//增加用户知识库
func (this *KbUserController) PutKb(ctx *gin.Context) goft.Json {
	req := &KbModel.KbInputRequest{}
	goft.Error(ctx.ShouldBindJSON(req), "输入不合法")
	req.CreatorID = ctx.GetInt("_userid")
	result := this.KbUserService.PutKb(req)
	return gin.H{"result": result, "code": 10002}
}

//通过知识库ID获取分组信息
func (this *KbUserController) GetGroupByID(ctx *gin.Context) goft.Json {
	kbID := ctx.Param("kb_id")
	id, err := strconv.Atoi(kbID)
	goft.Error(err, "知识库ID错误")
	userID := ctx.GetInt("_userid")
	result := this.KbUserService.GroupDetailByID(id, userID)
	return gin.H{"result": result, "code": 10002}
}

//删除分组
func (this *KbUserController) DeleteGroupByID(ctx *gin.Context) goft.Json {
	groupID := ctx.Param("group_id")
	id, err := strconv.Atoi(groupID)
	goft.Error(err, "分组ID错误")
	userID := ctx.GetInt("_userid")
	result := this.KbUserService.DeleteGroupByID(id, userID)
	return gin.H{"result": result, "code": 10002}
}

//更新分组信息
func (this *KbUserController) UpdateGroup(ctx *gin.Context) goft.Json {
	req := DocGrpModel.NewDocGroupInsertRequest()
	goft.Error(ctx.ShouldBindJSON(req), "提交数据错误")
	userID := ctx.GetInt("_userid")
	result := this.KbUserService.UpdateGroupByID(req, userID)
	return gin.H{"result": result, "code": 10003}
}

//插入子分组
func (this *KbUserController) InsertGroup(ctx *gin.Context) goft.Json {
	req := DocGrpModel.NewDocGroupInsertRequest()
	goft.Error(ctx.ShouldBindJSON(req), "提交数据错误")
	userID := ctx.GetInt("_userid")
	result := this.KbUserService.InsertGroup(req, userID)
	return gin.H{"result": result, "code": 10004}
}

func (this *KbUserController) Build(goft *goft.Goft) {
	//获取知识库列表
	goft.HandleWithFairing("POST", "/kns", this.KbsByUserID, middlewares.NewKbUserParamsCheck()).
		//增加知识库
		Handle("PUT", "/kns", this.PutKb).
		//获取知识库信息
		Handle("GET", "/kns/:id", this.KbDetailByID).
		//查看分组
		Handle("GET", "/group/:kb_id", this.GetGroupByID).
		//删除分组
		Handle("DELETE", "/group/:group_id", this.DeleteGroupByID).
		//修改
		Handle("POST", "/group", this.UpdateGroup).
		//增加
		Handle("PUT", "/group", this.InsertGroup)

}
