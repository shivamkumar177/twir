package tts

import (
	"context"
	"fmt"
	"strconv"

	"github.com/guregu/null"
	model "github.com/satont/twir/libs/gomodels"

	"github.com/samber/lo"
	"github.com/satont/twir/apps/parser/internal/types"
	"go.uber.org/zap"
)

var PitchCommand = &types.DefaultCommand{
	ChannelsCommands: &model.ChannelsCommands{
		Name:        "tts pitch",
		Description: null.StringFrom("Change tts pitch"),
		Module:      "TTS",
		IsReply:     true,
	},
	Handler: func(ctx context.Context, parseCtx *types.ParseContext) *types.CommandsHandlerResult {
		result := &types.CommandsHandlerResult{}
		channelSettings, channelModele := getSettings(ctx, parseCtx.Services.Gorm, parseCtx.Channel.ID, "")

		if channelSettings == nil {
			result.Result = append(result.Result, "TTS is not configured.")
			return result
		}

		userSettings, currentUserModel := getSettings(
			ctx,
			parseCtx.Services.Gorm,
			parseCtx.Channel.ID,
			parseCtx.Sender.ID,
		)

		if parseCtx.Text == nil {
			result.Result = append(
				result.Result,
				fmt.Sprintf(
					"Global pitch: %v | Your pitch: %v",
					channelSettings.Pitch,
					lo.IfF(
						userSettings != nil, func() int {
							return userSettings.Pitch
						},
					).Else(channelSettings.Pitch),
				),
			)
			return result
		}

		pitch, err := strconv.Atoi(*parseCtx.Text)
		if err != nil {
			result.Result = append(result.Result, "Pitch must be a number")
			return result
		}

		if pitch < 0 || pitch > 100 {
			result.Result = append(result.Result, "Pitch must be between 0 and 100")
			return result
		}

		if parseCtx.Channel.ID == parseCtx.Sender.ID {
			channelSettings.Pitch = pitch
			err := updateSettings(ctx, parseCtx.Services.Gorm, channelModele, channelSettings)
			if err != nil {
				zap.S().Error(err)
				result.Result = append(result.Result, "Error while updating settings")
				return result
			}
		} else {
			if userSettings == nil {
				_, _, err := createUserSettings(
					ctx,
					parseCtx.Services.Gorm,
					pitch,
					50,
					channelSettings.Voice,
					parseCtx.Channel.ID,
					parseCtx.Sender.ID,
				)
				if err != nil {
					zap.S().Error(err)
					result.Result = append(result.Result, "Error while creating settings")
					return result
				}
			} else {
				userSettings.Pitch = pitch
				err := updateSettings(ctx, parseCtx.Services.Gorm, currentUserModel, userSettings)
				if err != nil {
					zap.S().Error(err)
					result.Result = append(result.Result, "Error while updating settings")
					return result
				}
			}
		}

		result.Result = append(result.Result, fmt.Sprintf("Pitch changed to %v", pitch))

		return result
	},
}
