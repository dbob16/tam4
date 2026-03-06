package baskets

type Basket struct {
	Prefix string `json:"prefix"`
	BasketID int `json:"basket_id"`
	Description string `json:"description"`
	Donors string `json:"donors"`
	WinningTicket int `json:"winning_ticket"`
}
