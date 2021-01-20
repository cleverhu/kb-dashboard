package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"gorm.io/gorm"
	"knowledgeBase/src/middlewares"
	"knowledgeBase/src/models/KbUserModel"
	"knowledgeBase/src/services"
)

type KbUserController struct {
	db            *gorm.DB                `inject:"-"`
	KbUserService *services.KbUserService `inject:"-"`
}

func NewKbController() *KbUserController {
	return &KbUserController{}
}

func (this *KbUserController) Name() string {
	return "KbUserController"
}

func (this *KbUserController) KbsByUserID(ctx *gin.Context) goft.Json {

	k, _ := ctx.Get("_req")
	//fmt.Println(k)
	kbs := this.KbUserService.UserKbs(k.(*KbUserModel.GetKbsRequest))
	return gin.H{"result": kbs, "code": 10000}
}

func (this *KbUserController) KbDetailByID(ctx *gin.Context) goft.Json {
	kbID := ctx.Param("id")
	fmt.Println(kbID)
	return gin.H{"result": nil, "code": 10001}
}

func (this *KbUserController) Build(goft *goft.Goft) {
	//页尺寸参数size ，页面参数page
	//用户名参数在头里 X-User="xxxoo"
	//实际运行，网关会把这个头传给 /kns 。  测试时 ，自己设置一个头 ，头的值是明文的用户名
	//网关肯定是用traefik+k8s。  暂时还没部署 。 先不管 。手工设置头
	//2、/kns/id   代表查看知识库详细。 依然需要判断用户名（头）
	goft.HandleWithFairing("POST", "/kns", this.KbsByUserID, middlewares.NewKbUserParamsCheck()).
		HandleWithFairing("GET", "/kns/:id", this.KbDetailByID)

}
