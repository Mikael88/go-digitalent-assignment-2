package ordercontroller

import (
	"net/http"
	"strconv"

	"github.com/Mikael88/go-digitalent-assignment-2/models"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if result := models.DB.Create(&order); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"order": order})
}

func GetOrders(c *gin.Context) {
    var orders []models.Order
    if err := models.DB.Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func UpdateOrder(c *gin.Context) {
    orderId, err := strconv.ParseUint(c.Param("orderId"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }

    var updatedOrder models.Order
    if err := c.ShouldBindJSON(&updatedOrder); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var existingOrder models.Order
    if err := models.DB.First(&existingOrder, orderId).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    existingOrder.CustomerName = updatedOrder.CustomerName
    existingOrder.OrderedAt = updatedOrder.OrderedAt

    if err := models.DB.Save(&existingOrder).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"order": existingOrder})
}

func DeleteOrder(c *gin.Context) {
    orderId, err := strconv.ParseUint(c.Param("orderId"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }

    var order models.Order
    if err := models.DB.Preload("Items").First(&order, orderId).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    for _, item := range order.Items {
        if err := models.DB.Delete(&item).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    }

    if err := models.DB.Delete(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
