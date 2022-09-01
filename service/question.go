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

// QuestionSubmit
// @Tags 用户私有方法
// @Summary 发起提问
// @Param authorization header string true "authorization"
// @Param title formData string false "title"
// @Param content formData string false "content"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/submit [post]
func QuestionSubmit(ctx *gin.Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")

	if title == "" || content == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "文章标题或内容为空",
		})
		return
	}
	u, _ := ctx.Get("user")
	userClaim := u.(*helper.UserClaims)
	// var Name = "sss"
	qa, err := respository.CreateQA(title, content, userClaim.Name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "database createQA err : " + err.Error(),
		})
		return
	}
	//视图返回
	ctx.JSON(http.StatusOK, gin.H{
		"code": define.SUCCESS,
		"data": map[string]interface{}{
			"list": qa,
		},
	})
}

// GetProblemList
// @Tags 用户私有方法
// @Summary 查看本人所发出的问题
// @Param authorization header string true "authorization"
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/problem-list [get]
func GetProblemList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", define.DefaultPage)) //返回的结果是string，要进行转换

	size, _ := strconv.Atoi(ctx.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Fatal("stronv err : ", err)
	}
	//注意在这里page为1的时候其实对应到数据库上是0，所以还需要处理一下page
	page = (page - 1) * size

	keyword := ctx.Query("keyword")

	var list []respository.QuestionBasic
	var count int64

	u, _ := ctx.Get("user")
	userClaim := u.(*helper.UserClaims)

	tx := respository.GetProblemList(keyword, userClaim.Name)

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

// ProblemDelete
// @Tags 用户私有方法
// @Summary 删除问题
// @Param authorization header string true "authorization"
// @Param question_identity query string true "question_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/problem-delete [delete]
func ProblemDelete(ctx *gin.Context) {
	identity := ctx.Query("question_identity")
	if err := respository.DB.Where("identity = ?", identity).Delete(new(respository.QuestionBasic)).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "问题删除失败 :" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": define.SUCCESS,
		"data": "删除成功！",
	})

}

// ProblemUpdate
// @Tags 用户私有方法
// @Summary 修改问题
// @Param authorization header string true "authorization"
// @Param content formData string false "content"
// @Param title formData string false "title"
// @Param problem_identity formData string true "problem_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/problem-update [put]
func ProblemUpdate(ctx *gin.Context) {
	identity := ctx.PostForm("problem_identity")
	content := ctx.PostForm("content")
	title := ctx.PostForm("title")

	if content == "" || title == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "提问标题或内容不能为空",
		})
		return
	}

	if err := respository.DB.Model(new(respository.QuestionBasic)).Where("identity = ?", identity).Updates(map[string]interface{}{"title": title, "content": content}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "提问修改失败 :" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": define.SUCCESS,
		"data": "修改成功！",
	})
}
