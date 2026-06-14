package main

import (
	"distributed-job-scheduler/jobqueue/internals/api"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	api.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
