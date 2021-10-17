package orm_nlp

type DialogScenes struct {
	Id          uint64 `xorm:"pk notnull autoincr"`  // 对话场景ID
	Name        string `xorm:"notnull varchar(15)"`  // 对话场景名称
	Description string `xorm:"notnull varchar(127)"` // 对话场景描述
	Type        uint64 `xorm:"notnull"`              // 对话场景类别
}