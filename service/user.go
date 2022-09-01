package service

import (
	"context"
	"net/http"
	"online-QA-community/define"
	"online-QA-community/helper"
	"online-QA-community/respository"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserLogin
// @Tags 公共方法
// @Summary 用户登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user-login [post]
func UserLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if username == "" || password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "用户名或密码不能为空",
		})
	}

	//开始校验
	password = helper.GetMd5(password)
	data := new(respository.UserBasic)

	if err := respository.DB.Where("name = ? AND password = ?", username, password).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusOK, gin.H{
				"code": define.FAIL,
				"msg":  "用户名或密码错误",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "Get User Error(database):" + err.Error(),
		})
		return
	}

	//签发token
	token, err := helper.GenerateToken(data.Identity, username, data.IsAdmin)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "token Generate err" + err.Error(),
		})
		return
	}

	//传给前端
	ctx.JSON(http.StatusOK, gin.H{
		"code": define.SUCCESS,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

// UserRegister
// @Tags 公共方法
// @Summary 用户注册
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Param phone formData string false "phone"
// @Param mail formData string true "mail"
// @Param is_admin formData string false "is_admin"
// @Param code formData string false "code"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user-register [post]
func UserRegister(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	mail := ctx.PostForm("mail")
	code := ctx.PostForm("code")
	isAdmin, _ := strconv.Atoi(ctx.PostForm("is_admin"))

	//检验一下信息
	if username == "" || password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "用户名或密码不能为空",
		})
	}

	if code == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "验证码不能为空",
		})
	}

	var count int64
	if err := respository.DB.Where("name = ? ", username).Model(new(respository.UserBasic)).Count(&count).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "database count err :" + err.Error(),
		})
		return
	}

	if count != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "用户已经存在",
		})
		return
	}
	var rctx = context.Background()
	v, err := respository.RDB.Get(rctx, mail).Result()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "redis err :" + err.Error(),
		})
		return
	}

	if v != code {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "验证码不正确",
		})
		return
	}

	identity := helper.GenerateUUID()
	password = helper.GetMd5(password)

	user, err := respository.CreateUser(username, identity, password, phone, mail, isAdmin)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "User Create Error :" + err.Error(),
		})
	}

	//给前端返回用户信息
	ctx.JSON(http.StatusOK, gin.H{
		"code": define.SUCCESS,
		"data": user,
	})

}

// SendCode
// @Tags 公共方法
// @Summary 发送验证码
// @Param mail formData string false "mail"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /send-code [post]
func SendCode(ctx *gin.Context) {
	mail := ctx.PostForm("mail")
	if mail == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "邮箱不能为空",
		})

		return
	}

	code := helper.GetCode()
	err := helper.SendCode(mail, code)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": define.FAIL,
			"msg":  "验证码发送失败",
		})
		return
	}
	//将验证码和邮箱绑定到redis数据库中
	var rctx = context.Background()
	respository.RDB.Set(rctx, mail, code, time.Minute*5)
	ctx.JSON(http.StatusOK, gin.H{
		"code": define.FAIL,
		"msg":  "验证码发送成功",
	})

}
