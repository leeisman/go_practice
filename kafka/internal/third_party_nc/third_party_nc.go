package third_party_nc

import (
	"context"
	"nc/internal/third_party_nc/jys_kafka_nc"
	"nc/internal/third_party_nc/silkrode_nc"
	"nc/pkg/model"
)

type ThirdPartyNC interface {
	Push(context.Context, []*model.User, *model.Notification)
}

func NewThirdPartyNC(silkrodeNC *silkrode_nc.SilkrodeNC, jysNC *jys_kafka_nc.JysNC) []ThirdPartyNC {
	return []ThirdPartyNC{
		silkrodeNC,
		jysNC,
	}
}
