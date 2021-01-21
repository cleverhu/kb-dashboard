package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"gorm.io/gorm"
	"knowledgeBase/src/middlewares"
	"knowledgeBase/src/models/KbUserModel"
	"knowledgeBase/src/services"
	"strconv"
)

type KbUserController struct {
	db            *gorm.DB                `inject:"-"`
	KbUserService *services.KbUserService `inject:"-"`
}

func NewKbUserController() *KbUserController {
	return &KbUserController{}
}

func (this *KbUserController) Name() string {
	return "KbUserController"
}

func (this *KbUserController) KbsByUserID(ctx *gin.Context) goft.Json {

	k, _ := ctx.Get("_req")
	kbs := this.KbUserService.UserKbs(k.(*KbUserModel.GetKbsRequest))
	return gin.H{"result": kbs, "code": 10000}
}

func (this *KbUserController) KbDetailByID(ctx *gin.Context) goft.Json {
	kbID := ctx.Param("id")
	id, err := strconv.Atoi(kbID)
	goft.Error(err, "知识库ID错误")
	kbDetail := this.KbUserService.KbByID(id)
	return gin.H{"result": kbDetail, "code": 10001}
}

func (this *KbUserController) Build(goft *goft.Goft) {
	goft.HandleWithFairing("POST", "/kns", this.KbsByUserID, middlewares.NewKbUserParamsCheck()).
		HandleWithFairing("GET", "/kns/:id", this.KbDetailByID)

}
