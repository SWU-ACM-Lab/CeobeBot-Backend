package orm_nlp

type DialogTemplates struct {
	Id          uint64 `xorm:"pk notnull autoincr"`  // 对话模版ID
	ScenesId    uint64 `xorm:"notnull"`              // 对话场景ID -> DialogScenes: Id
	Template    string `xorm:"notnull varchar(255)"` // 模版内容
	DataCount   uint   `xorm:"notnull"`              // 需要数据数量
	Personality uint   `xorm:"notnull"`              // 对话人格偏好
}
