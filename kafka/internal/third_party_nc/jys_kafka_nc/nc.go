package jys_kafka_nc

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"gitlab.silkrode.com.tw/golang/gopher/db/v2/db"
	"gitlab.silkrode.com.tw/golang/gopher/delivery/kafka"
	nc2 "gitlab.silkrode.com.tw/team_golang/jys/infra/pb/nc"
	"gitlab.silkrode.com.tw/team_golang/jys/infra/topic"
	"gorm.io/gorm"
	"nc/internal/converter"
	"nc/pkg/model"
)

// 串接 交易所 web推播系統
type JysNC struct {
	Config        *kafka.Config
	Write         *gorm.DB
	KafkaProducer kafka.Producer
}

func NewJysNC(config *kafka.Config, conn *db.Connections, kafkaProducer kafka.Producer) *JysNC {
	return &JysNC{
		Config:        config,
		Write:         conn.WriteDB,
		KafkaProducer: kafkaProducer,
	}
}

// kafka producer
func (nc *JysNC) Push(ctx context.Context, users []*model.User, notification *model.Notification) {
	req := nc2.PushNotification{}
	jysNcReq := &model.JysNCReq{
		Users:          users,
		PublishedAt:    &notification.PublishedAt,
		Type:           notification.Type,
		Content:        notification.Content,
		Title:          notification.Title,
		NotificationID: notification.ID,
	}
	err := converter.JSON(jysNcReq, &req)
	if err != nil {
		nc.recordError(jysNcReq, err)
		return
	}
	err = nc.KafkaProducer.Pub(ctx, topic.TopicPushNotification, req)
	if err != nil {
		nc.recordError(jysNcReq, err)
		return
	}
}

// 錯誤紀錄
func (nc *JysNC) recordError(req *model.JysNCReq, err error) {
	pushNcError := model.ThirdPartyNCError{
		Type:     model.SilkrodeNC,
		Req:      fmt.Sprintf("%s", req),
		ErrorMsg: err.Error(),
	}
	err = nc.Write.Create(&pushNcError).Error
	if err != nil {
		fmt.Println("")
		log.Error().Msgf("recordError err: %s", err.Error())
		fmt.Println("")
	}
}
