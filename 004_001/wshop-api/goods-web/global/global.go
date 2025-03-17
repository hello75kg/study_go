package global

import (
	ut "github.com/go-playground/universal-translator"
	"wshop-api/goods-web/config"
	"wshop-api/goods-web/proto"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient
)
