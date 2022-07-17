package controllers

import (
	"github.com/labstack/echo/v4"
	"go-checkin/models"
	"go-checkin/service"
	"go-checkin/utils/session"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type AttendanceController struct {
	BaseBackendController
	service *service.AttendanceService
}

func NewAttendanceController(service *service.AttendanceService) AttendanceController {
	return AttendanceController{
		BaseBackendController: BaseBackendController{
			Menu:        "Attendance",
			BreadCrumbs: []map[string]interface{}{},
		},
		service: service,
	}
}

func (c *AttendanceController) PhotoIndex(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Photo",
		"link": "/check/photo",
	}
	userInfo, _ := session.Manager.Get(ctx, session.SessionId)
	return Render(ctx, "Home", "attendance/take_photo", c.Menu, session.GetFlashMessage(ctx),
		append(c.BreadCrumbs, breadCrumbs), userInfo)
}

func (c *AttendanceController) CheckIndex(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Check",
		"link": "/check/check",
	}
	userInfo, _ := session.Manager.Get(ctx, session.SessionId)
	return Render(ctx, "Home", "attendance/check_in", c.Menu, session.GetFlashMessage(ctx),
		append(c.BreadCrumbs, breadCrumbs), userInfo)
}

func (c *AttendanceController) Checkin(ctx echo.Context) error {
	var req models.Checkin
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}

	if err := ctx.Validate(&req); err != nil {
		var validationErrors []models.ValidationError
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors = models.WrapValidationErrors(errs)
		}
		return ctx.JSON(400, echo.Map{"message": "error validation", "errors": validationErrors})
	}
	// Get image name by user id

	// Face Compare
	isMatch, err, photo := c.service.PhotoCompare(req)
	if err != nil || !isMatch {
		return ctx.JSON(400, echo.Map{"message": "Not Matched", "errors": err})
	}

	// Save attendance
	err = c.service.Save(req, *photo)
	if err != nil {
		return ctx.JSON(400, echo.Map{"message": "Failed save to database", "errors": err})

	}

	return ctx.JSON(http.StatusOK, "BERHASIL")
}
