package handlers

import (
	"distributed-job-scheduler/jobqueue/internals/service"
	"distributed-job-scheduler/jobqueue/utils/structure"

	"github.com/labstack/echo/v4"
)

type JobHandler struct {
	service *service.JobService
}

func NewJobHandler(service *service.JobService) *JobHandler {
	return &JobHandler{
		service: service,
	}
}
func (h *JobHandler) CreateJobs(c echo.Context) error {
	body := structure.CreateJobs{}
	err := c.Bind(&body)
	if err != nil {
		structure.Resp400WithMessage(c, err.Error())
		return nil
	}
	h.service.CreateJobs(body)
	return c.JSON(200, "ok")
}
