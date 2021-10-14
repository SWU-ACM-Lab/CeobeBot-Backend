package orm_problem

import "time"

type ProblemSolvedRecords struct {
	Id        uint64 `xorm:"pk notnull autoincr"` // 记录ID
	UserId    uint64 `xorm:"notnull"`             // 用户ID(Q号) -> Users: Id
	ProblemId uint64 `xorm:"notnull"`             // 题目ID -> Problems: Id
	// SolvedStatus 1: 正确 2: 错误 0: 无效值
	SolvedStatus uint      `xorm:"notnull default(0)"` // 题目解决状态
	UpdateTime   time.Time `xorm:"notnull"`            // 记录更新时间
}