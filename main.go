package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"go-uds/auth"
	"go-uds/bootstrap"
	"go-uds/database"
	"go-uds/repositories"
	"go-uds/services"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"time"
)

const (
	accessTokenMaxAge  = 2 * time.Hour
	refreshTokenMaxAge = time.Hour
)

var g errgroup.Group

func NewServiceContext() *services.ServiceContext {
	db := database.ConnectDB()
	idRepos := repositories.NewIdentityRepository(*db)
	usrRepos := repositories.NewUserRepository(*db)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PWD"),
		DB:       0,
	})

	privateKey, publicKey := jwt.MustLoadRSA("private_key.pem", "public_key.pem")
	signer := jwt.NewSigner(jwt.RS256, privateKey, accessTokenMaxAge)
	verifier := jwt.NewVerifier(jwt.RS256, publicKey)

	return &services.ServiceContext{
		idRepos,
		usrRepos,
		*signer,
		*verifier,
		client,
	}
}

func main() {
	// 起动SASS服务
	app := bootstrap.New("用户目录", "csharp2002@hotmail.com")
	app.PrepareDefault()

	app.Party("/auth").ConfigureContainer(func(container *iris.APIContainer) {
		authService := services.NewAuthService(NewServiceContext())
		container.RegisterDependency(authService)
		auth.RegisterRoutes(container)
	})

	g.Go(app.Start)

	// 起动管理服务

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
