package global

import (
	"database/sql"
	"go.uber.org/zap"
	"zhihu/model/config"
)

var (
	Config *config.Config

	Logger *zap.Logger

	MysqlDB *sql.DB
)
