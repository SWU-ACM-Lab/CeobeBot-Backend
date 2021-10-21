package user_controller

import (
	"CeobeBot-Backend/module/orm_module/orm_user"
)

// UserAuthToken 验证User的Token是否正确，如果正确，返回User实体
func (c UserController) UserAuthToken(uid uint64, token string) (result bool, user orm_user.Users, err error) {
	// 获取uid对应的User实体
	users := orm_user.Users{}
	queryResult, queryError := c.db.Connection.ID(uid).Get(&users)

	// 给出认证
	if queryResult != true || queryError != nil {
		return false, orm_user.Users{}, queryError
	} else {
		if users.Token == token {
			return true, users, nil
		} else {
			return false, orm_user.Users{}, nil
		}
	}
}

// UserQuery 查询数据库中是否存在User，如果存在，返回TRUE
func (c UserController) UserQuery(uid uint64) (result bool, err error) {
	// 获取uid对应的User实体
	users := orm_user.Users{}
	queryResult, queryError := c.db.Connection.ID(uid).Get(&users)

	if queryResult != true || queryError != nil {
		return false, queryError
	} else {
		return true, nil
	}
}
