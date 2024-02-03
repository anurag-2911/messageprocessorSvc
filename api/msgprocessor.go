package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Content string `json:"content"`
}

func Processor() {
	router := gin.Default()
	router.POST("/process", func(c *gin.Context) {
		log.Println("request on /process end point of msgprocessorSvc")
		processorHandler(c)
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
