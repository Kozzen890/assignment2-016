package routes

import (
	orderControllers "github.com/Kozzen890/assignment2-016/controllers"
	"github.com/Kozzen890/assignment2-016/database"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	database.StartDB()

	router := gin.Default()

	router.POST("/order", orderControllers.CreateOrders)
	router.GET("/orders", orderControllers.GetOrders)
	router.PUT("/order/:id", orderControllers.UpdateOrderById)
	router.DELETE("/order/:id", orderControllers.DeleteOrderById)

	return router
}