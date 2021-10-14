package clients

import (
	"CeobeBot-Backend/middleware/database"
	"github.com/gin-gonic/gin"
)

type ClientInterface struct {

}

func (c ClientInterface) BindApi(engine *gin.Engine, db database.MysqlConnection) error {
	return nil
}
