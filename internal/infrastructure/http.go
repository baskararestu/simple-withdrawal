package infrastructure

import (
	"fmt"

	"simple-withdraw-api/internal/balance"
	"simple-withdraw-api/internal/docs"
	"simple-withdraw-api/internal/user"
	"simple-withdraw-api/internal/utilities/tools"
	"simple-withdraw-api/internal/withdrawal"
	"simple-withdraw-api/pkg/xlogger"
	"strings"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Run() {
	logger := xlogger.Logger

	app := fiber.New(fiber.Config{
		ProxyHeader:           cfg.ProxyHeader,
		DisableStartupMessage: true,
		ErrorHandler:          defaultErrorHandler,
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
		Fields: cfg.LogFields,
	}))
	app.Use(recover2.New())
	app.Use(etag.New())
	app.Use(requestid.New())

	api := app.Group("/api")
	docs.NewHttpHandler(api.Group("/docs"))
	withdrawal.NewWithdrawalHandler(api.Group("/withdraw"), WithdrawalSvc, cfg.SecretKey)
	user.NewUserHttpHandler(api.Group("/user"), UserSvc)
	tools.NewToolsHttpHandler(api.Group("/tools"), cfg.SecretKey)
	balance.NewBalanceHttpHandler(api.Group("/balance"),BalanceSvc, cfg.SecretKey)


	logger.Info().Msg("Registered Routes:")
	for _, routes := range app.Stack() {
		for _, route := range routes {
            if strings.HasPrefix(route.Path, "/api") && !strings.HasPrefix(route.Path, "/api/docs") {
                logger.Info().Msgf("Method: %s, Path: %s", route.Method, route.Path)
            }
		}
	}

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	logger.Info().Msgf("Server is running on address: %s", addr)
	if err := app.Listen(addr); err != nil {
		logger.Fatal().Err(err).Msg("Server failed to start")
	}
}
