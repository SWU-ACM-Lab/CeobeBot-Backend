package orm_problem

type Problems struct {
	Id        uint64 `xorm:"pk notnull autoincr"`   // 题目ID
	SubjectId uint   `xorm:"notnull"`               // 题目科目ID -> Subjects: Id
	Question  string `xorm:"notnull varchar(1023)"` // 题目内容
	// QuestionType: 1: 单选题 2: 多选题 0: 无数据
	QuestionType uint   `xorm:"notnull default(1)"`    // 题目类型
	OptionsA     string `xorm:"notnull varchar(127)"`  // 选项A
	OptionsB     string `xorm:"notnull varchar(127)"`  // 选项B
	OptionsC     string `xorm:"notnull varchar(127)"`  // 选项C
	OptionsD     string `xorm:"notnull varchar(127)"`  // 选项D
	Answer       string `xorm:"notnull varchar(4)"`    // 答案
	Tips         string `xorm:"notnull varchar(1023)"` // 解析
	TotalJudge   uint   `xorm:"default(0)"`            // 总做题人数
	CorrectJudge uint   `xorm:"default(0)"`            // 正确做题人数
}