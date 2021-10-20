package orm_user

import "time"

type UserStatisticData struct {
	Id                   uint64    `xorm:"pk notnull"`         // 用户ID -> Users: Id
	TotalProblemResolved uint64    `xorm:"notnull"`            // 用户做对题数
	TotalEnglishWord     uint64    `xorm:"notnull"`            // 用户词汇量
	TodayProblemResolved uint      `xorm:"notnull"`            // 用户今日做对题数
	TodayEnglishWord     uint      `xorm:"notnull"`            // 用户今日词汇量
	Intimacy             uint      `xorm:"notnull default(1)"` // 用户好感度
	ActiveDays           uint      `xorm:"notnull default(1)"` // 用户活跃天数
	UpdatedTime          time.Time `xorm:"notnull"`            // 数据有效日期
}
