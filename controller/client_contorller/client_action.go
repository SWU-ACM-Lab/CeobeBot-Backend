package client_contorller

import (
	"CeobeBot-Backend/middleware/encoder"
	"CeobeBot-Backend/module/orm_module/orm_client"
	"time"
)

// RefreshToken 更新客户端token和refresh_token
func (c ClientController) RefreshToken(cid uint64, secret, refreshToken string) (result bool, token, refresh string, err error) {
	// 获取ClientId对应的Client实体
	client := orm_client.Clients{}
	queryResult, queryError := c.db.Connection.ID(cid).Get(&client)

	// 判断是否合法
	if queryResult != true || queryError != nil || client.Secret != secret || client.RefreshToken != refreshToken {
		return false, "", "", queryError
	} else {
		// 刷新Token
		client.Token, _ = encoder.EncodeHash(client.Name + time.Now().String())
		client.RefreshToken, _ = encoder.EncodeHash(client.Token + time.Now().String())
		// 更新数据库
		if _, errs := c.db.Connection.ID(cid).Update(client); errs != nil {
			return false, "", "", errs
		} else {
			return true, client.Token, client.RefreshToken, nil
		}
	}
}

func (c ClientController) CallApi(cid uint64) bool {
	// 获取ClientId对应的Client实体
	client := orm_client.Clients{}
	queryResult, queryError := c.db.Connection.ID(cid).Get(&client)

	// 判断是否合法
	if queryResult != true || queryError != nil {
		return false
	} else {
		client.CallBalance -= 1
		// 更新到数据库
		_, _ = c.db.Connection.ID(client.Id).Update(client)
		return true
	}
}