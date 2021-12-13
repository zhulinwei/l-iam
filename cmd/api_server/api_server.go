package main

import (
	"l-iam/internal/api_server"
	"l-iam/internal/api_server/config/options"
	"l-iam/pkg/app"

	"math/rand"
	"time"
)

const (
	name    = "api-server"
	version = "0.1.0"
	desc    = `l-iam是极客时间专栏的学习项目，项目仅供学习，勿用于生产环境
关于l-iam的更多详情，可以访问：
	https://github.com/zhulinwei/l-iam/blob/main/README.md
`
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	opts := options.NewOptions()
	app.NewApp(name,
		app.WithDesc(desc),
		app.WithOptions(opts),
		app.WithVersion(version),
		// api_server.Run返回的func作为app的属性，在app构建完成执行Run方法后执行func
		app.WithRunFunc(api_server.Run(opts)),
	).Run()
}
