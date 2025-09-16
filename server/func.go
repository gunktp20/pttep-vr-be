package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"pttep-vr-api/pkg/config"
	"pttep-vr-api/pkg/constant/color"
	"pttep-vr-api/pkg/constant/state"
	"pttep-vr-api/pkg/repository"
	sAuthentication "pttep-vr-api/pkg/services/authentications"
	sGame "pttep-vr-api/pkg/services/game"
	sPermission "pttep-vr-api/pkg/services/permissions"
	sPing "pttep-vr-api/pkg/services/ping"
	sRole "pttep-vr-api/pkg/services/roles"
	sUser "pttep-vr-api/pkg/services/users"
	"pttep-vr-api/pkg/utils/errorMessage"
	"pttep-vr-api/pkg/utils/gormDB"
	hAuthentication "pttep-vr-api/server/handler/authentication"
	hGame "pttep-vr-api/server/handler/game"
	hPermission "pttep-vr-api/server/handler/permission"
	hPing "pttep-vr-api/server/handler/ping"
	hRegistration "pttep-vr-api/server/handler/registration"
	hRole "pttep-vr-api/server/handler/role"
	hUser "pttep-vr-api/server/handler/user"
	hVersion "pttep-vr-api/server/handler/version"
	"pttep-vr-api/server/recovery"
	"pttep-vr-api/server/response"
	"pttep-vr-api/server/route"
	"reflect"
	"runtime"
	"strings"
	"syscall"
	"time"

	// _ "pttep-vr-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"gorm.io/driver/mysql"
)

func Run() {

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	err = config.Init(pwd + "/pkg/config/config.yaml")
	if err != nil {
		panic(err)
		return
	}

	if err := errorMessage.Read(errorMessage.Raw); err != nil {
		panic(err)
		return
	}

	conf := config.Get()
	response.Init(conf)

	loc, err := time.LoadLocation(conf.App.Timezone)
	if err != nil {
		panic(err)
		return
	}
	time.Local = loc
	//rand.Seed(time.Now().UnixNano()) // for go < 1.19

	debug := false
	if conf.App.State == state.LOCAL {
		debug = true
	}

	app := fiber.New(fiber.Config{
		//DisableStartupMessage: true,
		BodyLimit: 50 * 1024 * 1024, // 5MB
		//EnablePrintRoutes: true,
		ProxyHeader:           fiber.HeaderXForwardedFor,
		ReadTimeout:           5 * time.Second,
		DisableStartupMessage: !debug,
		//...
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.New(context.Background(), errorMessage.Get(""), nil, fmt.Errorf("no path")))
		},
	})

	allowMethod := []string{
		fiber.MethodGet,
		//fiber.MethodHead,
		//fiber.MethodOptions,
		fiber.MethodPost,
		fiber.MethodPut,
		//fiber.MethodPatch,
		fiber.MethodDelete,
		//fiber.MethodTrace,
	}

	allowHeaders := []string{
		fiber.HeaderOrigin,
		fiber.HeaderAcceptLanguage,
		fiber.HeaderContentLength,
		fiber.HeaderContentType,
		fiber.HeaderAuthorization,
	}

	allowOrigins := strings.Join(conf.App.Config.Allows.Origins, ",")
	allowCredentials := true
	if strings.Contains(allowOrigins, "*") {
		allowOrigins = "*"
		allowCredentials = false
		fmt.Println(color.Red(fmt.Sprintf("  [WARNING] AllowOrigins will be set to [%s]", allowOrigins)))
		fmt.Println(color.Red(fmt.Sprintf("  [WARNING] AllowCredentials will be set to [%v]", allowCredentials)))
	}
	app.Use(cors.New(cors.Config{
		//Next:             nil,
		//AllowOriginsFunc: nil,
		AllowOrigins:     allowOrigins,
		AllowMethods:     strings.Join(allowMethod, ","),
		AllowHeaders:     strings.Join(allowHeaders, ","),
		AllowCredentials: allowCredentials,
		//ExposeHeaders:    "",
		MaxAge: 12 * 3600,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)
	log.Println("Swagger is enabled and available at /swagger/index.html")

	app.Use(recovery.New)
	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Set(fiber.HeaderCacheControl, "no-cache,no-store")
		ctx.Set(fiber.HeaderExpires, "0")
		ctx.Set(fiber.HeaderPragma, "no-cache")
		return ctx.Next()
	})

	route_ := app.Group(conf.App.Config.Path)
	//Set route to app
	if debug {
		fmt.Println("  [API]")
	}
	for _, r := range routes(conf) {
		route_ = route_.Add(r.Method, r.Path, append(r.Middleware, r.HandlerFunc)...)
		if debug {
			fmt.Printf("  %s %s%s\n", color.Blue(fmt.Sprintf("[%s]", r.Method)), color.Green(fmt.Sprintf("- %s%s", conf.App.Config.Path, r.Path)), color.Yellow(fmt.Sprintf("(%s)", runtime.FuncForPC(reflect.ValueOf(r.HandlerFunc).Pointer()).Name())))
		}
	}

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGINT)

	await := make(chan struct{})
	go func() {
		_ = <-osSignal
		fmt.Println("\nApp shutting down . . .")
		if err := app.Shutdown(); err != nil {
			panic(err)
		}
		await <- struct{}{}
	}()

	fmt.Printf("App running on %s:%d\n", conf.App.Config.Host, conf.App.Config.Port)
	if err := app.Listen(fmt.Sprintf("%s:%d", conf.App.Config.Host, conf.App.Config.Port)); err != nil {
		panic(err)
	}

	fmt.Println("App running clear task . . .")
	<-await

	fmt.Println("App close")
}

func routes(conf *config.Config) []*route.Route {

	//connector
	gormClient := gormDB.New(conf.Database.Host, conf.Database.Port, conf.Database.Username, conf.Database.Password, conf.Database.Name)
	if err := gormClient.Connect(mysql.Open); err != nil {
		panic(err)
	}

	//new repository
	repositories := repository.New(gormClient)

	//new ping service
	_sPing := sPing.New(conf)
	_sGame := sGame.New(conf, repositories.PTTEPVR())
	_sUser := sUser.New(conf, repositories.PTTEPVR())
	_sAuthentication := sAuthentication.New(conf, repositories.PTTEPVR())
	_sPermission := sPermission.New(conf, repositories.PTTEPVR())
	_sRole := sRole.New(conf, repositories.PTTEPVR())

	//new service and other package
	//service_ := service1.New(context.Background(), repositories)
	//if err != nil {
	//	panic(err)
	//}

	//add routes
	var list []*route.Route

	//TODO: add handler to route
	list = append(list, hPing.Route(_sPing)...)
	list = append(list, hGame.Route(_sGame)...)
	list = append(list, hAuthentication.Route(_sAuthentication)...)
	list = append(list, hRegistration.Route(_sUser)...)
	list = append(list, hPermission.Route(_sPermission)...)
	list = append(list, hRole.Route(_sRole)...)
	list = append(list, hUser.Route(_sUser)...)
	list = append(list, hVersion.Route(conf)...)

	return list
}
