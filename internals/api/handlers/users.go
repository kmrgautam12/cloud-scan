package handlers

import (
	"distributed-job-scheduler/jobqueue/utils/structure"

	"github.com/labstack/echo/v4"
)

func (h *JobHandler) LoginUser(c echo.Context) error {
	body := structure.LoginUser{}
	err := c.Bind(&body)
	if err != nil {
		structure.Resp400WithMessage(c, err)
		return nil
	}
	token, err := h.service.LoginUser(body)
	if err != nil {
		structure.Resp500WithMessage(c, err)
		return nil
	}
	structure.Resp200WithMessage(c, token)
	return nil
}

func (h *JobHandler) CreateUser(c echo.Context) error {
	body := structure.CreateUser{}
	err := c.Bind(&body)
	if err != nil {
		structure.Resp400WithMessage(c, err)
		return nil
	}
	err = h.service.CreateUser(body)
	if err != nil {
		structure.Resp500WithMessage(c, err)
		return nil
	}
	structure.Resp200WithMessage(c, "user created successfully")
	return nil
}
