package mysql_utils

import "time"

type PlanType string

const (
	ByDay       PlanType = "by_day"
	ByDataUsage PlanType = "by_usage"
)

type TopUpRecord struct {
	ID        int64     `bson:"-" json:"-" gorm:"column:id;type:bigserial;PRIMARY_KEY;AUTO_INCREMENT:true;NOT NULL"`
	UserID    int64     `bson:"user_id" json:"userID" gorm:"column:user_id;type:bigint;NOT NULL"`
	OrderID   int64     `bson:"order_id" json:"orderID" gorm:"column:order_id;type:bigint;UNIQUE_INDEX:idx_order_id_plan_type;NOT NULL"`
	PublishID string    `bson:"publish_id" json:"publishID" gorm:"column:publish_id;type:varchar(50);NOT NULL"`
	PlanType  PlanType  `bson:"plan_type" json:"planType" gorm:"column:plan_type;type:varchar(15);UNIQUE_INDEX:idx_order_id_plan_type;NOT NULL"`
	DataUsage int64     `bson:"data_usage,omitempty" json:"dataUsage,omitempty" gorm:"column:data_usage;type:bigint"`
	ExpiredAt time.Time `bson:"expired_at,omitempty" json:"expiredAt,omitempty" gorm:"column:expired_at;type:timestamptz"`
	CreatedAt time.Time `bson:"created_at" json:"createdAt" gorm:"column:created_at;type:timestamptz;NOT NULL;DEFAULT:now()"`
}

