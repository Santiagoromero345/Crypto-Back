package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	application "crypto-back/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AddAllowMethods("GET")

	r.Use(cors.New(corsConfig))

	r.GET("/coins", func(c *gin.Context) {
		coins, err := application.GetCoinsData()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		// Convert the string to a byte array
		c.Data(200, "application/json", []byte(coins))
	})

	r.Run()
}
