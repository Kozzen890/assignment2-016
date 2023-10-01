package orderControllers

import (
	"fmt"
	"net/http"

	"github.com/Kozzen890/assignment2-016/database"
	"github.com/Kozzen890/assignment2-016/models"
	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	var data []models.Order
	db := database.GetDatabase()

	if err := db.Preload("Items").Find(&data).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func UpdateOrderById(context *gin.Context) {
	db := database.GetDatabase()
	id := context.Param("id")

	var order models.Order

	err := db.Preload("Items").First(&order, id).Error

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": fmt.Sprintf("user with id %v not found", id),
		})
		return
	}

	if err = context.ShouldBindJSON(&order); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	for _, item := range order.Items {
		updatedItem := models.Item {
			ItemCode: item.ItemCode,
			Description: item.Description,
			Quantity: item.Quantity,
		}

		err = db.Model(&item).Where("id = ?", item.ID).Updates(updatedItem).Error

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
			return
		}
	}

	err = db.Model(&order).Where("id = ?", id).Updates(order).Error

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": order,
	})
}

func CreateOrders(c *gin.Context) {
	var order models.Order
	db := database.GetDatabase()

	if err := c.ShouldBindJSON(&order); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
  }

  if err := db.Create(&order).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusCreated, order)
}

func DeleteOrderById(context *gin.Context) {
	db := database.GetDatabase()
	id := context.Param("id")

	var order models.Order

	err := db.Preload("Items").First(&order, id).Error

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": fmt.Sprintf("user with id %v not found", id),
		})
		return
	}

	for _, item := range order.Items {
		db.Delete(&item)
	}

	db.Delete(&order)

	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("order with id %v has been successfully deleted", id),
	})
}