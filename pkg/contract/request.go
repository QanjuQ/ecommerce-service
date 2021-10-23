package contract

type CreateUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateItemRequest struct {
	Name string `json:"name"`
}

type AddItemToCartRequest struct {
	Id       uint `json:"id"`
	Quantity int `json:"quantity"`
}
