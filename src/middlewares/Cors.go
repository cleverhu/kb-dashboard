package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Cors struct {

}

func NewCors() *Cors {
	return &Cors{}
}


func (this *Cors) OnRequest(ctx *gin.Context) error {
	method := ctx.Request.Method
	if method != "" {
		ctx.Header("Access-Control-Allow-Origin", "*")  // 可将将 * 替换为指定的域名
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, X-User")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
	}
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}
	return nil
}

func (this *Cors) OnResponse(result interface{}) (interface{}, error) {
	return result, nil
}