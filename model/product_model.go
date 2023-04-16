package model

type CreateProductRequest struct {
	Title       string `json:"title" valid:"required~Title is required"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

type UpdateProductRequest struct {
	Title       string `json:"title" valid:"required~Title is required"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

type ProductResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}
