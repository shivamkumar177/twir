package permit

import (
	"fmt"
	"github.com/samber/do"
	config "github.com/satont/tsuwari/libs/config"
	"github.com/satont/tsuwari/libs/grpc/generated/tokens"
	"github.com/satont/tsuwari/libs/twitch"
	"go.uber.org/zap"

	"github.com/satont/tsuwari/apps/parser/internal/di"
	"strconv"
	"strings"

	"github.com/satont/tsuwari/apps/parser/internal/types"
	variables_cache "github.com/satont/tsuwari/apps/parser/internal/variablescache"

	model "github.com/satont/tsuwari/libs/gomodels"

	"github.com/samber/lo"
	"github.com/satont/go-helix/v2"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var Command = types.DefaultCommand{
	Command: types.Command{
		Name:        "permit",
		Description: lo.ToPtr("Permits user."),
		Permission:  "MODERATOR",
		Visible:     false,
		Module:      lo.ToPtr("MODERATION"),
		IsReply:     true,
	},
	Handler: func(ctx variables_cache.ExecutionContext) *types.CommandsHandlerResult {
		db := do.MustInvoke[gorm.DB](di.Provider)
		cfg := do.MustInvoke[config.Config](di.Provider)
		tokensGrpc := do.MustInvoke[tokens.TokensClient](di.Provider)

		twitchClient, err := twitch.NewAppClient(cfg, tokensGrpc)

		result := &types.CommandsHandlerResult{}

		count := 1
		params := strings.Split(*ctx.Text, " ")

		paramsLen := len(params)
		if paramsLen < 1 {
			result.Result = []string{"you have type user name to permit."}
			return result
		}

		if paramsLen == 2 {
			newCount, err := strconv.Atoi(params[1])
			if err == nil {
				count = newCount
			}
		}

		if count > 100 {
			result.Result = []string{"cannot create more then 100 permits."}
			return result
		}

		target, err := twitchClient.GetUsers(&helix.UsersParams{
			Logins: []string{params[0]},
		})

		if err != nil || target.StatusCode != 200 || len(target.Data.Users) == 0 {
			result.Result = []string{"user not found."}
			return result
		}

		db.Transaction(func(tx *gorm.DB) error {
			for i := 0; i < count; i++ {
				permit := model.ChannelsPermits{
					ID:        uuid.NewV4().String(),
					ChannelID: ctx.ChannelId,
					UserID:    target.Data.Users[0].ID,
				}
				err := tx.Create(&permit).Error
				if err != nil {
					zap.S().Error(err)
					return err
				}
			}
			return nil
		})

		result.Result = []string{
			fmt.Sprintf("✅ added %v permits to %s", count, target.Data.Users[0].DisplayName),
		}
		return result
	},
}
