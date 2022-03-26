package statistic

import (
	"CeobeBot-Backend/module/response_module"
)

type UserStatisticDataResponse struct {
	response_module.BaseResponse
	Id                   uint64 `json:"user_id"`                // 用户ID
	TotalProblemResolved uint64 `json:"total_problem_resolved"` // 用户总做对题目数
	TotalEnglishWord     uint64 `json:"total_english_word"`     // 用户总词汇量
	TodayProblemResolved uint   `json:"today_problem_resolved"` // 用户今日做对题目数
	TodayEnglishWord     uint   `json:"today_english_word"`     // 用户今日背单词数量
	Intimacy             uint   `json:"intimacy"`               // 用户亲密度
	ActiveDays           uint   `json:"active_days"`            // 用户活跃天数
	UpdatedTime          string `json:"updated_time"`           // 数据更新日期
}
