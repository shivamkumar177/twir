package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

type ChannelsIntegrationsData struct {
	Code     *string `json:"code,omitempty"`
	Name     *string `json:"name,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
	UserName *string `json:"username,omitempty"`
	Game     *string `json:"game,omitempty"`
	UserId   *string `json:"userId,omitempty"`
}

type ChannelsIntegrations struct {
	ID            string                    `gorm:"primaryKey;column:id;type:TEXT;default:uuid_generate_v4()" json:"id"`
	Enabled       bool                      `gorm:"column:enabled;type:BOOL;"        json:"enabled"`
	ChannelID     string                    `gorm:"column:channelId;type:TEXT;"      json:"channelId"`
	IntegrationID string                    `gorm:"column:integrationId;type:TEXT;"  json:"integrationId"`
	AccessToken   null.String               `gorm:"column:accessToken;type:TEXT;"    json:"accessToken"   swaggertype:"string"`
	RefreshToken  null.String               `gorm:"column:refreshToken;type:TEXT;"   json:"refreshToken"  swaggertype:"string"`
	ClientID      null.String               `gorm:"column:clientId;type:TEXT;"       json:"clientId"      swaggertype:"string"`
	ClientSecret  null.String               `gorm:"column:clientSecret;type:TEXT;"   json:"clientSecret"  swaggertype:"string"`
	APIKey        null.String               `gorm:"column:apiKey;type:TEXT;"         json:"apiKey"        swaggertype:"string"`
	Data          *ChannelsIntegrationsData `gorm:"column:data;type:JSONB;"          json:"data"`
	Integration   *Integrations             `gorm:"foreignKey:IntegrationID"         json:"-"`
}

func (c *ChannelsIntegrations) TableName() string {
	return "channels_integrations"
}

func (a ChannelsIntegrationsData) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *ChannelsIntegrationsData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
