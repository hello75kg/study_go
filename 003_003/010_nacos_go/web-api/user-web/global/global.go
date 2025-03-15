package global

import (
	ut "github.com/go-playground/universal-translator"
	"wshop-api/user-web/config"
	"wshop-api/user-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	UserSrvClient proto.UserClient
)
