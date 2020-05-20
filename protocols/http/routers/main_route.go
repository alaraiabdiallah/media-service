package routers

import (
	"fmt"
	"github.com/alaraiabdiallah/media-service/app"
	"github.com/alaraiabdiallah/media-service/protocols/http/handlers"
	"github.com/alaraiabdiallah/media-service/protocols/http/middlewares"
	"github.com/labstack/echo"
	"net/http"
)

func HttpRouters(e *echo.Echo){
	e.GET("/", func(c echo.Context) error {
		version := fmt.Sprintf("App version %v", app.Ctx().GetVersion())
		return c.String(http.StatusOK, version)
	})

	e.GET("/media", handlers.GetAllMedia, middlewares.APIAuth)
	e.GET("/media/:media_id", handlers.ShowMediaFile)
	e.POST("/media", handlers.UploadMedia, middlewares.APIAuth)
}
