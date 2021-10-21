package module

import (
	"CeobeBot-Backend/middleware/database"
	"CeobeBot-Backend/module/application_module"
	"github.com/gin-gonic/gin"
)

// 全局配置文件
var (
	DatabaseService   database.MysqlConnection             // 数据库服务
	ApiEngine         *gin.Engine                          // API服务
	BasePath          string                               // 基本路径
	DatabaseConfig    application_module.DatabaseConfig    // 全局数据库配置
	RedisConfig       application_module.RedisConfig       // 全局Redis配置
	ApplicationConfig application_module.ApplicationConfig // 全局应用配置
	RobotConfig       application_module.RobotConfig       // 全局机器人配置
)

// 全局变量
var (
	WordCount    uint64            // 单词数量缓存
	ProblemCount map[string]uint64 // 问题数量缓存
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
	Super  ScopeStatus = 111
	Client ScopeStatus = 011
	User   ScopeStatus = 001
	Guest  ScopeStatus = 000
)
