package model

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
