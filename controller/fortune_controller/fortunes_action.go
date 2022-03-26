package fortune_controller

import (
	"CeobeBot-Backend/middleware/random"
	"CeobeBot-Backend/module/orm_module/orm_client"
	orm_module "CeobeBot-Backend/module/orm_module/orm_fortune"
)

func (c FortuneController) AddFortunes(theme uint64, level uint, report, description string) (result bool, fortune orm_module.Fortunes, err error) {
	fortune = orm_module.Fortunes{
		Theme:        theme,
		FortuneLevel: level,
		Report:       report,
		Description:  description,
	}
	if _, errs := c.db.Connection.Table(orm_module.Fortunes{}).InsertOne(fortune); errs != nil {
		return false, orm_module.Fortunes{}, errs
	} else {
		return true, fortune, nil
	}
}

func (c FortuneController) GetFortunesWithAuth(cid uint64, token string, uid uint64, theme uint64) (result bool, fortune orm_module.Fortunes, err error) {
	// 获取ClientId对应的Client实体
	client := orm_client.Clients{}
	queryResult, queryError := c.db.Connection.ID(cid).Get(&client)

	// 给出认证
	if queryResult != true || queryError != nil {
		return false, orm_module.Fortunes{}, queryError
	}
	if client.Token != token {
		return false, orm_module.Fortunes{}, nil
	}

	// 生成签文
	level := random.DailyRandomUInt(uid+cid, 10)
	var fortunes []orm_module.Fortunes
	fortuneResult, fortuneError := c.db.Connection.Where("Theme = ?", theme).Where("FortuneLevel = ?", level).Get(&fortunes)
	if fortuneResult != true || fortuneError != nil {
		return false, orm_module.Fortunes{}, fortuneError
	} else {
		if len(fortunes) == 0 {
			return false, orm_module.Fortunes{}, nil
		} else {
			index := random.DailyRandomUInt(uid+cid, uint(len(fortunes)))
			realFortune := fortunes[index]
			return true, realFortune, nil
		}
	}
}
