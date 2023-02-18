package main

import (
	"douyinV/user/config"
	"douyinV/user/dal"
	"douyinV/user/global"
	"douyinV/user/kitex_gen/user/usersvr"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

const (
	ConfigFileName = "userConfig"
)

func Init() {
	// 初始化配置
	var err error
	err = config.InitConfig(&global.UserConfig, ConfigFileName)
	if err != nil {
		klog.Fatal("配置初始化失败")
		panic(err)
	}

	// 初始化数据库
	err = dal.Init(&global.UserConfig.DBConfig)
	if err != nil {
		klog.Fatal("数据库初始化失败")
		panic(err)
	}

}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{global.UserConfig.EtcdConfig.EtcdAddress})
	if err != nil {
		klog.Fatal("etcd注册失败")
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", global.UserConfig.EtcdConfig.ServiceAddress)
	if err != nil {

		panic(err)
	}
	svr := usersvr.NewServer(new(UserSvrImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: global.UserConfig.EtcdConfig.ServiceName}),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		//server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithRegistry(r),
	)
	if err = svr.Run(); err != nil {
		klog.Fatal(err)
		panic(err)
	}
}
