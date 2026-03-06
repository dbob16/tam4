package baskets

import (
	"errors"
	"database/sql"
	"slices"
	"maps"
	"cmp"

	"git.dilangilluly.us/dbob16/tam4/apiserver/apikeys"
	"git.dilangilluly.us/dbob16/tam4/apiserver/localdb"
)

func GetAllBaskets(api_key string) ([]Basket, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return nil, errors.New("Invalid API Key.")
	}

	var rtnBaskets []Basket

	results, err := db.Query("SELECT prefix, basket_id, description, donors, winning_ticket FROM baskets")
	if err != nil {
		return nil, errors.New("Error getting data from database.")
	}

	for results.Next() {
		var basket Basket
		results.Scan(&basket.Prefix, &basket.BasketID, &basket.Description, &basket.Donors, &basket.WinningTicket)
		rtnBaskets = append(rtnBaskets, basket)
	}

	return rtnBaskets, nil
}

func GetPrefixBaskets(api_key, prefix string) ([]Basket, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return nil, errors.New("Invalid API Key.")
	}

	var rtnBaskets []Basket

	results, err := db.Query("SELECT prefix, basket_id, description, donors, winning_ticket FROM baskets WHERE prefix = $1 ORDER BY basket_id ASC", prefix)
	if err != nil {
		return nil, errors.New("Error getting data from database.")
	}

	for results.Next() {
		var basket Basket
		results.Scan(&basket.Prefix, &basket.BasketID, &basket.Description, &basket.Donors, &basket.WinningTicket)
		rtnBaskets = append(rtnBaskets, basket)
	}

	return rtnBaskets, nil
}

func GetSingleBasket(api_key, prefix string, basket_id int) (Basket, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return Basket{}, errors.New("Invalid API Key.")
	}

	var rtnBasket Basket

	err := db.QueryRow("SELECT prefix, basket_id, description, donors, winning_ticket FROM baskets WHERE prefix = $1 AND basket_id = $2", prefix, basket_id).Scan(&rtnBasket.Prefix, &rtnBasket.BasketID, &rtnBasket.Description, &rtnBasket.Donors, &rtnBasket.WinningTicket)
	if err != nil {
		if err == sql.ErrNoRows {
			return Basket{Prefix: prefix, BasketID: basket_id, Description: "", Donors: "", WinningTicket: 0}, nil
		} else {
			return Basket{}, err
		}
	}

	return rtnBasket, nil
}

func GetRangeBaskets(api_key, prefix string, id_from, id_to int) ([]Basket, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return nil, errors.New("Invalid API Key.")
	}

	mapBaskets := make(map[int]Basket)

	for i := id_from; i <= id_to; i++ {
		mapBaskets[i] = Basket{Prefix: prefix, BasketID: i, Description: "", Donors: "", WinningTicket: 0}
	}

	results, err := db.Query(`SELECT prefix, basket_id, description, donors, winning_ticket FROM baskets
		WHERE prefix = $1 AND basket_id BETWEEN $2 AND $3`, prefix, id_from, id_to)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var basket Basket
		results.Scan(&basket.Prefix, &basket.BasketID, &basket.Description, &basket.Donors, &basket.WinningTicket)
		mapBaskets[basket.BasketID] = basket
	}

	rtnBaskets := slices.Collect(maps.Values(mapBaskets))
	slices.SortFunc(rtnBaskets, func(a, b Basket) int {
		return cmp.Compare(a.BasketID, b.BasketID)
	})

	return rtnBaskets, nil
}

func PostBaskets(api_key string, baskets []Basket) ([]Basket, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return nil, errors.New("Invalid API Key.")
	}

	for _, b := range baskets {
		_, err := db.Exec(`INSERT INTO baskets (prefix, basket_id, description, donors, winning_ticket)
			VALUES ($1, $2, $3, $4, $5) ON CONFLICT (prefix, basket_id)
			DO UPDATE SET description = EXCLUDED.description, donors = EXCLUDED.donors`, b.Prefix, b.BasketID, b.Description, b.Donors, b.WinningTicket)
		if err != nil {
			return nil, err
		}
	}

	return baskets, nil
}
