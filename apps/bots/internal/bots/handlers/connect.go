package handlers

import (
	"sync"
	"time"

	model "github.com/satont/tsuwari/libs/gomodels"

	ratelimiting "github.com/aidenwallis/go-ratelimiting/local"
	"github.com/samber/lo"
	"github.com/satont/go-helix/v2"
)

func (c *Handlers) OnConnect() {
	c.logger.Sugar().
		Infow("Bot connected to twitch", "botId", c.BotClient.Model.ID, "botName", c.BotClient.TwitchUser.Login)
	c.BotClient.RateLimiters.Channels = make(map[string]ratelimiting.SlidingWindow)

	twitchUsers := []helix.User{}
	twitchUsersMU := sync.Mutex{}

	botChannels := []model.Channels{}
	c.db.
		Where(
			`"botId" = ? AND "isEnabled" = ? AND "isBanned" = ? AND "isTwitchBanned" = ?`,
			c.BotClient.Model.ID,
			true,
			false,
			false,
		).Find(&botChannels)

	channelsChunks := lo.Chunk(botChannels, 100)
	wg := sync.WaitGroup{}
	wg.Add(len(channelsChunks))

	for _, chunk := range channelsChunks {
		go func(chunk []model.Channels) {
			defer wg.Done()
			usersIds := lo.Map(chunk, func(item model.Channels, _ int) string {
				return item.ID
			})

			twitchUsersReq, err := c.BotClient.Api.Client.GetUsers(&helix.UsersParams{
				IDs: usersIds,
			})
			if err != nil {
				panic(err)
			}
			twitchUsersMU.Lock()
			twitchUsers = append(twitchUsers, twitchUsersReq.Data.Users...)
			twitchUsersMU.Unlock()
		}(chunk)
	}

	wg.Wait()

	for _, u := range twitchUsers {
		botModRequest, err := c.BotClient.Api.Client.GetChannelMods(&helix.GetChannelModsParams{
			BroadcasterID: u.ID,
			UserID:        c.BotClient.Model.ID,
		})

		var limiter ratelimiting.SlidingWindow
		if err != nil && len(botModRequest.Data.Mods) == 1 {
			l, _ := ratelimiting.NewSlidingWindow(20, 30*time.Second)
			limiter = l
		} else {
			l, _ := ratelimiting.NewSlidingWindow(1, 2*time.Second)
			limiter = l
		}

		c.BotClient.RateLimiters.Channels[u.Login] = limiter
		c.BotClient.Join(u.Login)
	}
}
