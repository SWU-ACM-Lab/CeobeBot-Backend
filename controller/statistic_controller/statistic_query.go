package statistic_controller

import (
	"CeobeBot-Backend/module/orm_module/orm_user"
)

func (c StatisticController) QueryUserStatisticWithAuth (cid uint64, token string, uid uint64, uname string) (result bool, data orm_user.UserStatisticData, err error) {
	if updateResult, updateData, updateError := c.updateStatistic(cid, token, uid, uname); updateResult != true || updateError != nil {
		return false, orm_user.UserStatisticData{}, updateError
	} else {
		return true, updateData, nil
	}
}
