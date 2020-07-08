package model

import "time"

type Action int8

const (
	Create = Action(0)
	Read   = Action(1)
	Update = Action(2)
	Delete = Action(3)
)

// 所有模型实现此接口
type Model interface {
	TableName() string            //定义表名
	BeforeSave() error            //保存前进行一些处理
	Prepare()                     //预处理
	Validate(action Action) error //对增删改查进行验证
}

// 所有模型继承 就可以有这些字段了
type TimeModel struct {
	CreatedAt time.Time  `gorm:"column:created_at;type:DATETIME;" json:"created_at"` //[ 4] created_at                                     DATETIME             null: false  primary: false  auto: false
	UpdatedAt time.Time  `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"` //[16] updated_at                                     DATETIME             strue   primary: false  auto: false
	DeletedAt *time.Time `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"` //[ 6] deleted_at                                     DATETIME             null: true   primary: false  auto: false
}
