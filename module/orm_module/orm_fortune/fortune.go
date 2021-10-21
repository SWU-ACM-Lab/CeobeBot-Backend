package orm_module

type Fortunes struct {
	Id          uint64 `xorm:"pk notnull autoincr"`  // 运势ID
	Report      string `xorm:"notnull varchar(15)"`  // 综合运势
	Description string `xorm:"notnull varchar(255)"` // 运势详解
}
