package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"strconv"
)

type UserIDCheck struct {
}

func NewKbUserIDCheck() *UserIDCheck {
	return &UserIDCheck{}
}

func (this *UserIDCheck) OnRequest(ctx *gin.Context) error {
	userID := ctx.GetHeader("X-User")
	id, err := strconv.Atoi(userID)
	if err != nil {
		goft.Error(err, "用户ID错误")
	}
	ctx.Set("_userid", id)

	return nil
}

func (this *UserIDCheck) OnResponse(result interface{}) (interface{}, error) {
	return result, nil
}

