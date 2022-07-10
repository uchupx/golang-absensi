package webservice

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/uchupx/golang-absensi/cmd/webservice/middleware"
	"github.com/uchupx/golang-absensi/cmd/webservice/router"
	"github.com/uchupx/golang-absensi/internal/config"
	"github.com/uchupx/golang-absensi/internal/repository"
	"github.com/uchupx/golang-absensi/internal/service"
)

type ServiceInitParams struct {
	Logger     *logrus.Entry
	Config     *config.Config
	Repository repository.Repository
	Service    service.Service
	Middleware middleware.Middleware
}

func StartServer(param *ServiceInitParams) {
	param.Logger.Infof("[StartServer] starting server in webserver mode")

	ec := echo.New()

	router.Init(&router.RouterInitParams{
		Ec:         ec,
		Conf:       param.Config,
		Service:    param.Service,
		Repository: param.Repository,
		Middleware: param.Middleware,
	})

	if err := ec.Start(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))); err != nil {
		param.Logger.Errorf("[StartServer] error listening to %s: %+v", param.Config.AppAddress, err)
	}
}
