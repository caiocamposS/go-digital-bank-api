package response

type AccountResponse struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	Number  string `json:"number"`
	Balance int64  `json:"balance"`
	Type    string `json:"type"`
}
