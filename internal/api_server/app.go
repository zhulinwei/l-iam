package api_server

import (
	"context"
	"encoding/json"
	"fmt"
	"l-iam/internal/api_server/config"
	"l-iam/internal/api_server/config/options"
	"l-iam/internal/api_server/dao"
	"l-iam/internal/api_server/router"
	"l-iam/pkg/app"
	"l-iam/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(opts *options.Options) app.RunFunc {
	return func(name string) error {
		log.Init(opts.Log)
		defer log.Flush()
		optsBytes, _ := json.Marshal(opts)
		log.Info(string(optsBytes))
		apiServer := NewAPIServer(config.NewConfig(opts))
		if err := apiServer.PrepareRun(); err != nil {
			return err
		}
		// todo 有待思考合理性
		if err := apiServer.BeforeStop(); err != nil {
			return err
		}
		return apiServer.Run()
	}
}

type APIServer struct {
	route  *gin.Engine
	server *http.Server
	config *config.Config
}

func NewAPIServer(cfg *config.Config) *APIServer {
	return &APIServer{config: cfg}
}

func (a *APIServer) PrepareRun() error {
	// 初始化数据库
	// 使用mysql实现的存储层，如果有需求可以直接在此处替换其他实现
	factory, err := dao.NewApiServerFactory(a.config.MySQL)
	if err != nil {
		return err
	}
	dao.SetClient(factory)

	// 初始化路由
	gin.SetMode(a.config.Server.Mode)
	a.route = gin.New()
	router.InitRouter(a.route)

	return nil
}

func (a *APIServer) Run() error {
	// grpc run
	// web run
	a.server = &http.Server{
		// 监听的TCP地址
		Addr: a.config.Server.Address + ":" + strconv.Itoa(a.config.Server.Port),
		// http句柄，用于处理程序响应的HTTP请求
		Handler: a.route,
		// 等待的最大时间
		IdleTimeout: 6 * time.Minute,
		// 允许读取的最大时间
		ReadTimeout: 30 * time.Second,
		// 允许写入的最大时间
		WriteTimeout: 30 * time.Second,
		// 请求头的最大字节数
		MaxHeaderBytes: 1 << 20,
	}
	// 启动rpc服务
	if err := a.server.ListenAndServe(); err != nil {
		fmt.Println("err:", err.Error())
	}
	return nil
}

func (a *APIServer) BeforeStop() error {
	// web close
	// grpc close
	// mysql close
	// redis close

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := a.server.Shutdown(context.Background()); err != nil {
		}
	}()

	return nil
}
