package client_controller

import (
	"CeobeBot-Backend/middleware/encoder"
	"CeobeBot-Backend/module"
	"CeobeBot-Backend/module/orm_module/orm_client"
	"CeobeBot-Backend/module/orm_module/orm_user"
	"time"
)

func (c ClientController) AddClient(name, email string, scope uint, admin uint64, token string) (result bool, err error) {
	// 获取admin对应的User实体
	user := orm_user.Users{}
	queryResult, queryError := c.db.Connection.ID(admin).Get(&user)
	if queryResult != true || queryError != nil {
		return false, queryError
	} else {
		// 验证admin是否有AddClient的权限
		if user.Scope < uint(module.ScopeStatus(module.Super)) || user.Token != token {
			return false, nil
		} else {
			// 生成Client信息
			client := orm_client.Clients{
				Name:        name,
				AdminEmail:  email,
				Status:      uint(ClientStatus(normal)),
				Scope:       scope,
				CallBalance: 10000,
				UpdateTime:  time.Now(),
				CreateTime:  time.Now(),
			}
			client.Secret, _ = encoder.EncodeHash(client.AdminEmail + time.Now().String())
			client.Token, _ = encoder.EncodeHash(client.Name + time.Now().String())
			client.RefreshToken, _ = encoder.EncodeHash(client.Token + time.Now().String())
			// 更新到数据库
			if _, errs := c.db.Connection.Table(client).InsertOne(client); errs != nil {
				return false, errs
			} else {
				return true, nil
			}
		}
	}
}

func (c ClientController) UpdateClient(cid uint64, client orm_client.Clients, admin uint64, token string) (result bool, err error) {
	// 获取admin对应的User实体
	user := orm_user.Users{}
	queryResult, queryError := c.db.Connection.ID(admin).Get(&user)
	if queryResult != true || queryError != nil {
		return false, queryError
	} else {
		// 验证admin是否有UpdateClient的权限
		if user.Scope < uint(module.ScopeStatus(module.Super)) || user.Token != token {
			return false, nil
		} else {
			// 验证cid是否合法
			if queryResult, _ := c.db.Connection.ID(cid).Get(new(orm_client.Clients)); queryResult != true {
				return false, nil
			} else {
				// 更新到数据库
				_, errs := c.db.Connection.Table(client).Update(client)
				if errs != nil {
					return false, errs
				} else {
					return true, nil
				}
			}
		}
	}
}

func (c ClientController) GetClient(cid uint64, admin uint64, token string) (result bool, client orm_client.Clients, err error) {
	// 获取admin对应的User实体
	user := orm_user.Users{}
	queryResult, queryError := c.db.Connection.ID(admin).Get(&user)
	if queryResult != true || queryError != nil {
		return false, orm_client.Clients{}, queryError
	} else {
		// 验证admin是否有GetClient的权限
		if user.Scope < uint(module.ScopeStatus(module.Super)) || user.Token != token {
			return false, orm_client.Clients{}, nil
		} else {
			// 进行数据库查询
			queryClient := orm_client.Clients{}
			clientResult, clientError := c.db.Connection.ID(cid).Get(&queryClient)
			if clientError != nil || clientResult != true {
				return false, orm_client.Clients{}, queryError
			} else {
				return true, queryClient, nil
			}
		}
	}
}
