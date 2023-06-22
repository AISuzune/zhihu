package api

import (
	"github.com/gin-gonic/gin"
	"zhihu/api/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.POST("/api/user/register", register) // 注册
	r.POST("/api/user/login", login)       // 登录

	UserRouter := r.Group("/api/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.GET("/get", getUsernameFromToken)
	}

	r.POST("/question", addQuestion)
	r.POST("/answer", addAnswer)
	r.GET("/question", getQuestions)
	r.GET("/answer", getAnswers)
	r.DELETE("/question", deleteQuestion)
	r.DELETE("/answer", deleteAnswer)
	r.PUT("/question", updateQuestion)
	r.PUT("/answer", updateAnswer)

	err := r.Run(":8088")
	if err != nil {
		return
	} // 跑在 8088 端口上
}
