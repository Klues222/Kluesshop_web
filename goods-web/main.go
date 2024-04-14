package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"

	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/initialize"
	"mxshop-api/goods-web/utils"
	"mxshop-api/goods-web/utils/register/consul"
)

func main() {

	//1.初始化logger
	initialize.InitLogger()
	//2.初始化配置文件
	initialize.InitConfig()
	//3.初始化routers
	Router := initialize.Routers()

	//4.初始化翻译
	err := initialize.InitTrans("zh")
	if err != nil {
		panic(err)
	}
	//5.初始化srv连接
	initialize.InitSrvConn()
	viper.AutomaticEnv()
	//如果是本地开发环境端口号固定，线上环境获取端口号
	debug := viper.GetBool("Debug")
	if debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	/**
	1.s可以获取一个全局的Sugar对象
	2.日志是分级别的，debug info warn error fetal
	3.s函数和l函数很有用
	*/

	register_client := consul.NewRegister(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprint("%s", uuid.NewV4())
	err = register_client.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败", err.Error())
	}
	zap.S().Infof("启动服务器， 端口： %d", global.ServerConfig.Port)
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败", err.Error())
		}
	}()
	//接受终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = register_client.DeRegister(serviceId)
	if err != nil {
		zap.S().Info("注销失败", err.Error())
	} else {
		zap.S().Info("注销成功")
	}

}
