package execute

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExecuteRequestBody struct {
	Context   []string `json:"context" binding:"required"`
	Namespace string   `json:"namespace" binding:"required"`
	Query     string   `json:"query" binding:"required"`
}

func Execute(c *gin.Context) {
	var requestBody ExecuteRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(requestBody)

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}
