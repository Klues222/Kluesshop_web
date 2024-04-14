package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"
)

func InitSrvConn() {
	userconn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", global.ServerConfig.ConsulInfo.Host,
			global.ServerConfig.ConsulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接用户服务失败")
	}
	goodsSrvClient := proto.NewGoodsClient(userconn)
	global.GoodsSrvClient = goodsSrvClient
}
