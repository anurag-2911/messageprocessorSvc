package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
	"log"
	"net/http"
)

type Message struct {
	Content string `json:"content"`
}

func Processor() {
	router := gin.Default()
	/*
		This code adds Prometheus middleware to the Gin application,
		which will automatically expose metrics at /metrics endpoint for Prometheus to scrape
	*/
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	router.POST("/process", func(c *gin.Context) {
		log.Println("request on /process end point of msgprocessorSvc")
		processorHandler(c)
	})
	router.GET("/ping", func(c *gin.Context) {
		log.Println("ping request on msgprocessorSvc")
		c.JSON(http.StatusOK, gin.H{"message": "ping working on msgprocessorSvc"})
	})
	router.Run(":8081")

}
func processorHandler(c *gin.Context) {
	var msg Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	processedMessage := " processed >> " + msg.Content
	c.JSON(http.StatusOK, gin.H{"message": processedMessage})
}
