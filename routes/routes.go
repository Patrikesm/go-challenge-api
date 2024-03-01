package routes

import (
	"challenge-week-one/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()

	r.GET("/depoimentos", controllers.GetAll)
	r.POST("/depoimentos", controllers.Create)
	r.PATCH("/depoimentos/:id", controllers.Update)
	r.DELETE("/depoimentos/:id", controllers.Delete)

	r.GET("/depoimentos-home", controllers.GetRandomDeclaration)

	r.Use(cors.Default())
	r.Run(":8080")
}
