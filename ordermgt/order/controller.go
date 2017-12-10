package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	OrderService Service `inject:"inline"`
}

func (ctrl *Controller) CreateOrderController(c *gin.Context) {

	var request CreateOrderRequest

	c.BindJSON(&request)

	fmt.Println(request)

	response, err := ctrl.OrderService.CreateOrderService(request)

	if err != nil {
		// TODO Handle error
	}

	c.JSON(http.StatusCreated, response)
}

func (ctrl *Controller) GetOrderById(c *gin.Context) {

	response, err := ctrl.OrderService.GetOrderByIdService(c.Param("id"))

	if err != nil {
		// TODO Handle error
	}

	c.JSON(http.StatusOK, response)
}
