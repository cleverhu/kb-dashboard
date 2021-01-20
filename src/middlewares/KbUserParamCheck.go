package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"knowledgeBase/src/models/KbUserModel"
)

type KbUserParamsCheck struct {
}

func NewKbUserParamsCheck() *KbUserParamsCheck {
	return &KbUserParamsCheck{}
}

func (this *KbUserParamsCheck) OnRequest(ctx *gin.Context) error {
	k := KbUserModel.NewGetKbsRequest()
	goft.Error(ctx.ShouldBindJSON(k))

	k.UserID = ctx.GetInt("_userid")
	ctx.Set("_req", k)

	return nil
}

func (this *KbUserParamsCheck) OnResponse(result interface{}) (interface{}, error) {
	return result, nil
}
