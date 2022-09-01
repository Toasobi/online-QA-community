package middlewares

import (
	"net/http"
	"online-QA-community/helper"

	"github.com/gin-gonic/gin"
)

//判断用户
func AuthUserCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//check user
		auth := ctx.GetHeader("Authorization")
		userClaims, err := helper.AnalyseToken(auth)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized",
			})
			return
		}
		if userClaims.IsAdmin != 0 && userClaims.IsAdmin != 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized",
			})
			return
		}
		ctx.Set("user", userClaims)
		ctx.Next()

	}

}
