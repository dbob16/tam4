package tickets

type Ticket struct {
	Prefix string `json:"prefix"`
	TicketID int `json:"ticket_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Preference string `json:"preference"`
}
