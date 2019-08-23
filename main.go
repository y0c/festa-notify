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

func main() {
	godotenv.Load()
	addr := ":4000"
	stage := os.Getenv("GIN_MODE")

	g := gin.New()
	g.POST("/sendMail", handler.SendMailHandler)

	if stage == "release" {
		log.Fatal(gateway.ListenAndServe(addr, g))
	} else {
		log.Fatal(http.ListenAndServe(addr, g))
	}
}
