package client_contorller

import (
	"CeobeBot-Backend/module"
	"CeobeBot-Backend/module/orm_module/orm_client"
)

// ClientAuthSecret 验证客户端密钥
func (c ClientController) ClientAuthSecret(cid uint64, secret string) (result bool, err error) {
	// 获取ClientId对应的Client实体
	client := orm_client.Clients{}
	queryResult, queryError := c.db.Connection.ID(cid).Get(&client)

	// 给出认证
	if queryResult != true || queryError != nil {
		return false, queryError
	} else {
		return client.Secret == secret, nil
	}
}

// ClientAuthToken 验证客户端token
func (c ClientController) ClientAuthToken(cid uint64, token string) (result bool, err error) {
	// 获取ClientId对应的Client实体
	client := orm_client.Clients{}
	queryResult, queryError := c.db.Connection.ID(cid).Get(&client)

	// 给出认证
	if queryResult != true || queryError != nil {
		return false, queryError
	} else {
		return client.Token == token, nil
	}
}

// ClientAuthCall 验证客户端能否调用API
func (c ClientController) ClientAuthCall(cid uint64) (result bool, err error) {
	// 获取ClientId对应的Client实体
	client := orm_client.Clients{}
	queryResult, queryError := c.db.Connection.ID(cid).Get(&client)

	// 判断是否合法
	if queryResult != true || queryError != nil {
		return false, queryError
	} else {
		return client.CallBalance > 0 && client.Status != uint(ClientStatus(banned)) && client.Scope >= uint(module.ScopeStatus(module.Client)), nil
	}
}
