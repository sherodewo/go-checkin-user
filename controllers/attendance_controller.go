package controllers

import (
	"github.com/labstack/echo/v4"
	"go-checkin/models"
	"go-checkin/service"
	"go-checkin/utils/session"
	"gopkg.in/go-playground/validator.v9"
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

func (c *AttendanceController) PhotoCheck(ctx echo.Context) error {
	var req models.PhotoRequest
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
	// Get User Session
	result, err := session.Manager.Get(ctx, session.SessionId)
	if err != nil {
		return ctx.JSON(400, echo.Map{"message": "error get session", "errors": err})
	}
	userInfo := result.(session.UserInfo)

	// Compare 2 Face
	isSimiliar, err := c.service.PhotoCompare(req, userInfo.UserID)
	if err != nil {
		return ctx.JSON(400, echo.Map{"message": "error Compare photo", "errors": err})
	}

	if isSimiliar {
		return ctx.JSON(200, isSimiliar)
	}
	return nil
}

func (c *AttendanceController) Checkin(ctx echo.Context) error {
	return nil
}
