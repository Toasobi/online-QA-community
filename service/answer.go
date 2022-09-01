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

// QusetionAnswer
// @Tags 用户私有方法
// @Summary 回答提问
// @Param authorization header string true "authorization"
// @Param content formData string false "content"
// @Param question_identity formData string true "question_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/answer [post]
func QusetionAnswer(ctx *gin.Context) {
	content := ctx.PostForm("content")
	QaId := ctx.PostForm("question_identity")

	if content == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "回答内容不能为空",
		})
		return
	}

	u, _ := ctx.Get("user")
	userClaim := u.(*helper.UserClaims)

	answer, err := respository.CreateAnswer(QaId, content, userClaim.Name)
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
			"answer": answer,
		},
	})

}

// GetAnswerList
// @Tags 用户私有方法
// @Summary 查看本人所有回答
// @Param authorization header string true "authorization"
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/answer-list [get]
func GetAnswerList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", define.DefaultPage)) //返回的结果是string，要进行转换

	size, _ := strconv.Atoi(ctx.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Fatal("stronv err : ", err)
	}
	//注意在这里page为1的时候其实对应到数据库上是0，所以还需要处理一下page
	page = (page - 1) * size

	var list []respository.Answer
	var count int64

	u, _ := ctx.Get("user")
	userClaim := u.(*helper.UserClaims)

	tx := respository.GetAnswerList(userClaim.Name)

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

// AnswerDelete
// @Tags 用户私有方法
// @Summary 删除回答
// @Param authorization header string true "authorization"
// @Param answer_identity query string true "answer_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/answer-delete [delete]
func AnswerDelete(ctx *gin.Context) {
	identity := ctx.Query("answer_identity")
	if err := respository.DB.Where("identity = ?", identity).Delete(new(respository.Answer)).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "回答删除失败 :" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": define.SUCCESS,
		"data": "删除成功！",
	})

}

// AnswerUpdate
// @Tags 用户私有方法
// @Summary 修改回答
// @Param authorization header string true "authorization"
// @Param content formData string false "content"
// @Param answer_identity formData string true "answer_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/answer-update [put]
func AnswerUpdate(ctx *gin.Context) {
	identity := ctx.PostForm("answer_identity")
	content := ctx.PostForm("content")

	if content == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "回答内容不能为空",
		})
		return
	}

	if err := respository.DB.Model(new(respository.Answer)).Where("identity = ?", identity).Update("content", content).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "回答修改失败 :" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": define.SUCCESS,
		"data": "修改成功！",
	})
}
