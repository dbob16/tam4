package tickets

import (
	"database/sql"
	"errors"
	"maps"
	"slices"

	"git.dilangilluly.us/dbob16/tam4/apiserver/apikeys"
	"git.dilangilluly.us/dbob16/tam4/apiserver/localdb"
)

func GetAllTickets(api_key string) ([]Ticket, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return nil, errors.New("Invalid API Key.")
	}

	var rtnTickets []Ticket

	results, err := db.Query("SELECT prefix, ticket_id, first_name, last_name, phone_number, preference FROM tickets")
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var ticket Ticket
		results.Scan(&ticket.Prefix, &ticket.TicketID, &ticket.FirstName, &ticket.LastName, &ticket.PhoneNumber, &ticket.Preference)
		rtnTickets = append(rtnTickets, ticket)
	}

	return rtnTickets, nil
}

func GetPrefixTickets(api_key string, prefix string) ([]Ticket, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return nil, errors.New("Invalid API Key.")
	}

	var rtnTickets []Ticket

	results, err := db.Query("SELECT prefix, ticket_id, first_name, last_name, phone_number, preference FROM tickets WHERE prefix = $1", prefix)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var ticket Ticket
		results.Scan(&ticket.Prefix, &ticket.TicketID, &ticket.FirstName, &ticket.LastName, &ticket.PhoneNumber, &ticket.Preference)
		rtnTickets = append(rtnTickets, ticket)
	}

	return rtnTickets, nil
}

func GetSingleTicket(api_key string, prefix string, ticket_id int) (Ticket, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return Ticket{}, errors.New("Invalid API Key.")
	}

	var rtnTicket Ticket

	if err := db.QueryRow(`SELECT prefix, ticket_id, first_name, last_name, phone_number, preference
		FROM tickets WHERE prefix = $1 AND ticket_id = $2`, prefix, ticket_id).Scan(&rtnTicket.Prefix, &rtnTicket.TicketID,
			&rtnTicket.FirstName, &rtnTicket.LastName, &rtnTicket.PhoneNumber, &rtnTicket.Preference); err != nil {
		if err == sql.ErrNoRows {
			return Ticket{Prefix: prefix, TicketID: ticket_id}, nil
		} else {
			return Ticket{}, err
		}
	}

	return rtnTicket, nil
}

func GetRangeTickets(api_key string, prefix string, id_from int, id_to int) ([]Ticket, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return nil, errors.New("Invalid API Key.")
	}

	rtnTickets := make(map[int]Ticket)

	for i := id_from; i <= id_to; i++ {
		rtnTickets[i] = Ticket{Prefix: prefix, TicketID: i, FirstName: "", LastName: "", PhoneNumber: "", Preference: ""}
	}

	results, err := db.Query("SELECT prefix, ticket_id, first_name, last_name, phone_number, preference FROM tickets WHERE prefix = $1", prefix)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var ticket Ticket
		results.Scan(&ticket.Prefix, &ticket.TicketID, &ticket.FirstName, &ticket.LastName, &ticket.PhoneNumber, &ticket.Preference)
		rtnTickets[ticket.TicketID] = ticket
	}

	rtnValues := slices.Collect(maps.Values(rtnTickets))

	return rtnValues, nil
}

func PostTickets(api_key string, tickets []Ticket) ([]Ticket, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return nil, errors.New("Invalid API Key.")
	}

	for _, t := range tickets {
		_, err := db.Exec(`INSERT INTO tickets (prefix, ticket_id, first_name, last_name, phone_number, preference) VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT (prefix, ticket_id) DO UPDATE SET first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name, phone_number = EXCLUDED.phone_number,
			preference = EXCLUDED.preference`, t.Prefix, t.TicketID, t.FirstName, t.LastName, t.PhoneNumber, t.Preference)
		if err != nil {
			return nil, err
		}
	}

	return tickets, nil
}
