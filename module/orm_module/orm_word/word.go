package orm_word

type Word struct {
	Id         string `xorm:"pk notnull char(10)"` // 单词ID
	Spell      string `xorm:"notnull varchar(64)"` // 单词拼写
	Phonetic   string `xorm:"notnull longtext"`    // 单词音标
	Pos        string `xorm:"notnull longtext"`    // 单词释义
	WordForm   string `xorm:"notnull longtext"`    // 单词形式
	Paraphrase string `xorm:"text"`                // 单词英英释义
}
