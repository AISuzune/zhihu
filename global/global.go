package global

import (
	"go.uber.org/zap"
	"zhihu/model/config"
)

var (
	Config *config.Config

	Logger *zap.Logger
)
