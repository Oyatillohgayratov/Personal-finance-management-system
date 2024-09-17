package v1

import "github.com/gin-gonic/gin"

func (h *handlerV1) GetRepoting(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get Reporting"})
}


func (h *handlerV1) GetRepotingByCategory(c *gin.Context) {
    c.JSON(200, gin.H{"message": "Get Expenses"})
}