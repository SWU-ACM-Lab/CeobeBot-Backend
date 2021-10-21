package orm_client

import "time"

type Clients struct {
	Id           uint64    `xorm:"pk notnull autoincr"`  // 客户端ID
	Name         string    `xorm:"notnull varchar(15)"`  // 客户端名称
	Secret       string    `xorm:"notnull char(64)"`     // 客户端密钥
	Token        string    `xorm:"notnull char(64)"`     // 客户端TOKEN
	RefreshToken string    `xorm:"notnull char(64)"`     // TOKEN刷新TOKEN
	AdminEmail   string    `xorm:"notnull varchar(63)"`  // 客户端管理员邮箱
	Status       uint      `xorm:"notnull default(1)"`   // 客户端状态
	Scope        uint      `xorm:"notnull default(111)"` // 客户端权限
	CallBalance  uint      `xorm:"notnull default(100)"` // 客户端调用余额
	UpdateTime   time.Time `xorm:"notnull"`              // 客户端更新时间
	CreateTime   time.Time `xorm:"notnull"`              // 客户端创建时间
}
