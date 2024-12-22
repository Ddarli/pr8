package models

import "time"

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
	Order struct {
		ID       int
		Customer string
		Date     time.Time
		Total    float32
	}
)
