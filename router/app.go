package router

import (
	"online-QA-community/middlewares"
	"online-QA-community/service"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "online-QA-community/docs"
)

func Router() *gin.Engine {
	r := gin.Default()

	//配置一下swaggo

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//公共方法
	//用户登录
	r.POST("/user-login", service.UserLogin)
	//用户注册
	r.POST("/user-register", service.UserRegister)
	//发送验证码
	r.POST("send-code", service.SendCode)
	//显示所有问题
	// r.GET("/question-list", service.GetQusetionList)

	//用户私有方法
	users := r.Group("/user", middlewares.AuthUserCheck())
	{
		users.POST("/submit", service.QuestionSubmit)
		users.POST("/answer", service.QusetionAnswer)
		users.POST("/comment", service.Comment)
		users.GET("/problem-list", service.GetProblemList)
		users.GET("/answer-list", service.GetAnswerList)
		users.GET("/comment-list", service.GetCommentList)
		users.DELETE("/answer-delete", service.AnswerDelete)
		users.PUT("/answer-update", service.AnswerUpdate)
		users.DELETE("/problem-delete", service.ProblemDelete)
		users.PUT("/problem-update", service.ProblemUpdate)
		users.POST("/answer-like", service.AnswerLike)
		users.POST("/attention", service.UserAttention)
	}
	return r
}
