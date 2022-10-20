package profilesdto

type CreateProfileRequest struct {
	Image  string `json:"image"`
	UserID int    `json:"user_id"`
}

type UpdateProfileRequest struct {
	Image string `json:"image"`
}
