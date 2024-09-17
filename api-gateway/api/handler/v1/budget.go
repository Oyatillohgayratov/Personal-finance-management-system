package v1

import (
	"api-gateway/protos/budget"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) CreateBudgetForCategory(c *gin.Context) {

	h.service.Budget().GetBudgets(c,&budget.GetBudgetsRequest{})
	c.JSON(200, gin.H{"message": "category created successfully"})
}

func (h *handlerV1) GetBudgets(c *gin.Context) {
	c.JSON(200, gin.H{"message": "get budgets successfully"})
}

func (h *handlerV1) PutBudget(c *gin.Context) {
	id := c.Param("budgetsId")
	c.JSON(200, gin.H{"message": id})
}
