package controllers

import (
	"challenge-week-one/database"
	"challenge-week-one/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	var declarations []models.Declaration

	database.DB.Find(&declarations)

	if len(declarations) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Declarations not found",
		})
		return
	}

	c.JSON(http.StatusOK, declarations)
}

func Create(c *gin.Context) {
	var declaration models.Declaration

	err := c.ShouldBindJSON(&declaration)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error in field binding",
		})
		return
	}

	database.DB.Create(&declaration)

	c.JSON(http.StatusOK, declaration)
}

func Update(c *gin.Context) {
	id := c.Params.ByName("id")

	var declaration models.Declaration

	database.DB.First(&declaration, id)

	if declaration.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Student not found",
		})
		return
	}

	var declarationUpdated models.Declaration

	err := c.ShouldBindJSON(&declarationUpdated)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	declaration.Author = declarationUpdated.Author
	declaration.Testimony = declarationUpdated.Testimony
	declaration.Image = declarationUpdated.Image

	database.DB.Save(&declaration)

	c.JSON(http.StatusOK, declaration)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	declaration := GetOneDeclarationById(id)

	database.DB.Delete(&declaration)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Declaration deleted successfully",
	})
}

func GetRandomDeclaration(c *gin.Context) {
	var declarations []models.Declaration

	database.DB.Limit(3).Order("random()").Find(&declarations)

	c.JSON(http.StatusOK, declarations)
}

func GetOneDeclarationById(id string) models.Declaration {
	var declaration models.Declaration

	database.DB.First(&declaration, id)

	return declaration
}
