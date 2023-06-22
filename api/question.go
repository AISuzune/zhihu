package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zhihu/dao"
	"zhihu/utils"
)

func addQuestion(c *gin.Context) {
	content := c.PostForm("content")
	username := c.PostForm("username")

	err := dao.AddQuestion(username, content)
	if err != nil {
		utils.RespFail(c, "Failed to add question")
		return
	}

	utils.RespSuccess(c, "Question added successfully")
}

func getQuestions(c *gin.Context) {
	username := c.PostForm("username")
	_, err := dao.GetUserQuestions(username)
	if err != nil {
		utils.RespFail(c, "Failed to get questions")
		return
	}

	utils.RespSuccess(c, "Questions got successfully")
}

func deleteQuestion(c *gin.Context) {
	username := c.PostForm("username")
	questionID := c.PostForm("question_id")

	// 将字符串类型的questionID转换为整数类型
	questionIDInt, err := strconv.Atoi(questionID)
	if err != nil {
		utils.RespFail(c, "Invalid question ID")
		return
	}

	err = dao.DeleteQuestion(username, questionIDInt)
	if err != nil {
		utils.RespFail(c, "Failed to delete question")
		return
	}

	utils.RespSuccess(c, "Question deleted successfully")
}

func updateQuestion(c *gin.Context) {
	username := c.PostForm("username")
	questionID := c.PostForm("question_id")
	content := c.PostForm("content")

	// 将字符串类型的questionID转换为整数类型
	questionIDInt, err := strconv.Atoi(questionID)
	if err != nil {
		utils.RespFail(c, "Invalid question ID")
		return
	}

	err = dao.UpdateQuestion(username, questionIDInt, content)
	if err != nil {
		utils.RespFail(c, "Failed to update question")
		return
	}

	utils.RespSuccess(c, "Question updated successfully")
}
