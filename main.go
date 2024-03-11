package main

import (
	"github.com/Mikael88/go-digitalent-assignment-2/controllers/productcontroller"
	"github.com/Mikael88/go-digitalent-assignment-2/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.POST("/api/orders", productcontroller.CreateOrder)
	r.GET("/api/orders", productcontroller.GetOrders)
	r.GET("/api/orders/:orderId", productcontroller.GetOrderByID)
	r.PUT("/api/orders/:orderId", productcontroller.UpdateOrder)
	r.DELETE("/api/orders/:orderId", productcontroller.DeleteOrder)

	r.Run(":8080")
}
