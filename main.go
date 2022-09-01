package main

import (
	"online-QA-community/config"
	"online-QA-community/respository"
	"online-QA-community/router"
)

// @title 蓝山工作室考核项目接口测试 -- 实现一个类似知乎的问答社区
// @version 1.0
func main() {
	config.InitConfig()     //加载配置文件
	respository.InitDB()    //初始化数据库
	respository.InitRedis() //初始化redis数据库

	engine := router.Router()
	engine.Run("0.0.0.0:8080") //启动路由

}
