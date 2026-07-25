package handlers

import (
	"distributed-job-scheduler/jobqueue/utils/structure"
	"errors"

	"github.com/labstack/echo/v4"
)

func (h *JobHandler) UploadFile(e echo.Context) error {
	var upload structure.UploadFile
	err := e.Bind(&upload)
	if err != nil {
		structure.Resp400WithMessage(e, err)
	}
	if upload.FileName == "" {
		structure.Resp400WithMessage(e, errors.New("filename can't be empty"))
		return nil
	}
	presignUrl, err := h.service.UploadFile(upload)
	if err != nil {
		structure.Resp500WithMessage(e, err)
		return nil
	}
	respBody := &structure.UploadFileResponse{
		FileName:   upload.FileName,
		PresignUrl: presignUrl,
	}
	structure.Resp200WithMessage(e, respBody)
	return nil
}
