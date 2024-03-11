package main

import (
	"github.com/Mikael88/go-digitalent-assignment-2/controllers/ordercontroller"
	"github.com/Mikael88/go-digitalent-assignment-2/models"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    models.ConnectDatabase()

    r.POST("/api/orders", ordercontroller.CreateOrder)
    r.GET("/api/orders", ordercontroller.GetOrders)
    r.PUT("/api/orders/:orderId", ordercontroller.UpdateOrder)
    r.DELETE("/api/orders/:orderId", ordercontroller.DeleteOrder)

    r.Run()
}
