package auth

import (
	"net/http"

	"github.com/anonychun/ecorp/internal/api"
	"github.com/anonychun/ecorp/internal/consts"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(c echo.Context) error {
	req := LoginRequest{
		IpAddress: c.RealIP(),
		UserAgent: c.Request().UserAgent(),
	}
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	res, err := h.usecase.Login(c.Request().Context(), req)
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     consts.CookieAdminSession,
		Value:    res.Token,
		Path:     "/",
		HttpOnly: true,
	})

	return api.NewResponse(c).SetData(res).Send()
}

func (h *Handler) Logout(c echo.Context) error {
	cookie, err := c.Cookie(consts.CookieAdminSession)
	if err != nil {
		return err
	}

	req := LogoutRequest{
		Token: cookie.Value,
	}

	err = h.usecase.Logout(c.Request().Context(), req)
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     consts.CookieAdminSession,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) Me(c echo.Context) error {
	res, err := h.usecase.Me(c.Request().Context())
	if err != nil {
		return err
	}

	return api.NewResponse(c).SetData(res).Send()
}
