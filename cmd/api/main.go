package main

import (
	"distributed-job-scheduler/jobqueue/internals/api"
	"distributed-job-scheduler/jobqueue/internals/api/handlers"
	"distributed-job-scheduler/jobqueue/internals/database"
	"distributed-job-scheduler/jobqueue/internals/repository"
	"distributed-job-scheduler/jobqueue/internals/service"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, err := database.InitialiseSchedulerStore()
	if err != nil {
		panic(err)
	}
	repo := repository.NewJobRepository(db)
	service := service.NewJobService(repo)
	handler := handlers.NewJobHandler(service)
	api.RegisterRoutes(e, handler)

	e.Logger.Fatal(e.Start(":8080"))
}
