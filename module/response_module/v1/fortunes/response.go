package fortunes

import "CeobeBot-Backend/module/response_module"

type FortuneResponse struct {
	response_module.BaseResponse
	Id           uint64 `json:"id"`            // 运势ID
	Theme        uint64 `json:"theme"`         // 运势主题
	FortuneLevel string `json:"fortune_level"` // 运势得分
	Report       string `json:"report"`        // 综合运势
	Description  string `json:"description"`   // 运势详解
}
