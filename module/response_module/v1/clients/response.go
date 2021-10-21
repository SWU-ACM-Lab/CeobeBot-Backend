package clients

import "CeobeBot-Backend/module/response_module"

type GetClientResponse struct {
	response_module.BaseResponse
	ClientId           uint64 `json:"client_id"`
	ClientName         string `json:"client_name"`
	AdminEmail         string `json:"admin_email"`
	ClientStatus       uint   `json:"client_status"`
	ClientScope        uint   `json:"client_scope"`
	ClientToken        string `json:"client_token"`
	ClientRefreshToken string `json:"client_refresh_token"`
}
