package bootstrap

import (
	stdContext "context"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/websocket"
	"os"
	"time"
)

const (
	StaticAssets = "./assets/"
	Favicon      = "favicon.ico"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
}

// 创建引导程序实例
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}

// 执行外部配置，用于进行配置扩展
func (b *Bootstrapper) Configure(cfgs ...Configurator) *Bootstrapper {
	for _, cfg := range cfgs {
		cfg(b)
	}
	return b
}

// 从环境变量中读取配置信息
func (b *Bootstrapper) LoadEnv() *Bootstrapper {
	env := os.Getenv("IRIS_ENV")
	env_file := ".env"
	if "" != env {
		env_file = env_file + env
		err := godotenv.Load(env_file)
		if err != nil {
			b.Logger().Fatal(err)
		}
	}

	err := godotenv.Load()
	if err != nil {
		b.Logger().Fatal(err)
	}

	return b
}

// 配置视图
func (b *Bootstrapper) SetupViews(viewDir string) *Bootstrapper {
	b.RegisterView(iris.HTML(viewDir, ".html").Layout("shared/layout.html"))
	return b
}

// 引导启动,执行必要的配置
func (b *Bootstrapper) PrepareDefault() *Bootstrapper {
	b.Logger().SetLevel("debug")
	b.Use(recover.New())
	b.Use(logger.New())
	return b.SetupViews("./views").
		SetupStaticFiles("./assets").
		LoadEnv().
		SetupErrorHandlers().
		SetupHealthChecker()
}

// 优雅地启动与关闭
func (b *Bootstrapper) Start() error {
	// 优雅地关闭系统
	idleConnsClosed := make(chan struct{})
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		//关闭所有主机
		if err := b.Shutdown(ctx) != nil; err {
			panic(err)
		}
		close(idleConnsClosed)
	})

	err := b.Listen(":8080", iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
	<-idleConnsClosed
	return err
}

// 准备WebSocket服务器
func (b *Bootstrapper) SetupWebsockets(endpoint string, handler websocket.ConnHandler) *Bootstrapper {
	ws := websocket.New(websocket.DefaultGorillaUpgrader, handler)
	b.Get(endpoint, websocket.Handler(ws))
	return b
}

// 配置健康度检查
func (b *Bootstrapper) SetupHealthChecker() *Bootstrapper {
	b.Get("/health", func(ctx *context.Context) {
		ctx.StatusCode(iris.StatusOK)
	})
	return b
}

// 配置静态资源路径
func (b *Bootstrapper) SetupStaticFiles(staticDir string) *Bootstrapper {
	b.Favicon(StaticAssets + Favicon)
	b.HandleDir(staticDir, iris.Dir(StaticAssets))
	return b
}

// 配置非正常的异常处理
func (b *Bootstrapper) SetupErrorHandlers() *Bootstrapper {
	b.OnAnyErrorCode(func(c iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  c.GetStatusCode(),
			"message": c.Values().GetString("message"),
		}

		if jsonOutput := c.GetHeader("Content Type") == "application/json"; jsonOutput {
			c.JSON(err)
			return
		}

		c.ViewData("Err", err)
		c.ViewData("Title", "未知错误")
		if er := c.View("shared/error.html") != nil; er {
			panic(er)
		}
	})
	return b
}
