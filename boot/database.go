package boot

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	g "zhihu/global"
)

func MysqlDBSetup() {
	//链接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/zhihu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	g.MysqlDB = db
}
