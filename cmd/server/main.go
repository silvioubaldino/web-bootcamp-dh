package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"os"
	"web-bootcamp-dh/cmd/server/controller"
	"web-bootcamp-dh/docs"
	"web-bootcamp-dh/internal/products"

	"github.com/joho/godotenv"
)

//@title MELI Bootcampo API
//@version 1.0
//@description This API handle MELI products

//@contact.name API Support
//@contact.url https://developers.mercadolibre.com.ar.support
func main() {
	err := godotenv.Load("/Users/sineto/Documents/web-bootcamp-dh/env.env")
	if err != nil {
		log.Fatal(err)
	}

	Repository := products.NewUserRepository()
	Service := products.NewUserService(Repository)
	Controller := controller.NewUserController(Service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := router.Group("/user")
	group.GET("/", Controller.GetAll())
	group.POST("/", Controller.Save())
	group.PUT("/:id", Controller.Update())
	group.DELETE("/:id", Controller.Delete())
	group.PATCH("/:id", Controller.UpdateSobrenomeIdade())
	router.Run()
}
