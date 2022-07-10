package main

import (
	"flag"

	"github.com/uchupx/golang-absensi/cmd/webservice"
	"github.com/uchupx/golang-absensi/cmd/webservice/middleware"
	"github.com/uchupx/golang-absensi/internal/config"
	"github.com/uchupx/golang-absensi/internal/repository"
	"github.com/uchupx/golang-absensi/internal/service"
	"github.com/uchupx/golang-absensi/pkg/database"
	"github.com/uchupx/golang-absensi/pkg/redis"
	"github.com/uchupx/golang-absensi/pkg/util/crypt"
	"github.com/uchupx/golang-absensi/pkg/util/logutil"
)

var webserviceMode bool

func init() {
	flag.BoolVar(&webserviceMode, "webservice-mode", true, "run app in webservice mode")
}

func main() {
	flag.Parse()

	config.Init()
	conf := config.Get()

	logger := logutil.NewLogger(logutil.NewLoggerParams{
		ServiceName: conf.AppName,
	})

	db, err := database.InitMariaDB(&database.InitMariaDBParams{
		Conf:   &conf.DBConfig,
		Logger: logger,
	})

	if err != nil {
		logger.Fatalf("[main] error initialize database, error:%+v", err)
	}

	repo := repository.New(repository.Params{
		Logger: logger,
		DB:     db,
	})

	cryptSvc := crypt.NewCryptService(crypt.Params{
		Logger: logger,
		Conf:   conf,
	})

	redis, err := redis.InitRedis(&conf.RedisConfig)
	if err != nil {
		logger.Fatalf("[main] error init redis: %+v", err)
	}

	svc := service.New(service.Params{
		Repository:   repo,
		Logger:       logger,
		Conf:         conf,
		Redis:        redis,
		CryptService: cryptSvc,
	})

	middlewareSvc := middleware.NewMiddleware(middleware.Params{
		Repository:   repo,
		CryptService: cryptSvc,
		Redis:        redis,
	})

	if webserviceMode {
		webservice.StartServer(&webservice.ServiceInitParams{
			Logger:     logger,
			Service:    svc,
			Config:     conf,
			Repository: repo,
			Middleware: middlewareSvc,
		})
	}
}
