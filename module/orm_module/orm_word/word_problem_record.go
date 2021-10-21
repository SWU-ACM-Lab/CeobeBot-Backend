package orm_word

import "time"

type WordProblemRecord struct {
	Id         uint64    `xorm:"pk notnull autoincr"` // 单词记录ID
	UserId     uint64    `xorm:"notnull"`             // 单词问题请求者 -> Users: Id
	UserType   uint      `xorm:"notnull default(1)"`  // 单词问题请求者类型: 1-个人; 2-群聊;
	UpdateTime time.Time `xorm:"notnull"`             // 单词问题更新时间
	Answer     int       `xorm:"notnull"`             // 单词问题答案
	Status     uint      `xorm:"notnull default(1)"`  // 单词问题解决状态: 1-未解决; 2-已解决
}
