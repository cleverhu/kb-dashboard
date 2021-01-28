package configuration

import (
	"gorm.io/gorm"
	"knowledgeBaseNuxt/src/common"
)

type DBConfig struct {
}

func NewDBConfig() *DBConfig {
	return &DBConfig{}
}

func (this *DBConfig) GormDB() *gorm.DB {
	return common.Orm
}
