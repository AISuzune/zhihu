package api

import (
	"github.com/gin-gonic/gin"
	"zhihu/api/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/api/register", register) // 注册
	r.POST("/api/login", login)       // 登录

	UserRouter := r.Group("/api/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	QuestionRouter := r.Group("/api")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		QuestionRouter.POST("/question", addQuestion)
		QuestionRouter.GET("/question", getQuestions)
		QuestionRouter.DELETE("/question", deleteQuestion)
		QuestionRouter.PUT("/question", updateQuestion)
	}

	AnswerRouter := r.Group("/api")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		AnswerRouter.POST("/answer", addAnswer)
		AnswerRouter.GET("/answer", getAnswers)
		AnswerRouter.DELETE("/answer", deleteAnswer)
		AnswerRouter.PUT("/answer", updateAnswer)
	}

	err := r.Run(":8088")
	if err != nil {
		return
	} // 跑在 8088 端口上
}
