package users

import (
	"CeobeBot-Backend/controller/statistic_controller"
	"CeobeBot-Backend/middleware/database"
	"CeobeBot-Backend/module/response_module"
	statistic2 "CeobeBot-Backend/module/response_module/v1/statistic"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type UserInterface struct {
}

func (u UserInterface) BindApi(engine *gin.Engine, db database.MysqlConnection) error {
	route := engine.Group("/v1/users")

	u.queryStatistic(route, db)

	return nil
}

func (u UserInterface) queryStatistic(route *gin.RouterGroup, db database.MysqlConnection) {
	route.GET("/statistic", func(context *gin.Context) {
		cid := context.Query("cid")
		token := context.Query("token")
		uid := context.Query("uid")
		uname := context.Query("uname")
		if cid == "" || token == "" || uid == "" || uname == "" {
			context.JSON(400, nil)
			return
		}

		intCid, err1 := strconv.ParseUint(cid, 10, 64)
		intUid, err2 := strconv.ParseUint(uid, 10, 64)
		if err1 != nil || err2 != nil {
			context.JSON(400, nil)
			return
		}

		statistic := statistic_controller.StatisticController{}
		statistic.Init(db)
		result, data, err := statistic.QueryUserStatisticWithAuth(intCid, token, intUid, uname)
		if result != true || err != nil {
			context.JSON(400, nil)
			return
		} else {
			context.JSON(200, statistic2.UserStatisticDataResponse{
				BaseResponse: response_module.BaseResponse{
					Message: "Success",
					Time:    time.Now().String(),
				},
				Id:                   data.Id,
				TotalProblemResolved: data.TotalProblemResolved,
				TotalEnglishWord:     data.TotalEnglishWord,
				TodayProblemResolved: data.TodayProblemResolved,
				TodayEnglishWord:     data.TodayEnglishWord,
				Intimacy:             data.Intimacy,
				ActiveDays:           data.ActiveDays,
				UpdatedTime:          data.UpdatedTime.String(),
			})
		}

	})
}
