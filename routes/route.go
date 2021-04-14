package routes

import (
	"example.com/sumit/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes() {
	r := gin.Default()
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook) // new
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()
}
