package orm_user

type UserTaskList struct {
	Id     uint64 `xorm:"pk notnull autoincr"` // 用户任务ID
	UserId uint64 `xorm:"notnull"`             // 用户ID

}
