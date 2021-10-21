package orm_task

type Tasks struct {
	Id            uint64 `xorm:"pk notnull autoincr"` // 任务ID
	Name          string `xorm:"notnull varchar(15)"` // 任务名称
	Description   string `xorm:"notnull varchar(63)"` // 任务描述
	ExceptedValue uint   `xorm:"notnull"`             // 任务目标
}
