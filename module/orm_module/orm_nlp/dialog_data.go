package orm_nlp

type DialogData struct {
	Id         uint64 `xorm:"pk notnull autoincr"` // 对话数据ID
	TemplateId uint64 `xorm:"notnull"`             // 对话模版ID DialogTemplates: Id
	Index      uint   `xorm:"notnull"`             // 对话数据索引
	Data       string `xorm:"notnull varchar(63)"` // 对话数据
}
