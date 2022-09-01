package service

import (
	"log"
	"net/http"
	"online-QA-community/define"
	"online-QA-community/helper"
	"online-QA-community/respository"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Comment
// @Tags 用户私有方法
// @Summary 发个评论
// @Param authorization header string true "authorization"
// @Param content formData string false "content"
// @Param question_identity formData string true "question_identity"
// @Param answer_identity formData string true "answer_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/comment [post]
func Comment(ctx *gin.Context) {
	content := ctx.PostForm("content")
	question_identity := ctx.PostForm("question_identity")
	answer_identity := ctx.PostForm("answer_identity")

	if content == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "评论内容不能为空",
		})
		return
	}

	u, _ := ctx.Get("user")
	userClaim := u.(*helper.UserClaims)

	comment, err := respository.CreateComment(content, question_identity, answer_identity, userClaim.Name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "CreateAnswer err : " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": define.SUCCESS,
		"data": map[string]interface{}{
			"comment": comment,
		},
	})

}

// GetCommentList
// @Tags 用户私有方法
// @Summary 查看本人所有评论
// @Param authorization header string true "authorization"
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/comment-list [get]
func GetCommentList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", define.DefaultPage)) //返回的结果是string，要进行转换

	size, _ := strconv.Atoi(ctx.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Fatal("stronv err : ", err)
	}
	//注意在这里page为1的时候其实对应到数据库上是0，所以还需要处理一下page
	page = (page - 1) * size

	var list []respository.Comment
	var count int64

	u, _ := ctx.Get("user")
	userClaim := u.(*helper.UserClaims)

	tx := respository.GetCommentList(userClaim.Name)

	if err = tx.Count(&count).Offset(page).Limit(size).Find(&list).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "数据库寄了 " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})

}
