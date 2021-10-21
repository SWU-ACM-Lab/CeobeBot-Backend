package user_controller

import (
	"CeobeBot-Backend/middleware/encoder"
	"CeobeBot-Backend/module/orm_module/orm_user"
	"time"
)

// RefreshToken 刷新User的Token，如果成功，返回User实体
func (c UserController) RefreshToken(uid uint64, token string) (result bool, user orm_user.Users, err error) {
	users := orm_user.Users{}
	queryResult, queryError := c.db.Connection.ID(uid).Get(&users)

	// 给出认证
	if queryResult != true || queryError != nil {
		return false, orm_user.Users{}, queryError
	} else {
		// 用户鉴权
		if users.Token == token {
			// 刷新Token
			users.Token, _ = encoder.EncodeHash(users.Name + time.Now().String())
			users.UpdateTime = time.Now()

			// 写入数据库
			if _, errs := c.db.Connection.ID(uid).Update(users); errs != nil {
				return false, orm_user.Users{}, errs
			} else {
				return true, users, nil
			}
		} else {
			return false, orm_user.Users{}, nil
		}
	}
}

// SignUp 注册新用户，如果成功，返回User实体
func (c UserController) SignUp(uid uint64, uname string) (result bool, user orm_user.Users, err error) {
	// 查询用户是否存在
	queryResult, queryError := c.UserQuery(uid)
	if queryResult == true || queryError != nil {
		return false, orm_user.Users{}, queryError
	} else {
		// 填写新注册User字段
		users := orm_user.Users{
			Id:          uid,
			Name:        uname,
			Personality: 1,
			UpdateTime:  time.Now(),
			Scope:       001,
		}
		users.Token, _ = encoder.EncodeHash(uname + time.Now().String())

		// 更新数据库
		if _, errs := c.db.Connection.Table(new(orm_user.Users)).InsertOne(users); errs != nil {
			return false, orm_user.Users{}, errs
		} else {
			return true, users, nil
		}
	}
}

// Update 更新用户信息Name或者Personality，如果成功，返回User实体
func (c UserController) Update(uid uint64, uname string, personality uint, token string) (result bool, user orm_user.Users, err error) {
	// 获取uid对应的User实体
	users := orm_user.Users{}
	queryResult, queryError := c.db.Connection.ID(uid).Get(&users)

	if queryResult != true || queryError != nil {
		return false, orm_user.Users{}, queryError
	} else if users.Token == token {
		// 更新用户信息
		users.Name = uname
		users.Personality = personality
		// 更新数据库
		if _, errs := c.db.Connection.ID(uid).Update(users); errs != nil {
			return false, orm_user.Users{}, errs
		} else {
			return true, users, nil
		}
	} else {
		return false, orm_user.Users{}, nil
	}
}
