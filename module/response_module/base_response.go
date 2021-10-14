package response_module

type BaseResponse struct {
	Message string `json:"message"` // 响应信息
	Time    string `json:"time"`    // 响应时间
}