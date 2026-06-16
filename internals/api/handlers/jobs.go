package handlers

import (
	"distributed-job-scheduler/jobqueue/internals/service"

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
	return c.JSON(200, "ok")
}
