package global

import (
	ut "github.com/go-playground/universal-translator"
	"wshop-api/userop-web/config"
	"wshop-api/userop-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient

	MessageClient proto.MessageClient
	AddressClient proto.AddressClient
	UserFavClient proto.UserFavClient
)
