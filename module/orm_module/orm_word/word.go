package orm_word

type Words struct {
	Id           string `xorm:"pk notnull char(10)"`              // 单词ID
	Spell        string `xorm:"notnull varchar(64)"`              // 单词拼写
	Phonetic     string `xorm:"notnull longtext"`                 // 单词音标
	Pos          string `xorm:"notnull longtext"`                 // 单词释义
	WordForms    string `xorm:"'word_forms' longtext"`            // 单词形式
	AudioSources string `xorm:"'audio_sources' notnull longtext"` // 单词读音URL
	Paraphrase   string `xorm:"text"`                             // 单词英英释义
}
