package handlers

import (
	"api/zeus/models"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {

	users, err := models.GetAllUsers(ctx)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatus(http.StatusNotFound)
	} else if users != nil {
		ctx.JSON(http.StatusOK, gin.H{"result": users})
	} else {
		ctx.JSON(http.StatusOK, json.RawMessage(`[]`))
	}

}

func GetUserByID(ctx *gin.Context) {

	users, err := models.GetUserByID(ctx)
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "non existant ID",
		})
	} else if len(users) < 1 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "ID not found",
		})
	} else {
		ctx.JSON(http.StatusOK, users)
	}

}

func CreateUser(ctx *gin.Context) {

	user, err := models.CreateUser(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Null value") {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Don't accept null value",
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "fail creating user",
			})
		}

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}

}

func UpdateUserByID(ctx *gin.Context) {

	user, err := models.UpdateUserByID(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Null value") {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Don't accept null value",
			})
		} else {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "fail updating user",
			})
		}

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"user":    user,
			"message": "user updated with success",
		})
	}
}

func DeleteUserByID(ctx *gin.Context) {

	err := models.DeleteUserByID(ctx)

	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "Non existant ID") {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "User deleted with success",
		})
	}
}

func PartiallyUpdatingUser(ctx *gin.Context) {

	user, err := models.PartiallyUpdatingUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "fail updating user",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"user":    user,
			"message": "user updated with success",
		})
	}
}
