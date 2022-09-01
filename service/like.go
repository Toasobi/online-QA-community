package service

import (
	"net/http"
	"online-QA-community/define"
	"online-QA-community/helper"
	"online-QA-community/respository"

	"github.com/gin-gonic/gin"
)

// AnswerLike
// @Tags 用户私有方法
// @Summary 点赞功能
// @Param authorization header string true "authorization"
// @Param answer_identity formData string true "answer_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/answer-like [post]
func AnswerLike(ctx *gin.Context) {
	identity := ctx.PostForm("answer_identity")
	u, _ := ctx.Get("user")
	userClaim := u.(*helper.UserClaims)

	var count int64
	//已经点赞过了，再点即为取消点赞
	if err := respository.DB.Where("name = ? AND identity = ?", userClaim.Name, identity).Model(new(respository.IsLike)).Count(&count).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "database count err :" + err.Error(),
		})
		return
	}

	if count != 0 {
		err := respository.AnswerDisLike(identity, userClaim.Name)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": define.FAIL,
				"msg":  "操作失败",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"data": "取消点赞",
		})
		return
	}

	//点赞
	err := respository.AnswerLike(identity, userClaim.Name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "操作失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": define.FAIL,
		"data": "点赞成功",
	})

}
