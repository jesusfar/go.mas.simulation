package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrderController(c *gin.Context) {

	var request CreateOrderRequest

	c.BindJSON(&request)

	fmt.Println(request)

	response, err := CreateOrderService(request)

	if err != nil {
		// return response
	}

	c.JSON(http.StatusCreated, response)
}
