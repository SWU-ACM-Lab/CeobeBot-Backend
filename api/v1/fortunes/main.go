package fortunes

import (
	"CeobeBot-Backend/controller/client_controller"
	"CeobeBot-Backend/controller/fortune_controller"
	"CeobeBot-Backend/middleware/database"
	"CeobeBot-Backend/module/request_module/v1/fortunes"
	"CeobeBot-Backend/module/response_module"
	fortunes2 "CeobeBot-Backend/module/response_module/v1/fortunes"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type FortuneInterface struct {
}

func (c FortuneInterface) BindApi(engine *gin.Engine, db database.MysqlConnection) error {
	route := engine.Group("/v1/fortunes")

	c.addFortune(route, db)
	c.getFortune(route, db)
	return nil
}

func (c FortuneInterface) addFortune(route *gin.RouterGroup, db database.MysqlConnection) {
	route.POST("", func(context *gin.Context) {
		// 查询上下文
		cid := context.Query("cid")
		token := context.Query("token")
		request := fortunes.AddFortuneRequest{}
		err1 := context.BindJSON(&request)
		if cid == "" || token == "" {
			context.JSON(400, nil)
			return
		}
		intCid, err2 := strconv.ParseUint(cid, 10, 64)
		if err1 != nil || err2 != nil {
			context.JSON(400, nil)
			return
		}

		//验证client
		client := client_controller.ClientController{}
		fortunes := fortune_controller.FortuneController{}
		client.Init(db)
		fortunes.Init(db)
		if clientResult, clientError := client.ClientAuthToken(intCid, token); clientResult != true || clientError != nil {
			context.JSON(403, nil)
			return
		}

		// 添加签文
		if result, fortune, err := fortunes.AddFortunes(request.Theme, request.Level, request.Report, request.Description); result != true || err != nil {
			fmt.Println(result, err)
			context.JSON(500, nil)
		} else {
			context.JSON(200, fortunes2.FortuneResponse{
				BaseResponse: response_module.BaseResponse{
					Message: "Success",
					Time:    time.Now().String(),
				},
				Theme:        fortune.Theme,
				FortuneLevel: fortune.LevelToString(),
				Report:       fortune.Report,
				Description:  fortune.Description,
			})
		}

	})
}

func (c FortuneInterface) getFortune(route *gin.RouterGroup, db database.MysqlConnection) {
	route.GET("", func(context *gin.Context) {
		// 查询上下文
		cid := context.Query("cid")
		token := context.Query("token")
		theme := context.Query("theme")
		uid := context.Query("uid")
		if cid == "" || token == "" || theme == "" || uid == "" {
			context.JSON(400, nil)
			return
		}
		intCid, err0 := strconv.ParseUint(cid, 10, 64)
		intTheme, err1 := strconv.ParseUint(theme, 10, 64)
		intUid, err2 := strconv.ParseUint(uid, 10, 64)
		if err0 != nil || err1 != nil || err2 != nil {
			context.JSON(400, nil)
			return
		}

		// 生成签文
		fortunes := fortune_controller.FortuneController{}
		fortunes.Init(db)
		result, fortune, err := fortunes.GetFortunesWithAuth(intCid, token, intUid, intTheme)
		if result != true || err != nil {
			context.JSON(500, nil)
			return
		} else {
			context.JSON(200, fortunes2.FortuneResponse{
				BaseResponse: response_module.BaseResponse{
					Message: "Success",
					Time:    time.Now().String(),
				},
				Theme:        fortune.Theme,
				FortuneLevel: fortune.LevelToString(),
				Report:       fortune.Report,
				Description:  fortune.Description,
			})
			return
		}
	})
}
