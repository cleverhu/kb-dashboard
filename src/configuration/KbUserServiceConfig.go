package configuration

import (
	"knowledgeBase/src/daos"
	"knowledgeBase/src/services"
)

type KbUserServiceConfig struct {
}

func NewKbUserServiceConfig() *KbUserServiceConfig {
	return &KbUserServiceConfig{}
}

func (this *KbUserServiceConfig) KbUserDAO() *daos.KbUserDAO {
	return daos.NewKbUserDao()
}

func (this *KbUserServiceConfig) KbUserService() *services.KbUserService {
	return services.NewKbUserService()
}



