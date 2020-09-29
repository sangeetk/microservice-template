package rest

import (
	"context"
	"fmt"

	"github.com/sangeetk/microservice-template/config"
	"github.com/sangeetk/microservice-template/data"
	"github.com/sangeetk/microservice-template/service"

	"github.com/gin-gonic/gin"
)

var s service.Service

func Server(ctx context.Context, c config.AppConfig) {
	r := gin.Default()

	r.POST("/greet", GreetEndpoint)

	r.Run(fmt.Sprintf(":%d", c.ApiServerPort))
}

func GreetEndpoint(c *gin.Context) {
	var req data.GreetRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, nil)
		return
	}

	res, err := s.Greet(c, &req)
	if err != nil {
		c.JSON(400, nil)
		return
	}

	c.JSON(200, res)
}
