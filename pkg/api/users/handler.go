package user

import (
	"net/http"
	"strconv"

	goerrors "github.com/go-errors/errors"
	"github.com/labstack/echo"
	utility "github.com/malkhandi-anibrata-tft/self.best-temp/pkg/api/utility"
)

type Handler interface {
	CreateUserInfo(c echo.Context) error
	GetUserInfo(c echo.Context) error
	UpdateDetails(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateUserInfo(c echo.Context) error {
	req := CreateDetailsRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utility.CreateErrorResponse(goerrors.New(err)))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utility.CreateErrorResponse(goerrors.New(err)))
	}

	approved, err := h.service.CreateUser(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.CreateErrorResponse(err))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"Approved": approved,
	})
}

func (h *handler) GetUserInfo(c echo.Context) error {
	userid, err := strconv.ParseUint(c.QueryParam("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.CreateErrorResponse(goerrors.New(err)))
	}

	resp, err := h.service.GetUserInfo(userid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.CreateErrorResponse(err))
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) UpdateDetails(c echo.Context) error {
	req := UpdateDetailsRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utility.CreateErrorResponse(goerrors.New(err)))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utility.CreateErrorResponse(goerrors.New(err)))
	}
	resp, err := h.service.UpdateDetails(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.CreateErrorResponse(err))
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) DeleteUser(c echo.Context) error {

	req := DeleteUserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utility.CreateErrorResponse(goerrors.New(err)))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utility.CreateErrorResponse(goerrors.New(err)))
	}

	if err := h.service.DeleteUser(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utility.CreateErrorResponse(err))
	}

	return c.NoContent(http.StatusOK)
}
