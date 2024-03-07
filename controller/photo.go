package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-api/database"
	"golang-api/helpers"
	"golang-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ListPhoto(c *gin.Context) {
	var photos []models.Photo
	db := database.GetDB()
	err := db.Preload("User").Find(&photos).Error

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": photos,
	})
}

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	var Photo models.Photo
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		if err := c.ShouldBindJSON(&Photo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&Photo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err":     "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	Photo.UserID = userID

	if err := db.Debug().Create(&Photo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("photoId")

	var Photo models.Photo
	if err := db.First(&Photo, id).Error; err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var NewPhoto models.Photo
	if err := json.NewDecoder(c.Request.Body).Decode(&NewPhoto); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Photo.Title = NewPhoto.Title
	Photo.Caption = NewPhoto.Caption
	Photo.PhotoUrl = NewPhoto.PhotoUrl

	if err := db.Save(&Photo).Error; err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"updated_at": Photo.UpdatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("photoId")

	var Photo models.Photo
	if err := db.First(&Photo, id).Error; err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Delete(&Photo).Error; err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
