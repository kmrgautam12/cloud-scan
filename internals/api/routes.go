package api

import (
	"distributed-job-scheduler/jobqueue/internals/api/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, handler *handlers.JobHandler) {
	v1 := e.Group("/api/v1")
	v1.POST("/jobs", handler.CreateJobs)
}
