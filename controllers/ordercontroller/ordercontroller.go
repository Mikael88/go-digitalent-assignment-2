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
    if err := models.DB.Preload("Items").Find(&orders).Error; err != nil {
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
    if err := models.DB.Preload("Items").First(&existingOrder, orderId).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    if err := updateOrderItems(&existingOrder, updatedOrder.Items); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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


// Helper functions
func updateOrderItems(order *models.Order, updatedItems []models.Item) error {
    for i, requestItem := range updatedItems {
        if i < len(order.Items) {
            if err := updateItem(&order.Items[i], requestItem); err != nil {
                return err
            }
        } else {
            if err := createNewItem(order, requestItem); err != nil {
                return err
            }
        }
    }

    if len(updatedItems) < len(order.Items) {
        if err := deleteRemainingItems(order, updatedItems); err != nil {
            return err
        }
    }

    return nil
}

func updateItem(item *models.Item, updatedItem models.Item) error {
    item.ItemCode = updatedItem.ItemCode
    item.Description = updatedItem.Description
    item.Quantity = updatedItem.Quantity

    if err := models.DB.Save(item).Error; err != nil {
        return err
    }

    return nil
}

func createNewItem(order *models.Order, requestItem models.Item) error {
    newItem := models.Item{
        ItemCode:    requestItem.ItemCode,
        Description: requestItem.Description,
        Quantity:    requestItem.Quantity,
        OrderID:     order.ID,
    }

    if err := models.DB.Create(&newItem).Error; err != nil {
        return err
    }

    order.Items = append(order.Items, newItem)

    return nil
}

func deleteRemainingItems(order *models.Order, updatedItems []models.Item) error {
    for i := len(updatedItems); i < len(order.Items); i++ {
        item := &order.Items[i]
        if err := models.DB.Delete(item).Error; err != nil {
            return err
        }
    }

    order.Items = order.Items[:len(updatedItems)]

    return nil
}