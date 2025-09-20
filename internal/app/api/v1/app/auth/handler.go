package auth

import (
	"net/http"

	"github.com/anonychun/ecorp/internal/api"
	"github.com/anonychun/ecorp/internal/consts"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SignUp(c echo.Context) error {
	req := SignUpRequest{
		IpAddress: c.RealIP(),
		UserAgent: c.Request().UserAgent(),
	}

	err := c.Bind(&req)
	if err != nil {
		return err
	}

	res, err := h.usecase.SignUp(c.Request().Context(), req)
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     consts.CookieUserSession,
		Value:    res.Token,
		Path:     "/",
		HttpOnly: true,
	})

	return api.NewResponse(c).SendOk()
}

func (h *Handler) SignIn(c echo.Context) error {
	req := SignInRequest{
		IpAddress: c.RealIP(),
		UserAgent: c.Request().UserAgent(),
	}

	err := c.Bind(&req)
	if err != nil {
		return err
	}

	res, err := h.usecase.SignIn(c.Request().Context(), req)
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     consts.CookieUserSession,
		Value:    res.Token,
		Path:     "/",
		HttpOnly: true,
	})

	return api.NewResponse(c).SendOk()
}

func (h *Handler) SignOut(c echo.Context) error {
	cookie, err := c.Cookie(consts.CookieUserSession)
	if err != nil {
		return err
	}

	req := SignOutRequest{
		Token: cookie.Value,
	}

	err = h.usecase.SignOut(c.Request().Context(), req)
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     consts.CookieUserSession,
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
