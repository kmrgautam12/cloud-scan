package handlers

import (
	"distributed-job-scheduler/jobqueue/internals/service"
	"distributed-job-scheduler/jobqueue/utils/structure"
	"fmt"

	"github.com/labstack/echo/v4"
)

func CreateJobs(e echo.Context) error {
	var job structure.CreateJobs
	err := e.Bind(&job)
	if err != nil {
		e.Logger().Error(fmt.Errorf("Error binding request %v", err))
		structure.Resp400WithMessage(e, "Error binding input request")
		return nil
	}
	s := service.NewService()
	s.CreateJobs()
	structure.Resp200WithMessage(e, "Job added to queue")
	return nil
}
