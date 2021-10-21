package statistic_controller

import (
	"CeobeBot-Backend/controller/client_contorller"
	"CeobeBot-Backend/controller/user_controller"
	"CeobeBot-Backend/module"
	"CeobeBot-Backend/module/orm_module/orm_user"
	"time"
)

func (c StatisticController) updateStatistic(cid uint64, token string, uid uint64, uname string) (result bool, data orm_user.UserStatisticData, err error) {
	clientController := client_contorller.ClientController{}
	userController := user_controller.UserController{}
	userController.Init(c.db)
	clientController.Init(c.db)

	// 验证客户端是否合法
	if results, errs := clientController.ClientAuthToken(cid, token); errs != nil || results != true {
		return false, orm_user.UserStatisticData{}, errs
	}

	// 判断用户是否存在
	if results, errs := userController.UserQuery(uid); errs != nil {
		return false, orm_user.UserStatisticData{}, errs
	} else if results != true {
		// 自动注册用户
		if signResult, _, signError := userController.SignUp(uid, uname); signResult != true || signError != nil {
			return false, orm_user.UserStatisticData{}, signError
		}
	}

	// 获取uid对应的UserData
	userData := orm_user.UserStatisticData{}
	if dataResult, _ := c.db.Connection.ID(uid).Get(&userData); dataResult != true {
		// 如果没有用户信息数据，则创建用户信息数据
		userData = orm_user.UserStatisticData{
			Id:                   uid,
			TotalProblemResolved: 0,
			TotalEnglishWord:     0,
			TodayProblemResolved: 0,
			TodayEnglishWord:     0,
			Intimacy:             1,
			ActiveDays:           1,
			UpdatedTime:          time.Now(),
		}
		if _, insertError := c.db.Connection.Table(orm_user.UserStatisticData{}).InsertOne(userData); insertError != nil {
			return false, orm_user.UserStatisticData{}, insertError
		}
	}

	// 判断用户活跃时间
	d, _ := time.ParseDuration("-24h")
	if time.Now().Day() == userData.UpdatedTime.Day() {
		// 今天登录了，啥也不干
		userData.UpdatedTime = time.Now()
		c.db.Connection.ID(userData.Id).Update(userData)
		return true, userData, nil
	} else if time.Now().Add(d).Day() == userData.UpdatedTime.Day() {
		// 昨天登录了，更新亲密度，重置今日数据
		userData.Intimacy += uint(module.RobotConfig.ActiveIntimacy)
		userData.UpdatedTime = time.Now()
		userData.ActiveDays += 1
		userData.TodayEnglishWord = 0
		userData.TodayProblemResolved = 0
		c.db.Connection.ID(userData.Id).Update(userData)
		return true, userData, nil
	} else {
		// 没有连续登录，重置活跃时间与今日数据
		userData.UpdatedTime = time.Now()
		userData.ActiveDays = 0
		userData.TodayEnglishWord = 0
		userData.TodayProblemResolved = 0
		c.db.Connection.ID(userData.Id).Update(userData)
		return true, userData, nil
	}
}

func (c StatisticController) SolveWordProblem(cid uint64, token string, uid uint64, uname string) {
	if result, data, err := c.updateStatistic(cid, token, uid, uname); err != nil || result != true {
		return
	} else {
		data.TodayEnglishWord += 1
		data.TotalEnglishWord += 1
		data.Intimacy += 1
		//data.Intimacy += uint(module.RobotConfig.RememberWordIntimacy)  // todo: FIX ini loader
		c.db.Connection.ID(data.Id).Update(data)
	}
}
