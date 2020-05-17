package http

import (
	"fmt"
	"github.com/alaraiabdiallah/media-service/protocols/http/routers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func Run(){
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	routers.HttpRouters(e)
	port := ":2807"
	if envport := os.Getenv("HTTP_PROT_PORT"); envport != "" {
		port = fmt.Sprintf(":%v",envport)
	}
	e.Logger.Fatal(e.Start(port))
}