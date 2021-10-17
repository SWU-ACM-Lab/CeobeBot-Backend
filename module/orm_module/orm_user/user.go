package orm_user

import "time"

type Users struct {
	Id          uint64    `xorm:"pk notnull"`           // 用户ID -> QQ号
	Name        string    `xorm:"notnull varchar(63)"`  // 用户昵称
	Personality uint      `xorm:"notnull"`              // 用户偏好人格
	Token       string    `xorm:"char(64)"`             // 用户Token
	UpdateTime  time.Time `xorm:"notnull"`              // 用户更新时间
	Scope       uint      `xorm:"notnull default(111)"` // 用户权限
}