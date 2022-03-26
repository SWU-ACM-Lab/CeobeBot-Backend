package clients

import (
	"CeobeBot-Backend/controller/client_controller"
	"CeobeBot-Backend/middleware/database"
	"CeobeBot-Backend/module/response_module"
	"CeobeBot-Backend/module/response_module/v1/clients"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type ClientInterface struct {
}

func (c ClientInterface) BindApi(engine *gin.Engine, db database.MysqlConnection) error {
	route := engine.Group("/v1/clients")

	c.getClient(route, db)
	c.addClient(route, db)

	return nil
}

func (c ClientInterface) getClient(route *gin.RouterGroup, db database.MysqlConnection) {
	route.GET("", func(context *gin.Context) {
		cid := context.Query("cid")
		admin := context.Query("uid")
		token := context.Query("token")

		// 验证参数是否足够
		if cid == "" || admin == "" || token == "" {
			context.JSON(400, nil)
			return
		} else {
			controller := client_controller.ClientController{}
			intCid, err1 := strconv.ParseUint(cid, 10, 64)
			intUid, err2 := strconv.ParseUint(admin, 10, 64)
			if controller.Init(db) != nil || err1 != nil || err2 != nil {
				context.JSON(500, nil)
				return
			} else {
				// 获取Client
				result, client, err := controller.GetClient(intCid, intUid, token)
				if err != nil {
					context.JSON(500, nil)
					return
				} else if result != true {
					context.JSON(404, nil)
					return
				} else {
					// 返回Client
					context.JSON(200, clients.GetClientResponse{
						BaseResponse: response_module.BaseResponse{
							Message: "Success",
							Time:    time.Now().String(),
						},
						ClientId:           client.Id,
						ClientName:         client.Name,
						AdminEmail:         client.AdminEmail,
						ClientStatus:       client.Status,
						ClientScope:        client.Scope,
						ClientToken:        client.Token,
						ClientRefreshToken: client.RefreshToken,
					})
					return
				}
			}
		}
	})
}

func (c ClientInterface) addClient(route *gin.RouterGroup, db database.MysqlConnection) {
	route.POST("", func(context *gin.Context) {
		client := client_controller.ClientController{}
		client.Init(db)
	})
}
