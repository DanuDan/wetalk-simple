package messagesdto

type CreateMessageRequest struct {
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
}

type UpdateMessageRequest struct {
	Message string `json:"message"`
}
