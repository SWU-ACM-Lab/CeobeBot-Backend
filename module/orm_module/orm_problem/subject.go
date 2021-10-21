package orm_problem

type Subjects struct {
	Id          uint64 `xorm:"pk notnull autoincr"`  // 科目ID
	Name        string `xorm:"notnull varchar(63)"`  // 科目名称
	Description string `xorm:"notnull varchar(255)"` // 科目描述
}
