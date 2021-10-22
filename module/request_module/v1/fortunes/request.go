package fortunes

import "CeobeBot-Backend/module/request_module"

type AddFortuneRequest struct {
	request_module.BaseRequest
	Level       uint   `json:"level"`       // 签文等级
	Theme       uint64 `json:"theme"`       // 签文主题
	Report      string `json:"report"`      // 签文概述
	Description string `json:"description"` // 签文描述
}
