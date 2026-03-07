package drawing

import (
	"errors"
	"slices"
	"cmp"
	"maps"

	"git.dilangilluly.us/dbob16/tam4/apiserver/localdb"
	"git.dilangilluly.us/dbob16/tam4/apiserver/apikeys"
)

func GetRangeDrawing(api_key, prefix string, id_from, id_to int) ([]DrawingEntry, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return nil, errors.New("Invalid API Key.")
	}

	rtnDrawing := make(map[int]DrawingEntry)

	for i := id_from; i <= id_to; i++ {
		rtnDrawing[i] = DrawingEntry{Prefix: prefix, BasketID: i, Description: "", WinningTicket: 0, WinnerName: "No Winner", PhoneNumber: ""}
	}

	results, err := db.Query(`SELECT prefix, basket_id, description, winning_ticket, winner_name, phone_number FROM drawing
		WHERE prefix = $1 AND basket_id BETWEEN $2 AND $3`, prefix, id_from, id_to)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var drawing DrawingEntry
		results.Scan(&drawing.Prefix, &drawing.BasketID, &drawing.Description, &drawing.WinningTicket, &drawing.WinnerName, &drawing.PhoneNumber)
		rtnDrawing[drawing.BasketID] = drawing
	}

	sliceDrawing := slices.Collect(maps.Values(rtnDrawing))
	slices.SortFunc(sliceDrawing, func(a, b DrawingEntry) int {
		return cmp.Compare(a.BasketID, b.BasketID)
	})

	return sliceDrawing, nil
}

func PostDrawingEntries(api_key string, drawings []DrawingEntry) ([]DrawingEntry, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return nil, errors.New("Invalid API Key.")
	}

	for _, d := range drawings {
		_, err := db.Exec(`INSERT INTO baskets (prefix, basket_id, winning_ticket) VALUES ($1, $2, $3)
			ON CONFLICT (prefix, basket_id) DO UPDATE SET winning_ticket = EXCLUDED.winning_ticket`, d.Prefix, d.BasketID, d.WinningTicket)
		if err != nil {
			return nil, err
		}
	}

	return drawings, nil
}
