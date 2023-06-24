package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	g "zhihu/global"
	"zhihu/model"
)

// AddQuestion 向数据库中添加问题
func AddQuestion(username string, content string) error {
	// 准备插入语句
	insertQuery := "INSERT INTO questions (content, username, create_time, update_time) VALUES ( ?, ?, ?, ?)"

	// 获取当前时间
	now := time.Now()

	// 执行插入操作
	_, err = g.MysqlDB.Exec(insertQuery, content, username, now, now)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Question added successfully")
	return nil
}

// GetUserQuestions 获取指定用户名下的所有问题
func GetUserQuestions(username string) ([]model.Question, error) {
	// 准备查询语句
	selectQuery := "SELECT * FROM questions WHERE username = ?"

	// 执行查询操作
	rows, err := g.MysqlDB.Query(selectQuery, username)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	// 遍历查询结果并构建Question切片
	questions := make([]model.Question, 0) // 使用空的切片
	for rows.Next() {
		var question model.Question
		err := rows.Scan(
			&question.ID,
			&question.Content,
			&question.Username,
			&question.CreatedAt,
			&question.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		questions = append(questions, question)
	}

	if len(questions) == 0 {
		fmt.Println("No questions found for the given username") // 没有找到问题记录
	} else {
		fmt.Println(questions) // 打印问题切片
	}

	log.Println("Questions got successfully")
	return questions, nil
}

// DeleteQuestion 删除指定用户名下的问题及相关的回答
func DeleteQuestion(username string, questionID int) error {
	// 开启事务
	tx, err := g.MysqlDB.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	// 删除问题
	deleteQuestionQuery := "DELETE FROM questions WHERE id = ? AND username = ?"
	_, err = tx.Exec(deleteQuestionQuery, questionID, username)
	if err != nil {
		log.Println(err)
		_ = tx.Rollback()
		return err
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Question and its answers deleted successfully")
	return nil
}

// UpdateQuestion 修改指定用户名下的问题内容
func UpdateQuestion(username string, questionID int, content string) error {
	// 准备更新语句
	updateQuery := "UPDATE questions SET content = ?, update_time = ? WHERE id = ? AND username = ?"

	// 获取当前时间
	now := time.Now()

	// 执行更新操作
	_, err = g.MysqlDB.Exec(updateQuery, content, now, questionID, username)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Question updated successfully")
	return nil
}
