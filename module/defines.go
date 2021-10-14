package module

import (
	"CeobeBot-Backend/middleware/database"
	"CeobeBot-Backend/module/application_module"
	"github.com/gin-gonic/gin"
)

var (
	DatabaseService database.MysqlConnection
	ApiEngine *gin.Engine
	BasePath string
	DatabaseConfig application_module.DatabaseConfig
	RedisConfig application_module.RedisConfig
	ApplicationConfig application_module.ApplicationConfig
	RobotConfig application_module.RobotConfig
)

// ScopeStatus 数据库Scope字段
//
// 第一位： manage端口的权限
//
// 第二位： client端口的权限
//
// 第三位： user端口的权限
type ScopeStatus uint

const (
	Super ScopeStatus = 111
	Client ScopeStatus = 011
	User ScopeStatus = 001
	Guest ScopeStatus = 000
)