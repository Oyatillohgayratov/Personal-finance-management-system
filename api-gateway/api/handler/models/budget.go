package models

type CreateBudgetForCategoryReq struct {
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type CreateBudgetForCategoryResp struct {
	Message string `json:"message"`
	BudgeId int    `json:"budge_id"`
}

type GetBudgetsResp struct {
	BudgetId int     `json:"budget_id"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Spent    int     `json:"spent"`
	Currency string  `json:"currency"`
}