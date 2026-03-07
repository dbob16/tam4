package drawing

type DrawingEntry struct {
	Prefix string `json:"prefix"`
	BasketID int `json:"basket_id"`
	Description string `json:"description"`
	WinningTicket int `json:"winning_ticket"`
	WinnerName string `json:"winner_name"`
	PhoneNumber string `json:"phone_number"`
}
