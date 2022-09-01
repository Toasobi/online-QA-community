package service

import (
	"net/http"
	"online-QA-community/define"
	"online-QA-community/helper"
	"online-QA-community/respository"

	"github.com/gin-gonic/gin"
)

// UserAttention
// @Tags 用户私有方法
// @Summary 关注功能
// @Param authorization header string true "authorization"
// @Param user_identity formData string true "user_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/attention [post]
func UserAttention(ctx *gin.Context) {
	identity := ctx.PostForm("user_identity")
	u, _ := ctx.Get("user")
	userClaim := u.(*helper.UserClaims)

	var count int64
	//已经关注过了，再点即为取消关注
	if err := respository.DB.Where("name = ? AND identity = ?", userClaim.Name, identity).Model(new(respository.IsAttention)).Count(&count).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "database count err :" + err.Error(),
		})
		return
	}

	if count != 0 {
		err := respository.DisAttention(identity, userClaim.Name)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": define.FAIL,
				"msg":  "操作失败",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"data": "取消关注",
		})
		return
	}

	//点赞
	err := respository.Attention(identity, userClaim.Name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "操作失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": define.FAIL,
		"data": "关注成功",
	})

}
