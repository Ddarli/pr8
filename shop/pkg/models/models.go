package models

type (
	OrderRequest struct {
		ID       int
		Name     string
		Quantity int
	}
	OrderResponse struct {
		ID      int
		Success bool
	}
)
