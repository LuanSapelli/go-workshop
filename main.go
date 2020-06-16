package main

import (
	"fmt"

	"go-workshop/src"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.New()
	health := engine.Group("/")

	userRoute := engine.Group("/user")
	debtRoute := engine.Group("/debt")

	//User
	userRoute.GET("/", src.GetUsers)
	userRoute.GET("/:id", src.GetUser)
	userRoute.GET("/:id/debts", src.GetUserDebt)
	userRoute.POST("/", src.PostUser)
	userRoute.PUT("/:id", src.PutUser)
	userRoute.DELETE("/:id", src.DeleteUser)

	//debt
	debtRoute.GET("/", src.GetDebts)
	debtRoute.GET("/:id", src.GetDebt)
	debtRoute.POST("/", src.PostDebt)
	debtRoute.PUT("/:id", src.PutDebt)
	debtRoute.DELETE("/:id", src.DeleteDebt)

	src.AutoMigration()

	health.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Go healthy!",
		})
	})

	engine.Run(fmt.Sprintf(":8088"))
}
