package auth

import (
	"slices"

	"github.com/anonychun/ecorp/internal/consts"
	"github.com/anonychun/ecorp/internal/current"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) AuthenticateAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bypassedPaths := []string{
			"/api/v1/admin/auth/login",
		}

		if slices.Contains(bypassedPaths, c.Request().URL.Path) {
			return next(c)
		}

		cookie, err := c.Cookie(consts.CookieAdminSession)
		if err != nil {
			return consts.ErrUnauthorized
		}

		adminSession, err := m.repository.AdminSession.FindByToken(c.Request().Context(), cookie.Value)
		if err != nil {
			return consts.ErrUnauthorized
		}

		admin, err := m.repository.Admin.FindById(c.Request().Context(), adminSession.AdminId.String())
		if err != nil {
			return consts.ErrUnauthorized
		}

		ctx := current.SetAdmin(c.Request().Context(), admin)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
