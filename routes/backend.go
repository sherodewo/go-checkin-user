package routes

import (
	"errors"
	"github.com/foolin/goview/supports/echoview-v4"
	"go-checkin/config"
	"html/template"
	"strconv"
	"time"

	"github.com/foolin/goview"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-checkin/controllers"
	"go-checkin/middleware"
	"go-checkin/models"
	"go-checkin/utils"
	"go-checkin/utils/session"
	"gorm.io/gorm"
)

func BackendRoute(e *echo.Echo, db *gorm.DB) {

	//=========== Backend ===========//
	var userInfo session.UserInfo
	//new middleware
	mv := echoview.NewMiddleware(goview.Config{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/master",
		Partials: []string{
			"partials/asside",
			"partials/brand-menu",
			"partials/chart",
			"partials/chat",
			"partials/demo-panel",
			"partials/header-mobile",
			"partials/language",
			"partials/notification",
			"partials/quick-action",
			"partials/quick-panel",
			"partials/quick-panel-toogle",
			"partials/scrolltop",
			"partials/search",
			"partials/sticky-toolbar",
			"partials/sub-header",
			"partials/user-bar",
		},
		Funcs: template.FuncMap{
			"user": func(ctx echo.Context) models.User {
				userModel, err := utils.GetUserInfoFromContext(ctx, db)
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						return models.User{}
					}
					return models.User{}
				}
				return userModel
			},
			"floatToString": func(value *float64) string {
				if value == nil {
					return ""
				}
				return strconv.FormatFloat(*value, 'f', -1, 64)
			},
			"floatNotPointerToString": func(value float64) string {
				return strconv.FormatFloat(value, 'f', -1, 64)
			},
			"formatDate": func(date time.Time, layout string) string {
				return date.Format(layout)
			},
			"maritalStatus": func(data string) string {
				var result string
				if data == "S" {
					result = "Belum Pernah Menikah"
				}
				return result
			},
			"MenuParent": func(ctx echo.Context) []map[string]interface{} {
				var dataMenu map[string]interface{}
				var listOfMenu []map[string]interface{}
				result, err := session.Manager.Get(ctx, session.SessionId)
				if err != nil {
					panic(err)
				}
				userInfo = result.(session.UserInfo)
				//var menu models.Menu
				//var menuRole []models.MenuRole
				//if err := db.Raw("select * from web_menu_role where user_role_id = ? and active = ?", userInfo.UserRoleID, 1 ).
				//Scan(&menuRole); err != nil {
				//}

				var listParentMenu []models.Menu
				if err := db.Raw("select * from web_menu where is_active = ? and menu_type = ? ",
					2, "parent_menu").Scan(&listParentMenu).Error; err != nil {
				}

				for _, v := range listParentMenu {
					var menus []models.Menu
					if err := db.Raw("select * from web_menu where parent_id = ? and is_active = ? ",
						v.ID, 1).Scan(&menus).Error; err != nil {
						log.Info("ERROR GET MENU BY ROLE ", err.Error())
					}
					dataMenu = map[string]interface{}{
						"Name":  v.Name,
						"Icon":  v.IconClass,
						"Menus": menus,
					}
					listOfMenu = append(listOfMenu, dataMenu)
				}

				return listOfMenu
			},
			"Menu": func(ctx echo.Context) []map[string]interface{} {
				var dataMenu map[string]interface{}
				var listOfMenu []map[string]interface{}
				result, err := session.Manager.Get(ctx, session.SessionId)
				if err != nil {
					panic(err)
				}
				userInfo = result.(session.UserInfo)

				var menus []models.Menu
				if err := db.Raw("select * from web_menu where parent_id = ? and is_active = ? and menu_type = ? ",
					0, 2, "menu").Scan(&menus).Error; err != nil {
					log.Info("ERROR GET MENU BY ROLE ", err.Error())
				}
				dataMenu = map[string]interface{}{
					"Menus": menus,
				}
				listOfMenu = append(listOfMenu, dataMenu)
				return listOfMenu
			},
		},
		DisableCache: true,
	})

	// You should use helper func `Middleware()` to set the supplied
	// TemplateEngine and make `HTML()` work validly.
	backendGroup := e.Group("/check", mv, middleware.SessionMiddleware(session.Manager))
	authorizationMiddleware := middleware.NewAuthorizationMiddleware(db)

	var menus []models.Menu
	if err := db.Raw("select * from web_menu where is_active = ? ",
		1).Scan(&menus).Error; err != nil {
		log.Info("ERROR GET MENU BY ROLE ", err.Error())
	}

	homeController := controllers.NewHomeController()
	backendGroup.GET("/home", homeController.Index)

	authController := config.InjectAuthController(db)
	backendGroup.POST("/logout", authController.Logout)

	//userController
	userController := config.InjectUserController(db)
	userGroup := backendGroup.Group("/register")
	userGroup.GET("", userController.Index)
	userGroup.POST("/store", userController.Store)
	userGroup.GET("/add", userController.Add)
	userGroup.GET("/profile", userController.Profile)
	userGroup.GET("/datatable", userController.List)
	userGroup.DELETE("/delete/:id", userController.Delete)
	userGroup.GET("/detail/:id", userController.View)
	userGroup.GET("/edit/:id", userController.Edit)
	userGroup.POST("/update/:id", userController.Update)

	//ConfigController
	configController := config.InjectConfigController(db)
	configGroup := backendGroup.Group("/config", authorizationMiddleware.AuthorizationMiddleware(menus, "config"))
	configGroup.GET("", configController.Index)
	configGroup.POST("/store", configController.Store)
	configGroup.GET("/datatable", configController.Datatable)
	configGroup.POST("/update/:id", configController.Update)
	configGroup.GET("/:id", configController.Detail)
	configGroup.POST("/delete/:id", configController.Delete)
	backendGroup.POST("/check/config/set-active/:id", configController.SetActive)
	backendGroup.POST("/check/config/set-inactive/:id", configController.SetInactive)

	attendanceController := config.InjectAttendanceController(db)
	attendanceGroup := backendGroup.Group("/attend")
	attendanceGroup.GET("/photo", attendanceController.PhotoIndex)
	attendanceGroup.GET("/checkin", attendanceController.CheckIndex)
	attendanceGroup.POST("/checkin", attendanceController.Checkin)
	attendanceGroup.GET("/datatable", attendanceController.Datatable)

}
