package router

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/golang-absensi/cmd/webservice/handler"
	"github.com/uchupx/golang-absensi/cmd/webservice/middleware"
	"github.com/uchupx/golang-absensi/internal/config"
	"github.com/uchupx/golang-absensi/internal/repository"
	"github.com/uchupx/golang-absensi/internal/service"
)

type RouterInitParams struct {
	Ec         *echo.Echo
	Conf       *config.Config
	Service    service.Service
	Repository repository.Repository
	Middleware middleware.Middleware
}

func Init(params *RouterInitParams) {
	// public Path
	params.Ec.GET(PingPath, handler.HandlerPing(params.Service.Ping))

	params.Ec.POST(UserPath, handler.HandlerCreateUser(params.Service.CreateUser))
	params.Ec.POST(LoginPath, handler.HandlerLogin(params.Service.Login))

	// need access token
	params.Ec.POST(LogoutPath, handler.HandlerLogout(params.Service.Logout), params.Middleware.Auth())

	params.Ec.PUT(UserPath, handler.HandlerUpdateUser(params.Service.EditUser), params.Middleware.Auth())

	params.Ec.POST(AttendanceInPath, handler.HandlerCheckIn(params.Service.CheckIn), params.Middleware.Auth())
	params.Ec.POST(AttendanceOutPath, handler.HandlerCheckOut(params.Service.CheckOut), params.Middleware.Auth())
	params.Ec.GET(AttendancePath, handler.HandlerFindAttendance(params.Service.FindAttendance), params.Middleware.Auth())

	params.Ec.GET(ActivityPath, handler.HandlerGetActvity(params.Service.FindActivityByUserId), params.Middleware.Auth())
	params.Ec.POST(ActivityPath, handler.HandlerCreateActivity(params.Service.CreateActivity), params.Middleware.Auth())
	params.Ec.PUT(ActivityPath+"/:id", handler.HandlerUpdateActivity(params.Service.EditActivity), params.Middleware.Auth())
	params.Ec.DELETE(ActivityPath+"/:id", handler.HandlerDeleteActivity(params.Service.DeleteActivity), params.Middleware.Auth())
}
