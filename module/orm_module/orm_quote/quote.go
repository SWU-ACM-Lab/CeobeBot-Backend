package orm_quote

type Quotes struct {
	Id   uint64 `xorm:"pk notnull autoincr"`   // 名言ID
	Data string `xorm:"notnull varchar(1023)"` // 名言内容
	From string `xorm:"notnull varchar(127)"`  // 名言出处
}
