package server

import (
	"github.com/anonychun/ecorp/internal/api"
	"github.com/anonychun/ecorp/internal/app"
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/anonychun/ecorp/internal/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/samber/do"
)

func namespace(e *echo.Group, path string, f func(e *echo.Group)) {
	f(e.Group(path))
}

func routes(e *echo.Echo) error {
	m := do.MustInvoke[*middleware.Middleware](bootstrap.Injector)
	h := do.MustInvoke[*app.Handler](bootstrap.Injector)

	e.HTTPErrorHandler = api.HttpErrorHandler
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.Logger())

	apiRouter := e.Group("/api")
	namespace(apiRouter, "/v1", func(e *echo.Group) {
		namespace(e, "/admin", func(e *echo.Group) {
			e.Use(m.Auth.AuthenticateAdmin)

			e.POST("/auth/login", h.Api.V1.Admin.Auth.Login)
			e.POST("/auth/logout", h.Api.V1.Admin.Auth.Logout)
			e.GET("/auth/me", h.Api.V1.Admin.Auth.Me)

			e.GET("/admin", h.Api.V1.Admin.Admin.FindAll)
			e.GET("/admin/:id", h.Api.V1.Admin.Admin.FindById)
			e.POST("/admin", h.Api.V1.Admin.Admin.Create)
			e.PUT("/admin/:id", h.Api.V1.Admin.Admin.Update)
			e.DELETE("/admin/:id", h.Api.V1.Admin.Admin.Delete)
		})

		namespace(e, "/app", func(e *echo.Group) {
		})

		namespace(e, "/landing", func(e *echo.Group) {
		})

	})

	return nil
}
