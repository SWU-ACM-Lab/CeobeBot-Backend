package orm_user

import "time"

type UserTaskList struct {
	Id         uint64    `xorm:"pk notnull autoincr"` // 用户任务ID
	UserId     uint64    `xorm:"notnull"`             // 用户ID
	TaskIdA    uint64    `xorm:"notnull"`             // 用户任务1ID Tasks -> Id
	TaskValueA uint      `xorm:"notnull default(0)"`  // 用户任务1完成度
	TaskIdB    uint64    `xorm:"notnull"`             // 用户任务2ID Tasks -> Id
	TaskValueB uint      `xorm:"notnull default(0)"`  // 用户任务2完成度
	UpdateTime time.Time `xorm:"notnull"`             // 用户任务更新时间
}
