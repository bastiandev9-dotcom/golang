package models

type Service struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
	ImageURL    *string `json:"image_url,omitempty"`
	CreatedAt   string  `json:"created_at"`
}

type CreateServiceInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
	ImageURL    *string `json:"image_url"`
}

type UpdateServiceInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
	ImageURL    *string `json:"image_url"`
}
