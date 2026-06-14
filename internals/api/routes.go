package api

import (
	"distributed-job-scheduler/jobqueue/internals/api/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	// e.GET("/health", Health)
	v1 := e.Group("/api/v1")
	v1.POST("/jobs", handlers.CreateJobs)
	// v1.GET("/jobs/:id", GetJobs)
	// v1.GET("/jobs", ListJobs)
}
