package profilesdto

type ProfileResponse struct {
	ID     int    `json:"id"`
	Image  string `json:"image"`
	UserID int    `json:"user_id"`
}
