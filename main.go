package main

import (
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/y0c/festa-notify/handler"
	"log"
	"net/http"
	"os"
)

func handleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				c.JSON(500, gin.H{
					"message": "Internal Server Error",
				})
			}
		}()
		c.Next()
	}
}
func main() {
	godotenv.Load()
	addr := ":4000"
	stage := os.Getenv("GIN_MODE")

	r := gin.New()
	r.Use(handleErrors())
	r.POST("/sendMail", handler.SendMailHandler)

	if stage == "release" {
		log.Fatal(gateway.ListenAndServe(addr, r))
	} else {
		log.Fatal(http.ListenAndServe(addr, r))
	}
}
