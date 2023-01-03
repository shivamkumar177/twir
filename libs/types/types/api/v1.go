package api

import (
	"github.com/satont/tsuwari/libs/types/types/api/bot"
	"github.com/satont/tsuwari/libs/types/types/api/modules"
)

type Modules struct {
	YouTube modules.YouTube
}

type Channels struct {
	MODULES Modules
	BOT     bot.Bot
}

type V1 struct {
	CHANNELS Channels
}