package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	log.Println("Starting recsys server...")
	r.Run(":8080")
}
