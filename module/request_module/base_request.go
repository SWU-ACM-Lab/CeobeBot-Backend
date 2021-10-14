package request_module

type BaseRequest struct {
	ClientId    uint64 `json:"client_id"`    // 客户端ID
	ClientToken string `json:"client_token"` // 客户端token
}
