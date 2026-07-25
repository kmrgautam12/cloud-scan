package api

import (
	"distributed-job-scheduler/jobqueue/internals/api/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, handler *handlers.JobHandler) {
	v1 := e.Group("/api/v1")
	jobs := v1.Group("/jobs")
	jobs.POST("/jobs", handler.CreateJobs)

	// User apis
	user := v1.Group("/user")
	user.POST("", handler.CreateUser)
	user.POST("/login", handler.LoginUser)

	//submit to s3
	v1.Use()

	upload := v1.Group("/upload")
	upload.POST("/", handler.UploadFile)
}
