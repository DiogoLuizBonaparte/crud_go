package main

import (
	"log"

	"github.com/DiogoLuizBonaparte/go_api.git/src/controller/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
