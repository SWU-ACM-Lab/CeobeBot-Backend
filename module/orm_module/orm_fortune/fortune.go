package orm_module

type Fortunes struct {
	Id           uint64 `xorm:"pk notnull autoincr"`  // 运势ID
	Theme        uint64 `xorm:"notnull"`              // 运势主题
	FortuneLevel uint   `xorm:"notnull"`              // 运势得分
	Report       string `xorm:"notnull varchar(15)"`  // 综合运势
	Description  string `xorm:"notnull varchar(255)"` // 运势详解
}

func (c Fortunes) LevelToString() string {
	var fortuneMapper = [...]string{
		"大凶",
		"凶",
		"小凶",
		"下平",
		"中平",
		"上平",
		"小吉",
		"吉",
		"大吉",
	}
	return fortuneMapper[c.FortuneLevel]
}