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

// AddAnswer 向数据库中添加回答
func AddAnswer(username string, questionID int, content string) error {
	// 准备插入语句
	insertQuery := "INSERT INTO answers (qid, content, username, create_time, update_time) VALUES (?, ?, ?, ?, ?)"

	// 获取当前时间
	now := time.Now()

	// 执行插入操作
	_, err = g.MysqlDB.Exec(insertQuery, questionID, content, username, now, now)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Answer added successfully")
	return nil
}

// GetUserAnswers 获取指定用户名下的所有回答
func GetUserAnswers(username string) ([]model.Answer, error) {
	// 准备查询语句
	selectQuery := "SELECT * FROM answers WHERE username = ?"

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

	// 遍历查询结果并构建Answer切片
	answers := make([]model.Answer, 0) // 使用空的切片
	for rows.Next() {
		var answer model.Answer
		err := rows.Scan(
			&answer.ID,
			&answer.QuestionID,
			&answer.Content,
			&answer.Username,
			&answer.CreatedAt,
			&answer.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		answers = append(answers, answer)
	}

	if len(answers) == 0 {
		fmt.Println("No answers found for the given username") // 没有找到回答记录
	} else {
		fmt.Println(answers) // 打印回答切片
	}

	log.Println("Answers got successfully")
	return answers, nil
}

// DeleteAnswer 删除指定用户名下的回答
func DeleteAnswer(username string, answerID int) error {
	// 准备删除语句
	deleteQuery := "DELETE FROM answers WHERE id = ? AND username = ?"

	// 执行删除操作
	_, err = g.MysqlDB.Exec(deleteQuery, answerID, username)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Answer deleted successfully")
	return nil
}

// UpdateAnswer 修改指定用户名下的回答内容
func UpdateAnswer(username string, answerID int, content string) error {
	// 准备更新语句
	updateQuery := "UPDATE answers SET content = ?, update_time = ? WHERE id = ? AND username = ?"

	// 获取当前时间
	now := time.Now()

	// 执行更新操作
	_, err = g.MysqlDB.Exec(updateQuery, content, now, answerID, username)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Answer updated successfully")
	return nil
}
