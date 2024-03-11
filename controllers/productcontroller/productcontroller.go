package productcontroller

import (
	"net/http"

	"github.com/Mikael88/go-digitalent-assignment-2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&order)
	c.JSON(http.StatusOK, gin.H{"order": order})
}

func GetOrders(c *gin.Context) {
	var orders []models.Order

	models.DB.Preload("Items").Find(&orders)
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func GetOrderByID(c *gin.Context) {
	var order models.Order
	id := c.Param("orderId")

	if err := models.DB.Preload("Items").First(&order, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Order not found"})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func UpdateOrder(c *gin.Context) {
	var order models.Order
	id := c.Param("orderId")

	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&order).Where("id = ?", id).Updates(&order).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("orderId")
	if models.DB.Delete(&models.Order{}, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to delete order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
