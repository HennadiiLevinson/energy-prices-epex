package main

import (
	"challenge.zaehlerfreunde.com/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/energy_cost", handlers.EnergyCostHandler)

	r.Run() // Listen and serve on 0.0.0.0:8080
}