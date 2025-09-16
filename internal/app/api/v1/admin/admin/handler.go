package admin

import (
	"net/http"

	"github.com/anonychun/ecorp/internal/api"
	"github.com/labstack/echo/v4"
)

func (h *Handler) FindAll(c echo.Context) error {
	res, err := h.usecase.FindAll(c.Request().Context())
	if err != nil {
		return err
	}

	return api.NewResponse(c).SetData(res).Send()
}

func (h *Handler) FindById(c echo.Context) error {
	var req FindByIdRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	res, err := h.usecase.FindById(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return api.NewResponse(c).SetData(res).Send()
}

func (h *Handler) Create(c echo.Context) error {
	var req CreateRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	res, err := h.usecase.Create(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return api.NewResponse(c).SetStatus(http.StatusCreated).SetData(res).Send()
}

func (h *Handler) Update(c echo.Context) error {
	var req UpdateRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	res, err := h.usecase.Update(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return api.NewResponse(c).SetData(res).Send()
}

func (h *Handler) Delete(c echo.Context) error {
	var req DeleteRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = h.usecase.Delete(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
