package handlers

import (
	"github.com/alaraiabdiallah/media-service/app"
	"github.com/alaraiabdiallah/media-service/data_sources/mongods"
	"github.com/alaraiabdiallah/media-service/helpers"
	"github.com/alaraiabdiallah/media-service/models"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetAllMedia(c echo.Context) error {
	var results interface{}
	only_link := false
	if par := c.QueryParam("only-link"); par == "true" {only_link = true}
	filter := models.MediaFilter{OnlyLink: only_link, Query: echo.Map{}}
	if err := app.GetAllMedia(filter,&results); err != nil {return err}
	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"message": "",
		"data": results,
	})
}

func ShowMediaFile(c echo.Context) error {
	var result models.MediaDS
	media_id := c.Param("media_id")
	id, _ :=primitive.ObjectIDFromHex(media_id)
	filter := echo.Map{"_id": id}
	if err := mongods.FindMedia(filter,&result); err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(http.StatusNotFound, helpers.FailedJsonMessage("Not Found"))
		}
		return err
	}
	return c.File(result.Filepath)
}

func UploadMedia(c echo.Context) error {
	file, err := c.FormFile("file")
	var file_data models.MediaDS
	if file == nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedJsonMessage("File body not defined"))
	}
	if err != nil { return c.JSON(http.StatusInternalServerError, helpers.FailedJsonMessage(err))}
	if err := app.SaveMedia(file, &file_data); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedJsonMessage(err))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
		"message": "Successfully to upload media",
		"data": file_data.Id,
	})
}