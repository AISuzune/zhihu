package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"zhihu/dao"
	"zhihu/utils"
)

func addAnswer(c *gin.Context) {
	content := c.PostForm("content")
	username := c.PostForm("username")
	questionID := c.PostForm("question_id")

	// 将字符串类型的questionID转换为整数类型
	questionIDInt, err := strconv.Atoi(questionID)
	if err != nil {
		utils.RespFail(c, "Invalid question ID")
		return
	}

	err = dao.AddAnswer(username, questionIDInt, content)
	if err != nil {
		utils.RespFail(c, "Failed to add answer")
		return
	}

	utils.RespSuccess(c, "Answer added successfully")
}

func getAnswers(c *gin.Context) {
	username := c.PostForm("username")

	_, err := dao.GetUserAnswers(username)
	if err != nil {
		utils.RespFail(c, "Failed to get answers")
		return
	}

	utils.RespSuccess(c, "Answers got successfully")
}

func deleteAnswer(c *gin.Context) {
	username := c.PostForm("username")
	answerID := c.PostForm("answer_id")

	// 将字符串类型的answerID转换为整数类型
	answerIDInt, err := strconv.Atoi(answerID)
	if err != nil {
		utils.RespFail(c, "Invalid answer ID")
		return
	}

	err = dao.DeleteAnswer(username, answerIDInt)
	if err != nil {
		utils.RespFail(c, "Failed to delete answer")
		return
	}

	utils.RespSuccess(c, "Answer deleted successfully")
}

func updateAnswer(c *gin.Context) {
	username := c.PostForm("username")
	answerID := c.PostForm("answer_id")
	content := c.PostForm("content")

	// 将字符串类型的answerID转换为整数类型
	answerIDInt, err := strconv.Atoi(answerID)
	if err != nil {
		utils.RespFail(c, "Invalid answer ID")
		return
	}

	err = dao.UpdateAnswer(username, answerIDInt, content)
	if err != nil {
		utils.RespFail(c, "Failed to update answer")
		return
	}

	utils.RespSuccess(c, "Answer updated successfully")
}
